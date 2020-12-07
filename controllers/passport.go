package controllers

import (
	"encoding/base64"
	"github.com/astaxie/beego/client/orm"
	"io/ioutil"
	"net/http"
	"obapi/models"
	"obapi/Tools"
)

type PassportController struct {
	BaseController
}

//第三方登录
func (p *PassportController) OauthLogin() {
	reg_from := p.GetString("reg_from","")
	nick_name := p.GetString("nickname","")
	head_img := p.GetString("headimgurl","")
	unionid := p.GetString("unionid","")
	if nick_name == "" || head_img == "" || unionid == "" {
		p.ResponseData(1,"参数异常",nil)
	}
	res,err := http.Get(head_img)
	if err !=nil {
		p.ResponseData(1,"图片读取失败",nil)
	}
	defer res.Body.Close()
	img_data,_ := ioutil.ReadAll(res.Body)
	imageBase64 := base64.StdEncoding.EncodeToString(img_data)
	img_path,err := Tools.UpImg(imageBase64)
	if err != nil {
		p.ResponseData(1,"头像上传失败",nil)
	}
	user := &models.User{

	}

}

func (p *PassportController) Upinfo() {
	nike_name := p.GetString("nikename","")
	summary := p.GetString("summary","")
	file := p.GetString("file","")
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

