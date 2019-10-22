package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func ConstruirNovedadAdicionPost(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadAdicion := make(map[string]interface{})
	NovedadAdicion = novedad

	NovedadAdicionPost := make(map[string]interface{})
	contratoid, _ := strconv.ParseInt(NovedadAdicion["contrato"].(string), 10, 32)
	numerocdpid, _ := strconv.ParseInt(NovedadAdicion["numerocdp"].(string), 10, 32)
	numerosolicitudentero := NovedadAdicion["numerosolicitud"].(float64)
	numerosolicitud := strconv.FormatFloat(numerosolicitudentero, 'f', -1, 64)
	vigencia, _ := strconv.ParseInt(NovedadAdicion["vigencia"].(string), 10, 32)
	vigenciacdp, _ := strconv.ParseInt(NovedadAdicion["vigencia"].(string), 10, 32)

	fmt.Println(NovedadAdicion["contrato"], NovedadAdicion["numerocdp"], NovedadAdicion["numerosolicitud"], NovedadAdicion["vigencia"], NovedadAdicion["vigencia"])
	fmt.Println("\n", contratoid, numerocdpid, numerosolicitud, vigencia, vigenciacdp, "\n")

	NovedadAdicionPost["NovedadPoscontractual"] = map[string]interface{}{
		"Aclaracion":        nil,
		"Activo":            true,
		"ContratoId":        contratoid,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"Motivo":            NovedadAdicion["motivo"],
		"NumeroCdpId":       numerocdpid,
		"NumeroSolicitud":   numerosolicitud,
		"Observacion":       nil,
		"TipoNovedad":       6,
		"Vigencia":          vigencia,
		"VigenciaCdp":       vigenciacdp,
	}

	fechas := make([]map[string]interface{}, 0)

	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadAdicion["fechasolicitud"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 7,
		},
	})
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadAdicion["fechaadicion"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 1,
		},
	})

	NovedadAdicionPost["Fechas"] = fechas

	propiedades := make([]map[string]interface{}, 0)
	propiedades = append(propiedades, map[string]interface{}{
		"Activo":            true,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoPropiedad": map[string]interface{}{
			"Id": 6,
		},
		"propiedad": NovedadAdicion["valoradicion"],
	})

	propiedades = append(propiedades, map[string]interface{}{
		"Activo":            true,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoPropiedad": map[string]interface{}{
			"Id": 2,
		},
		"propiedad": NovedadAdicion["cesionario"],
	})

	NovedadAdicionPost["Propiedad"] = propiedades

	fmt.Println(NovedadAdicionPost)

	return NovedadAdicionPost
}

func GetNovedadAdicion(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadAdicion := make(map[string]interface{})
	var fechas []map[string]interface{}
	var propiedades []map[string]interface{}
	NovedadAdicion = novedad
	NovedadAdicionGet := make(map[string]interface{})
	var fechaadicion interface{}
	var fechasolicitud interface{}
	var cesionario interface{}
	var valoradicion interface{}

	error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/fechas/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &fechas)
	error1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/propiedad/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &propiedades)

	for _, fecha := range fechas {
		tipofecha := fecha["IdTipoFecha"].(map[string]interface{})
		nombrefecha := tipofecha["Nombre"]
		if nombrefecha == "FechaAdicion" {
			fechaadicion = fecha["Fecha"]
		}
		if nombrefecha == "FechaSolicitud" {
			fechasolicitud = fecha["Fecha"]
		}
		//fmt.Println(fechaadicion, fechasolicitud)
	}
	for _, propiedad := range propiedades {
		tipopropiedad := propiedad["IdTipoPropiedad"].(map[string]interface{})
		nombrepropiedad := tipopropiedad["Nombre"]
		if nombrepropiedad == "Cesionario" {
			cesionario = propiedad["Propiedad"]
		}
		if nombrepropiedad == "ValorAdicion" {
			valoradicion = propiedad["Propiedad"]
		}
		//fmt.Println(cesionario, valoradicion)
	}

	NovedadAdicionGet = map[string]interface{}{
		"id":                         NovedadAdicion["Id"].(float64),
		"aclaracion":                 NovedadAdicion["Aclaracion"],
		"camposaclaracion":           "",
		"camposmodificacion":         "",
		"camposmodificados":          "",
		"cedente":                    "",
		"cesionario":                 cesionario,
		"contrato":                   NovedadAdicion["ContratoId"],
		"fechaadicion":               fechaadicion,
		"fechacesion":                "",
		"fechaliquidacion":           "",
		"fechaprorroga":              "",
		"fecharegistro":              "",
		"fechareinicio":              "",
		"fechasolicitud":             fechasolicitud,
		"fechasuspension":            "",
		"fechaterminacionanticipada": "",
		"motivo":                     NovedadAdicion["Motivo"],
		"numeroactaentrega":          "",
		"numerocdp":                  NovedadAdicion["NumeroCdpId"],
		"numerooficioestadocuentas":  "",
		"numerosolicitud":            NovedadAdicion["NumeroSolicitud"],
		"observacion":                NovedadAdicion["Observacion"],
		"periodosuspension":          "",
		"plazoactual":                "",
		"poliza":                     "",
		"tiempoprorroga":             "",
		"tiponovedad":                NovedadAdicion["TipoNovedad"],
		"valoradicion":               valoradicion,
		"valorfinalcontrato":         "",
		"vigencia":                   NovedadAdicion["Vigencia"],
	}

	fmt.Println(error, error1)

	return NovedadAdicionGet
}

func FormatAdmAmazonNovedad(novedad []map[string]interface{}) (novedadformatted map[string]interface{}) {
	var NovedadesAdicion []map[string]interface{}
	var fechas []map[string]interface{}
	var propiedades []map[string]interface{}
	NovedadesAdicion = novedad
	NovedadAdicionGet := make(map[string]interface{})
	var fechaadicion interface{}
	var fechasolicitud interface{}
	var cesionario interface{}
	var valoradicion interface{}
	var id interface{}

	fmt.Println(NovedadesAdicion)

	for _, NovedadAdicion := range NovedadesAdicion {

		id = NovedadAdicion["Id"]
		fmt.Println(NovedadAdicion)

		fmt.Println("Aqui se muestra el id luego de ser pasado por el for \n", id)

		error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/fechas/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &fechas)
		error1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/propiedad/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &propiedades)

		for _, fecha := range fechas {
			tipofecha := fecha["IdTipoFecha"].(map[string]interface{})
			nombrefecha := tipofecha["Nombre"]
			if nombrefecha == "FechaAdicion" {
				fechaadicion = fecha["Fecha"]
			}
			if nombrefecha == "FechaSolicitud" {
				fechasolicitud = fecha["Fecha"]
			}
			//fmt.Println(fechaadicion, fechasolicitud)
		}
		for _, propiedad := range propiedades {
			tipopropiedad := propiedad["IdTipoPropiedad"].(map[string]interface{})
			nombrepropiedad := tipopropiedad["Nombre"]
			if nombrepropiedad == "Cesionario" {
				cesionario = propiedad["Propiedad"]
			}
			if nombrepropiedad == "ValorAdicion" {
				valoradicion = propiedad["Propiedad"]
			}
			//fmt.Println(cesionario, valoradicion)
		}

		NovedadAdicionGet = map[string]interface{}{
			"id":                         NovedadAdicion["Id"].(float64),
			"aclaracion":                 NovedadAdicion["Aclaracion"],
			"camposaclaracion":           "",
			"camposmodificacion":         "",
			"camposmodificados":          "",
			"cedente":                    "",
			"cesionario":                 cesionario,
			"contrato":                   NovedadAdicion["ContratoId"],
			"fechaadicion":               fechaadicion,
			"fechacesion":                "",
			"fechaliquidacion":           "",
			"fechaprorroga":              "",
			"fecharegistro":              "",
			"fechareinicio":              "",
			"fechasolicitud":             fechasolicitud,
			"fechasuspension":            "",
			"fechaterminacionanticipada": "",
			"motivo":                     NovedadAdicion["Motivo"],
			"numeroactaentrega":          "",
			"numerocdp":                  NovedadAdicion["NumeroCdpId"],
			"numerooficioestadocuentas":  "",
			"numerosolicitud":            NovedadAdicion["NumeroSolicitud"],
			"observacion":                NovedadAdicion["Observacion"],
			"periodosuspension":          "",
			"plazoactual":                "",
			"poliza":                     "",
			"tiempoprorroga":             "",
			"tiponovedad":                NovedadAdicion["TipoNovedad"],
			"valoradicion":               valoradicion,
			"valorfinalcontrato":         "",
			"vigencia":                   NovedadAdicion["Vigencia"],
		}
		fmt.Println(error, error1)
	}

	return NovedadAdicionGet
}
