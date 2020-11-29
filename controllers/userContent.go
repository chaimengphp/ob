package controllers

import (
	"github.com/astaxie/beego/client/orm"
	beego "github.com/astaxie/beego/server/web"
	"obapi/models"
	"strconv"
)

type UserContentController struct {
	beego.Controller
}

func (uc *UserContentController) Detail() {
	id:= uc.Ctx.Input.Param(":id")
	content_id,_ := strconv.ParseInt(id,0,10)
	if content_id <= 0 {
		resp := reponseJson{1,"内容不存在",nil}
		uc.Data["json"] = resp
		uc.ServeJSON()
	}
	o := orm.NewOrm()
	user_content_info := models.UserContent{}
	qs := o.QueryTable(new(models.UserContent))
	err := qs.Filter("id",id).One(&user_content_info,"id","content","imglist")
	if err == orm.ErrNoRows {
		resp := reponseJson{1,"内容不存在",nil}
		uc.Data["json"] = resp
		uc.ServeJSON()
	}

	resp := reponseJson{0,"suss",user_content_info}
	uc.Data["json"] = resp
	uc.ServeJSON()


}
