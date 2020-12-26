package models

import (
	"github.com/astaxie/beego/client/orm"
	"obapi/Tools"
)


var (
	PetList map[string]*Pet
)

func init() {
	orm.RegisterModel(new(Pet))
}


type Pet struct {
	Id int64 `json:"id"`
	HeadImg string `json:"head_img"`
	PetName string `json:"pet_name"`
	PetAge int64 `json:"pet_age"`
	PetSex string `json:"pet_sex"`
	PetType *PetType `orm:"rel(fk)" json:"pet_type"`
	User *User `orm:"rel(fk)" json:"user"`
}

func (u *Pet) TableName() string {
	return "ob_pet"
}

func (pt *Pet) GetOne(id int64) (pet Pet,err error) {
	o := orm.NewOrm()
	pet = Pet{Id:id}
	err = o.Read(&pet)
	return
}

func (pt *Pet) IsExists(name string) bool {
	o := orm.NewOrm()
	exists := o.QueryTable(pt).Filter("pet_name",name).Exist()
	return exists
}

func (pt *Pet) GetCount(user_id int64) (num int64) {
	o := orm.NewOrm()
	num,_ = o.QueryTable(pt).Filter("user_id",user_id).Count()
	return
}

func (pt *Pet) GetList(user_id int64) (pet []*Pet,err error) {
	o := orm.NewOrm()
	//user := User{Id:user_id}
	num,err := o.QueryTable(pt).Filter("user_id",user_id).RelatedSel().All(&pet)
	if err == nil && num>0 {
		for _,v := range pet {
			v.HeadImg = Tools.HandelImg(v.HeadImg)
		}
	}

	return
}


