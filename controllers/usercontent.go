package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/client/orm"
	"obapi/Tools"
	"obapi/models"
	"strconv"
	"time"
)

type UserContentController struct {
	BaseController
}
//发布内容
func (uc *UserContentController) PubContent() {
	content := uc.GetString("content","")
	img_list := uc.GetStrings("imglist")
	video_url := uc.GetString("imglist","")
	if(content == "" && video_url == "" && len(img_list) <=0) {
		uc.ResponseData(1,"数据异常",nil)
	}
	user_content := new(models.UserContent)
	if(content != "") {
		user_content.Content = content
	}
	if(video_url != "") {
		user_content.VideoUrl = video_url
	}
	if(len(img_list) > 0) {
		jdata,_ := json.Marshal(img_list)
		user_content.Imglist = string(jdata)
	}
	user_content.Pubtime = time.Now().Unix()
	user_content.UserId = uc.getUid()
	o := orm.NewOrm()
	if num,err := o.Insert(&user_content);err ==nil && num>0 {
		uc.ResponseData(0,"suss",user_content)
	}else{
		uc.ResponseData(1,"发布失败",nil)
	}

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

//删除内容
func (uc *UserContentController) Del() {
	id,_ := uc.GetInt64("id",0)
	uid  := uc.getUid()
	if id <=0 || uid <=0 {
		uc.ResponseData(1,"数据异常",nil)
	}
	user_content := models.UserContent{
		Id:id,UserId:uid,
	}
	o := orm.NewOrm()
	if num,err := o.Delete(&user_content);err == nil && num <=0 {
		uc.ResponseData(1,"删除失败",nil)
	}else {
		uc.ResponseData(0,"suss",nil)
	}
}

//上传图像
func (uc *UserContentController) UploadImg() {
	img := uc.GetString("img","")
	if img == "" {
		uc.ResponseData(1,"",nil)
	}
	head_img_path,err := Tools.UpImg(img)
	if(err != nil) {
		uc.ResponseData(1,"上传失败",nil)
	}
	if head_img_path!="" {
		uc.ResponseData(0,"suss", struct {
			ImgPath string
		}{ImgPath:head_img_path})
	}else{
		uc.ResponseData(1,"图片上传异常",nil)
	}
}