package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func ConstruirNovedadAdProrrogaPost(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadAdProrroga := make(map[string]interface{})
	NovedadAdProrroga = novedad

	NovedadAdProrrogaPost := make(map[string]interface{})
	contratoid, _ := strconv.ParseInt(NovedadAdProrroga["contrato"].(string), 10, 32)
	numerocdpid, _ := strconv.ParseInt(NovedadAdProrroga["numerocdp"].(string), 10, 32)
	numerosolicitudentero := NovedadAdProrroga["numerosolicitud"].(float64)
	numerosolicitud := strconv.FormatFloat(numerosolicitudentero, 'f', -1, 64)
	vigencia, _ := strconv.ParseInt(NovedadAdProrroga["vigencia"].(string), 10, 32)
	vigenciacdp, _ := strconv.ParseInt(NovedadAdProrroga["vigencia"].(string), 10, 32)

	NovedadAdProrrogaPost["NovedadPoscontractual"] = map[string]interface{}{
		"Aclaracion":        nil,
		"Activo":            true,
		"ContratoId":        contratoid,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"Motivo":            NovedadAdProrroga["motivo"],
		"NumeroCdpId":       numerocdpid,
		"NumeroSolicitud":   numerosolicitud,
		"Observacion":       nil,
		"TipoNovedad":       8,
		"Vigencia":          vigencia,
		"VigenciaCdp":       vigenciacdp,
	}

	fechas := make([]map[string]interface{}, 0)

	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadAdProrroga["fechasolicitud"],
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
		"Fecha":             NovedadAdProrroga["fechaadicion"],
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
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadAdProrroga["fechaprorroga"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 4,
		},
	})

	NovedadAdProrrogaPost["Fechas"] = fechas

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
		"propiedad": NovedadAdProrroga["valoradicion"],
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
			"Id": 5,
		},
		"propiedad": NovedadAdProrroga["tiempoprorroga"],
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
		"propiedad": NovedadAdProrroga["cesionario"],
	})

	NovedadAdProrrogaPost["Propiedad"] = propiedades

	fmt.Println(NovedadAdProrrogaPost)

	return NovedadAdProrrogaPost
}

func GetNovedadAdProrroga(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadAdicion := make(map[string]interface{})
	var fechas []map[string]interface{}
	var propiedades []map[string]interface{}
	NovedadAdicion = novedad
	NovedadAdicionGet := make(map[string]interface{})
	var fechaadicion interface{}
	var fechasolicitud interface{}
	var fechaprorroga interface{}
	var cesionario interface{}
	var valoradicion interface{}
	var tiempoprorroga interface{}

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
		if nombrefecha == "FechaProrroga" {
			fechaprorroga = fecha["Fecha"]
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
		if nombrepropiedad == "TiempoProrroga" {
			tiempoprorroga = propiedad["Propiedad"]
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
		"fechaprorroga":              fechaprorroga,
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
		"tiempoprorroga":             tiempoprorroga,
		"tiponovedad":                NovedadAdicion["TipoNovedad"],
		"valoradicion":               valoradicion,
		"valorfinalcontrato":         "",
		"vigencia":                   NovedadAdicion["Vigencia"],
	}

	fmt.Println(error, error1)

	return NovedadAdicionGet
}

func FormatAdmAmazonNovedadAdProrroga(novedad []map[string]interface{}) (novedadformatted map[string]interface{}) {
	var NovedadesAdicion []map[string]interface{}
	var fechas []map[string]interface{}
	var propiedades []map[string]interface{}
	var ultimasnovedades []map[string]interface{}
	NovedadesAdicion = novedad
	NovedadAdicionGet := make(map[string]interface{})
	var fechaadicion interface{}
	var fechasolicitud interface{}
	var fechaprorroga interface{}
	var cesionario interface{}
	var valoradicion interface{}
	var tiempoprorroga interface{}
	var id interface{}

	var idultimanovedad interface{}

	fmt.Println(fechaadicion)
	fmt.Println(NovedadesAdicion)

	for _, NovedadAdicion := range NovedadesAdicion {

		id = NovedadAdicion["Id"]
		fmt.Println(NovedadAdicion)

		fmt.Println("Aqui se muestra el id luego de ser pasado por el for \n", id)

		error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/fechas/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &fechas)
		error1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/propiedad/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &propiedades)
		error2 := request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+"/novedad_postcontractual/?sortby=id&order=desc&limit=1", &ultimasnovedades)

		for _, ultimanovedad := range ultimasnovedades {

			idultimanovedad = ultimanovedad["Id"]

		}
		fmt.Println("Aquí se muestra el id de la última novedad en admamazon \n", idultimanovedad)

		for _, fecha := range fechas {
			tipofecha := fecha["IdTipoFecha"].(map[string]interface{})
			nombrefecha := tipofecha["Nombre"]
			if nombrefecha == "FechaAdicion" {
				fechaadicion = fecha["Fecha"]
			}
			if nombrefecha == "FechaSolicitud" {
				fechasolicitud = fecha["Fecha"]
			}
			if nombrefecha == "FechaProrroga" {
				fechaprorroga = fecha["Fecha"]
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
			if nombrepropiedad == "TiempoProrroga" {
				tiempoprorroga = propiedad["Propiedad"]
			}
			//fmt.Println(cesionario, valoradicion)
		}

		NovedadAdicionGet = map[string]interface{}{
			"Id":              (idultimanovedad.(float64) + 1),
			"NumeroContrato":  strconv.FormatFloat(NovedadAdicion["ContratoId"].(float64), 'f', -1, 64),
			"Vigencia":        NovedadAdicion["Vigencia"].(float64),
			"TipoNovedad":     220,
			"FechaInicio":     fechaprorroga,
			"FechaFin":        "2019-10-08T18:26:20.046Z",
			"FechaRegistro":   fechasolicitud,
			"Contratista":     cesionario.(float64),
			"NumeroCdp":       NovedadAdicion["NumeroCdpId"].(float64),
			"VigenciaCdp":     NovedadAdicion["Vigencia"].(float64),
			"PlazoEjecucion":  tiempoprorroga.(float64),
			"UnidadEjecucion": 205,
			"ValorNovedad":    valoradicion.(float64),
		}

		fmt.Println(error, error1, error2)
	}

	return NovedadAdicionGet
}
