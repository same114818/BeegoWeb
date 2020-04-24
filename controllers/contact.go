package controllers

import (
	"BeegoWeb/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type ContactController struct {
	beego.Controller
}

func (c *ContactController) URLMapping() {
	c.Mapping("SendMessage", c.SendMessage)
}

// @router /contact [post]
func (c *ContactController) SendMessage() {
	name := c.GetString("name")
	email := c.GetString("email")
	subject := c.GetString("subject")
	message := c.GetString("message")
	contact := models.Contact{
		Name:    name,
		Email:   email,
		Subject: subject,
		Message: message,
	}
	err := models.CreateContact(&contact)
	if err == nil {
		jsonBytes, _ := json.Marshal(contact)
		c.Data["json"] = string(jsonBytes)
	}
	c.ServeJSON()
}
