// Package routers define versions
// @APIVersion 1.0.0
// @Title poc-beego-api
// @Description poc-beego-api
package routers

import (
	"poc-beego-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/job",
			beego.NSInclude(
				&controllers.JobController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
