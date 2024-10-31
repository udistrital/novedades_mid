package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/time_bogota"
)

func ConstruirNovedadCesion(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadCesion := make(map[string]interface{})
	NovedadCesion = novedad

	NovedadCesionPost := make(map[string]interface{})
	contratoid, _ := strconv.ParseInt(NovedadCesion["contrato"].(string), 10, 32)
	numerocdpid, _ := strconv.ParseInt(NovedadCesion["numerocdp"].(string), 10, 32)
	// numerosolicitudentero := NovedadCesion["numerosolicitud"].(float64)
	// numerosolicitud := strconv.FormatFloat(numerosolicitudentero, 'f', -1, 64)
	vigencia, _ := strconv.ParseInt(NovedadCesion["vigencia"].(string), 10, 32)
	vigenciacdp, _ := strconv.ParseInt(NovedadCesion["vigenciacdp"].(string), 10, 32)

	codEstado := ""

	var estadoNovedad map[string]interface{}
	error3 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro?query=TipoParametroId.CodigoAbreviacion:ENOV,CodigoAbreviacion:"+NovedadCesion["estado"].(string), &estadoNovedad)

	if error3 == nil {
		if len(estadoNovedad) != 0 {
			inter := estadoNovedad["Data"].([]interface{})
			data := inter[0].(map[string]interface{})
			idEstado, _ := data["Id"].(float64)
			codEstado = strconv.FormatFloat(idEstado, 'f', -1, 64)
		}
	}

	NovedadCesionPost["NovedadPoscontractual"] = map[string]interface{}{
		"Aclaracion":        NovedadCesion["aclaracion"],
		"Activo":            true,
		"ContratoId":        contratoid,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"Motivo":            NovedadCesion["motivo"],
		"NumeroCdpId":       numerocdpid,
		"NumeroSolicitud":   NovedadCesion["numerosolicitud"],
		"Observacion":       NovedadCesion["observacion"],
		"TipoNovedad":       2,
		"Vigencia":          vigencia,
		"VigenciaCdp":       vigenciacdp,
		"OficioSupervisor":  NovedadCesion["numerooficiosupervisor"],
		"OficioOrdenador":   NovedadCesion["numerooficioordenador"],
		"Estado":            codEstado,
		"EnlaceDocumento":   NovedadCesion["enlace"],
	}

	fechas := make([]map[string]interface{}, 0)

	loc, _ := time.LoadLocation("America/Bogota")

	f_solicitud, _ := time.Parse("2006-01-02T15:04:05Z07:00", NovedadCesion["fechasolicitud"].(string))
	// f_oficio, _ := time.Parse("2006-01-02T15:04:05Z07:00", NovedadCesion["fechaoficio"].(string))
	f_registro, _ := time.Parse("2006-01-02T15:04:05Z07:00", NovedadCesion["fecharegistro"].(string))

	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadCesion["fechacesion"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 2,
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
		"Fecha":             f_registro.In(loc),
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
		"Fecha":             NovedadCesion["fechafinefectiva"],
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
		"Fecha":             NovedadCesion["fechaoficiosupervisor"],
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
		"Fecha":             NovedadCesion["fechaoficioordenador"],
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
		"Fecha":             NovedadCesion["fechaexpedicion"],
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

	NovedadCesionPost["Fechas"] = fechas

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
			"Id": 1,
		},
		"propiedad": NovedadCesion["cedente"],
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
		"propiedad": NovedadCesion["cesionario"],
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
		"propiedad": NovedadCesion["valor_desembolsado"],
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
			"Id": 16,
		},
		"propiedad": NovedadCesion["valor_a_favor"],
	})
	NovedadCesionPost["Propiedad"] = propiedades

	poliza := make([]map[string]interface{}, 0)

	poliza = append(poliza, map[string]interface{}{
		"Activo":               true,
		"EntidadAseguradoraId": NovedadCesion["entidadaseguradora"],
		"FechaCreacion":        nil,
		"FechaModificacion":    nil,
		"Id":                   0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"NumeroPolizaId": NovedadCesion["poliza"],
	})

	NovedadCesionPost["Poliza"] = poliza

	return NovedadCesionPost
}

func GetNovedadCesion(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadAdicion := make(map[string]interface{})
	var fechas []map[string]interface{}
	var propiedades []map[string]interface{}
	var poliza []map[string]interface{}
	NovedadAdicion = novedad
	NovedadAdicionGet := make(map[string]interface{})
	var fechacesion interface{}
	var fecharegistro interface{}
	var fechasolicitud interface{}
	var fechafinefectiva interface{}
	var fechaexpedicion interface{}
	var tiponovedad []map[string]interface{}
	var tipoNovedadNombre string
	var estadoNovedad map[string]interface{}
	var nombreEstadoNov string
	var codEstado string

	var cedente interface{}
	var cesionario interface{}

	// var polizas interface{}
	// var entidadaseguradora interface{}

	error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/fechas/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &fechas)
	error1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/propiedad/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &propiedades)
	error2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/poliza/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &poliza)
	error3 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/tipo_novedad/?query=Id:"+strconv.FormatFloat((NovedadAdicion["TipoNovedad"]).(float64), 'f', -1, 64), &tiponovedad)
	error4 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro/"+NovedadAdicion["Estado"].(string), &estadoNovedad)

	if len(fechas[0]) > 0 {
		for _, fecha := range fechas {
			tipofecha := fecha["IdTipoFecha"].(map[string]interface{})
			nombrefecha := tipofecha["Nombre"]
			if nombrefecha == "FechaCesion" {
				fechacesion = fecha["Fecha"]
			}
			if nombrefecha == "FechaRegistro" {
				fecharegistro = fecha["Fecha"]
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
		}
	}
	if len(propiedades[0]) > 0 {
		for _, propiedad := range propiedades {
			tipopropiedad := propiedad["IdTipoPropiedad"].(map[string]interface{})
			nombrepropiedad := tipopropiedad["Nombre"]
			if nombrepropiedad == "Cedente" {
				cedente = propiedad["Propiedad"]
			}
			if nombrepropiedad == "Cesionario" {
				cesionario = propiedad["Propiedad"]
			}
		}
	}
	// if len(poliza[0]) > 0 {
	// 	for _, poliz := range poliza {
	// 		polizas = poliz["NumeroPolizaId"]
	// 		entidadaseguradora = poliz["EntidadAseguradoraId"]
	// 	}
	// }

	if error3 == nil {
		if len(tiponovedad[0]) != 0 {
			tipoNovedadNombre = tiponovedad[0]["Nombre"].(string)
		}
	}

	if error4 == nil {
		if len(estadoNovedad) != 0 {
			data := estadoNovedad["Data"].(map[string]interface{})
			nombreEstadoNov = data["Nombre"].(string)
			codEstado = data["CodigoAbreviacion"].(string)
		}
	}

	NovedadAdicionGet = map[string]interface{}{
		"Id":                         NovedadAdicion["Id"].(float64),
		"Aclaracion":                 "",
		"Cedente":                    cedente,
		"Cesionario":                 cesionario,
		"Contrato":                   NovedadAdicion["ContratoId"],
		"EntidadAseguradora":         0,
		"FechaAdicion":               "",
		"FechaCesion":                fechacesion,
		"FechaLiquidacion":           "",
		"FechaProrroga":              "",
		"FechaRegistro":              fecharegistro,
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
		"TiempoProrroga":             0,
		"TipoNovedad":                NovedadAdicion["TipoNovedad"],
		"NombreTipoNovedad":          tipoNovedadNombre,
		"CodAbreviacionTipo":         "NP_CES",
		"ValorAdicion":               0,
		"ValorFinalContrato":         0,
		"Vigencia":                   NovedadAdicion["Vigencia"],
		"NumeroOficioSupervisor":     NovedadAdicion["OficioSupervisor"],
		"NumeroOficioOrdenador":      NovedadAdicion["OficioOrdenador"],
		"Estado":                     codEstado,
		"NombreEstado":               nombreEstadoNov,
		"Enlace":                     NovedadAdicion["EnlaceDocumento"],
	}

	fmt.Println(error, error1, error2)

	return NovedadAdicionGet
}

func FormatAdmAmazonNovedadCesion(novedad []map[string]interface{}) (novedadformatted map[string]interface{}) {
	var NovedadesAdicion []map[string]interface{}
	//NovedadAdicion := make(map[string]interface{})
	var fechas []map[string]interface{}
	var propiedades []map[string]interface{}
	var poliza []map[string]interface{}
	NovedadesAdicion = novedad
	NovedadAdicionGet := make(map[string]interface{})
	var fechaadicion string
	var fechacesion string
	var fechaliquidacion string
	var fechaprorroga string
	var fecharegistro string
	var fechareinicio string
	var fechasuspension string
	var fechaterminacionanticipada string
	var fechasolicitud string
	var fechaoficio string

	var cedente interface{}
	var cesionario interface{}
	var numeroactaentrega interface{}
	var numerooficioestadocuentas interface{}
	var periodosuspension interface{}
	var plazoactual interface{}
	var tiempoprorroga interface{}
	var valoradicion interface{}
	var valorfinalcontrato interface{}

	var polizas interface{}
	var entidadaseguradora interface{}

	fmt.Println(cedente, numeroactaentrega, numerooficioestadocuentas, periodosuspension, plazoactual, tiempoprorroga, valoradicion, valorfinalcontrato, polizas, entidadaseguradora)

	for _, NovedadAdicion := range NovedadesAdicion {

		error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/fechas/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &fechas)
		error1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/propiedad/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &propiedades)
		error2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/poliza/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &poliza)

		for _, fecha := range fechas {
			tipofecha := fecha["IdTipoFecha"].(map[string]interface{})
			nombrefecha := tipofecha["Nombre"]
			if nombrefecha == "FechaAdicion" {
				fechaadicion = fecha["Fecha"].(string)
				fechaadicion = time_bogota.TiempoCorreccionFormato(fechaadicion)
			}
			if nombrefecha == "FechaCesion" {
				fechacesion = fecha["Fecha"].(string)
				fechacesion = time_bogota.TiempoCorreccionFormato(fechacesion)
			}
			if nombrefecha == "FechaLiquidacion" {
				fechaliquidacion = fecha["Fecha"].(string)
				fechaliquidacion = time_bogota.TiempoCorreccionFormato(fechaliquidacion)
			}
			if nombrefecha == "FechaProrroga" {
				fechaprorroga = fecha["Fecha"].(string)
				fechaprorroga = time_bogota.TiempoCorreccionFormato(fechaprorroga)
			}
			if nombrefecha == "FechaRegistro" {
				fecharegistro = fecha["Fecha"].(string)
				fecharegistro = time_bogota.TiempoCorreccionFormato(fecharegistro)
			}
			if nombrefecha == "FechaReinicio" {
				fechareinicio = fecha["Fecha"].(string)
				fechareinicio = time_bogota.TiempoCorreccionFormato(fechareinicio)
			}
			if nombrefecha == "FechaSuspension" {
				fechasuspension = fecha["Fecha"].(string)
				fechasuspension = time_bogota.TiempoCorreccionFormato(fechasuspension)
			}
			if nombrefecha == "FechaSolicitud" {
				fechasolicitud = fecha["Fecha"].(string)
				fechasolicitud = time_bogota.TiempoCorreccionFormato(fechasolicitud)
			}
			if nombrefecha == "FechaTerminacionAnticipada" {
				fechaterminacionanticipada = fecha["Fecha"].(string)
				fechaterminacionanticipada = time_bogota.TiempoCorreccionFormato(fechaterminacionanticipada)
			}
			if nombrefecha == "FechaOficio" {
				fechaoficio = fecha["Fecha"].(string)
				fechaoficio = time_bogota.TiempoCorreccionFormato(fechaoficio)
			}
		}
		for _, propiedad := range propiedades {
			tipopropiedad := propiedad["IdTipoPropiedad"].(map[string]interface{})
			nombrepropiedad := tipopropiedad["Nombre"]
			if nombrepropiedad == "Cedente" {
				cedente = propiedad["Propiedad"]
			}
			if nombrepropiedad == "Cesionario" {
				cesionario = propiedad["Propiedad"]
			}
			if nombrepropiedad == "NumeroActaEntrega" {
				numeroactaentrega = propiedad["Propiedad"]
			}
			if nombrepropiedad == "NumeroOficioEstadoCuentas" {
				numerooficioestadocuentas = propiedad["Propiedad"]
			}
			if nombrepropiedad == "PeriodoSuspension" {
				periodosuspension = propiedad["Propiedad"]
			}
			if nombrepropiedad == "PlazoActual" {
				plazoactual = propiedad["Propiedad"]
			}
			if nombrepropiedad == "TiempoProrroga" {
				tiempoprorroga = propiedad["Propiedad"]
			}
			if nombrepropiedad == "ValorAdicion" {
				valoradicion = propiedad["Propiedad"]
			}
			if nombrepropiedad == "ValorFinalContrato" {
				valorfinalcontrato = propiedad["Propiedad"]
			}
		}

		for _, poliz := range poliza {
			polizas = poliz["NumeroPolizaId"]
			entidadaseguradora = poliz["EntidadAseguradoraId"]
		}

		NovedadAdicionGet = map[string]interface{}{
			//"Id":              (idultimanovedad.(float64) + 1),
			"NumeroContrato":  strconv.FormatFloat(NovedadAdicion["ContratoId"].(float64), 'f', -1, 64),
			"Vigencia":        NovedadAdicion["Vigencia"].(float64),
			"TipoNovedad":     219,
			"FechaInicio":     nil,
			"FechaFin":        "0001-01-01T00:00:00Z",
			"FechaRegistro":   fechasolicitud,
			"Contratista":     cesionario.(float64),
			"NumeroCdp":       nil,
			"VigenciaCdp":     nil,
			"PlazoEjecucion":  plazoactual.(float64),
			"UnidadEjecucion": 205,
			"ValorNovedad":    nil,
		}

		fmt.Println(error, error1, error2)
	}

	return NovedadAdicionGet
}
