package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BeegoWeb/controllers:ContactController"] = append(beego.GlobalControllerRouter["BeegoWeb/controllers:ContactController"],
        beego.ControllerComments{
            Method: "SendMessage",
            Router: `/contact`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeegoWeb/controllers:IndexController"] = append(beego.GlobalControllerRouter["BeegoWeb/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetIndex",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeegoWeb/controllers:IndexController"] = append(beego.GlobalControllerRouter["BeegoWeb/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetAbout",
            Router: `/about`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeegoWeb/controllers:IndexController"] = append(beego.GlobalControllerRouter["BeegoWeb/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetBlog",
            Router: `/blog`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeegoWeb/controllers:IndexController"] = append(beego.GlobalControllerRouter["BeegoWeb/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetContact",
            Router: `/contact`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
