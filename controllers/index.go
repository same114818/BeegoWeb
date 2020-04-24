package controllers

import (
	"BeegoWeb/models"

	"github.com/astaxie/beego"
)

// Index api
type IndexController struct {
	beego.Controller
}

func (c *IndexController) URLMapping() {
	c.Mapping("GetIndex", c.GetIndex)
	c.Mapping("GetAbout", c.GetAbout)
	c.Mapping("GetBlog", c.GetBlog)
	c.Mapping("GetContact", c.GetContact)
}

// @router / [get]
func (c *IndexController) GetIndex() {
	c.TplName = "index.html"
}

// @router /about [get]
func (c *IndexController) GetAbout() {
	user := models.QueryUser()
	userDetail := models.QueryUserDetail(user)
	userSkill := models.QueryUserSkill(user)
	c.Data["User"] = user
	c.Data["UserDetail"] = userDetail
	c.Data["UserSkill"] = userSkill
	c.TplName = "about.html"
}

// @router /blog [get]
func (c *IndexController) GetBlog() {
	c.TplName = "blog.html"
}

// @router /contact [get]
func (c *IndexController) GetContact() {
	c.TplName = "contact.html"
}
