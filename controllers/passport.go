package controllers

import (
	"encoding/base64"
	"github.com/astaxie/beego/client/orm"
	"io/ioutil"
	"net/http"
	"obapi/models"
	"obapi/Tools"
	"time"
)

type PassportController struct {
	BaseController
}

//第三方登录
func (p *PassportController) OauthLogin() {
	reg_from := p.GetString("reg_from","")
	nick_name := p.GetString("nickname","")
	head_img := p.GetString("headimgurl","")
	oauth_uid := p.GetString("oauth_uid","")
	if nick_name == "" || head_img == "" || oauth_uid == "" {
		p.ResponseData(1,"参数异常",nil)
	}
	if is_intval := Tools.In_array(reg_from,[]string{"wx","qq"});is_intval == false {
		p.ResponseData(1,"来源错误",nick_name)
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
		DeviceId:p.GetDeviceId(),
		NikeName:nick_name,
		HeadImg:img_path,
		RegFrom:reg_from,
		OauthUid:oauth_uid,
		RegTime:time.Now().Unix(),
		LoginTime:time.Now().Unix(),
	}
	o := orm.NewOrm()
	user_info := p.GetOauthOne(reg_from,oauth_uid)
	if user_info !=nil {
		user.Id = user_info.Id
		_,errs := o.Update(&user)
		if errs != nil {
			p.ResponseData(1,"登录异常",nil)
		}
	}else{
		num,errs := o.Insert(&user)
		user.Id = num
		if errs != nil {
			p.ResponseData(1,"登录异常",nil)
		}
	}
	p.ResponseData(0,"登录成功",user)
}

func (p *PassportController) GetOauthOne(reg_form string,oauth_uid string) (info *models.User) {
	user := &models.User{RegFrom:reg_form,OauthUid:oauth_uid}
	o := orm.NewOrm()
	if err := o.Read(&user); err != nil {
		return nil
	}
	return user
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
		user.HeadImg = img_pth
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

