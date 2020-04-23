package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Contact struct {
	Id      int
	Name    string
	Email   string
	Subject string
	Message string
}

func init() {
	// register model
	orm.RegisterModel(new(Contact))
}

func CreateContact(contact *Contact) error {
	o := orm.NewOrm()

	_, err := o.Insert(contact)
	if err != nil {
		beego.Info(err)
	}
	return err
}
