package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func ConstruirNovedadSuspension(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadSuspension := make(map[string]interface{})
	NovedadSuspension = novedad

	NovedadSuspensionPost := make(map[string]interface{})
	contratoid, _ := strconv.ParseInt(NovedadSuspension["contrato"].(string), 10, 32)
	// numerosolicitudentero := NovedadSuspension["numerosolicitud"].(float64)
	// numerosolicitud := strconv.FormatFloat(numerosolicitudentero, 'f', -1, 64)
	vigencia, _ := strconv.ParseInt(NovedadSuspension["vigencia"].(string), 10, 32)

	codEstado := ""

	var estadoNovedad map[string]interface{}
	error3 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro?query=TipoParametroId.CodigoAbreviacion:ENOV,CodigoAbreviacion:"+NovedadSuspension["estado"].(string), &estadoNovedad)

	if error3 == nil {
		if len(estadoNovedad) != 0 {
			inter := estadoNovedad["Data"].([]interface{})
			data := inter[0].(map[string]interface{})
			idEstado, _ := data["Id"].(float64)
			codEstado = strconv.FormatFloat(idEstado, 'f', -1, 64)
		}
	}

	NovedadSuspensionPost["NovedadPoscontractual"] = map[string]interface{}{
		"Aclaracion":        nil,
		"Activo":            true,
		"ContratoId":        contratoid,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"Motivo":            NovedadSuspension["motivo"],
		"NumeroCdpId":       0,
		"NumeroSolicitud":   NovedadSuspension["numerosolicitud"],
		"Observacion":       nil,
		"TipoNovedad":       1,
		"Vigencia":          vigencia,
		"OficioSupervisor":  NovedadSuspension["numerooficiosupervisor"],
		"OficioOrdenador":   NovedadSuspension["numerooficioordenador"],
		"Estado":            codEstado,
		"EnlaceDocumento":   NovedadSuspension["enlace"],
	}

	fechas := make([]map[string]interface{}, 0)

	loc, _ := time.LoadLocation("America/Bogota")
	f_solicitud, _ := time.Parse("2006-01-02T15:04:05Z07:00", NovedadSuspension["fechasolicitud"].(string))

	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadSuspension["fecharegistro"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 5,
		},
	})
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
		"Fecha":             NovedadSuspension["fechasuspension"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 8,
		},
	})
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadSuspension["fechareinicio"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 6,
		},
	})
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadSuspension["fechafinsuspension"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 11,
		},
	})
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadSuspension["fechafinefectiva"],
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
		"Fecha":             NovedadSuspension["fechaoficiosupervisor"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 10,
		},
	})
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadSuspension["fechaoficioordenador"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 13,
		},
	})
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadSuspension["fechaexpedicion"],
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

	NovedadSuspensionPost["Fechas"] = fechas

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
			"Id": 3,
		},
		"propiedad": NovedadSuspension["periodosuspension"],
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
		"propiedad": NovedadSuspension["cesionario"],
	})

	NovedadSuspensionPost["Propiedad"] = propiedades

	return NovedadSuspensionPost
}

func GetNovedadSuspension(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadAdicion := make(map[string]interface{})
	var fechas []map[string]interface{}
	var propiedades []map[string]interface{}
	var tiponovedad []map[string]interface{}
	NovedadAdicion = novedad
	NovedadAdicionGet := make(map[string]interface{})
	var fecharegistro interface{}
	var fechareinicio interface{}
	var fechasolicitud interface{}
	var fechasuspension interface{}
	var fechafinsuspension interface{}
	var fechafinefectiva interface{}
	var fechaexpedicion interface{}
	// var numerooficiosupervisor interface{}
	// var numerooficioordenador interface{}
	var cesionario interface{}
	var periodosuspension interface{}
	var tipoNovedadNombre string
	var estadoNovedad map[string]interface{}
	var codEstado string
	var nombreEstadoNov string

	error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/fechas/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &fechas)
	error1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/propiedad/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &propiedades)
	error2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/tipo_novedad/?query=Id:"+strconv.FormatFloat((NovedadAdicion["TipoNovedad"]).(float64), 'f', -1, 64), &tiponovedad)
	error3 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro/"+NovedadAdicion["Estado"].(string), &estadoNovedad)

	if error == nil {
		if len(fechas[0]) != 0 {
			for _, fecha := range fechas {
				tipofecha := fecha["IdTipoFecha"].(map[string]interface{})
				nombrefecha := tipofecha["Nombre"]
				if nombrefecha == "FechaRegistro" {
					fecharegistro = fecha["Fecha"]
				}
				if nombrefecha == "FechaSolicitud" {
					fechasolicitud = fecha["Fecha"]
				}
				if nombrefecha == "FechaReinicio" {
					fechareinicio = fecha["Fecha"]
				}
				if nombrefecha == "FechaSuspension" {
					fechasuspension = fecha["Fecha"]
				}
				if nombrefecha == "FechaFinSuspension" {
					fechafinsuspension = fecha["Fecha"]
				}
				if nombrefecha == "FechaFinEfectiva" {
					fechafinefectiva = fecha["Fecha"]
				}
				if nombrefecha == "FechaExpedicion" {
					fechaexpedicion = fecha["Fecha"]
				}
			}
		}
	}

	if error1 == nil {
		if len(propiedades[0]) != 0 {
			for _, propiedad := range propiedades {
				tipopropiedad := propiedad["IdTipoPropiedad"].(map[string]interface{})
				nombrepropiedad := tipopropiedad["Nombre"]
				if nombrepropiedad == "Cesionario" {
					cesionario = propiedad["Propiedad"]
				}
				if nombrepropiedad == "PeriodoSuspension" {
					periodosuspension = propiedad["Propiedad"]
				}
				// if nombrepropiedad == "NumeroOficioSupervisor" {
				// 	numerooficiosupervisor = propiedad["Propiedad"]
				// }
				// if nombrepropiedad == "NumeroOficioOrdenador" {
				// 	numerooficioordenador = propiedad["Propiedad"]
				// }
			}
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
		"Id":                         NovedadAdicion["Id"].(float64),
		"Aclaracion":                 NovedadAdicion["Aclaracion"],
		"Cedente":                    0,
		"Cesionario":                 cesionario,
		"Contrato":                   NovedadAdicion["ContratoId"],
		"EntidadAseguradora":         0,
		"FechaAdicion":               "",
		"FechaCesion":                "",
		"FechaLiquidacion":           "",
		"FechaProrroga":              "",
		"FechaRegistro":              fecharegistro,
		"FechaReinicio":              fechareinicio,
		"FechaSolicitud":             fechasolicitud,
		"FechaExpedicion":            fechaexpedicion,
		"FechaSuspension":            fechasuspension,
		"FechaFinSuspension":         fechafinsuspension,
		"FechaFinEfectiva":           fechafinefectiva,
		"FechaTerminacionAnticipada": "",
		"Motivo":                     NovedadAdicion["Motivo"],
		"NumeroActaEntrega":          "",
		"NumeroCdp":                  NovedadAdicion["NumeroCdpId"],
		"NumeroSolicitud":            NovedadAdicion["NumeroSolicitud"],
		"Observacion":                NovedadAdicion["Observacion"],
		"PeriodoSuspension":          periodosuspension,
		"PlazoActual":                0,
		"Poliza":                     "",
		"TiempoProrroga":             0,
		"TipoNovedad":                NovedadAdicion["TipoNovedad"],
		"NombreTipoNovedad":          tipoNovedadNombre,
		"CodAbreviacionTipo":         "NP_SUS",
		"ValorAdicion":               0,
		"ValorFinalContrato":         0,
		"Vigencia":                   NovedadAdicion["Vigencia"],
		"VigenciaCdp":                NovedadAdicion["VigenciaCdp"],
		"NumeroOficioSupervisor":     NovedadAdicion["OficioSupervisor"],
		"NumeroOficioOrdenador":      NovedadAdicion["OficioOrdenador"],
		"Estado":                     codEstado,
		"NombreEstado":               nombreEstadoNov,
		"Enlace":                     NovedadAdicion["EnlaceDocumento"],
	}

	return NovedadAdicionGet
}
