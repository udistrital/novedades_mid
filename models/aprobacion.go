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
	// var response []map[string]interface{}

	var tiponovedad []map[string]interface{}
	var estadoNovedad map[string]interface{}

	var aprobFirmas []map[string]interface{}
	var documentoActa string

	var estados []string

	var r1 []map[string]interface{}
	var r2 []map[string]interface{}
	var r3 []map[string]interface{}

	fmt.Println("Rol: ", rol)
	switch rol {
	case "SUPERVISOR":
		err1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:4536", &r1)
		if err1 == nil && len(r1[0]) > 0 {
			novedades = append(novedades, r1...)
		}
		err2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:4537", &r2)
		if err2 == nil && len(r2[0]) > 0 {
			novedades = append(novedades, r2...)
		}
		err3 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:4538", &r3)
		if err3 == nil && len(r3[0]) > 0 {
			novedades = append(novedades, r3...)
		}
		// estados = append(estados, "4536", "4537", "4538")
	case "ORDENADOR_DEL_GASTO":
		err1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:4539", &r1)
		if err1 == nil && len(r1[0]) > 0 {
			novedades = append(novedades, r1...)
		}
		err2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:4540", &r2)
		if err2 == nil && len(r2[0]) > 0 {
			novedades = append(novedades, r2...)
		}
		// estados = append(estados, "4539", "4540")
	case "ASISTENTE_JURIDICA":
		err1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:4541", &r1)
		if err1 == nil && len(r1[0]) > 0 {
			novedades = append(novedades, r1...)
		}
		err2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:4542", &r2)
		if err2 == nil && len(r2[0]) > 0 {
			novedades = append(novedades, r2...)
		}
		// estados = append(estados, "4518")
	default:
		fmt.Println("El rol no coincide con alguno registrado!")
	}
	fmt.Println("Estados:", estados)

	fmt.Println("Novedades: ", novedades)
	// for i := 0; i < len(estados); i++ {
	// 	error1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:"+estados[i], &response)
	// 	if error1 == nil && len(response[0]) > 0 {
	// 		// fmt.Println("Response: ", response)
	// 		novedades = append(novedades, response...)
	// 	}
	// 	fmt.Println("Novedades: ", novedades)
	// }

	// for _, estado := range estados {
	// 	fmt.Println(estado)
	// 	error1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:"+estado, &response)
	// 	if error1 == nil && len(response[0]) > 0 {
	// 		// fmt.Println("Response: ", response)
	// 		for i := 0; i < len(response); i++ {
	// 			novedades = append(novedades, response[i])
	// 		}
	// 	}
	// 	// fmt.Println("Paso: ", novedades)
	// }

	if len(novedades) != 0 {
		for _, novedad := range novedades {
			var urltn string
			if reflect.TypeOf(novedad["TipoNovedad"]).String() == "string" {
				urltn = novedad["TipoNovedad"].(string)
			} else {
				urltn = strconv.FormatFloat((novedad["TipoNovedad"]).(float64), 'f', -1, 64)
			}
			error2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/tipo_novedad/?query=Id:"+urltn, &tiponovedad)
			error3 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro/"+novedad["Estado"].(string), &estadoNovedad)
			url := "/aprobacionfirma/?query=id_novedades_poscontractuales:" + strconv.FormatFloat((novedad["Id"]).(float64), 'f', -1, 64) + "&limit=0&sortby=FechaCreacion&order=asc"
			error4 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+url, &aprobFirmas)
			if error2 == nil && len(tiponovedad[0]) > 0 {
				novedad["TipoNovedad"] = tiponovedad[0]["Nombre"].(string)
			}
			// fmt.Println("Novedad: ", novedad)
			// fmt.Println("estadoNovedad", estadoNovedad)
			if error3 == nil && len(estadoNovedad) > 0 {
				data := estadoNovedad["Data"].(map[string]interface{})
				novedad["EstadoNombre"] = data["Nombre"].(string)
			}
			if error4 == nil && len(aprobFirmas[0]) != 0 {
				documentoActa = aprobFirmas[0]["DocumentoActa"].(string)
				novedad["Documento"] = documentoActa
			} else {
				novedad["Documento"] = ""
			}
		}
	} else {
		outputError = map[string]interface{}{"funcion": "/ReplicaSuspension", "err": "No se encontraron novedades registradas!"}
		return nil, outputError
	}
	return novedades, nil
}

func ActualizarEstadoNovedad(id string, info map[string]interface{}) (status interface{}, outputError interface{}) {

	var novedad map[string]interface{}
	var aprobStruct map[string]interface{}
	var resultadoRegistro map[string]interface{}

	err := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/"+id, &novedad)
	if err == nil {
		novedad["Estado"] = info["Estado"]
		novedad["FechaCreacion"] = info["FechaProceso"]
		novedad["FechaModificacion"] = info["FechaProceso"]
		errRegNovedad := request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/"+id, "PUT", &resultadoRegistro, novedad)
		fmt.Println("Error", errRegNovedad)
		if errRegNovedad == nil {
			fmt.Println("Novedad actualizada!!")
			var idInt, _ = strconv.Atoi(id)
			aprobStruct = map[string]interface{}{
				"Activo":            true,
				"FechaCreacion":     novedad["FechaCreacion"],
				"FechaModificacion": novedad["FechaCreacion"],
				"Id":                0,
				"IdNovedadesPoscontractuales": map[string]interface{}{
					"Id": idInt,
				},
				"Proceso":          info["Estado"],
				"FechaProceso":     info["FechaProceso"],
				"DocumentoPersona": info["DocPersona"],
				"NombrePersona":    info["NombrePersona"],
				"DocumentoActa":    info["Doc"],
			}
			errAprobacion := request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/aprobacionfirma", "POST", &resultadoRegistro, aprobStruct)
			fmt.Println("Error2", errAprobacion)
			if errAprobacion == nil {
				return resultadoRegistro, nil
			} else {
				outputError = map[string]interface{}{"funcion": "/ActualizarEstadoNovedad", "err": errAprobacion}
				return nil, outputError
			}
		} else {
			outputError = map[string]interface{}{"funcion": "/ActualizarEstadoNovedad", "err": errRegNovedad}
			return nil, outputError
		}
	} else {
		outputError = map[string]interface{}{"funcion": "/ActualizarEstadoNovedad", "err": err}
		return nil, outputError
	}
}
