package models

import (
	"fmt"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type User struct {
	Id       int
	UserName string
	Password string
	Name     string
	Profile  string
	Email    string
	Phone    string
	Status   int `orm:"default(1)"`
}

type UserSkill struct {
	Id      int
	UserId  int
	SkillId int
	RateId  int
}

type Skill struct {
	Id        int
	SkillName string
}
type Rate struct {
	Id          int
	RatePercent string
	RateNumber  int
	Process     string `orm:"type(text)"`
}

type UserDetail struct {
	Id     int
	UserId int
	Detail string
}

type UserSkillSearchResult struct {
	SkillName   string
	RatePercent string
	Process     string
}

func init() {
	// register model
	orm.RegisterModel(new(User), new(UserSkill), new(Skill), new(Rate), new(UserDetail))
}

func QueryUser() (user User) {
	o := orm.NewOrm()
	user = User{Id: 1}
	err := o.Read(&user)
	if err != nil {
		beego.Info(err)
	}
	return
}

func QueryUserDetail(user User) (detailList []string) {
	var details []UserDetail
	o := orm.NewOrm()
	//qs:=o.QueryTable
	userDetail := new(UserDetail)
	_, _ = o.QueryTable(userDetail).Filter("user_id", user.Id).All(&details)
	for _, item := range details {
		detailList = append(detailList, item.Detail)
	}
	return
}

func QueryUserSkill(user User) (userskills []UserSkillSearchResult) {
	o := orm.NewOrm()
	num, _ := o.Raw("SELECT skill_name,rate_percent,rate_number,process FROM `user_skill` INNER JOIN skill on user_skill.id = skill.id INNER JOIN rate on user_skill.rate_id=rate.id where user_id=?", user.Id).QueryRows(&userskills)
	fmt.Println(num)
	return
}
