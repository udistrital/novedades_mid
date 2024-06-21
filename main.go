package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/udistrital/auditoria"
	"github.com/udistrital/novedades_mid/models"
	_ "github.com/udistrital/novedades_mid/routers"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/xray"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	AllowedOrigins := []string{"https://*.udistrital.edu.co", "http://api.intranetoas.udistrital.edu.co:*", "http://api2.intranetoas.udistrital.edu.co:*"}
	if beego.BConfig.RunMode != "production" {
		AllowedOrigins = []string{"*"}
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//AllowOrigins: []string{"*"},
		//AllowOrigins: []string{"https://*.portaloas.udistrital.edu.co"},
		AllowOrigins: AllowedOrigins,
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
	go models.Temporizador()
	auditoria.InitMiddleware()
	xray.InitXRay()
	apistatus.Init()
	beego.Run()
}
