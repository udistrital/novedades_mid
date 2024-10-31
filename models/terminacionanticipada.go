package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func ConstruirNovedadTAnticipada(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadTAnticipada := make(map[string]interface{})
	NovedadTAnticipada = novedad

	NovedadTAnticipadaPost := make(map[string]interface{})
	contratoid, _ := strconv.ParseInt(NovedadTAnticipada["contrato"].(string), 10, 32)
	//numerocdpid, _ := strconv.ParseInt(NovedadSuspension["numerocdp"].(string), 10, 32)
	// numerosolicitudentero := NovedadTAnticipada["numerosolicitud"].(float64)
	// numerosolicitud := strconv.FormatFloat(numerosolicitudentero, 'f', -1, 64)
	vigencia, _ := strconv.ParseInt(NovedadTAnticipada["vigencia"].(string), 10, 32)
	// vigenciacdp, _ := strconv.ParseInt(NovedadTAnticipada["vigenciacdp"].(string), 10, 32)

	codEstado := ""

	var estadoNovedad map[string]interface{}
	error3 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro?query=TipoParametroId.CodigoAbreviacion:ENOV,CodigoAbreviacion:"+NovedadTAnticipada["estado"].(string), &estadoNovedad)

	if error3 == nil {
		if len(estadoNovedad) != 0 {
			inter := estadoNovedad["Data"].([]interface{})
			data := inter[0].(map[string]interface{})
			idEstado, _ := data["Id"].(float64)
			codEstado = strconv.FormatFloat(idEstado, 'f', -1, 64)
		}
	}

	NovedadTAnticipadaPost["NovedadPoscontractual"] = map[string]interface{}{
		"Aclaracion":        nil,
		"Activo":            true,
		"ContratoId":        contratoid,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"Motivo":            NovedadTAnticipada["motivo"],
		"NumeroCdpId":       0,
		"NumeroSolicitud":   NovedadTAnticipada["numerosolicitud"],
		"Observacion":       nil,
		"TipoNovedad":       5,
		"Vigencia":          vigencia,
		"VigenciaCdp":       0,
		"OficioSupervisor":  NovedadTAnticipada["numerooficiosupervisor"],
		"OficioOrdenador":   NovedadTAnticipada["numerooficioordenador"],
		"Estado":            codEstado,
		"EnlaceDocumento":   NovedadTAnticipada["enlace"],
	}

	fechas := make([]map[string]interface{}, 0)

	loc, _ := time.LoadLocation("America/Bogota")
	f_solicitud, _ := time.Parse("2006-01-02T15:04:05Z07:00", NovedadTAnticipada["fechasolicitud"].(string))

	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadTAnticipada["fecharegistro"],
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
		"Fecha":             NovedadTAnticipada["fecha_terminacion_anticipada"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 9,
		},
	})
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadTAnticipada["fechafinefectiva"],
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
		"Fecha":             NovedadTAnticipada["fechaoficiosupervisor"],
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
		"Fecha":             NovedadTAnticipada["fechaoficioordenador"],
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
		"Fecha":             NovedadTAnticipada["fechaexpedicion"],
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

	NovedadTAnticipadaPost["Fechas"] = fechas

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
			"Id": 2,
		},
		"propiedad": NovedadTAnticipada["cesionario"],
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
			"Id": 9,
		},
		"propiedad": NovedadTAnticipada["numerooficioestadocuentas"],
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
			"Id": 10,
		},
		"propiedad": NovedadTAnticipada["valor_desembolsado"],
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
			"Id": 11,
		},
		"propiedad": NovedadTAnticipada["saldo_contratista"],
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
			"Id": 12,
		},
		"propiedad": NovedadTAnticipada["saldo_universidad"],
	})

	NovedadTAnticipadaPost["Propiedad"] = propiedades

	// fmt.Println(NovedadTAnticipadaPost)

	return NovedadTAnticipadaPost
}

func GetNovedadTAnticipada(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadAdicion := make(map[string]interface{})
	var fechas []map[string]interface{}
	var propiedades []map[string]interface{}
	NovedadAdicion = novedad
	NovedadAdicionGet := make(map[string]interface{})
	var cesionario interface{}
	var fecharegistro interface{}
	var fechaterminacionanticipada interface{}
	var fechasolicitud interface{}
	var fechafinefectiva interface{}
	var fechaexpedicion interface{}
	var tiponovedad []map[string]interface{}
	var tipoNovedadNombre string
	var estadoNovedad map[string]interface{}
	var nombreEstadoNov string
	var codEstado string

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
				if nombrefecha == "FechaTerminacionAnticipada" {
					fechaterminacionanticipada = fecha["Fecha"]
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
		"Aclaracion":                 "",
		"Cedente":                    0,
		"Cesionario":                 cesionario,
		"Contrato":                   NovedadAdicion["ContratoId"],
		"EntidadAseguradora":         0,
		"FechaAdicion":               "",
		"FechaCesion":                "",
		"FechaLiquidacion":           "",
		"FechaProrroga":              "",
		"FechaRegistro":              fecharegistro,
		"FechaReinicio":              "",
		"FechaSolicitud":             fechasolicitud,
		"FechaSuspension":            "",
		"FechaFinSuspension":         "",
		"FechaFinEfectiva":           fechafinefectiva,
		"FechaTerminacionAnticipada": fechaterminacionanticipada,
		"FechaExpedicion":            fechaexpedicion,
		"Motivo":                     NovedadAdicion["Motivo"],
		"NumeroActaEntrega":          "",
		"NumeroCdp":                  NovedadAdicion["NumeroCdpId"],
		"NumeroSolicitud":            NovedadAdicion["NumeroSolicitud"],
		"Observacion":                "",
		"PeriodoSuspension":          0,
		"PlazoActual":                0,
		"Poliza":                     "",
		"TiempoProrroga":             0,
		"TipoNovedad":                NovedadAdicion["TipoNovedad"],
		"NombreTipoNovedad":          tipoNovedadNombre,
		"CodAbreviacionTipo":         "NP_TER",
		"ValorAdicion":               0,
		"ValorFinalContrato":         0,
		"Vigencia":                   NovedadAdicion["Vigencia"],
		"NumeroOficioSupervisor":     NovedadAdicion["OficioSupervisor"],
		"NumeroOficioOrdenador":      NovedadAdicion["OficioOrdenador"],
		"Estado":                     codEstado,
		"NombreEstado":               nombreEstadoNov,
		"Enlace":                     NovedadAdicion["EnlaceDocumento"],
	}

	return NovedadAdicionGet
}
