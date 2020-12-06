package models

import "github.com/astaxie/beego/client/orm"

var (
	PetList map[string]*User
)

type Pet struct {
	Id int64 `json:"id"`
	UserId int64 `json:"user_id"`
	HeadImg string `json:"head_img"`
	PetName string `json:"pet_name"`
	PetType  string `json:"pet_type"`
	PetAge int64 `json:"pet_age"`
	PetSex string `json:"pet_sex"`
	User *User `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(User))
}
