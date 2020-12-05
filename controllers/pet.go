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
	pet_type := p.GetString("pet_type","")
	pet_age,_ := p.GetInt64("pet_age",0)
	pet_sex := p.GetString("","")
	pet := new(models.Pet)
	user_id := p.getUid()
	if(user_id <= 0) {
		p.ResponseData(1,"uid异常",nil)
	}

	if(head_img == "" || pet_name == "" || pet_type== "" || pet_age == 0 || pet_sex == "") {
		p.ResponseData(1,"数据异常",nil)
	}

	head_img_path,err := Tools.UpImg(head_img)
	if(err != nil) {
		p.ResponseData(1,"上传失败",nil)
	}

	pet.UserId = int64(user_id)
	pet.HeadImg = head_img_path
	pet.PetName = pet_name
	pet.PetAge = pet_age
	pet.PetType = pet_type
	pet.PetSex = pet_sex

	o := orm.NewOrm()
	if num,err := o.Insert(&pet);err !=nil {
		p.ResponseData(1,"宠物添加失败",nil)
	}else{
		if(num > 0) {
			p.ResponseData(0,"suss",pet)
		}else{
			p.ResponseData(1,"宠物添加失败",pet)
		}
	}

}
