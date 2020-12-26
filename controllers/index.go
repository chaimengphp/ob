package controllers

import (
	"github.com/astaxie/beego/client/orm"
	"math"
	"obapi/Tools"
	"obapi/models"
)
type IndexController struct {
	BaseController
}


//信息数据块结构
type itemContent struct {
	Id int64 `json:"id"`
	Type string `json:"type"`
	Content string `json:"content"`
	Video string `json:"video"`
	Imglist []string `json:"imglist"`
	DateFormat string `json:"date_format"`
}

//用户块数据结构
type indexItemUser struct {
	Uid int64 `json:"uid"`
	NikeName string `json:"nike_name"`
	HeadImage string `json:"head_image"`
}

//图文单条数据结构
type ImgTextItem struct {
	IShowTpl string `json:"i_show_tpl"`
	Data Content `json:"data"`
	//User indexItemUser `json:"user"`
}

type Content struct {
	Content itemContent `json:"content"`
}


type indexListData struct {
	TotalPage int `json:"total_page"`
	List []*ImgTextItem `json:"list"`
}


func (this *IndexController) Index() {
	o := orm.NewOrm()
	var user_contents []*models.UserContent
	var listData []*ImgTextItem
	page,_ := this.GetInt("page",1)
	total,_ := o.QueryTable(new(models.UserContent)).Count()
	_,err := o.QueryTable(new(models.UserContent)).Filter("user_id",this.GetUid()).Offset((page-1)*10).Limit(10).OrderBy("-pubtime").All(&user_contents)
	if err == nil {
		for _,data := range user_contents {
			//图文结构
			item := &ImgTextItem{
				IShowTpl:"imgtext_tpl",
				Data:Content{Content:itemContent{
					Id:data.Id,
					Content:data.Content,
					Imglist:Tools.HandelHeadImg(data.ImgList),
					DateFormat:Tools.DateFormat(data.Pubtime),
				}},
				//User:indexItemUser{
				//	Uid:data.User.Id,
				//	NikeName:data.User.NikeName,
				//	HeadImage:data.User.HeadImg,
				//},
			}
			listData = append(listData,item)
		}
		indexList := indexListData{
			TotalPage:int(math.Ceil(float64(total/10))),
			List:listData,
		}
		this.ResponseData(0,"suss",indexList)
	}


}