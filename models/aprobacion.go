package models

import (
	"fmt"
	// "strconv"
	// "github.com/astaxie/beego"
	// "github.com/udistrital/utils_oas/request"
)

func ConsultaTablaAprobacion(rol string) (result []map[string]interface{}, resultErr error) {
	fmt.Println("Rol: ", rol)
	// var url = ""

	// if rol == "CONTRATISTA" {
	// 	url = ""
	// }

	// err = request.GetJson(beego.AppConfig.String("NovedadesCrudService") + url, )

	// error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/fechas/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &fechas)
	// error1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/propiedad/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &propiedades)
	// error2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/tipo_novedad/?query=Id:"+strconv.FormatFloat((NovedadAdicion["TipoNovedad"]).(float64), 'f', -1, 64), &tiponovedad)
	// error3 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro/"+NovedadAdicion["Estado"].(string), &estadoNovedad)

	// if error == nil {
	// 	if len(fechas[0]) != 0 {
	// 		for _, fecha := range fechas {
	// 			tipofecha := fecha["IdTipoFecha"].(map[string]interface{})
	// 			nombrefecha := tipofecha["Nombre"]
	// 			if nombrefecha == "FechaRegistro" {
	// 				fecharegistro = fecha["Fecha"]
	// 			}
	// 			if nombrefecha == "FechaSolicitud" {
	// 				fechasolicitud = fecha["Fecha"]
	// 			}
	// 			if nombrefecha == "FechaReinicio" {
	// 				fechareinicio = fecha["Fecha"]
	// 			}
	// 			if nombrefecha == "FechaSuspension" {
	// 				fechasuspension = fecha["Fecha"]
	// 			}
	// 			if nombrefecha == "FechaFinSuspension" {
	// 				fechafinsuspension = fecha["Fecha"]
	// 			}
	// 			if nombrefecha == "FechaFinEfectiva" {
	// 				fechafinefectiva = fecha["Fecha"]
	// 			}
	// 		}
	// 	}
	// }

	// if error1 == nil {
	// 	if len(propiedades[0]) != 0 {
	// 		for _, propiedad := range propiedades {
	// 			tipopropiedad := propiedad["IdTipoPropiedad"].(map[string]interface{})
	// 			nombrepropiedad := tipopropiedad["Nombre"]
	// 			if nombrepropiedad == "Cesionario" {
	// 				cesionario = propiedad["Propiedad"]
	// 			}
	// 			if nombrepropiedad == "PeriodoSuspension" {
	// 				periodosuspension = propiedad["Propiedad"]
	// 			}
	// 		}
	// 	}
	// }
	return nil, nil
}
