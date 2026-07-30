package lib

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509/pkix"
	"encoding/hex"
	"io"
	"os"

	"github.com/beego/beego/v2/core/logs"
	"github.com/glaolle/openvpn-ui/models"
	"github.com/glaolle/openvpn-ui/state"
	"github.com/kemsta/go-easyrsa/v2/pki"
)

// SelectCurve returns the standard elliptic.Curve based on a string name
func SelectCurve(name string) elliptic.Curve {
	switch name {
	case "P-256":
		return elliptic.P256()
	case "P-384":
		return elliptic.P384()
	case "P-521":
		return elliptic.P521()
	default:
		return nil
	}
}

func CreateDefaultPKI(pkiPath string) (*pki.PKI, error) {
	cfg := models.EasyRSAConfig{Profile: "default"}
	err := cfg.Read("Profile")
	if err != nil {
		logs.Error(err)
	}

	// EasyRSAReqEmail:    "sweet@home.net",
	curve := SelectCurve(cfg.EasyRSACurve)

	myConfig := pki.Config{
		// Key generation defaults
		KeyAlgo: pki.KeyAlgo(cfg.EasyRSAAlgo), // rsa | ecdsa | ed25519
		KeySize: cfg.EasyRSAKeySize,           // RSA key size
		Curve:   curve,

		// Certificate validity
		DefaultDays:   cfg.EasyRSACertExpire, // End-entity certificates
		CADays:        cfg.EasyRSACaExpire,   // CA certificate
		CRLDays:       cfg.EasyRSACrlDays,    // CRL validity
		PreExpiryDays: cfg.EasyRSACertRenew,  // ShowExpiring window

		// Subject DN
		//DNMode: pki.DNModeCNOnly, // cn_only - DNModeCNOnly | org - DNModeOrg
		//DNMode: pki.DNModeOrg, // cn_only - DNModeCNOnly | org - DNModeOrg
		DNMode: pki.DNMode(cfg.EasyRSADN), // cn_only - DNModeCNOnly | org - DNModeOrg
		SubjTemplate: pkix.Name{ // Template for org mode
			Country:            []string{cfg.EasyRSAReqCountry},
			Organization:       []string{cfg.EasyRSAReqOrg},
			OrganizationalUnit: []string{cfg.EasyRSAReqOu},
			Locality:           []string{cfg.EasyRSAReqCity},
			Province:           []string{cfg.EasyRSAReqProvince},
			CommonName:         cfg.EasyRSAReqCn,
		},

		// Key protection
		NoPass:        true, // true = store keys unencrypted
		CAPassphrase:  "",   // CA key passphrase
		KeyPassphrase: "",   // Default key passphrase

		// Serial numbers
		SequentialSerial: false, // false = random 128-bit (default)

		// Storage
		CAName: "ca", // CA entity name in storage
	}

	p, err := pki.NewWithFS(pkiPath, myConfig)
	if err != nil {
		logs.Error(err)
	}
	logs.Info("Successfully open PKI")

	return p, nil
}

func GenerateTA(taPath string) {
	// Create the ta.key file
	file, err := os.Create(taPath)
	if err != nil {
		logs.Error("Error creating file:", err)
		return
	}
	defer file.Close()

	// Generate the OpenVPN static key header
	header := []byte(
		"#\n" +
			"# 2048 bit OpenVPN static cipher seed\n" +
			"#\n" +
			"-----BEGIN OpenVPN Static key V1-----\n")
	_, err = file.Write(header)
	if err != nil {
		logs.Error("Error writing header:", err)
		return
	}

	// Generate exactly 2048 bits (256 bytes) of random data
	keyBytes := make([]byte, 256)
	_, err = io.ReadFull(rand.Reader, keyBytes)
	if err != nil {
		logs.Error("Error generating random bytes:", err)
		return
	}

	// Format bytes into OpenVPN's hex block
	hexString := hex.EncodeToString(keyBytes)
	for i := 0; i < len(hexString); i += 32 {
		end := i + 32
		if end > len(hexString) {
			end = len(hexString)
		}
		file.WriteString(hexString[i:end] + "\n")
	}

	// Write the OpenVPN footer
	footer := []byte("-----END OpenVPN Static key V1-----\n")
	_, err = file.Write(footer)
	if err != nil {
		logs.Error("Error writing footer:", err)
		return
	}

	logs.Info("Successfully generated ta.key")
}

func GenerateCA() {
	// Build a CA
	p := state.GlobalPKI
	ca, err := p.BuildCA()
	if err != nil {
		logs.Error(err)
	}
	logs.Info("CA created: %s", ca.Name)
}

func GenerateServerCert() {
	p := state.GlobalPKI
	// Issue a server certificate
	server, err := p.BuildServerFull("server")
	if err != nil {
		logs.Error(err)
	}
	logs.Info("Server cert issued: %s", server.Name)
}

func GenerateCRL() {
	// Generate CRL
	p := state.GlobalPKI
	_, err := p.GenCRL()
	if err != nil {
		logs.Error(err)
	}
	logs.Info("Generate CRL")
}

func GenerateClientCert(name string) {
	// Issue a client certificate
	p := state.GlobalPKI
	client, err := p.BuildClientFull(name)
	if err != nil {
		logs.Error(err)
	}
	logs.Info("Client cert issued: %s\n", client.Name)
}
