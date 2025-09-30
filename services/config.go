package services

import (
	"os"
	"strings"

	"github.com/astaxie/beego"
)

func titanCrudBase() string {
	if b := beego.AppConfig.String("TitanCrudService"); b != "" {
		return strings.TrimRight(b, "/")
	}
	if b := beego.AppConfig.String("UrlTitanCrud"); b != "" {
		return strings.TrimRight(b, "/")
	}
	if b := os.Getenv("TITAN_CRUD_SERVICE"); b != "" {
		return strings.TrimRight(b, "/")
	}
	beego.Error("Falta configurar TitanCrudService / UrlTitanCrud / TITAN_CRUD_SERVICE")
	return ""
}

func titanMidBase() string {
	if b := beego.AppConfig.String("TitanMidService"); b != "" {
		return strings.TrimRight(b, "/")
	}
	if b := beego.AppConfig.String("UrlTitanMid"); b != "" {
		return strings.TrimRight(b, "/")
	}
	if b := os.Getenv("TITAN_MID_SERVICE"); b != "" {
		return strings.TrimRight(b, "/")
	}
	beego.Error("Falta configurar TitanMidService / UrlTitanMid / TITAN_MID_SERVICE")
	return ""
}
