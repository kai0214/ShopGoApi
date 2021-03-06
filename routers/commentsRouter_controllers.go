package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ShopGoApi/controllers:AccountController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:AccountController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:AccountController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:AccountController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:CartController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:CartController"],
        beego.ControllerComments{
            Method: "PostAdd",
            Router: "/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:CartController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:CartController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "PostSub",
            Router: "/addSub",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:FileController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:FileController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:GoodsController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "AddGoods",
            Router: "/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:GoodsController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "GetCategoryGoodsList",
            Router: "/categoryGoodsList",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:GoodsController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "GetDetail",
            Router: "/detail",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:GoodsController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "GetPageList",
            Router: "/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:UserController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:UserController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:UserController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:UserController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:uid",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:UserController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:uid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:UserController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ShopGoApi/controllers:UserController"] = append(beego.GlobalControllerRouter["ShopGoApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
