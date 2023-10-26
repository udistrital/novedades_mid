package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func ConstruirNovedadAdicionPost(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadAdicion := make(map[string]interface{})
	NovedadAdicion = novedad

	NovedadAdicionPost := make(map[string]interface{})
	contratoid, _ := strconv.ParseInt(NovedadAdicion["contrato"].(string), 10, 32)
	numerocdpid, _ := strconv.ParseInt(NovedadAdicion["numerocdp"].(string), 10, 32)
	numerorp, _ := strconv.ParseInt(NovedadAdicion["numerorp"].(string), 10, 32)
	// numerosolicitudentero := NovedadAdicion["numerosolicitud"].(float64)
	// numerosolicitud := strconv.FormatFloat(numerosolicitudentero, 'f', -1, 64)
	vigencia, _ := strconv.ParseInt(NovedadAdicion["vigencia"].(string), 10, 32)
	vigenciacdp, _ := strconv.ParseInt(NovedadAdicion["vigenciacdp"].(string), 10, 32)
	vigenciarp, _ := strconv.ParseInt(NovedadAdicion["vigenciarp"].(string), 10, 32)

	codEstado := ""

	var estadoNovedad map[string]interface{}
	error3 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro?query=TipoParametroId.CodigoAbreviacion:ENOV,CodigoAbreviacion:"+NovedadAdicion["estado"].(string), &estadoNovedad)

	if error3 == nil {
		if len(estadoNovedad) != 0 {
			inter := estadoNovedad["Data"].([]interface{})
			data := inter[0].(map[string]interface{})
			idEstado, _ := data["Id"].(float64)
			codEstado = strconv.FormatFloat(idEstado, 'f', -1, 64)
		}
	}

	// fmt.Println(NovedadAdicion["contrato"], NovedadAdicion["numerocdp"], NovedadAdicion["numerosolicitud"], NovedadAdicion["vigencia"], NovedadAdicion["vigencia"])
	// fmt.Println("\n", contratoid, numerocdpid, numerorp, numerosolicitud, vigencia, vigenciacdp, vigenciarp, "\n")

	NovedadAdicionPost["NovedadPoscontractual"] = map[string]interface{}{
		"Aclaracion":        nil,
		"Activo":            true,
		"ContratoId":        contratoid,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"Motivo":            NovedadAdicion["motivo"],
		"NumeroCdpId":       numerocdpid,
		"NumeroSolicitud":   NovedadAdicion["numerosolicitud"],
		"Observacion":       nil,
		"TipoNovedad":       6,
		"Vigencia":          vigencia,
		"VigenciaCdp":       vigenciacdp,
		"NumeroRp":          numerorp,
		"VigenciaRp":        vigenciarp,
		"OficioSupervisor":  NovedadAdicion["numerooficiosupervisor"],
		"OficioOrdenador":   NovedadAdicion["numerooficioordenador"],
		"Estado":            codEstado,
		"EnlaceDocumento":   NovedadAdicion["enlace"],
	}

	fechas := make([]map[string]interface{}, 0)

	loc, _ := time.LoadLocation("America/Bogota")
	f_solicitud, _ := time.Parse("2006-01-02T15:04:05Z07:00", NovedadAdicion["fechasolicitud"].(string))

	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             f_solicitud.In(loc),
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
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadAdicion["fechafinefectiva"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 12,
		},
	})
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadAdicion["fechaexpedicion"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 14,
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

	propiedades = append(propiedades, map[string]interface{}{
		"Activo":            true,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoPropiedad": map[string]interface{}{
			"Id": 14,
		},
		"propiedad": numerorp,
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
			"Id": 15,
		},
		"propiedad": vigenciarp,
	})

	NovedadAdicionPost["Propiedad"] = propiedades

	// fmt.Println(NovedadAdicionPost)

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
	var fechafinefectiva interface{}
	var fechaexpedicion interface{}
	var cesionario interface{}
	var valoradicion interface{}
	var tiponovedad []map[string]interface{}
	var tipoNovedadNombre string
	var estadoNovedad map[string]interface{}
	var nombreEstadoNov string
	var codEstado string

	error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/fechas/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &fechas)
	error1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/propiedad/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &propiedades)
	error2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/tipo_novedad/?query=Id:"+strconv.FormatFloat((NovedadAdicion["TipoNovedad"]).(float64), 'f', -1, 64), &tiponovedad)
	error3 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro/"+NovedadAdicion["Estado"].(string), &estadoNovedad)

	if len(fechas[0]) != 0 {
		for _, fecha := range fechas {
			tipofecha := fecha["IdTipoFecha"].(map[string]interface{})
			nombrefecha := tipofecha["Nombre"]
			if nombrefecha == "FechaAdicion" {
				fechaadicion = fecha["Fecha"]
			}
			if nombrefecha == "FechaSolicitud" {
				fechasolicitud = fecha["Fecha"]
			}
			if nombrefecha == "FechaFinEfectiva" {
				fechafinefectiva = fecha["Fecha"]
			}
			if nombrefecha == "FechaExpedicion" {
				fechaexpedicion = fecha["Fecha"]
			}
			//fmt.Println(fechaadicion, fechasolicitud)
		}
	}
	if len(propiedades[0]) != 0 {
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
	}

	if error2 == nil {
		if len(tiponovedad[0]) != 0 {
			tipoNovedadNombre = tiponovedad[0]["Nombre"].(string)
		}
	}

	if error3 == nil {
		if len(estadoNovedad) != 0 {
			data := estadoNovedad["Data"].(map[string]interface{})
			nombreEstadoNov = data["Nombre"].(string)
			codEstado = data["CodigoAbreviacion"].(string)
		}
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
		"fechaexpedicion":            fechaexpedicion,
		"fechasuspension":            "",
		"fechaterminacionanticipada": "",
		"motivo":                     NovedadAdicion["Motivo"],
		"numeroactaentrega":          "",
		"numerocdp":                  NovedadAdicion["NumeroCdpId"],
		"numerosolicitud":            NovedadAdicion["NumeroSolicitud"],
		"observacion":                NovedadAdicion["Observacion"],
		"periodosuspension":          "",
		"plazoactual":                "",
		"poliza":                     "",
		"tiempoprorroga":             "",
		"tiponovedad":                NovedadAdicion["TipoNovedad"],
		"nombreTipoNovedad":          tipoNovedadNombre,
		"valoradicion":               valoradicion,
		"valorfinalcontrato":         "",
		"vigencia":                   NovedadAdicion["Vigencia"],
		"fechafinefectiva":           fechafinefectiva,
		"nombreEstado":               nombreEstadoNov,
		"numerooficiosupervisor":     NovedadAdicion["OficioSupervisor"],
		"numerooficioordenador":      NovedadAdicion["OficioOrdenador"],
		"estado":                     codEstado,
		"enlace":                     NovedadAdicion["EnlaceDocumento"],
	}

	fmt.Println(error, error1)

	return NovedadAdicionGet
}
