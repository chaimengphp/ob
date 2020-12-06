package controllers

import (
	"github.com/astaxie/beego/client/orm"
	"obapi/models"
	"strconv"
)

type UserContentController struct {
	BaseController
}

func (uc *UserContentController) Detail() {
	id:= uc.Ctx.Input.Param(":id")
	content_id,_ := strconv.ParseInt(id,0,10)
	if content_id <= 0 {
		uc.ResponseData(1,"参数为空",nil)
	}
	o := orm.NewOrm()
	user_content_info := models.UserContent{}
	qs := o.QueryTable(new(models.UserContent))
	err := qs.Filter("id",id).One(&user_content_info,"id","content","imglist")
	if err == orm.ErrNoRows {
		uc.ResponseData(1,"内容不存在",nil)
	}
	uc.ResponseData(0,"suss",user_content_info)

}
