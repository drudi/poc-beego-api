package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["poc-beego-api/controllers:JobController"] = append(beego.GlobalControllerRouter["poc-beego-api/controllers:JobController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["poc-beego-api/controllers:JobController"] = append(beego.GlobalControllerRouter["poc-beego-api/controllers:JobController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["poc-beego-api/controllers:JobController"] = append(beego.GlobalControllerRouter["poc-beego-api/controllers:JobController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

}
