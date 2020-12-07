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
	var r *http.Request
	uid := r.Header.Get("uid")
	user_id,_ := strconv.Atoi(uid)
	return int64(user_id)
}

func (b *BaseController) GetDeviceId() string {
	var r *http.Request
	device_id := r.Header.Get("device_id")
	return device_id
}


func (b *BaseController) ResponseData(code int64,message string,Result interface{}) {
	type ReponseJson struct {
		Code int64
		Message string
		Result interface{}
	}
	reponseJson := ReponseJson{
		Code:0,Message:"suss",Result:Result,
	}
	b.Data["json"] = reponseJson
	b.ServeJSON()
}
