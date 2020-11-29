package routers

import (
	beego "github.com/astaxie/beego/server/web"
	"github.com/astaxie/beego/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["obapi/controllers:UserController"] = append(beego.GlobalControllerRouter["obapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
