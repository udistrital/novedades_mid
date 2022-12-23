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

func ReplicafechaAnterior(informacionReplica map[string]interface{}) (outputError map[string]interface{}) {

	ArgoNovedadPost := make(map[string]interface{})
	TitanNovedadPost := make(map[string]interface{})
	var resultPostArgo map[string]interface{}
	var resultPostTitan map[string]interface{}

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

	url := "/novedad_postcontractual"
	if err := SendJson(beego.AppConfig.String("AdministrativaAmazonService")+url, "POST", &resultPostArgo, &ArgoNovedadPost); err == nil {
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
		if err := SendJson(beego.AppConfig.String("TitanMidService")+url, "POST", &resultPostTitan, &TitanNovedadPost); err == nil {
			fmt.Println("Registro en Titan exitoso!")
			return nil
		} else {
			outputError = map[string]interface{}{"funcion": "/ReplicafechaAnterior", "err": err.Error(), "status": "502"}
			fmt.Println(err)
		}
	} else {
		outputError = map[string]interface{}{"funcion": "/ReplicafechaAnterior", "err": err.Error(), "status": "502"}
	}
	return outputError
}

func ReplicaFechaPosterior() {

	currentDate := time.Now()
	lastMonth := currentDate.AddDate(0, -1, 0)
	goneDayMonth := lastMonth.Day()
	firstDay := lastMonth.AddDate(0, 0, -goneDayMonth+1)
	timeLayout := "2006-01-02"
	fechaReferencia := currentDate.Format(timeLayout)

	var fechasResponse []map[string]interface{}

	url := "/fechas?limit=-1"
	if err := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+url, &fechasResponse); err == nil {
		for _, fechaRegistro := range fechasResponse {
			fechaParse, _ := time.Parse("2006-01-02 15:04:05 +0000 +0000", fmt.Sprint(fechaRegistro["Fecha"]))
			if fechaParse.After(firstDay) {
				fecha := fechaParse.Format(timeLayout)
				if fecha == fechaReferencia {
					idTipoFecha := fechaRegistro["IdTipoFecha"].(map[string]interface{})
					tipoFecha := idTipoFecha["Id"]
					if tipoFecha != 3 && tipoFecha != 5 && tipoFecha != 7 && tipoFecha != 9 && tipoFecha != 10 {
						novedad := fechaRegistro["IdNovedadesPoscontractuales"].(map[string]interface{})
						ConsultarTipoNovedad(novedad)
					}
				}
			}
		}
	} else {
		fmt.Println(err)
	}
}

func ConsultarTipoNovedad(novedad map[string]interface{}) (structura map[string]interface{}) {

	tipoNovedad := int(novedad["TipoNovedad"].(float64))

	var propiedades []map[string]interface{}

	url := "/propiedad?query=IdNovedadesPoscontractuales.Id:" + fmt.Sprintf("%v", novedad["Id"]) + "&limit=0"
	if error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+url, &propiedades); error == nil {
		switch tipoNovedad {
		case 1:
			ReplicaSuspension(novedad, propiedades)
		case 2:
			ReplicaCesion(novedad, propiedades)
		// case 3:
		// 	ReplicaReinicio(novedad, propiedades)
		case 6:
			ReplicaAdicionProrroga(novedad, propiedades)
		case 7:
			ReplicaAdicionProrroga(novedad, propiedades)
		case 8:
			ReplicaAdicionProrroga(novedad, propiedades)
		}
	}

	return nil
}

func ReplicaSuspension(novedad map[string]interface{}, propiedades []map[string]interface{}) {

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
	TitanSuspensionPost["NovedadPoscontractual"] = map[string]interface{}{
		"Documento":      contratistaDoc,
		"FechaFin":       fechaFin.Format("2006-01-02 15:04:05"),
		"FechaInicio":    fechaInicio.Format("2006-01-02 15:04:05"),
		"NumeroContrato": strconv.Itoa(numContrato),
		"Vigencia":       strconv.Itoa(vigencia),
	}
	fmt.Println("TitanOtrosiPost:", TitanSuspensionPost)

	var resultPost map[string]interface{}

	url = "/novedad_postcontractual"
	if err := SendJson(beego.AppConfig.String("AdministrativaAmazonService")+url, "POST", &resultPost, &ArgoSuspensionPost); err == nil {
		url = "/novedadCPS/suspender_contrato"
		if err := SendJson(beego.AppConfig.String("TitanMidService")+url, "POST", &resultPost, &TitanSuspensionPost); err == nil {
			fmt.Println("Registro en Titan exitoso!")
		}
	}
}

func ReplicaCesion(novedad map[string]interface{}, propiedades []map[string]interface{}) {

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
	fmt.Println("ArgoCesionPost", ArgoCesionPost)

	TitanSuspensionPost := make(map[string]interface{})
	TitanSuspensionPost["NovedadPoscontractual"] = map[string]interface{}{
		"DocumentoActual": cedenteDoc,
		"DocumentoNuevo":  cesionarioDoc,
		"FechaInicio":     fechaInicio.Format("2006-01-02 15:04:05"),
		"NombreCompleto":  nombreCesionario,
		"NumeroContrato":  strconv.Itoa(numContrato),
		"Vigencia":        strconv.Itoa(vigencia),
	}
	fmt.Println("TitanOtrosiPost:", TitanSuspensionPost)

	var resultPost map[string]interface{}

	url = "/novedad_postcontractual"
	if err := SendJson(beego.AppConfig.String("AdministrativaAmazonService")+url, "POST", &resultPost, &ArgoCesionPost); err == nil {
		url = "/novedadCPS/ceder_contrato"
		if err := SendJson(beego.AppConfig.String("TitanMidService")+url, "POST", &resultPost, &TitanSuspensionPost); err == nil {
			fmt.Println("Registro en Titan exitoso!")
		}
	}
}

func ReplicaReinicio(novedad map[string]interface{}, idStr string) (result map[string]interface{}, outputError map[string]interface{}) {

	// var result map[string]interface{}

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

	// TitanReinicioPost := make(map[string]interface{})
	// TitanReinicioPost["NovedadPoscontractual"] = map[string]interface{}{
	// 	"DocumentoActual": cedenteDoc,
	// 	"DocumentoNuevo":  cesionarioDoc,
	// 	"FechaInicio":     fechaInicio.Format("2006-01-02 15:04:05"),
	// 	"NombreCompleto":  nombreCesionario,
	// 	"NumeroContrato":  strconv.Itoa(numContrato),
	// 	"Vigencia":        strconv.Itoa(vigencia),
	// }

	url := "/novedad_postcontractual/" + idStr
	if err := SendJson(beego.AppConfig.String("AdministrativaAmazonService")+url, "PUT", &result, &ArgoReinicioPost); err == nil {
		return result, nil
	} else {
		outputError = map[string]interface{}{"funcion": "/ReplicaReinicio", "err": err.Error()}
		return nil, outputError
	}
}

func ReplicaAdicionProrroga(novedad map[string]interface{}, propiedades []map[string]interface{}) {

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
	fmt.Println("ArgoOtrosiPost:", ArgoOtrosiPost)

	TitanOtrosiPost := make(map[string]interface{})
	TitanOtrosiPost["NovedadPoscontractual"] = map[string]interface{}{
		"Documento":      contratistaDoc,
		"FechaFin":       fechaFin.Format("2006-01-02 15:04:05"),
		"NumeroContrato": strconv.Itoa(numContrato),
		"Vigencia":       strconv.Itoa(vigencia),
	}
	fmt.Println("TitanOtrosiPost:", TitanOtrosiPost)

	var resultPost map[string]interface{}

	url = "/novedad_postcontractual"
	if err := SendJson(beego.AppConfig.String("AdministrativaAmazonService")+url, "POST", &resultPost, &ArgoOtrosiPost); err == nil {
		url = "/novedadCPS/otrosi_contrato"
		if err := SendJson(beego.AppConfig.String("TitanMidService")+url, "POST", &resultPost, &TitanOtrosiPost); err == nil {
			fmt.Println("Registro en Titan exitoso!")
		}
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
	fmt.Println(fecha)
	if fechaParse, err := time.Parse("2006-01-02T15:04:05.000Z", fecha); err == nil {
		fechaFormat := fechaParse.Format("2006-01-02T15:04:05.000Z")
		// ("2006-01-02T15:04:05.000Z")
		fechaTitan = fechaFormat
	}
	return fechaTitan
}
