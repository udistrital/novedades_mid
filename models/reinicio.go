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
	// umerocdpid, _ := strconv.ParseInt(NovedadSuspension["numerocdp"].(string), 10, 32)
	// numerosolicitudentero := NovedadReinicio["numerosolicitud"].(float64)
	// numerosolicitud := strconv.FormatFloat(numerosolicitudentero, 'f', -1, 64)
	vigencia, _ := strconv.ParseInt(NovedadReinicio["vigencia"].(string), 10, 32)

	codEstado := ""

	var estadoNovedad map[string]interface{}
	error3 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro?query=TipoParametroId.CodigoAbreviacion:ENOV,CodigoAbreviacion:"+NovedadReinicio["estado"].(string), &estadoNovedad)

	if error3 == nil {
		if len(estadoNovedad) != 0 {
			inter := estadoNovedad["Data"].([]interface{})
			data := inter[0].(map[string]interface{})
			idEstado, _ := data["Id"].(float64)
			codEstado = strconv.FormatFloat(idEstado, 'f', -1, 64)
		}
	}

	NovedadReinicioPost["NovedadPoscontractual"] = map[string]interface{}{
		"Aclaracion":        nil,
		"Activo":            true,
		"ContratoId":        contratoid,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"Motivo":            NovedadReinicio["motivo"],
		"NumeroCdpId":       0,
		"NumeroSolicitud":   NovedadReinicio["numerosolicitud"],
		"Observacion":       NovedadReinicio["observacion"],
		"TipoNovedad":       3,
		"Vigencia":          vigencia,
		"OficioSupervisor":  NovedadReinicio["numerooficiosupervisor"],
		"OficioOrdenador":   NovedadReinicio["numerooficioordenador"],
		"Estado":            codEstado,
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
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadReinicio["fechaexpedicion"],
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

	return NovedadReinicioPost
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
	var fechaexpedicion interface{}
	var fechafinsuspension interface{}
	var fechafinefectiva interface{}
	var tiponovedad []map[string]interface{}
	var tipoNovedadNombre string
	var estadoNovedad map[string]interface{}
	var nombreEstadoNov string
	var codEstado string

	var cesionario interface{}
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
		"FechaReinicio":              fechareinicio,
		"FechaSolicitud":             fechasolicitud,
		"FechaSuspension":            fechasuspension,
		"FechaFinSuspension":         fechafinsuspension,
		"FechaFinEfectiva":           fechafinefectiva,
		"FechaTerminacionAnticipada": "",
		"FechaExpedicion":            fechaexpedicion,
		"Motivo":                     NovedadAdicion["Motivo"],
		"NumeroActaEntrega":          "",
		"NumeroCdp":                  NovedadAdicion["NumeroCdpId"],
		"NumeroSolicitud":            NovedadAdicion["NumeroSolicitud"],
		"Observacion":                "",
		"PeriodoSuspension":          periodosuspension,
		"PlazoActual":                0,
		"Poliza":                     "",
		"TiempoProrroga":             0,
		"TipoNovedad":                NovedadAdicion["TipoNovedad"],
		"NombreTipoNovedad":          tipoNovedadNombre,
		"CodAbreviacionTipo":         "NP_REI",
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

	fmt.Println(error, error1)

	return NovedadAdicionGet
}

func ReplicaReinicio(novedad map[string]interface{}, idStr string) (map[string]interface{}, map[string]interface{}) {

	resultPostArgo := make(map[string]interface{})
	resultPostTitan := make(map[string]interface{})
	var outputError map[string]interface{}

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

	// fmt.Println("ArgoReinicioPost: ", ArgoReinicioPost)
	// fmt.Println("TitanReinicioPost: ", TitanReinicioPost)

	url := "/novedad_postcontractual/" + idStr
	if err := SendJson(beego.AppConfig.String("AdministrativaAmazonService")+url, "PUT", &resultPostArgo, &ArgoReinicioPost); err == nil {
		url = "/novedadCPS/reiniciar_contrato"
		if err := SendJson(beego.AppConfig.String("TitanMidService")+url, "POST", &resultPostTitan, &TitanReinicioPost); err == nil {
			if len(resultPostTitan) > 0 {
				status := resultPostTitan["Status"]
				if status == "201" {
					fmt.Println("Registro en Titan exitoso!")
					return resultPostTitan, nil
				} else {
					outputError = map[string]interface{}{"funcion": "/PostReplica_Titan_Status", "err": "Falló el registro en Titan"}
					return nil, outputError
				}
			} else {
				outputError = map[string]interface{}{"funcion": "/PostReplica_Titan", "err": "Falló el registro en Titan"}
				return nil, outputError
			}
		} else {
			outputError = map[string]interface{}{"funcion": "/ReplicaReinicio", "err": err.Error()}
			return nil, outputError
		}
	} else {
		outputError = map[string]interface{}{"funcion": "/ReplicaReinicio", "err": err.Error()}
		return nil, outputError
	}
}
