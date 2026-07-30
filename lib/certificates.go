package lib

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"

	//"github.com/glaolle/openvpn-ui/lib"
	"github.com/glaolle/openvpn-ui/models"
	"github.com/glaolle/openvpn-ui/state"
	"github.com/kemsta/go-easyrsa/v2/cert"
	"github.com/kemsta/go-easyrsa/v2/pki"
)

// Cert
// https://groups.google.com/d/msg/mailing.openssl.users/gMRbePiuwV0/wTASgPhuPzkJ
type Cert struct {
	EntryType   string
	Expiration  string
	ExpirationT time.Time
	IsExpiring  bool
	Revocation  string
	RevocationT time.Time
	Serial      string
	FileName    string
	Details     *Details
}

type Details struct {
	Name             string
	CN               string
	Country          string
	State            string
	City             string
	Organisation     string
	OrganisationUnit string
	Email            string
	LocalIP          string
	TFAName          string
}

func ReadCerts(path string) ([]*Cert, error) {
	certs := make([]*Cert, 0)
	text, err := os.ReadFile(path)
	if err != nil {
		return certs, err
	}
	lines := strings.Split(trim(string(text)), "\n")
	for _, line := range lines {
		fields := strings.Split(trim(line), "\t")
		if len(fields) != 6 {
			return certs,
				fmt.Errorf("incorrect number of lines in line: \n%s\n. Expected %d, found %d",
					line, 6, len(fields))
		}
		expT, _ := time.Parse("060102150405Z", fields[1])
		expTA := time.Now().AddDate(0, 0, 30).After(expT) // If cer will expire in 30 days, raise this flag
		//logs.Debug("ExpirationT: %v, IsExpiring: %v", expT, expTA) // logging
		revT, _ := time.Parse("060102150405Z", fields[2])
		c := &Cert{
			EntryType:   fields[0],
			Expiration:  fields[1],
			ExpirationT: expT,
			IsExpiring:  expTA,
			Revocation:  fields[2],
			RevocationT: revT,
			Serial:      fields[3],
			FileName:    fields[4],
			Details:     parseDetails(fields[5]),
		}
		certs = append(certs, c)
	}

	return certs, nil
}

func parseDetails(d string) *Details {
	details := &Details{}
	lines := strings.Split(trim(d), "/")
	for _, line := range lines {
		if strings.Contains(line, "") {
			fields := strings.Split(trim(line), "=")
			switch fields[0] {
			case "name":
				details.Name = fields[1]
			case "CN":
				details.CN = fields[1]
			case "C":
				details.Country = fields[1]
			case "ST":
				details.State = fields[1]
			case "L":
				details.City = fields[1]
			case "O":
				details.Organisation = fields[1]
			case "OU":
				details.OrganisationUnit = fields[1]
			case "emailAddress":
				details.Email = fields[1]
			case "LocalIP":
				details.LocalIP = fields[1]
			case "2FAName":
				details.TFAName = fields[1]
			default:
				if line != "" && !strings.Contains(line, "name") && !strings.Contains(line, "LocalIP") {
					logs.Warn(fmt.Sprintf("Undefined entry: %s", line))
				}
			}
		}
	}
	return details
}

func trim(s string) string {
	return strings.Trim(strings.Trim(s, "\r\n"), "\n")
}

// func rewriteIndex(name string, appendSubstr string) {
// 	filename := state.GlobalCfg.OVConfigPath + "/pki/index.txt"
// 	appendSubstr = "/LocalIP=" + appendSubstr

// 	// 1. Читаем весь файл
// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		logs.Error("Не удалось прочитать файл: %s", err)
// 	}

// 	// 2. Разбиваем содержимое файла на строки
// 	lines := strings.Split(string(data), "\n")
// 	var updatedLines []string

// 	// 3. Перебираем строки и ищем подстроку
// 	for _, line := range lines {
// 		if strings.Contains(line, name) {
// 			// Добавляем подстроку к найденной строке
// 			line += appendSubstr
// 		}
// 		updatedLines = append(updatedLines, line)
// 	}

// 	// 4. Собираем строки обратно в один текст
// 	output := strings.Join(updatedLines, "\n")

// 	// 5. Перезаписываем файл новыми данными
// 	err = os.WriteFile(filename, []byte(output), 0644)
// 	if err != nil {
// 		logs.Error("Не удалось записать файл: %s", err)
// 	}

// 	logs.Info("Файл успешно обновлен!")
// }

func CreateCertificate(name string, staticip string, passphrase string, expiredays string, email string, country string, province string, city string, org string, orgunit string, tfaname string, tfaissuer string) error {
	logs.Info("Lib: Creating certificate with parameters: name=%s, staticip=%s, passphrase=%s, expiredays=%s, email=%s, country=%s, province=%s, city=%s, org=%s, orgunit=%s, tfaname=%s, tfaissuer=%s", name, staticip, passphrase, expiredays, email, country, province, city, org, orgunit, tfaname, tfaissuer)
	path := state.GlobalCfg.OVConfigPath + "/pki/index.txt"
	haveip := staticip != ""
	if staticip == "" {
		staticip = "dynamic.pool"
	}
	pass := passphrase != ""
	//logs.Info("Org set to: %v", org)

	existsError := errors.New("Error! There is already a valid or invalid certificate for the name \"" + name + "\"")
	certs, err := ReadCerts(path)
	if err != nil {
		logs.Error(err)
	}

	for _, v := range certs {
		if v.Details.Name == name {
			return existsError
		}
	}

	//Dump(certs)
	p := state.GlobalPKI

	if !pass { // if no passphrase
		logs.Info("No password")

		if !haveip {
			client, err := p.BuildClientFull(name)
			if err != nil {
				logs.Error(err)
			}
			logs.Info("Client cert issued: %s\n", client.Name)
		} else {
			client, err := p.BuildClientFull(name, pki.WithIPAddresses(net.ParseIP(staticip)))
			if err != nil {
				logs.Error(err)
			}
			logs.Info("Client cert issued: %s\n", client.Name)
		}
	} else { // if passphrase
		logs.Info("Password")

		if !haveip {
			client, err := p.BuildClientFull(name, pki.WithPassphrase(passphrase))
			if err != nil {
				logs.Error(err)
			}
			logs.Info("Client cert issued: %s\n", client.Name)
		} else {
			client, err := p.BuildClientFull(name, pki.WithPassphrase(passphrase), pki.WithIPAddresses(net.ParseIP(staticip)))
			if err != nil {
				logs.Error(err)
			}
			logs.Info("Client cert issued: %s\n", client.Name)
		}
	}

	if haveip {
		logs.Info("Client have IP")

		text := "ifconfig-push " + staticip + " 255.255.255.0"
		err = os.WriteFile(filepath.Join(state.GlobalCfg.OVConfigPath, "ccd", name), []byte(text), 0644)
		if err != nil {
			logs.Error(err)
		}
	}

	cert := models.StaticIP{
		CertName: name,
		StaticIP: staticip,
	}

	o := orm.NewOrm()
	if created, _, err := o.ReadOrCreate(&cert, "Name"); err == nil {
		if created {
			logs.Info("New cert \"" + cert.CertName + "\" created successfully.")
		} else {
			logs.Debug(cert)
		}
	} else {
		logs.Error(err)
	}

	//rewriteIndex(name, staticip)
	return nil
}

func RevokeCertificate(name string, serial string, tfaname string) error {
	p := state.GlobalPKI
	// Revoke by name
	err := p.Revoke(name, cert.ReasonKeyCompromise)
	if err != nil {
		logs.Error(err)
	}
	logs.Info("Client cert revoke: %s\n", name)
	// Regenerate CRL after revocation (automatic in Revoke*)
	//crlPEM, err := p.GenCRL()

	// Check if a certificate is revoked
	//revoked, err := p.IsRevoked(serial)
	return nil
}

func Restart() error {
	// TODO
	logs.Info("Restart OpenVPN server")
	return nil
}

func BurnCertificate(CN string, serial string, tfaname string) error {
	// TODO
	logs.Info("Delete cert")
	return nil
}

func RenewCertificate(name string, localip string, serial string, tfaname string) error {
	p := state.GlobalPKI
	// Renew a certificate (new cert, same key)
	client, err := p.Renew(name)
	if err != nil {
		logs.Error(err)
	}
	logs.Info("Client cert issued: %s\n", client.Name)
	// Find certificates expiring within 30 days
	//expiring, err := p.ShowExpiring(30)

	// Mark expired certs in the index
	//err = p.UpdateDB()
	return nil
}
