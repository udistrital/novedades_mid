package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/time_bogota"
)

func ConstruirNovedadAdProrrogaPost(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadAdProrroga := make(map[string]interface{})
	NovedadAdProrroga = novedad

	NovedadAdProrrogaPost := make(map[string]interface{})
	contratoid, _ := strconv.ParseInt(NovedadAdProrroga["contrato"].(string), 10, 32)
	numerocdpid, _ := strconv.ParseInt(NovedadAdProrroga["numerocdp"].(string), 10, 32)
	numerorp, _ := strconv.ParseInt(NovedadAdProrroga["numerorp"].(string), 10, 32)
	// numerosolicitudentero := NovedadAdProrroga["numerosolicitud"].(float64)
	// numerosolicitud := strconv.FormatFloat(numerosolicitudentero, 'f', -1, 64)
	vigencia, _ := strconv.ParseInt(NovedadAdProrroga["vigencia"].(string), 10, 32)
	vigenciacdp, _ := strconv.ParseInt(NovedadAdProrroga["vigenciacdp"].(string), 10, 32)
	vigenciarp, _ := strconv.ParseInt(NovedadAdProrroga["vigenciarp"].(string), 10, 32)

	codEstado := ""

	var estadoNovedad map[string]interface{}
	error3 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro?query=TipoParametroId.CodigoAbreviacion:ENOV,CodigoAbreviacion:"+NovedadAdProrroga["estado"].(string), &estadoNovedad)

	if error3 == nil {
		if len(estadoNovedad) != 0 {
			inter := estadoNovedad["Data"].([]interface{})
			data := inter[0].(map[string]interface{})
			idEstado, _ := data["Id"].(float64)
			codEstado = strconv.FormatFloat(idEstado, 'f', -1, 64)
		}
	}

	NovedadAdProrrogaPost["NovedadPoscontractual"] = map[string]interface{}{
		"Aclaracion":        nil,
		"Activo":            true,
		"ContratoId":        contratoid,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"Motivo":            NovedadAdProrroga["motivo"],
		"NumeroCdpId":       numerocdpid,
		"NumeroSolicitud":   NovedadAdProrroga["numerosolicitud"],
		"Observacion":       nil,
		"TipoNovedad":       8,
		"Vigencia":          vigencia,
		"VigenciaCdp":       vigenciacdp,
		"NumeroRp":          numerorp,
		"VigenciaRp":        vigenciarp,
		"OficioSupervisor":  NovedadAdProrroga["numerooficiosupervisor"],
		"OficioOrdenador":   NovedadAdProrroga["numerooficioordenador"],
		"Estado":            codEstado,
		"EnlaceDocumento":   NovedadAdProrroga["enlace"],
	}

	fechas := make([]map[string]interface{}, 0)
	loc, _ := time.LoadLocation("America/Bogota")
	f_solicitud, _ := time.Parse("2006-01-02T15:04:05Z07:00", NovedadAdProrroga["fechasolicitud"].(string))

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
	// fmt.Println("FechaFinEfectiva: ", NovedadAdProrroga["fechafinefectiva"])
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadAdProrroga["fechafinefectiva"],
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
		"Fecha":             NovedadAdProrroga["fechaexpedicion"],
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

	NovedadAdProrrogaPost["Propiedad"] = propiedades

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
	var fechafinefectiva interface{}
	var fechaexpedicion interface{}
	var cesionario interface{}
	var valoradicion interface{}
	var tiempoprorroga interface{}
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
			if nombrefecha == "FechaProrroga" {
				fechaprorroga = fecha["Fecha"]
			}
			if nombrefecha == "FechaFinEfectiva" {
				fechafinefectiva = fecha["Fecha"]
			}
			if nombrefecha == "FechaExpedicion" {
				fechaexpedicion = fecha["Fecha"]
			}
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
			if nombrepropiedad == "TiempoProrroga" {
				tiempoprorroga = propiedad["Propiedad"]
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
		"FechaAdicion":               fechaadicion,
		"FechaCesion":                "",
		"FechaLiquidacion":           "",
		"FechaProrroga":              fechaprorroga,
		"FechaRegistro":              "",
		"FechaReinicio":              "",
		"FechaSolicitud":             fechasolicitud,
		"FechaSuspension":            "",
		"FechaFinSuspension":         "",
		"FechaFinEfectiva":           fechafinefectiva,
		"FechaTerminacionAnticipada": "",
		"FechaExpedicion":            fechaexpedicion,
		"Motivo":                     NovedadAdicion["Motivo"],
		"NumeroActaEntrega":          "",
		"NumeroCdp":                  NovedadAdicion["NumeroCdpId"],
		"NumeroSolicitud":            NovedadAdicion["NumeroSolicitud"],
		"Observacion":                "",
		"PeriodoSuspension":          0,
		"PlazoActual":                0,
		"Poliza":                     "",
		"TiempoProrroga":             tiempoprorroga,
		"TipoNovedad":                NovedadAdicion["TipoNovedad"],
		"NombreTipoNovedad":          tipoNovedadNombre,
		"CodAbreviacionTipo":         "NP_ADPRO",
		"ValorAdicion":               valoradicion,
		"ValorFinalContrato":         0,
		"Vigencia":                   NovedadAdicion["Vigencia"],
		"VigenciaCdp":                NovedadAdicion["VigenciaCdp"],
		"NumeroOficioSupervisor":     NovedadAdicion["OficioSupervisor"],
		"NumeroOficioOrdenador":      NovedadAdicion["OficioOrdenador"],
		"Estado":                     codEstado,
		"NombreEstado":               nombreEstadoNov,
		"Enlace":                     NovedadAdicion["EnlaceDocumento"],
	}

	fmt.Println(error, error1)

	return NovedadAdicionGet
}

func FormatAdmAmazonNovedadAdProrroga(novedad []map[string]interface{}) (novedadformatted map[string]interface{}) {
	var NovedadesAdicion []map[string]interface{}
	var fechas []map[string]interface{}
	var propiedades []map[string]interface{}

	NovedadesAdicion = novedad
	NovedadAdicionGet := make(map[string]interface{})
	var fechaadicion string
	var fechasolicitud string
	var fechaprorroga string
	var cesionario interface{}
	var valoradicion interface{}
	var tiempoprorroga interface{}
	var id interface{}

	fmt.Println(fechasolicitud, fechaprorroga, cesionario, valoradicion, tiempoprorroga)

	fmt.Println(fechaadicion)
	fmt.Println(NovedadesAdicion)

	for _, NovedadAdicion := range NovedadesAdicion {

		id = NovedadAdicion["Id"]

		fmt.Println("Aqui se muestra el id luego de ser pasado por el for \n", id)

		error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/fechas/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &fechas)
		error1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/propiedad/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &propiedades)

		if len(fechas) != 0 {
			for _, fecha := range fechas {
				tipofecha := fecha["IdTipoFecha"].(map[string]interface{})
				nombrefecha := tipofecha["Nombre"]
				if nombrefecha == "FechaAdicion" {
					fechaadicion = fecha["Fecha"].(string)
					fechaadicion = time_bogota.TiempoCorreccionFormato(fechaadicion)
				}
				if nombrefecha == "FechaSolicitud" {
					fechasolicitud = fecha["Fecha"].(string)
					fechasolicitud = time_bogota.TiempoCorreccionFormato(fechasolicitud)
				}
				if nombrefecha == "FechaProrroga" {
					fechaprorroga = fecha["Fecha"].(string)
					fechaprorroga = time_bogota.TiempoCorreccionFormato(fechaprorroga)
				}
			}
		}
		if len(propiedades) != 0 {
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
			}
		}

		NovedadAdicionGet = map[string]interface{}{
			//"Id":              (idultimanovedad.(float64) + 1),
			"NumeroContrato":  strconv.FormatFloat(NovedadAdicion["ContratoId"].(float64), 'f', -1, 64),
			"Vigencia":        NovedadAdicion["Vigencia"].(float64),
			"TipoNovedad":     220,
			"FechaInicio":     nil,
			"FechaFin":        "0001-01-01T00:00:00Z",
			"FechaRegistro":   fechasolicitud,
			"Contratista":     nil,
			"NumeroCdp":       NovedadAdicion["NumeroCdpId"].(float64),
			"VigenciaCdp":     NovedadAdicion["Vigencia"].(float64),
			"PlazoEjecucion":  tiempoprorroga.(float64),
			"UnidadEjecucion": 205,
			"ValorNovedad":    valoradicion.(float64),
		}

		fmt.Println(error, error1)
	}

	return NovedadAdicionGet
}
