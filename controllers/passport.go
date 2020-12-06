package controllers

import (
	"github.com/astaxie/beego/client/orm"
	"obapi/models"
	"obapi/Tools"
)

type PassportController struct {
	BaseController
}

//第三方登录
func (p *PassportController) OauthLogin() {

}

func (p *PassportController) Upinfo() {
	nike_name := p.GetString("nikename","")
	summary := p.GetString("summary","")
	file := p.GetString("head_img","")
	if file == "" && nike_name == "" && summary == "" {
		p.ResponseData(1,"数据无改动",nil)
	}
	user := new(models.User)
	if file != "" {
		img_pth,err := Tools.UpImg(file)
		if err != nil {
			p.ResponseData(1,"头像上传失败",nil)
		}
		user.HeadImg = img_pth;
	}

	if nike_name != "" {
		user.NikeName = nike_name
	}

	if summary != "" {
		user.Summary = summary
	}

	o:=orm.NewOrm()
	if num,err := o.Update(&user);err != nil {
		if num >0 {
			p.ResponseData(0,"suss",user)
		}
	}else{
		p.ResponseData(1,"修改失败",user)
	}

}

