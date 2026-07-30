package models

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
	//Sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

type StaticIP struct {
	Id int64
	//Profile string `orm:"size(64);unique" form:"Profile" valid:"Required;"`

	CertName string `orm:"size(64);unique" form:"CertName" valid:"Required;"`
	StaticIP string `orm:"size(64);" form:"StaticIP" valid:"Required;"`
}

// Insert wrapper
func (s *StaticIP) Insert() error {
	if _, err := orm.NewOrm().Insert(s); err != nil {
		return err
	}
	return nil
}

// Read wrapper
func (s *StaticIP) Read(fields ...string) error {
	if err := orm.NewOrm().Read(s, fields...); err != nil {
		return err
	}
	return nil
}

// Update wrapper
func (s *StaticIP) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(s, fields...); err != nil {
		return err
	}
	return nil
}

// Delete wrapper
func (s *StaticIP) Delete() error {
	if _, err := orm.NewOrm().Delete(s); err != nil {
		return err
	}
	return nil
}

func GetCertByCN(certname string) (*StaticIP, error) {
	cert := &StaticIP{CertName: certname}
	err := cert.Read("CertName")
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, errors.New("cert not found")
		}
		return nil, err
	}
	return cert, nil
}
