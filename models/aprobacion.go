package models

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func ConsultaTablaAprobacion(rol string) (result []map[string]interface{}, outputError interface{}) {

	var novedades []map[string]interface{}
	var response []map[string]interface{}

	var tiponovedad []map[string]interface{}
	var estadoNovedad map[string]interface{}

	var estados []string

	fmt.Println("Rol: ", rol)
	switch rol {
	case "SUPERVISOR":
		estados = append(estados, "4536", "4537", "4538")
	case "ORDENADOR_DEL_GASTO":
		estados = append(estados, "4539", "4540")
	case "ASISTENTE_JURIDICA":
		estados = append(estados, "4518")
	default:
		fmt.Println("El rol no coincide con alguno registrado!")
	}
	fmt.Println("Estados:", estados)

	for _, estado := range estados {
		error1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:"+estado, &response)
		if error1 == nil && len(response[0]) > 0 {
			novedades = append(novedades, response...)
		}
	}
	if len(novedades) != 0 {
		for _, novedad := range novedades {
			fmt.Println("Nov:", reflect.TypeOf(novedad["TipoNovedad"]))
			error2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/tipo_novedad/?query=Id:"+strconv.FormatFloat((novedad["TipoNovedad"]).(float64), 'f', -1, 64), &tiponovedad)
			error3 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro/"+novedad["Estado"].(string), &estadoNovedad)
			if error2 == nil && len(tiponovedad[0]) > 0 {
				novedad["TipoNovedad"] = tiponovedad[0]["Nombre"].(string)
			}
			fmt.Println("estadoNovedad", estadoNovedad)
			if error3 == nil && len(estadoNovedad) > 0 {
				data := estadoNovedad["Data"].(map[string]interface{})
				novedad["Estado"] = data["Nombre"].(string)
			}
		}
	} else {
		outputError = map[string]interface{}{"funcion": "/ReplicaSuspension", "err": "No se encontraron novedades registradas!"}
		return nil, outputError
	}
	fmt.Println("Novedades: ", novedades)
	return novedades, nil
}
