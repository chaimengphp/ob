package controllers

import (
	"github.com/astaxie/beego/client/orm"
	beego "github.com/astaxie/beego/server/web"
	"math"
	"obapi/Tools"
	"obapi/models"
)
type IndexController struct {
	beego.Controller
}

type reponseJson struct {
	Code int64 `json:"code"`
	Message string `json:"message"`
	Result interface{} `json:"result"`
}

//信息数据块结构
type indexItemData struct {
	Id int64 `json:"id"`
	Content string `json:"content"`
	Imglist []string `json:"imglist"`
	Pubtime string `json:"pubtime"`
}

//用户块数据结构
type indexItemUser struct {
	Uid int64 `json:"uid"`
	NikeName string `json:"nike_name"`
	HeadImage string `json:"head_image"`
}

//单条数据结构
type indexItem struct {
	Data indexItemData `json:"data"`
	User indexItemUser `json:"user"`
}

type indexListData struct {
	TotalPage int `json:"total_page"`
	List []*indexItem `json:"list"`
}


func (this *IndexController) Index() {
	o := orm.NewOrm()
	var user_contents []*models.UserContent
	var listData []*indexItem
	page,_ := this.GetInt("page",1)
	total,_ := o.QueryTable(new(models.UserContent)).Count()
	_,err := o.QueryTable(new(models.UserContent)).RelatedSel().Offset((page-1)*10).Limit(10).OrderBy("-pubtime").All(&user_contents)
	if err == nil {
		for _,data := range user_contents {
			item := &indexItem{
				Data:indexItemData{
					Id:data.Id,
					Content:data.Content,
					Imglist:Tools.HandelHeadImg(data.Imglist),
					Pubtime:Tools.DateFormat(data.Pubtime),
				},
				User:indexItemUser{
					Uid:data.User.Id,
					NikeName:data.User.NikeName,
					HeadImage:data.User.HeadImg,
				},
			}
			listData = append(listData,item)
		}
		indexList := indexListData{
			TotalPage:int(math.Ceil(float64(total/10))),
			List:listData,
		}
		reponseJson := reponseJson{
			0,"suss",indexList,
		}
		this.Data["json"] = reponseJson
		this.ServeJSON()
	}


}