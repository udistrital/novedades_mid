package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func ConstruirNovedadReinicio(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadReinicio := make(map[string]interface{})
	NovedadReinicio = novedad

	NovedadReinicioPost := make(map[string]interface{})
	contratoid, _ := strconv.ParseInt(NovedadReinicio["contrato"].(string), 10, 32)
	//numerocdpid, _ := strconv.ParseInt(NovedadSuspension["numerocdp"].(string), 10, 32)
	numerosolicitudentero := NovedadReinicio["numerosolicitud"].(float64)
	numerosolicitud := strconv.FormatFloat(numerosolicitudentero, 'f', -1, 64)
	vigencia, _ := strconv.ParseInt(NovedadReinicio["vigencia"].(string), 10, 32)

	NovedadReinicioPost["NovedadPoscontractual"] = map[string]interface{}{
		"Aclaracion":        nil,
		"Activo":            true,
		"ContratoId":        contratoid,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"Motivo":            NovedadReinicio["motivo"],
		"NumeroCdpId":       0,
		"NumeroSolicitud":   numerosolicitud,
		"Observacion":       NovedadReinicio["observacion"],
		"TipoNovedad":       3,
		"Vigencia":          vigencia,
		"Estado":            NovedadReinicio["estado"],
		"EnlaceDocumento":   NovedadReinicio["enlace"],
	}

	fechas := make([]map[string]interface{}, 0)

	loc, _ := time.LoadLocation("America/Bogota")
	f_solicitud, _ := time.Parse("2006-01-02T15:04:05Z07:00", NovedadReinicio["fechasolicitud"].(string))

	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadReinicio["fecharegistro"],
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
		"Fecha":             NovedadReinicio["fecha_terminacion_anticipada"],
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
		"Fecha":             NovedadReinicio["fechasuspension"],
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
		"Fecha":             NovedadReinicio["fechafinsuspension"],
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
		"Fecha":             NovedadReinicio["fechareinicio"],
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
		"Fecha":             NovedadReinicio["fechafinefectiva"],
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

	NovedadReinicioPost["Fechas"] = fechas

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
			"Id": 9,
		},
		"propiedad": NovedadReinicio["numerooficioestadocuentas"],
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
		"propiedad": NovedadReinicio["valor_desembolsado"],
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
		"propiedad": NovedadReinicio["saldo_contratista"],
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
		"propiedad": NovedadReinicio["saldo_universidad"],
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
		"propiedad": NovedadReinicio["periodosuspension"],
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
		"propiedad": NovedadReinicio["cesionario"],
	})

	NovedadReinicioPost["Propiedad"] = propiedades

	return
}

func GetNovedadReinicio(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadAdicion := make(map[string]interface{})
	var fechas []map[string]interface{}
	var propiedades []map[string]interface{}
	NovedadAdicion = novedad
	NovedadAdicionGet := make(map[string]interface{})
	var fecharegistro interface{}
	var fechareinicio interface{}
	var fechasolicitud interface{}
	var fechasuspension interface{}
	var fechaterminacionanticipada interface{}
	var fechafinefectiva interface{}
	var tiponovedad []map[string]interface{}
	var tipoNovedadNombre string
	var estadoNovedad map[string]interface{}
	var nombreEstadoNov string

	var cesionario interface{}
	var numerooficioestadocuentas interface{}
	var periodosuspension interface{}

	error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/fechas/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &fechas)
	error1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/propiedad/?query=id_novedades_poscontractuales:"+strconv.FormatFloat((NovedadAdicion["Id"]).(float64), 'f', -1, 64)+"&limit=0", &propiedades)
	error2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/tipo_novedad/?query=Id:"+strconv.FormatFloat((NovedadAdicion["TipoNovedad"]).(float64), 'f', -1, 64), &tiponovedad)
	error3 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro/"+NovedadAdicion["Estado"].(string), &estadoNovedad)

	if len(fechas[0]) != 0 {
		for _, fecha := range fechas {
			tipofecha := fecha["IdTipoFecha"].(map[string]interface{})
			nombrefecha := tipofecha["Nombre"]
			if nombrefecha == "FechaRegistro" {
				fecharegistro = fecha["Fecha"]
			}
			if nombrefecha == "FechaReinicio" {
				fechareinicio = fecha["Fecha"]
			}
			if nombrefecha == "FechaSolicitud" {
				fechasolicitud = fecha["Fecha"]
			}
			if nombrefecha == "FechaSolicitud" {
				fechasolicitud = fecha["Fecha"]
			}
			if nombrefecha == "FechaSuspension" {
				fechasuspension = fecha["Fecha"]
			}
			if nombrefecha == "FechaTerminacionAnticipada" {
				fechaterminacionanticipada = fecha["Fecha"]
			}
			if nombrefecha == "FechaFinEfectiva" {
				fechafinefectiva = fecha["Fecha"]
			}
		}
	}
	if len(propiedades[0]) != 0 {
		for _, propiedad := range propiedades {
			tipopropiedad := propiedad["IdTipoPropiedad"].(map[string]interface{})
			nombrepropiedad := tipopropiedad["Nombre"]
			if nombrepropiedad == "NumeroOficioEstadoCuentas" {
				numerooficioestadocuentas = propiedad["Propiedad"]
			}
			if nombrepropiedad == "Cesionario" {
				cesionario = propiedad["Propiedad"]
			}
			if nombrepropiedad == "PeriodoSuspension" {
				periodosuspension = propiedad["Propiedad"]
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
		}
	}

	NovedadAdicionGet = map[string]interface{}{
		"id":                         NovedadAdicion["Id"].(float64),
		"aclaracion":                 "",
		"camposaclaracion":           "",
		"camposmodificacion":         "",
		"camposmodificados":          "",
		"cedente":                    "",
		"cesionario":                 cesionario,
		"contrato":                   NovedadAdicion["ContratoId"],
		"fechaadicion":               "",
		"fechacesion":                "",
		"fechaliquidacion":           "",
		"fechaprorroga":              "",
		"fecharegistro":              fecharegistro,
		"fechareinicio":              fechareinicio,
		"fechasolicitud":             fechasolicitud,
		"fechasuspension":            fechasuspension,
		"fechaterminacionanticipada": fechaterminacionanticipada,
		"motivo":                     NovedadAdicion["Motivo"],
		"numeroactaentrega":          "",
		"numerocdp":                  "",
		"numerooficioestadocuentas":  numerooficioestadocuentas,
		"numerosolicitud":            NovedadAdicion["NumeroSolicitud"],
		"observacion":                NovedadAdicion["Observacion"],
		"periodosuspension":          periodosuspension,
		"plazoactual":                "",
		"poliza":                     "",
		"tiempoprorroga":             "",
		"tiponovedad":                NovedadAdicion["TipoNovedad"],
		"nombreTipoNovedad":          tipoNovedadNombre,
		"valoradicion":               "",
		"valorfinalcontrato":         "",
		"vigencia":                   NovedadAdicion["Vigencia"],
		"fechafinefectiva":           fechafinefectiva,
		"estado":                     nombreEstadoNov,
		"enlace":                     NovedadAdicion["EnlaceDocumento"],
	}

	fmt.Println(error, error1)

	return NovedadAdicionGet
}

func ReplicaReinicio(novedad map[string]interface{}, idStr string) (result map[string]interface{}, outputError map[string]interface{}) {

	ArgoReinicioPost := make(map[string]interface{})
	ArgoReinicioPost = map[string]interface{}{
		"NumeroContrato":  novedad["NumeroContrato"],
		"Vigencia":        novedad["Vigencia"],
		"FechaRegistro":   novedad["FechaRegistro"],
		"PlazoEjecucion":  novedad["PlazoEjecucion"],
		"FechaInicio":     novedad["FechaInicio"],
		"FechaFin":        novedad["FechaFin"],
		"UnidadEjecucion": novedad["UnidadEjecucion"],
		"TipoNovedad":     novedad["TipoNovedad"],
	}

	TitanReinicioPost := make(map[string]interface{})
	TitanReinicioPost = map[string]interface{}{
		"Documento":      novedad["Documento"],
		"FechaReinicio":  FormatFechaReplica(novedad["FechaReinicio"].(string), "2006-01-02T15:04:05.000Z"),
		"NumeroContrato": novedad["NumeroContrato"],
		"Vigencia":       novedad["Vigencia"],
	}

	fmt.Println("ArgoReinicioPost: ", ArgoReinicioPost)
	fmt.Println("TitanReinicioPost: ", TitanReinicioPost)

	url := "/novedad_postcontractual/" + idStr
	if err := SendJson(beego.AppConfig.String("AdministrativaAmazonService")+url, "PUT", &result, &ArgoReinicioPost); err == nil {
		url = "/novedadCPS/reiniciar_contrato"
		if err := SendJson(beego.AppConfig.String("TitanMidService")+url, "POST", &result, &TitanReinicioPost); err == nil {
			return result, nil
		} else {
			outputError = map[string]interface{}{"funcion": "/ReplicaReinicio", "err": err.Error()}
			return nil, outputError
		}
	} else {
		outputError = map[string]interface{}{"funcion": "/ReplicaReinicio", "err": err.Error()}
		return nil, outputError
	}
}
