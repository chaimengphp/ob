package models

import "github.com/astaxie/beego/client/orm"

var (
	PetTypeList map[string]*PetType
)

type PetType struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Pets []*Pet `orm:"reverse(many)",json:"pets"`
}

func init() {
	orm.RegisterModel(new(PetType))
}

func (u *PetType) TableName() string {
	return "ob_pet_type"
}

func (pt *PetType) GetOne(id int64) (pettype PetType,err error) {
	o := orm.NewOrm()
	pettype = PetType{Id:id}
	err = o.Read(&pettype)
	return
}