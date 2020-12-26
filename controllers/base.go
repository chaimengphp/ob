package controllers

import (
	beego "github.com/astaxie/beego/server/web"
	"net/http"
	"strconv"
)

type BaseController struct {
	beego.Controller
}


func (b *BaseController) GetUid() int64 {
	uid := b.Ctx.Request.Header.Get("uid")
	var user_id int64
	if uid != "" {
		userid,_ := strconv.Atoi(uid)
		user_id = int64(userid)
	}
	return user_id
}

func (b *BaseController) GetDeviceId() string {
	var r *http.Request
	device_id := r.Header.Get("device_id")
	return device_id
}


func (b *BaseController) ResponseData(code int64,message string,Result interface{}) {
	type ReponseJson struct {
		Code int64 `json:"code"`
		Message string `json:"message"`
		Result interface{} `json:"result"`
	}
	reponseJson := ReponseJson{
		Code:code,Message:message,Result:Result,
	}
	b.Data["json"] = reponseJson
	b.ServeJSON()
}
