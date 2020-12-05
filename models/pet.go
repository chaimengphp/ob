package models

import "github.com/astaxie/beego/client/orm"

var (
	PetList map[string]*User
)

type Pet struct {
	Id int64
	UserId int64
	HeadImg string
	PetName string
	PetType  string
	PetAge int64
	PetSex string
	User *User `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(User))
}
