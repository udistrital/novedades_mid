package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/udistrital/auditoria"
	_ "github.com/udistrital/novedades_mid/routers"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// AllowOrigins: []string{"*"},
		AllowOrigins: []string{"https://*.portaloas.udistrital.edu.co"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	auditoria.InitMiddleware()
	apistatus.Init()
	beego.Run()

}
