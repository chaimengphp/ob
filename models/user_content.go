package models

import "github.com/astaxie/beego/client/orm"

func init() {
	orm.RegisterModel(new(UserContent))
}

type UserContent struct {
	Id int64 `json:"id"`
	Content string `json:"content"`
	Imglist string `json:"imglist"`
	Pubtime int64 `json:"pubtime"`
	User *User `orm:"rel(fk)"`

}

func (uc *UserContent) TableName() string {
	return "ob_user_content"
}