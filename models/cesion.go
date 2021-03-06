package models

import (
	"fmt"
	"strconv"

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
	numerosolicitudentero := NovedadCesion["numerosolicitud"].(float64)
	numerosolicitud := strconv.FormatFloat(numerosolicitudentero, 'f', -1, 64)
	vigencia, _ := strconv.ParseInt(NovedadCesion["vigencia"].(string), 10, 32)
	vigenciacdp, _ := strconv.ParseInt(NovedadCesion["vigencia"].(string), 10, 32)

	NovedadCesionPost["NovedadPoscontractual"] = map[string]interface{}{
		"Aclaracion":        NovedadCesion["aclaracion"],
		"Activo":            true,
		"ContratoId":        contratoid,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"Motivo":            NovedadCesion["motivo"],
		"NumeroCdpId":       numerocdpid,
		"NumeroSolicitud":   numerosolicitud,
		"Observacion":       NovedadCesion["observacion"],
		"TipoNovedad":       2,
		"Vigencia":          vigencia,
		"VigenciaCdp":       vigenciacdp,
	}

	fechas := make([]map[string]interface{}, 0)

	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadCesion["fechaadicion"],
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
		"Fecha":             NovedadCesion["fechaliquidacion"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 3,
		},
	})
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadCesion["fechaprorroga"],
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
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadCesion["fechareinicio"],
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
		"Fecha":             NovedadCesion["fechasolicitud"],
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
		"Fecha":             NovedadCesion["fechasuspension"],
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
		"Fecha":             NovedadCesion["fechaterminacionanticipada"],
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
		"Fecha":             NovedadCesion["fechaoficio"],
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
		"Fecha":             NovedadCesion["fecharegistro"],
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
			"Id": 13,
		},
		"propiedad": NovedadCesion["numeroactaentrega"],
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
		"propiedad": NovedadCesion["numerooficioestadocuentas"],
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
			"Id": 3,
		},
		"propiedad": NovedadCesion["periodosuspension"],
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
			"Id": 4,
		},
		"propiedad": NovedadCesion["plazoactual"],
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
		"propiedad": NovedadCesion["tiempoprorroga"],
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
			"Id": 6,
		},
		"propiedad": NovedadCesion["valoradicion"],
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
			"Id": 8,
		},
		"propiedad": NovedadCesion["valorfinalcontrato"],
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

	// fmt.Println(NovedadCesionPost)

	return NovedadCesionPost
}

func GetNovedadCesion(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadAdicion := make(map[string]interface{})
	var fechas []map[string]interface{}
	var propiedades []map[string]interface{}
	var poliza []map[string]interface{}
	NovedadAdicion = novedad
	NovedadAdicionGet := make(map[string]interface{})
	var fechaadicion interface{}
	var fechacesion interface{}
	var fechaliquidacion interface{}
	var fechaprorroga interface{}
	var fecharegistro interface{}
	var fechareinicio interface{}
	var fechasuspension interface{}
	var fechaterminacionanticipada interface{}
	var fechasolicitud interface{}
	var fechaoficio interface{}

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

	error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/fechas/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &fechas)
	error1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/propiedad/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &propiedades)
	error2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/poliza/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &poliza)

	for _, fecha := range fechas {
		tipofecha := fecha["IdTipoFecha"].(map[string]interface{})
		nombrefecha := tipofecha["Nombre"]
		if nombrefecha == "FechaAdicion" {
			fechaadicion = fecha["Fecha"]
		}
		if nombrefecha == "FechaCesion" {
			fechacesion = fecha["Fecha"]
		}
		if nombrefecha == "FechaLiquidacion" {
			fechaliquidacion = fecha["Fecha"]
		}
		if nombrefecha == "FechaProrroga" {
			fechaprorroga = fecha["Fecha"]
		}
		if nombrefecha == "FechaRegistro" {
			fecharegistro = fecha["Fecha"]
		}
		if nombrefecha == "FechaReinicio" {
			fechareinicio = fecha["Fecha"]
		}
		if nombrefecha == "FechaSuspension" {
			fechasuspension = fecha["Fecha"]
		}
		if nombrefecha == "FechaSolicitud" {
			fechasolicitud = fecha["Fecha"]
		}
		if nombrefecha == "FechaTerminacionAnticipada" {
			fechaterminacionanticipada = fecha["Fecha"]
		}
		if nombrefecha == "FechaOficio" {
			fechaoficio = fecha["Fecha"]
		}
		//fmt.Println(fechaadicion, fechasolicitud)
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
		//fmt.Println(cesionario, valoradicion)
	}

	for _, poliz := range poliza {

		polizas = poliz["NumeroPolizaId"]
		entidadaseguradora = poliz["EntidadAseguradoraId"]
	}

	NovedadAdicionGet = map[string]interface{}{
		"id":                         NovedadAdicion["Id"].(float64),
		"aclaracion":                 NovedadAdicion["Aclaracion"],
		"camposaclaracion":           "",
		"camposmodificacion":         "",
		"camposmodificados":          "",
		"cedente":                    cedente,
		"cesionario":                 cesionario,
		"contrato":                   NovedadAdicion["ContratoId"],
		"fechaadicion":               fechaadicion,
		"fechacesion":                fechacesion,
		"fechaliquidacion":           fechaliquidacion,
		"fechaprorroga":              fechaprorroga,
		"fecharegistro":              fecharegistro,
		"fechareinicio":              fechareinicio,
		"fechasolicitud":             fechasolicitud,
		"fechasuspension":            fechasuspension,
		"fechaterminacionanticipada": fechaterminacionanticipada,
		"motivo":                     NovedadAdicion["Motivo"],
		"numeroactaentrega":          numeroactaentrega,
		"numerocdp":                  NovedadAdicion["NumeroCdpId"],
		"numerooficioestadocuentas":  numerooficioestadocuentas,
		"numerosolicitud":            NovedadAdicion["NumeroSolicitud"],
		"observacion":                NovedadAdicion["Observacion"],
		"periodosuspension":          periodosuspension,
		"plazoactual":                plazoactual,
		"poliza":                     polizas,
		"tiempoprorroga":             tiempoprorroga,
		"tiponovedad":                NovedadAdicion["TipoNovedad"],
		"valoradicion":               valoradicion,
		"valorfinalcontrato":         valorfinalcontrato,
		"vigencia":                   NovedadAdicion["Vigencia"],
		"fechaoficio":                fechaoficio,
		"entidadaseguradora":         entidadaseguradora,
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
			//fmt.Println(fechaadicion, fechasolicitud)
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
			//fmt.Println(cesionario, valoradicion)
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

			// "Id":              503,
			// "NumeroContrato":  "241",
			// "Vigencia":        2017,
			// "TipoNovedad":     219,
			// "FechaInicio":     "2018-06-01T00:00:00Z",
			// "FechaFin":        "2018-12-22T00:00:00Z",
			// "FechaRegistro":   "2018-06-13T00:00:00Z",
			// "Contratista":     11087,
			// "NumeroCdp":       0,
			// "VigenciaCdp":     0,
			// "PlazoEjecucion":  202,
			// "UnidadEjecucion": 205,
			// "ValorNovedad":    0,
		}

		fmt.Println(error, error1, error2)
	}

	return NovedadAdicionGet
}
