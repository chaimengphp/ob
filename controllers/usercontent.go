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
	content := uc.GetString("textcontent","")
	img_content_list := uc.GetStrings("imglist")
	video_url := uc.GetString("video_name","")
	addr := uc.GetString("addr","")
	pubtime := uc.GetString("pubtime","")

	if(content == "" && video_url == "" && len(img_content_list) <=0) {
		uc.ResponseData(1,"数据异常",nil)
	}
	//上传图片
	var img_list []string
	for _,imgcontent := range img_content_list {
		img_path,err := Tools.UpImg(imgcontent)
		if err == nil {
			img_list = append(img_list,img_path)
		}
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
		user_content.ImgList = string(jdata)
	}
	var Pubtime_stamp int64
	if pubtime == "" {
		Pubtime_stamp = time.Now().Unix()
	}else{
		tm,_ := time.Parse("2006-01-02 15:04:05",pubtime)
		Pubtime_stamp = tm.Unix()
	}
	user := models.User{Id:uc.GetUid()}
	user_content.Pubtime = Pubtime_stamp
	user_content.User = &user
	user_content.Addr = addr
	//fmt.Println(user_content)
	//return
	o := orm.NewOrm()
	if num,err := o.Insert(user_content);err ==nil && num>0 {
		item := &ImgTextItem{
			IShowTpl:"imgtext_tpl",
			Data:Content{Content:itemContent{
				Id:user_content.Id,
				Content:user_content.Content,
				Imglist:Tools.HandelHeadImg(user_content.ImgList),
				DateFormat:Tools.DFormat(user_content.Pubtime),
			}},
		}
		uc.ResponseData(0,"suss",item)
	}else{
		uc.ResponseData(1,"发布失败",nil)
	}

}

type DetailShowData struct {
	Id int64 `json:"id"`
	Content string `json:"content"`
	ImgList []string `json:"img_list"`
	Addr string `json:"addr"`
	DateFormat string `json:"date_format"`
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
	err := qs.Filter("id",id).One(&user_content_info,"id","content","imglist","pubtime")
	if err == orm.ErrNoRows {
		uc.ResponseData(1,"内容不存在",nil)
	}
	img_list := Tools.HandelHeadImg(user_content_info.ImgList)
	show_data := DetailShowData{
		Id:user_content_info.Id,
		Content:user_content_info.Content,
		ImgList:img_list,
		Addr:user_content_info.Addr,
		DateFormat:Tools.DFormat(user_content_info.Pubtime),

	}
	uc.ResponseData(0,"suss",show_data)
}

//删除内容
func (uc *UserContentController) Del() {
	id,_ := uc.GetInt64("id",0)
	uid  := uc.GetUid()
	if id <=0 || uid <=0 {
		uc.ResponseData(1,"数据异常",nil)
	}
	user := &models.User{Id:uid}
	user_content := models.UserContent{
		Id:id,User:user,
	}
	o := orm.NewOrm()
	if num,err := o.Delete(&user_content);err != nil || num <= 0 {
		uc.ResponseData(1,"删除失败",nil)
	}else if num > 0{
		uc.ResponseData(0,"suss",nil)
	}
}

//上传图像
func (uc *UserContentController) UploadImg() {
	img := uc.GetString("imgcontent","")

	if img == "" {
		uc.ResponseData(1,"请选择上传图片",nil)
	}
	head_img_path,err := Tools.UpImg(img)
	if(err != nil) {
		uc.ResponseData(1,err.Error(),nil)
	}
	if head_img_path!="" {
		uc.ResponseData(0,"suss", struct {
			ImgPath string `json:"img_path"`
		}{ImgPath:head_img_path})
	}else{
		uc.ResponseData(1,"图片上传异常",nil)
	}
}