package models

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func Temporizador() {
	tdr := time.Tick(86400 * time.Second)

	for horaActual := range tdr {
		ReplicaFechaPosterior()
		fmt.Println("Registro realizado en la fecha", horaActual)
	}
}

func ReplicafechaAnterior(informacionReplica map[string]interface{}) (result map[string]interface{}, outputError map[string]interface{}) {

	ArgoNovedadPost := make(map[string]interface{})
	TitanNovedadPost := make(map[string]interface{})
	var url = ""

	ArgoNovedadPost = map[string]interface{}{
		"NumeroContrato":  fmt.Sprintf("%v", informacionReplica["NumeroContrato"]),
		"Vigencia":        informacionReplica["Vigencia"],
		"FechaRegistro":   informacionReplica["FechaRegistro"],
		"Contratista":     informacionReplica["Contratista"],
		"PlazoEjecucion":  informacionReplica["PlazoEjecucion"],
		"FechaInicio":     informacionReplica["FechaInicio"],
		"FechaFin":        informacionReplica["FechaFin"],
		"UnidadEjecucion": informacionReplica["UnidadEjecucion"],
		"TipoNovedad":     informacionReplica["TipoNovedad"],
		"NumeroCdp":       informacionReplica["NumeroCdp"],
		"VigenciaCdp":     informacionReplica["VigenciaCdp"],
	}

	TitanNovedadPost = map[string]interface{}{
		"NumeroContrato": informacionReplica["NumeroContrato"],
		"Vigencia":       informacionReplica["Vigencia"],
	}

	// fmt.Println("ArgoNovedadPost", ArgoNovedadPost)
	// fmt.Println("TitanNovedadPost", TitanNovedadPost)

	// Elabora la estructura para Titán según el tipo de novedad
	if int(informacionReplica["TipoNovedad"].(float64)) == 216 {
		TitanNovedadPost["Documento"] = informacionReplica["Documento"]
		TitanNovedadPost["FechaInicio"] = FormatFechaTitan(informacionReplica["FechaInicio"].(string))
		TitanNovedadPost["FechaFin"] = FormatFechaTitan(informacionReplica["FechaFin"].(string))
		url = "/novedadCPS/suspender_contrato"
	}
	if int(informacionReplica["TipoNovedad"].(float64)) == 219 {
		TitanNovedadPost["DocumentoActual"] = informacionReplica["DocumentoActual"]
		TitanNovedadPost["DocumentoNuevo"] = informacionReplica["DocumentoNuevo"]
		TitanNovedadPost["FechaInicio"] = FormatFechaTitan(informacionReplica["FechaInicio"].(string))
		TitanNovedadPost["NombreCompleto"] = informacionReplica["NombreCompleto"]
		url = "/novedadCPS/ceder_contrato"
	}
	if int(informacionReplica["TipoNovedad"].(float64)) == 220 {
		TitanNovedadPost["Documento"] = informacionReplica["Documento"]
		TitanNovedadPost["FechaFin"] = FormatFechaTitan(informacionReplica["FechaFin"].(string))
		url = "/novedadCPS/otrosi_contrato"
	}
	if int(informacionReplica["TipoNovedad"].(float64)) == 218 {
		TitanNovedadPost["Documento"] = informacionReplica["Documento"]
		TitanNovedadPost["FechaCancelacion"] = FormatFechaTitan(informacionReplica["FechaFin"].(string))
		url = "/novedadCPS/cancelar_contrato"
	}
	if result, err := PostReplica(url, ArgoNovedadPost, TitanNovedadPost); err == nil {
		return result, nil
	} else {
		outputError = map[string]interface{}{"funcion": "/ReplicafechaAnterior", "err": err, "status": "502"}
		return nil, outputError
	}

}

func ReplicaFechaPosterior() {

	currentDate := time.Now()
	lastMonth := currentDate.AddDate(0, -1, 0)
	goneDayMonth := lastMonth.Day()
	firstDay := lastMonth.AddDate(0, 0, -goneDayMonth+1)
	timeLayout := "2006-01-02"
	fechaReferencia := currentDate.Format(timeLayout)

	var fechasResponse []map[string]interface{}
	var replicaResult map[string]interface{}
	var outputError map[string]interface{}

	url := "/fechas?limit=-1"
	if err := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+url, &fechasResponse); err == nil {
		for _, fechaRegistro := range fechasResponse {
			fechaParse, _ := time.Parse("2006-01-02 15:04:05 +0000 +0000", fmt.Sprint(fechaRegistro["Fecha"]))
			if fechaParse.After(firstDay) {
				fecha := fechaParse.Format(timeLayout)
				if fecha == fechaReferencia {
					idTipoFecha := fechaRegistro["IdTipoFecha"].(map[string]interface{})
					tipoFecha := idTipoFecha["Id"]
					if tipoFecha != 3 && tipoFecha != 5 && tipoFecha != 7 && tipoFecha != 10 && tipoFecha != 11 {
						novedad := fechaRegistro["IdNovedadesPoscontractuales"].(map[string]interface{})
						if replicaResult, outputError = ConsultarTipoNovedad(novedad); outputError == nil {
							fmt.Println("Replica realizada correctamente (Temporizador)")
							fmt.Println(replicaResult)
						} else {
							fmt.Println("Fallo al realizar la réplica (Temporizador)")
							fmt.Println(outputError)
						}
					}
				}
			}
		}
	} else {
		fmt.Println(err)
	}
}

func ConsultarTipoNovedad(novedad map[string]interface{}) (result map[string]interface{}, outputError map[string]interface{}) {

	tipoNovedad := int(novedad["TipoNovedad"].(float64))

	var propiedades []map[string]interface{}

	url := "/propiedad?query=IdNovedadesPoscontractuales.Id:" + fmt.Sprintf("%v", novedad["Id"]) + "&limit=0"
	if err := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+url, &propiedades); err == nil {
		switch tipoNovedad {
		case 1:
			return ReplicaSuspension(novedad, propiedades)
		case 2:
			return ReplicaCesion(novedad, propiedades)
		case 3:
			return ReplicaTempReinicio(novedad, propiedades)
		case 5:
			return ReplicaTerminacion(novedad, propiedades)
		case 6:
			return ReplicaAdicionProrroga(novedad, propiedades)
		case 7:
			return ReplicaAdicionProrroga(novedad, propiedades)
		case 8:
			return ReplicaAdicionProrroga(novedad, propiedades)
		default:
			outputError = map[string]interface{}{"funcion": "/ConsultarTipoNovedad", "err": "El tipo de novedad no es válido"}
			return nil, outputError
		}
	} else {
		outputError = map[string]interface{}{"funcion": "/ConsultarTipoNovedad", "err": err.Error()}
		return nil, outputError
	}
}

func ReplicaSuspension(novedad map[string]interface{}, propiedades []map[string]interface{}) (result map[string]interface{}, outputError map[string]interface{}) {

	numContrato := int(novedad["ContratoId"].(float64))
	vigencia := int(novedad["Vigencia"].(float64))

	var cesionario float64
	var contratistaDoc string
	var periodoSuspension float64
	var contratoSuscrito []map[string]interface{}
	var actaInicio []map[string]interface{}
	var informacion_proveedor []map[string]interface{}
	var fechaInicio time.Time
	var fechaFin time.Time
	var url = ""

	if len(propiedades[0]) != 0 {
		for _, propiedad := range propiedades {
			tipopropiedad := propiedad["IdTipoPropiedad"].(map[string]interface{})
			nombrepropiedad := tipopropiedad["Nombre"]
			if nombrepropiedad == "Cesionario" {
				cesionario = propiedad["Propiedad"].(float64)
				url = "/informacion_proveedor?query=Id:" + fmt.Sprintf("%v", propiedad["Propiedad"])
				if error := request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &informacion_proveedor); error == nil {
					contratistaDoc = informacion_proveedor[0]["NumDocumento"].(string)
				}
			}
			if nombrepropiedad == "PeriodoSuspension" {
				periodoSuspension = propiedad["Propiedad"].(float64)
			}
		}
	}

	url = "/contrato_suscrito?query=NumeroContratoSuscrito:" + fmt.Sprintf("%v", novedad["ContratoId"])
	if error := request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &contratoSuscrito); error == nil {
		numContratoActa := contratoSuscrito[len(contratoSuscrito)-1]["NumeroContrato"].(map[string]interface{})
		url = "/acta_inicio?query=NumeroContrato:" + fmt.Sprintf("%v", numContratoActa["Id"])
		if error = request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &actaInicio); error == nil {
			fechaInicio, _ = time.Parse("2006-01-02T00:00:00Z", fmt.Sprint(actaInicio[0]["FechaFin"]))
			fechaInicio = fechaInicio.AddDate(0, 0, 1)
			if fechaInicio.Day() == 31 {
				fechaInicio = fechaInicio.AddDate(0, 0, 1)
			}
			fechaFin = CalcularFechaFin(fechaInicio, periodoSuspension)
		}
	}

	ArgoSuspensionPost := make(map[string]interface{})
	ArgoSuspensionPost = map[string]interface{}{
		"NumeroContrato":  fmt.Sprintf("%v", numContrato),
		"Vigencia":        vigencia,
		"FechaRegistro":   time.Now().Format("2006-01-02"),
		"Contratista":     cesionario,
		"PlazoEjecucion":  periodoSuspension,
		"FechaInicio":     fechaInicio.Format("2006-01-02"),
		"FechaFin":        fechaFin.Format("2006-01-02"),
		"UnidadEjecucion": 205,
		"TipoNovedad":     216,
	}
	fmt.Println("ArgoOtrosiPost:", ArgoSuspensionPost)

	TitanSuspensionPost := make(map[string]interface{})
	TitanSuspensionPost = map[string]interface{}{
		"Documento":      contratistaDoc,
		"FechaFin":       fechaFin.Format("2006-01-02 15:04:05"),
		"FechaInicio":    fechaInicio.Format("2006-01-02 15:04:05"),
		"NumeroContrato": strconv.Itoa(numContrato),
		"Vigencia":       strconv.Itoa(vigencia),
	}
	fmt.Println("TitanOtrosiPost:", TitanSuspensionPost)

	url = "/novedadCPS/suspender_contrato"
	if result, err := PostReplica(url, ArgoSuspensionPost, TitanSuspensionPost); err == nil {
		return result, nil
	} else {
		outputError = map[string]interface{}{"funcion": "/ReplicaSuspension", "err": err}
		return nil, err
	}
}

func ReplicaCesion(novedad map[string]interface{}, propiedades []map[string]interface{}) (result map[string]interface{}, outputError map[string]interface{}) {

	numContrato := int(novedad["ContratoId"].(float64))
	vigencia := int(novedad["Vigencia"].(float64))

	var cesionario float64
	var cesionarioDoc string
	var nombreCesionario string
	var cedenteDoc string
	var contratoSuscrito []map[string]interface{}
	var actaInicio []map[string]interface{}
	var informacion_proveedor []map[string]interface{}
	var fechas []map[string]interface{}
	var fechaInicio time.Time
	var fechaFin time.Time
	var diasCesion float64
	var url = ""

	if len(propiedades[0]) != 0 {
		for _, propiedad := range propiedades {
			tipopropiedad := propiedad["IdTipoPropiedad"].(map[string]interface{})
			nombrepropiedad := tipopropiedad["Nombre"]
			if nombrepropiedad == "Cesionario" {
				cesionario = propiedad["Propiedad"].(float64)
				url = "/informacion_proveedor?query=Id:" + fmt.Sprintf("%v", propiedad["Propiedad"])
				if error := request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &informacion_proveedor); error == nil {
					cesionarioDoc = informacion_proveedor[0]["NumDocumento"].(string)
					nombreCesionario = informacion_proveedor[0]["NomProveedor"].(string)
				}
			}
			if nombrepropiedad == "Cedente" {
				url = "/informacion_proveedor?query=Id:" + fmt.Sprintf("%v", propiedad["Propiedad"])
				if error := request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &informacion_proveedor); error == nil {
					cedenteDoc = informacion_proveedor[0]["NumDocumento"].(string)
				}
			}
		}
	}

	url = "/fechas?query=IdNovedadesPoscontractuales.Id:" + fmt.Sprintf("%v", novedad["Id"]) + "&limit=0"
	if error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+url, &fechas); error == nil {
		for _, fecha := range fechas {
			tipoFecha := fecha["IdTipoFecha"].(map[string]interface{})
			nombreFecha := tipoFecha["Nombre"]
			if nombreFecha == "FechaCesion" {
				fechaInicio = fecha["Fecha"].(time.Time)
			}
		}
	}

	url = "/contrato_suscrito?query=NumeroContratoSuscrito:" + fmt.Sprintf("%v", novedad["ContratoId"])
	if error := request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &contratoSuscrito); error == nil {
		numContratoActa := contratoSuscrito[len(contratoSuscrito)-1]["NumeroContrato"].(map[string]interface{})
		url = "/acta_inicio?query=NumeroContrato:" + fmt.Sprintf("%v", numContratoActa["Id"])
		if error = request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &actaInicio); error == nil {
			fechaFin, _ = time.Parse("2006-01-02T00:00:00Z", fmt.Sprint(actaInicio[0]["FechaFin"]))
			diasCesion = fechaFin.Sub(fechaInicio).Hours() / 24
		}
	}

	ArgoCesionPost := make(map[string]interface{})
	ArgoCesionPost = map[string]interface{}{
		"NumeroContrato":  fmt.Sprintf("%v", numContrato),
		"Vigencia":        vigencia,
		"FechaRegistro":   time.Now().Format("2006-01-02"),
		"Contratista":     cesionario,
		"PlazoEjecucion":  int(diasCesion),
		"FechaInicio":     fechaInicio.Format("2006-01-02"),
		"FechaFin":        fechaFin.Format("2006-01-02"),
		"UnidadEjecucion": 205,
		"TipoNovedad":     219,
	}

	TitanCesionPost := make(map[string]interface{})
	TitanCesionPost = map[string]interface{}{
		"DocumentoActual": cedenteDoc,
		"DocumentoNuevo":  cesionarioDoc,
		"FechaInicio":     fechaInicio.Format("2006-01-02 15:04:05"),
		"NombreCompleto":  nombreCesionario,
		"NumeroContrato":  strconv.Itoa(numContrato),
		"Vigencia":        strconv.Itoa(vigencia),
	}

	url = "/novedadCPS/ceder_contrato"
	if result, err := PostReplica(url, ArgoCesionPost, TitanCesionPost); err == nil {
		return result, nil
	} else {
		outputError = map[string]interface{}{"funcion": "/ReplicaCesion", "err": err}
		return nil, err
	}
}

func ReplicaTempReinicio(novedad map[string]interface{}, propiedades []map[string]interface{}) (result map[string]interface{}, outputError map[string]interface{}) {

	numContrato := int(novedad["ContratoId"].(float64))
	vigencia := int(novedad["Vigencia"].(float64))

	var contratistaDoc string
	var informacion_proveedor []map[string]interface{}
	var fechas []map[string]interface{}
	var fechasuspension time.Time
	var fechaReinicio time.Time
	var periodoSuspension float64
	var url = ""
	var idStr = ""

	if len(propiedades[0]) != 0 {
		for _, propiedad := range propiedades {
			tipopropiedad := propiedad["IdTipoPropiedad"].(map[string]interface{})
			nombrepropiedad := tipopropiedad["Nombre"]
			if nombrepropiedad == "Cesionario" {
				idStr = fmt.Sprintf("%v", propiedad["Propiedad"])
				url = "/informacion_proveedor?query=Id:" + idStr
				if err := request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &informacion_proveedor); err == nil {
					contratistaDoc = informacion_proveedor[0]["NumDocumento"].(string)
				}
			}
			if nombrepropiedad == "PeriodoSuspension" {
				periodoSuspension = propiedad["Propiedad"].(float64)
			}
		}
	}

	url = "/fechas?query=IdNovedadesPoscontractuales.Id:" + fmt.Sprintf("%v", novedad["Id"]) + "&limit=0"
	if err := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+url, &fechas); err == nil {
		for _, fecha := range fechas {
			tipoFecha := fecha["IdTipoFecha"].(map[string]interface{})
			nombreFecha := tipoFecha["Nombre"]
			if nombreFecha == "FechaReinicio" {
				fechaReinicio = fecha["Fecha"].(time.Time)
			}
			if nombreFecha == "FechaSuspension" {
				fechasuspension = fecha["Fecha"].(time.Time)
			}
		}
	}

	ArgoReinicioPost := make(map[string]interface{})
	ArgoReinicioPost = map[string]interface{}{
		"NumeroContrato": fmt.Sprintf("%v", numContrato),
		"Vigencia":       vigencia,
		"FechaRegistro":  time.Now().Format("2006-01-02"),
		"Contratista":    idStr,
		"PlazoEjecucion": periodoSuspension,
		"FechaInicio":    fechasuspension.Format("2006-01-02"),
		// "FechaFin":        fechaFin.Format("2006-01-02"),
		"UnidadEjecucion": 205,
		"TipoNovedad":     216,
	}

	TitanReinicioPost := make(map[string]interface{})
	TitanReinicioPost = map[string]interface{}{
		"Documento":      contratistaDoc,
		"FechaReinicio":  fechaReinicio.Format("2006-01-02 15:04:05"),
		"NumeroContrato": strconv.Itoa(numContrato),
		"Vigencia":       strconv.Itoa(vigencia),
	}

	url = "/novedad_postcontractual/" + idStr
	if err := SendJson(beego.AppConfig.String("AdministrativaAmazonService")+url, "PUT", &result, &ArgoReinicioPost); err == nil {
		if err = SendJson(beego.AppConfig.String("TitanMidService")+"/novedadCPS/reiniciar_contrato", "POST", &result, &TitanReinicioPost); err == nil {
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

func ReplicaTerminacion(novedad map[string]interface{}, propiedades []map[string]interface{}) (result map[string]interface{}, outputError map[string]interface{}) {

	numContrato := int(novedad["ContratoId"].(float64))
	vigencia := int(novedad["Vigencia"].(float64))

	var contratistaDoc string
	var fechas []map[string]interface{}
	var fechaFin time.Time
	var url = ""

	url = "/fechas?query=IdNovedadesPoscontractuales.Id:" + fmt.Sprintf("%v", novedad["Id"]) + "&limit=0"
	if error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+url, &fechas); error == nil {
		for _, fecha := range fechas {
			tipoFecha := fecha["IdTipoFecha"].(map[string]interface{})
			nombreFecha := tipoFecha["Nombre"]
			if nombreFecha == "FechaTerminacionAnticipada" {
				fechaFin = fecha["Fecha"].(time.Time)
			}
		}
	}

	ArgoTerminacionPost := make(map[string]interface{})
	ArgoTerminacionPost = map[string]interface{}{
		"NumeroContrato": fmt.Sprintf("%v", numContrato),
		"Vigencia":       vigencia,
		"FechaRegistro":  time.Now().Format("2006-01-02"),
		"FechaFin":       fechaFin.Format("2006-01-02"),
		"TipoNovedad":    218,
	}

	TitanTerminacionPost := make(map[string]interface{})
	TitanTerminacionPost = map[string]interface{}{
		"Documento":        contratistaDoc,
		"FechaCancelacion": fechaFin.Format("2006-01-02 15:04:05"),
		"NumeroContrato":   numContrato,
		"Vigencia":         vigencia,
	}

	url = "/novedadCPS/cancelar_contrato"
	if result, err := PostReplica(url, ArgoTerminacionPost, TitanTerminacionPost); err == nil {
		return result, nil
	} else {
		outputError = map[string]interface{}{"funcion": "/ReplicaTerminacion", "err": err}
		return nil, err
	}
}

func ReplicaAdicionProrroga(novedad map[string]interface{}, propiedades []map[string]interface{}) (result map[string]interface{}, outputError map[string]interface{}) {

	numContrato := int(novedad["ContratoId"].(float64))
	vigencia := int(novedad["Vigencia"].(float64))
	numeroCdp := int(novedad["NumeroCdpId"].(float64))
	vigenciaCdp := int(novedad["VigenciaCdp"].(float64))

	var tipoNovedad float64
	if novedad["TipoNovedad"] == 6 {
		tipoNovedad = 248
	} else if novedad["TipoNovedad"] == 7 {
		tipoNovedad = 249
	} else if novedad["TipoNovedad"] == 8 {
		tipoNovedad = 220
	}
	var cesionario float64
	var valoradicion float64 = 0
	var tiempoprorroga float64 = 0
	var contratistaDoc string

	var contratoSuscrito []map[string]interface{}
	var actaInicio []map[string]interface{}
	var informacion_proveedor []map[string]interface{}
	var fechaInicio time.Time
	var fechaFin time.Time
	var url = ""

	if len(propiedades[0]) != 0 {
		for _, propiedad := range propiedades {
			tipopropiedad := propiedad["IdTipoPropiedad"].(map[string]interface{})
			nombrepropiedad := tipopropiedad["Nombre"]
			if nombrepropiedad == "Cesionario" {
				cesionario = propiedad["Propiedad"].(float64)
				url = "/informacion_proveedor?query=Id:" + fmt.Sprintf("%v", propiedad["Propiedad"])
				if error := request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &informacion_proveedor); error == nil {
					contratistaDoc = informacion_proveedor[0]["NumDocumento"].(string)
				}
			}
			if nombrepropiedad == "ValorAdicion" {
				valoradicion = propiedad["Propiedad"].(float64)
			}
			if nombrepropiedad == "TiempoProrroga" {
				tiempoprorroga = propiedad["Propiedad"].(float64)
			}
		}
	}

	url = "/contrato_suscrito?query=NumeroContratoSuscrito:" + fmt.Sprintf("%v", novedad["ContratoId"])
	if error := request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &contratoSuscrito); error == nil {
		numContratoActa := contratoSuscrito[len(contratoSuscrito)-1]["NumeroContrato"].(map[string]interface{})
		url = "/acta_inicio?query=NumeroContrato:" + fmt.Sprintf("%v", numContratoActa["Id"])
		if error = request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &actaInicio); error == nil {
			fechaInicio, _ = time.Parse("2006-01-02T00:00:00Z", fmt.Sprint(actaInicio[0]["FechaFin"]))
			fechaInicio = fechaInicio.AddDate(0, 0, 1)
			if fechaInicio.Day() == 31 {
				fechaInicio = fechaInicio.AddDate(0, 0, 1)
			}
			fechaFin = CalcularFechaFin(fechaInicio, tiempoprorroga)
		}
	}

	ArgoOtrosiPost := make(map[string]interface{})
	ArgoOtrosiPost = map[string]interface{}{
		"NumeroContrato":  fmt.Sprintf("%v", numContrato),
		"Vigencia":        vigencia,
		"FechaRegistro":   time.Now().Format("2006-01-02"),
		"Contratista":     cesionario,
		"PlazoEjecucion":  tiempoprorroga,
		"FechaInicio":     fechaInicio.Format("2006-01-02"),
		"FechaFin":        fechaFin.Format("2006-01-02"),
		"NumeroCdp":       numeroCdp,
		"VigenciaCdp":     vigenciaCdp,
		"ValorNovedad":    valoradicion,
		"UnidadEjecucion": 205,
		"TipoNovedad":     tipoNovedad,
	}

	TitanOtrosiPost := make(map[string]interface{})
	TitanOtrosiPost["NovedadPoscontractual"] = map[string]interface{}{
		"Documento":      contratistaDoc,
		"FechaFin":       fechaFin.Format("2006-01-02 15:04:05"),
		"NumeroContrato": strconv.Itoa(numContrato),
		"Vigencia":       strconv.Itoa(vigencia),
	}

	url = "/novedadCPS/otrosi_contrato"
	if result, err := PostReplica(url, ArgoOtrosiPost, TitanOtrosiPost); err == nil {
		return result, nil
	} else {
		outputError = map[string]interface{}{"funcion": "/ReplicaAdicionProrroga", "err": err}
		return nil, err
	}
}

func PostReplica(url string, ArgoOtrosiPost map[string]interface{}, TitanOtrosiPost map[string]interface{}) (resultPost map[string]interface{}, outputError map[string]interface{}) {
	if err := SendJson(beego.AppConfig.String("AdministrativaAmazonService")+"/novedad_postcontractual", "POST", &resultPost, &ArgoOtrosiPost); err == nil {
		if err := SendJson(beego.AppConfig.String("TitanMidService")+url, "POST", &resultPost, &TitanOtrosiPost); err == nil {
			fmt.Println("Registro en Titan exitoso!")
			return resultPost, nil
		} else {
			outputError = map[string]interface{}{"funcion": "/PostReplica", "err": err.Error()}
			return nil, outputError
		}
	} else {
		outputError = map[string]interface{}{"funcion": "/PostReplica", "err": err.Error()}
		return nil, outputError
	}
}

func CalcularFechaFin(fechaInicio time.Time, diasNovedad float64) (fechaFin time.Time) {
	var dias float64 = diasNovedad
	meses := dias / 30
	mesEntero := int(meses)
	decimal := meses - float64(mesEntero)
	numDias := decimal * 30

	if numDias+float64(fechaInicio.Day()) > 30 {
		dias = (numDias + float64(fechaInicio.Day())) / 30
		mesEntero += 1
		decimal = dias - 1
		numDias = math.Round(decimal * 30)
		if numDias == 1 || numDias == 0 {
			mesEntero -= 1
			numDias = 30
		} else {
			numDias -= 1
		}
		fechaFin = time.Date(fechaInicio.Year(), fechaInicio.Month()+time.Month(mesEntero), int(numDias), 0, 0, 0, 0, fechaInicio.Location())
	} else {
		fechaFin = time.Date(fechaInicio.Year(), fechaInicio.Month()+time.Month(mesEntero), fechaInicio.Day()-1, 0, 0, 0, 0, fechaInicio.Location())
	}
	return fechaFin
}

func FormatFechaTitan(fecha string) string {
	var fechaTitan = ""
	if fechaParse, err := time.Parse("2006-01-02T15:04:05.000Z", fecha); err == nil {
		fechaFormat := fechaParse.Format("2006-01-02T15:04:05.000Z")
		// ("2006-01-02T15:04:05.000Z")
		fechaTitan = fechaFormat
	}
	return fechaTitan
}
