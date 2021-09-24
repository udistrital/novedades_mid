// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/novedades_mid/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/registro_novedad",
			beego.NSInclude(
				&controllers.RegistroNovedadController{},
			),
		),
		beego.NSNamespace("/novedad",
			beego.NSInclude(
				&controllers.NovedadesController{},
			),
		),
		beego.NSNamespace("/argo_replica",
			beego.NSInclude(
				&controllers.ArgoReplicaController{},
			),
		),
		beego.NSNamespace("/gestor_documental",
			beego.NSInclude(
				&controllers.GestorDocumentalController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
