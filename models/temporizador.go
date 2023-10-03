package models

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

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
		"ValorNovedad":    informacionReplica["ValorNovedad"],
	}

	TitanNovedadPost = map[string]interface{}{
		"NumeroContrato": informacionReplica["NumeroContrato"],
		"Vigencia":       informacionReplica["Vigencia"],
	}

	// Elabora la estructura para Titán según el tipo de novedad
	if int(informacionReplica["TipoNovedad"].(float64)) == 216 {
		TitanNovedadPost["Documento"] = informacionReplica["Documento"]
		TitanNovedadPost["FechaInicio"] = FormatFechaReplica(informacionReplica["FechaInicio"].(string), "2006-01-02T15:04:05.000Z")
		TitanNovedadPost["FechaFin"] = FormatFechaReplica(informacionReplica["FechaFin"].(string), "2006-01-02T15:04:05.000Z")
		url = "/novedadCPS/suspender_contrato"
	}
	if int(informacionReplica["TipoNovedad"].(float64)) == 219 {
		TitanNovedadPost["DocumentoActual"] = informacionReplica["DocumentoActual"]
		TitanNovedadPost["DocumentoNuevo"] = informacionReplica["DocumentoNuevo"]
		TitanNovedadPost["FechaInicio"] = FormatFechaReplica(informacionReplica["FechaInicio"].(string), "2006-01-02T15:04:05.000Z")
		TitanNovedadPost["NombreCompleto"] = informacionReplica["NombreCompleto"]
		url = "/novedadCPS/ceder_contrato"
	}
	if int(informacionReplica["TipoNovedad"].(float64)) == 220 {
		TitanNovedadPost["Documento"] = informacionReplica["Documento"]
		TitanNovedadPost["FechaFin"] = FormatFechaReplica(informacionReplica["FechaFin"].(string), "2006-01-02T15:04:05.000Z")
		TitanNovedadPost["Valor"] = informacionReplica["ValorNovedad"]
		TitanNovedadPost["Cdp"] = informacionReplica["NumeroCdp"]
		TitanNovedadPost["Rp"] = 0
		url = "/novedadCPS/otrosi_contrato"
	}
	if int(informacionReplica["TipoNovedad"].(float64)) == 218 {
		TitanNovedadPost["Documento"] = informacionReplica["Documento"]
		TitanNovedadPost["FechaCancelacion"] = FormatFechaReplica(informacionReplica["FechaFin"].(string), "2006-01-02T15:04:05.000Z")
		url = "/novedadCPS/cancelar_contrato"
	}

	// fmt.Println("url: ", url)
	// fmt.Println("ArgoNovedadPost: ", ArgoNovedadPost)
	// fmt.Println("TitanNovedadPost: ", TitanNovedadPost)

	if result, err := PostReplica(url, ArgoNovedadPost, TitanNovedadPost); err == nil {
		return result, nil
	} else {
		outputError = map[string]interface{}{"funcion": "/ReplicafechaAnterior", "err": err, "status": "502"}
		return nil, outputError
	}
	// return nil, nil
}

func Temporizador() {

	dt := time.Now()
	until, _ := time.Parse(time.RFC3339, dt.String()[0:10]+"T20:45:00+00:00")
	// 18000
	tdr := time.Tick(10 * time.Second)
	for horaActual := range tdr {
		log.Printf("Temporizador ejecutándose")
		if dt.After(until) {
			ReplicaFechaPosterior(horaActual)
		}
	}

	// when we want to wait till
	// until, _ := time.Parse(time.RFC3339, "2023-06-22T15:04:05+02:00")

}

func ReplicaFechaPosterior(horaActual time.Time) {

	var novedadesResponse []map[string]interface{}
	var replicaResult map[string]interface{}
	var outputError map[string]interface{}

	codEstado := ""
	var estadoNovedad map[string]interface{}
	error3 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro?query=TipoParametroId.CodigoAbreviacion:ENOV,CodigoAbreviacion:ENTR", &estadoNovedad)

	if error3 == nil {
		if len(estadoNovedad) != 0 {
			inter := estadoNovedad["Data"].([]interface{})
			data := inter[0].(map[string]interface{})
			idEstado, _ := data["Id"].(float64)
			codEstado = strconv.FormatFloat(idEstado, 'f', -1, 64)
		}
	}

	url := "/novedades_poscontractuales?query=Estado:" + codEstado + "&limit=0"

	if err := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+url, &novedadesResponse); err == nil {
		for _, novedadRegistro := range novedadesResponse {
			if replicaResult, outputError = ConsultarTipoNovedad(novedadRegistro); outputError == nil {
				fmt.Println("Replica realizada correctamente (Temporizador)")
				fmt.Println("Registro realizado en la fecha", horaActual)
				fmt.Println(replicaResult)
			} else {
				fmt.Println("Fallo al realizar la réplica (Temporizador)")
				fmt.Println(outputError)
			}
		}
	} else {
		fmt.Println(err)
	}
}

func ConsultarTipoNovedad(novedad map[string]interface{}) (result map[string]interface{}, outputError map[string]interface{}) {

	currentDate := time.Now()
	timeLayout := "2006-01-02"
	fechaReferencia := currentDate.Format(timeLayout)
	// fmt.Println("fechaReferencia: ", fechaReferencia)

	tipoNovedad := int(novedad["TipoNovedad"].(float64))

	var fechasResponse []map[string]interface{}
	var propiedades []map[string]interface{}

	// fmt.Println("Id: ", fmt.Sprintf("%v", novedad["Id"]))
	url := "/fechas?query=IdNovedadesPoscontractuales.Id:" + fmt.Sprintf("%v", novedad["Id"]) + "&limit=0"
	if err := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+url, &fechasResponse); err == nil {
		url = "/propiedad?query=IdNovedadesPoscontractuales.Id:" + fmt.Sprintf("%v", novedad["Id"]) + "&limit=0"
		if err := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+url, &propiedades); err == nil {
			for _, fechaRegistro := range fechasResponse {
				idTipoFecha := fechaRegistro["IdTipoFecha"].(map[string]interface{})
				tipoFecha := int(idTipoFecha["Id"].(float64))
				fechaParse, _ := time.Parse("2006-01-02 15:04:05 +0000 +0000", fmt.Sprint(fechaRegistro["Fecha"]))
				fecha := fechaParse.Format(timeLayout)
				switch tipoNovedad {
				case 1:
					if tipoFecha == 8 {
						// if fecha == fechaReferencia || fechaParse.Before(currentDate) {
						if 1 < 2 {
							return ReplicaSuspension(novedad, propiedades, fechasResponse)
						}
					}
				case 2:
					if tipoFecha == 2 {
						if fecha == fechaReferencia || fechaParse.Before(currentDate) {
							return ReplicaCesion(novedad, propiedades, fechasResponse)
						}
					}
				case 3:
					if tipoFecha == 6 {
						if fecha == fechaReferencia || fechaParse.Before(currentDate) {
							return ReplicaTempReinicio(novedad, propiedades, fechasResponse)
						}
					}
				case 5:
					if tipoFecha == 9 {
						if fecha == fechaReferencia || fechaParse.Before(currentDate) {
							return ReplicaTerminacion(novedad, propiedades, fechasResponse)
						}
					}
				case 6:
					if tipoFecha == 1 {
						if fecha == fechaReferencia || fechaParse.Before(currentDate) {
							return ReplicaAdicionProrroga(novedad, propiedades, fechasResponse)
						}
					}
				case 7:
					if tipoFecha == 4 {
						if fecha == fechaReferencia || fechaParse.Before(currentDate) {
							return ReplicaAdicionProrroga(novedad, propiedades, fechasResponse)
						}
					}
				case 8:
					if tipoFecha == 4 {
						if fecha == fechaReferencia || fechaParse.Before(currentDate) {
							return ReplicaAdicionProrroga(novedad, propiedades, fechasResponse)
						}
					}
				}
			}
		} else {
			outputError = map[string]interface{}{"funcion": "/ConsultarTipoNovedad/GetPropiedades", "err": err.Error()}
			return nil, outputError
		}
	} else {
		outputError = map[string]interface{}{"funcion": "/ConsultarTipoNovedad/GetFechas", "err": err.Error()}
		return nil, outputError
	}
	return nil, nil
}

func ReplicaSuspension(novedad map[string]interface{}, propiedades []map[string]interface{}, fechas []map[string]interface{}) (result map[string]interface{}, outputError map[string]interface{}) {

	numContrato := int(novedad["ContratoId"].(float64))
	vigencia := int(novedad["Vigencia"].(float64))
	var idNovedad = fmt.Sprintf("%v", novedad["Id"])

	var cesionario float64
	var contratistaDoc string
	var periodoSuspension float64
	var informacion_proveedor []map[string]interface{}
	var fechaInicio string
	var fechaFin string
	var url = ""

	if len(propiedades[0]) != 0 {
		for _, propiedad := range propiedades {
			tipopropiedad := propiedad["IdTipoPropiedad"].(map[string]interface{})
			nombrepropiedad := tipopropiedad["Nombre"]
			if nombrepropiedad == "Cesionario" {
				cesionario = propiedad["Propiedad"].(float64)
				url = "/informacion_proveedor?query=Id:" + fmt.Sprintf("%v", propiedad["Propiedad"])
				if error := request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &informacion_proveedor); error == nil {
					if len(informacion_proveedor) > 0 {
						contratistaDoc = informacion_proveedor[0]["NumDocumento"].(string)
					}
				}
			}
			if nombrepropiedad == "PeriodoSuspension" {
				periodoSuspension = propiedad["Propiedad"].(float64)
			}
		}
	}

	for _, fecha := range fechas {
		tipoFecha := fecha["IdTipoFecha"].(map[string]interface{})
		nombreFecha := tipoFecha["Nombre"]
		if nombreFecha == "FechaSuspension" {
			fechaInicio = fecha["Fecha"].(string)
		}
		if nombreFecha == "FechaFinSuspension" {
			fechaFin = fecha["Fecha"].(string)
		}
	}

	ArgoSuspensionPost := make(map[string]interface{})
	ArgoSuspensionPost = map[string]interface{}{
		"NumeroContrato":  fmt.Sprintf("%v", numContrato),
		"Vigencia":        vigencia,
		"FechaRegistro":   time.Now().Format("2006-01-02T15:04:05.000Z"),
		"Contratista":     cesionario,
		"PlazoEjecucion":  periodoSuspension,
		"FechaInicio":     FormatFechaReplica(fechaInicio, "2006-01-02 15:04:05 +0000 +0000"),
		"FechaFin":        FormatFechaReplica(fechaFin, "2006-01-02 15:04:05 +0000 +0000"),
		"UnidadEjecucion": 205,
		"TipoNovedad":     216,
	}

	TitanSuspensionPost := make(map[string]interface{})
	TitanSuspensionPost = map[string]interface{}{
		"Documento":      contratistaDoc,
		"FechaFin":       FormatFechaReplica(fechaFin, "2006-01-02 15:04:05 +0000 +0000"),
		"FechaInicio":    FormatFechaReplica(fechaInicio, "2006-01-02 15:04:05 +0000 +0000"),
		"NumeroContrato": strconv.Itoa(numContrato),
		"Vigencia":       vigencia,
	}

	url = "/novedadCPS/suspender_contrato"
	if result, err := PostReplica(url, ArgoSuspensionPost, TitanSuspensionPost); err == nil {
		resultEstado, errEstado := CambioEstadoReplica(strconv.Itoa(numContrato), 2, idNovedad)
		if errEstado == nil {
			fmt.Println(resultEstado)
			return result, nil
		} else {
			outputError = map[string]interface{}{"funcion": "/ReplicaSuspension1", "err": errEstado}
			return nil, outputError
		}
	} else {
		outputError = map[string]interface{}{"funcion": "/ReplicaSuspension", "err": err}
		return nil, err
	}
}

func ReplicaCesion(novedad map[string]interface{}, propiedades []map[string]interface{}, fechas []map[string]interface{}) (result map[string]interface{}, outputError map[string]interface{}) {

	numContrato := int(novedad["ContratoId"].(float64))
	vigencia := int(novedad["Vigencia"].(float64))

	var cesionario float64
	var cesionarioDoc string
	var nombreCesionario string
	var cedenteDoc string
	var informacion_proveedor []map[string]interface{}
	var fechaInicio string
	var FechaFinEfectiva string
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
					if len(informacion_proveedor) > 0 {
						cesionarioDoc = informacion_proveedor[0]["NumDocumento"].(string)
						nombreCesionario = informacion_proveedor[0]["NomProveedor"].(string)
					}
				}
			}
			if nombrepropiedad == "Cedente" {
				url = "/informacion_proveedor?query=Id:" + fmt.Sprintf("%v", propiedad["Propiedad"])
				if error := request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &informacion_proveedor); error == nil {
					if len(informacion_proveedor) > 0 {
						cedenteDoc = informacion_proveedor[0]["NumDocumento"].(string)
					}
				}
			}
		}
	}

	for _, fecha := range fechas {
		tipoFecha := fecha["IdTipoFecha"].(map[string]interface{})
		nombreFecha := tipoFecha["Nombre"]
		if nombreFecha == "FechaCesion" {
			fechaInicio = fecha["Fecha"].(string)
		}
		if nombreFecha == "FechaFinEfectiva" {
			FechaFinEfectiva = fecha["Fecha"].(string)
		}
	}
	fechaCesion, _ := time.Parse("", fechaInicio)
	fechaFin, _ := time.Parse("", FechaFinEfectiva)
	diasCesion = float64(fechaFin.Sub(fechaCesion))

	ArgoCesionPost := make(map[string]interface{})
	ArgoCesionPost = map[string]interface{}{
		"NumeroContrato":  fmt.Sprintf("%v", numContrato),
		"Vigencia":        vigencia,
		"FechaRegistro":   time.Now().Format("2006-01-02T15:04:05.000Z"),
		"Contratista":     cesionario,
		"PlazoEjecucion":  int(diasCesion),
		"FechaInicio":     FormatFechaReplica(fechaInicio, "2006-01-02 15:04:05 +0000 +0000"),
		"FechaFin":        FormatFechaReplica(FechaFinEfectiva, "2006-01-02 15:04:05 +0000 +0000"),
		"UnidadEjecucion": 205,
		"TipoNovedad":     219,
	}

	TitanCesionPost := make(map[string]interface{})
	TitanCesionPost = map[string]interface{}{
		"DocumentoActual": cedenteDoc,
		"DocumentoNuevo":  cesionarioDoc,
		"FechaInicio":     FormatFechaReplica(fechaInicio, "2006-01-02 15:04:05 +0000 +0000"),
		"NombreCompleto":  nombreCesionario,
		"NumeroContrato":  strconv.Itoa(numContrato),
		"Vigencia":        vigencia,
	}

	url = "/novedadCPS/ceder_contrato"
	if result, err := PostReplica(url, ArgoCesionPost, TitanCesionPost); err == nil {
		return result, nil
	} else {
		outputError = map[string]interface{}{"funcion": "/ReplicaCesion", "err": err}
		return nil, err
	}
}

func ReplicaTempReinicio(novedad map[string]interface{}, propiedades []map[string]interface{}, fechas []map[string]interface{}) (result map[string]interface{}, outputError map[string]interface{}) {

	numContrato := int(novedad["ContratoId"].(float64))
	vigencia := int(novedad["Vigencia"].(float64))

	var contratistaDoc string
	var informacion_proveedor []map[string]interface{}
	var novedadesArgo []map[string]interface{}
	var idNovedadArgo float64
	var fechasuspension string
	var FechaFinSuspension string
	var fechaReinicio string
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
					if len(informacion_proveedor) > 0 {
						contratistaDoc = informacion_proveedor[0]["NumDocumento"].(string)
					}
				}
			}
			if nombrepropiedad == "PeriodoSuspension" {
				periodoSuspension = propiedad["Propiedad"].(float64)
			}
		}
	}

	for _, fecha := range fechas {
		tipoFecha := fecha["IdTipoFecha"].(map[string]interface{})
		nombreFecha := tipoFecha["Nombre"]
		if nombreFecha == "FechaSuspension" {
			fechasuspension = fecha["Fecha"].(string)
		}
		if nombreFecha == "FechaFinSuspension" {
			FechaFinSuspension = fecha["Fecha"].(string)
		}
		if nombreFecha == "FechaReinicio" {
			fechaReinicio = fecha["Fecha"].(string)
		}
	}

	ArgoReinicioPost := make(map[string]interface{})
	ArgoReinicioPost = map[string]interface{}{
		"NumeroContrato":  fmt.Sprintf("%v", numContrato),
		"Vigencia":        vigencia,
		"FechaRegistro":   time.Now().Format("2006-01-02T15:04:05.000Z"),
		"Contratista":     idStr,
		"PlazoEjecucion":  periodoSuspension,
		"FechaInicio":     FormatFechaReplica(fechasuspension, "2006-01-02 15:04:05 +0000 +0000"),
		"FechaFin":        FormatFechaReplica(FechaFinSuspension, "2006-01-02 15:04:05 +0000 +0000"),
		"UnidadEjecucion": 205,
		"TipoNovedad":     216,
	}

	TitanReinicioPost := make(map[string]interface{})
	TitanReinicioPost = map[string]interface{}{
		"Documento":      contratistaDoc,
		"FechaReinicio":  FormatFechaReplica(fechaReinicio, "2006-01-02 15:04:05 +0000 +0000"),
		"NumeroContrato": strconv.Itoa(numContrato),
		"Vigencia":       vigencia,
	}

	url = "/novedad_postcontractual?query=NumeroContrato:" + fmt.Sprintf("%v", numContrato) + ",Vigencia:" + strconv.Itoa(vigencia) + ",TipoNovedad:216&sortby=Id&order=desc&limit=1"

	if err := request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &novedadesArgo); err == nil {
		if len(novedadesArgo) > 0 {
			idNovedadArgo = novedadesArgo[0]["Id"].(float64)
		}
	}

	idNovedadString := strconv.FormatFloat(idNovedadArgo, 'f', -1, 64)

	url = "/novedad_postcontractual/" + idNovedadString

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

func ReplicaTerminacion(novedad map[string]interface{}, propiedades []map[string]interface{}, fechas []map[string]interface{}) (result map[string]interface{}, outputError map[string]interface{}) {

	numContrato := int(novedad["ContratoId"].(float64))
	vigencia := int(novedad["Vigencia"].(float64))

	var contratistaDoc string
	var informacion_proveedor []map[string]interface{}
	var fechaFin string
	var idStr = ""
	var url = ""

	for _, propiedad := range propiedades {
		tipopropiedad := propiedad["IdTipoPropiedad"].(map[string]interface{})
		nombrepropiedad := tipopropiedad["Nombre"]
		if nombrepropiedad == "Cesionario" {
			idStr = fmt.Sprintf("%v", propiedad["Propiedad"])
			url = "/informacion_proveedor?query=Id:" + idStr
			if err := request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &informacion_proveedor); err == nil {
				if len(informacion_proveedor) > 0 {
					contratistaDoc = informacion_proveedor[0]["NumDocumento"].(string)
				}
			}
		}
	}

	for _, fecha := range fechas {
		tipoFecha := fecha["IdTipoFecha"].(map[string]interface{})
		nombreFecha := tipoFecha["Nombre"]
		if nombreFecha == "FechaTerminacionAnticipada" {
			fechaFin = fecha["Fecha"].(string)
		}
	}

	ArgoTerminacionPost := make(map[string]interface{})
	ArgoTerminacionPost = map[string]interface{}{
		"NumeroContrato": fmt.Sprintf("%v", numContrato),
		"Vigencia":       vigencia,
		"FechaRegistro":  time.Now().Format("2006-01-02T15:04:05.000Z"),
		"FechaFin":       FormatFechaReplica(fechaFin, "2006-01-02 15:04:05 +0000 +0000"),
		"TipoNovedad":    218,
	}

	TitanTerminacionPost := make(map[string]interface{})
	TitanTerminacionPost = map[string]interface{}{
		"Documento":        contratistaDoc,
		"FechaCancelacion": FormatFechaReplica(fechaFin, "2006-01-02 15:04:05 +0000 +0000"),
		"NumeroContrato":   strconv.Itoa(numContrato),
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

func ReplicaAdicionProrroga(novedad map[string]interface{}, propiedades []map[string]interface{}, fechas []map[string]interface{}) (result map[string]interface{}, outputError map[string]interface{}) {

	numContrato := int(novedad["ContratoId"].(float64))
	vigencia := int(novedad["Vigencia"].(float64))
	numeroCdp := int(novedad["NumeroCdpId"].(float64))
	vigenciaCdp := int(novedad["VigenciaCdp"].(float64))

	var tipoNovedad int
	if int(novedad["TipoNovedad"].(float64)) == 6 {
		tipoNovedad = 248
	} else if int(novedad["TipoNovedad"].(float64)) == 7 {
		tipoNovedad = 249
	} else if int(novedad["TipoNovedad"].(float64)) == 8 {
		tipoNovedad = 220
	}
	var cesionario float64
	var valoradicion float64 = 0
	var tiempoprorroga float64 = 0
	var contratistaDoc string
	var informacion_proveedor []map[string]interface{}
	var fechaInicio string
	var fechaFin string
	var numeroRp float64
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
			if nombrepropiedad == "Numero_RP" {
				numeroRp = propiedad["Propiedad"].(float64)
			}
		}
	}

	for _, fecha := range fechas {
		tipoFecha := fecha["IdTipoFecha"].(map[string]interface{})
		nombreFecha := tipoFecha["Nombre"]
		if nombreFecha == "FechaProrroga" {
			fechaInicio = fecha["Fecha"].(string)
		}
		if nombreFecha == "FechaFinEfectiva" {
			fechaFin = fecha["Fecha"].(string)
		}
	}

	ArgoOtrosiPost := make(map[string]interface{})
	ArgoOtrosiPost = map[string]interface{}{
		"NumeroContrato":  fmt.Sprintf("%v", numContrato),
		"Vigencia":        vigencia,
		"FechaRegistro":   time.Now().Format("2006-01-02T15:04:05.000Z"),
		"Contratista":     cesionario,
		"PlazoEjecucion":  tiempoprorroga,
		"FechaInicio":     FormatFechaReplica(fechaInicio, "2006-01-02 15:04:05 +0000 +0000"),
		"FechaFin":        FormatFechaReplica(fechaFin, "2006-01-02 15:04:05 +0000 +0000"),
		"NumeroCdp":       numeroCdp,
		"VigenciaCdp":     vigenciaCdp,
		"ValorNovedad":    valoradicion,
		"UnidadEjecucion": 205,
		"TipoNovedad":     tipoNovedad,
	}

	TitanOtrosiPost := make(map[string]interface{})
	TitanOtrosiPost = map[string]interface{}{
		"NumeroContrato": strconv.Itoa(numContrato),
		"Documento":      contratistaDoc,
		"FechaFin":       FormatFechaReplica(fechaFin, "2006-01-02 15:04:05 +0000 +0000"),
		"Cdp":            numeroCdp,
		"Rp":             numeroRp,
		"Valor":          valoradicion,
		"Vigencia":       vigencia,
	}

	url = "/novedadCPS/otrosi_contrato"
	if result, err := PostReplica(url, ArgoOtrosiPost, TitanOtrosiPost); err == nil {
		return result, nil
	} else {
		outputError = map[string]interface{}{"funcion": "/ReplicaAdicionProrroga", "err": err}
		return nil, err
	}
}

func PostReplica(url string, ArgoOtrosiPost map[string]interface{}, TitanOtrosiPost map[string]interface{}) (map[string]interface{}, map[string]interface{}) {
	// fmt.Println("url: ", url)
	// fmt.Println("ArgoPost: ", ArgoOtrosiPost)
	// fmt.Println("TitanPost: ", TitanOtrosiPost)
	var resultPostArgo map[string]interface{}
	var resultPostTitan map[string]interface{}
	var outputError map[string]interface{}
	if err := SendJson(beego.AppConfig.String("AdministrativaAmazonService")+"/novedad_postcontractual", "POST", &resultPostArgo, &ArgoOtrosiPost); err == nil {
		if err := SendJson(beego.AppConfig.String("TitanMidService")+url, "POST", &resultPostTitan, &TitanOtrosiPost); err == nil {
			if len(resultPostTitan) > 0 {
				fmt.Println("Registro en Titan exitoso!")
				return resultPostTitan, nil
			} else {
				outputError = map[string]interface{}{"funcion": "/PostReplica", "err": err}
				return nil, outputError
			}
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

func FormatFechaReplica(fecha string, format string) string {
	// format Formato en el que se recibe la fecha
	var fechaReplica string
	if fechaParse, err := time.Parse(format, fecha); err == nil {
		fechaFormat := fechaParse.Format("2006-01-02T15:04:05.000Z")
		fechaReplica = fechaFormat
	} else {
		fmt.Println(err)
	}
	return fechaReplica
}

// func waitUntil(ctx context.Context, until time.Time) {
// 	timer := time.NewTimer(time.Until(until))
// 	defer timer.Stop()

// 	select {
// 	case <-timer.C:
// 		return
// 	case <-ctx.Done():
// 		return
// 	}
// }

func CambioEstadoReplica(numContrato string, estado int, idNovedad string) (map[string]interface{}, error) {

	var resultContrato []map[string]interface{}
	var resultadoEstadoAdmamazon map[string]interface{}
	var estadoNovedad map[string]interface{}
	error3 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro?query=TipoParametroId.CodigoAbreviacion:ENOV,CodigoAbreviacion:TERM", &estadoNovedad)
	var codEstado string
	if error3 == nil {
		if len(estadoNovedad) != 0 {
			inter := estadoNovedad["Data"].([]interface{})
			data := inter[0].(map[string]interface{})
			idEstado, _ := data["Id"].(float64)
			codEstado = strconv.FormatFloat(idEstado, 'f', -1, 64)
		}
	}

	errContrato := request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+"/contrato_suscrito?query=NumeroContratoSuscrito:"+numContrato, &resultContrato)
	if errContrato == nil {

		var novedad map[string]interface{}
		var resultadoRegistro map[string]interface{}

		result := resultContrato[0]
		numeroContrato := result["NumeroContrato"].(map[string]interface{})
		num_contrato_id := numeroContrato["Id"].(string)
		usuario := result["Usuario"].(string)

		body := make(map[string]interface{})
		body = map[string]interface{}{
			"FechaRegistro":  time.Now().Format("2006-01-02T15:04:05.000Z"),
			"NumeroContrato": num_contrato_id,
			"Usuario":        usuario,
			"Vigencia":       2023,
		}

		body["Estado"] = map[string]interface{}{
			"Id": estado,
		}

		url := beego.AppConfig.String("AdministrativaAmazonService") + "/contrato_estado"
		errEstado := request.SendJson(url, "POST", &resultadoEstadoAdmamazon, &body)
		if errEstado == nil {
			err := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/"+idNovedad, &novedad)
			if err == nil {
				novedad["Estado"] = codEstado
				errRegNovedad := request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/"+idNovedad, "PUT", &resultadoRegistro, novedad)
				if errRegNovedad == nil {
					fmt.Println("Estado de novedad actualizado!!")
				}
			}
			return resultadoEstadoAdmamazon, nil
		} else {
			return nil, errEstado
		}
	} else {
		return nil, errContrato
	}
}
