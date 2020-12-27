// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"obapi/controllers"
	beego "github.com/astaxie/beego/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/main.json",&controllers.IndexController{},"get:Index"),
		beego.NSRouter("/detail/:id:int.json",&controllers.UserContentController{},"get:Detail"),
		beego.NSRouter("/cont/del",&controllers.UserContentController{},"post:Del"),
		beego.NSRouter("/upload/img",&controllers.UserContentController{},"post:UploadImg"),
		beego.NSRouter("/pubcontent",&controllers.UserContentController{},"post:PubContent"),
		beego.NSRouter("/pet/add",&controllers.PetController{},"post:AddPet"),
		beego.NSRouter("/pet/list",&controllers.PetController{},"get:PetList"),

	)
	beego.AddNamespace(ns)
}
