package controllers

import (
	"github.com/astaxie/beego/client/orm"
	"obapi/Tools"
	"obapi/models"
)

type PetController struct {
	BaseController
}


func (p *PetController) AddPet() {
	head_img := p.GetString("head_img","")
	pet_name := p.GetString("pet_name","")
	pet_type_id,_ := p.GetInt64("pet_type",0)
	pet_age,_ := p.GetInt64("pet_age",0)
	pet_sex := p.GetString("pet_sex","")
	pet := new(models.Pet)
	pettype := new(models.PetType)
	user_id := p.GetUid()
	if(user_id <= 0) {
		p.ResponseData(1,"uid异常",nil)
		return
	}

	if(head_img == "" || pet_name == "" || pet_type_id == 0 || pet_age == 0 || pet_sex == "") {
		p.ResponseData(1,"数据异常",nil)
		return
	}

	if is_exists := pet.IsExists(pet_name);is_exists == true {
		p.ResponseData(1,"已存在相同的宠物名",nil)
		return
	}

	if pet_num := pet.GetCount(user_id);pet_num >=10 {
		p.ResponseData(1,"宠物数量已达上限",nil)
		return
	}

	pet_type_info,err := pettype.GetOne(pet_type_id)
	if err != nil {
		p.ResponseData(1,"非法的宠物品种类型",nil)
		return
	}

	head_img_path,err := Tools.UpImg(head_img)
	if(err != nil) {
		p.ResponseData(1,"宠物头像上传失败",nil)
		return
	}
	user := models.User{Id:user_id}
	pet.User = &user
	pet.HeadImg = head_img_path
	pet.PetName = pet_name
	pet.PetAge = pet_age
	pet.PetType = &pet_type_info
	pet.PetSex = pet_sex

	o := orm.NewOrm()
	if num,err := o.Insert(pet);err !=nil {
		p.ResponseData(1,"宠物添加失败",nil)
		return
	}else{
		if(num > 0) {
			pet.HeadImg = Tools.HandelImg(head_img_path)
			p.ResponseData(0,"suss",pet)
		}else{
			p.ResponseData(1,"宠物添加失败",nil)
		}
	}

}

func (p *PetController) PetList() {
	user_id := p.GetUid()
	if(user_id <= 0) {
		p.ResponseData(1,"请登录",nil)
		return
	}

	petmodel := new(models.Pet)
	petlist,err := petmodel.GetList(user_id)
	if err !=nil {
		p.ResponseData(0,"suss",nil)
		return
	}
	p.ResponseData(0,"suss",petlist)
}
