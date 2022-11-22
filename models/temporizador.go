package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func ConsultarFechaNovedad() (novedad map[string]interface{}) {

	currentDate := time.Now()
	lastMonth := currentDate.AddDate(0, -1, 0)
	goneDayMonth := lastMonth.Day()
	firstDay := lastMonth.AddDate(0, 0, -goneDayMonth+1)
	timeLayout := "2006-01-02"
	fechaReferencia := currentDate.Format(timeLayout)

	// tdr := time.Tick(86400 * time.Second)

	// for horaActual := range tdr {
	// 	fmt.Println("La hora es", horaActual)
	// }

	var fechasResponse []map[string]interface{}

	url := "/fechas?limit=-1"
	if err := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+url, &fechasResponse); err == nil {
		for _, fechaRegistro := range fechasResponse {
			// fmt.Println("FechaRegistro:", fechaRegistro["Fecha"])
			fechaParse, _ := time.Parse("2006-01-02 15:04:05 +0000 +0000", fmt.Sprint(fechaRegistro["Fecha"]))
			// fmt.Println("FechaParse:", fechaParse)
			if fechaParse.After(firstDay) {
				fecha := fechaParse.Format(timeLayout)
				if fecha == fechaReferencia {
					novedad := fechaRegistro["IdNovedadesPoscontractuales"].(map[string]interface{})
					ConsultarTipoNovedad(novedad)
				}
			}
		}
		fmt.Println("Sale del for")
	} else {
		fmt.Println(err)
	}
	return nil
}

func ConsultarTipoNovedad(novedad map[string]interface{}) (structura map[string]interface{}) {

	tipoNovedad := int(novedad["TipoNovedad"].(float64))
	fmt.Println("TipoNovedad:", tipoNovedad)

	var propiedades []map[string]interface{}
	var fechas []map[string]interface{}

	// TitanCesionPost := make(map[string]interface{})

	url := "/propiedad?query=IdNovedadesPoscontractuales.Id:" + fmt.Sprintf("%v", novedad["Id"]) + "&limit=0"
	if error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+url, &propiedades); error == nil {
		fmt.Println("1")
		url = "/fechas?query=IdNovedadesPoscontractuales.Id:" + fmt.Sprintf("%v", novedad["Id"]) + "&limit=0"
		if error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+url, &fechas); error == nil {
			fmt.Println("2")

			switch tipoNovedad {
			case 1:
				ReplicaSuspension(novedad, propiedades)
			case 2:
				ReplicaCesion(novedad, propiedades, fechas)
			case 3:
			// 	ReplicaReinicio(novedad, propiedades, fechas)
			case 6:
				ReplicaAdicionProrroga(novedad, propiedades)
			case 7:
				ReplicaAdicionProrroga(novedad, propiedades)
			case 8:
				ReplicaAdicionProrroga(novedad, propiedades)
			}

			// TitanCesionPost["NovedadPoscontractual"] = map[string]interface{}{
			// 	"DocumentoActual": 0,
			// 	"DocumentoNuevo":  2022,
			// 	"FechaInicio":     tdr,
			// 	"NombreCompleto":  0,
			// 	"NumeroContrato":  0,
			// 	"Vigencia":        tdr,
			// }

			// var result map[string]interface{}

			// url = ""
			// if err := SendJson(beego.AppConfig.String("AdministrativaAmazonService")+url, "POST", &result, &ArgoCesionPost); err == nil {
			// 	url = ""
			// 	if err := SendJson(beego.AppConfig.String("AdministrativaAmazonService")+url, "POST", &result, &ArgoCesionPost); err == nil {
			// 		fmt.Println("Registro en Titan exitoso!")
			// 	}
			// }
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
			var dias float64 = periodoSuspension
			meses := dias / 30
			mesEntero := int(meses)
			decimal := meses - float64(mesEntero)
			numDias := decimal * 30
			fechaFin = fechaInicio.AddDate(0, mesEntero, int(numDias))
			fmt.Println("Fecha Inicio Novedad:", fechaInicio)
			fmt.Println("Fecha Fin Novedad:", fechaFin)
		}
	}

	ArgoSuspensionPost := make(map[string]interface{})
	ArgoSuspensionPost = map[string]interface{}{
		"NumeroContrato":  numContrato,
		"Vigencia":        vigencia,
		"FechaRegistro":   time.Now(),
		"Contratista":     cesionario,
		"PlazoEjecucion":  periodoSuspension,
		"FechaInicio":     fechaInicio,
		"FechaFin":        fechaFin,
		"UnidadEjecucion": 205,
		"TipoNovedad":     216,
	}
	fmt.Println("ArgoOtrosiPost:", ArgoSuspensionPost)

	TitanSuspensionPost := make(map[string]interface{})
	TitanSuspensionPost["NovedadPoscontractual"] = map[string]interface{}{
		"Documento":      contratistaDoc,
		"FechaFin":       fechaFin,
		"FechaInicio":    fechaInicio,
		"NumeroContrato": strconv.Itoa(numContrato),
		"Vigencia":       vigencia,
	}
	fmt.Println("TitanOtrosiPost:", TitanSuspensionPost)

}

func ReplicaCesion(novedad map[string]interface{}, propiedades []map[string]interface{}, fechas []map[string]interface{}) {

	numContrato := int(novedad["ContratoId"].(float64))
	vigencia := int(novedad["Vigencia"].(float64))

	var cesionario float64
	var contratistaDoc string
	var cedente float64
	var contratoSuscrito []map[string]interface{}
	var actaInicio []map[string]interface{}
	var informacion_proveedor []map[string]interface{}
	var fechaInicio time.Time
	var fechaFin time.Time
	var url = ""

	for _, fecha := range fechas {
		tipoFecha := fecha["IdTipoFecha"].(map[string]interface{})
		nombreFecha := tipoFecha["Nombre"]
		if nombreFecha == "FechaCesion" {
			fechaInicio = fecha["Fecha"].(time.Time)
		}
	}

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
			if nombrepropiedad == "Cedente" {
				cedente = propiedad["Propiedad"].(float64)
			}
		}
	}

	url = "/contrato_suscrito?query=NumeroContratoSuscrito:" + fmt.Sprintf("%v", novedad["ContratoId"])
	if error := request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &contratoSuscrito); error == nil {
		numContratoActa := contratoSuscrito[len(contratoSuscrito)-1]["NumeroContrato"].(map[string]interface{})
		url = "/acta_inicio?query=NumeroContrato:" + fmt.Sprintf("%v", numContratoActa["Id"])
		if error = request.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+url, &actaInicio); error == nil {
			fechaFin, _ = time.Parse("2006-01-02T00:00:00Z", fmt.Sprint(actaInicio[0]["FechaFin"]))
			fmt.Println("Fecha Inicio Novedad:", fechaInicio)
			fmt.Println("Fecha Fin Novedad:", fechaFin)
		}
	}

	ArgoCesionPost := make(map[string]interface{})
	ArgoCesionPost = map[string]interface{}{
		"NumeroContrato":  numContrato,
		"Vigencia":        vigencia,
		"FechaRegistro":   time.Now(),
		"Contratista":     cesionario,
		"PlazoEjecucion":  0,
		"FechaInicio":     fechaInicio,
		"FechaFin":        fechaFin,
		"UnidadEjecucion": 205,
		"TipoNovedad":     219,
	}
	fmt.Println("ArgoCesionPost", ArgoCesionPost)

	TitanSuspensionPost := make(map[string]interface{})
	TitanSuspensionPost["NovedadPoscontractual"] = map[string]interface{}{
		"DocumentoActual": cedente,
		"DocumentoNuevo":  contratistaDoc,
		"FechaInicio":     fechaInicio,
		"NombreCompleto":  "",
		"NumeroContrato":  strconv.Itoa(numContrato),
		"Vigencia":        vigencia,
	}
	fmt.Println("TitanOtrosiPost:", TitanSuspensionPost)
}

// func ReplicaReinicio(novedad map[string]interface{}, propiedades []map[string]interface{}, fechas []map[string]interface{}) {

// 	numContrato, _ := strconv.ParseInt(novedad["ContratoId"].(string), 10, 32)
// 	vigencia, _ := strconv.ParseInt(novedad["Vigencia"].(string), 10, 32)

// 	ArgoCesionPost := make(map[string]interface{})

// }

func ReplicaAdicionProrroga(novedad map[string]interface{}, propiedades []map[string]interface{}) {

	numContrato := int(novedad["ContratoId"].(float64))
	vigencia := int(novedad["Vigencia"].(float64))
	numeroCdp := int(novedad["NumeroCdpId"].(float64))
	vigenciaCdp := int(novedad["VigenciaCdp"].(float64))

	// var contratoGeneral map[string]interface{}
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
			var dias float64 = tiempoprorroga
			meses := dias / 30
			mesEntero := int(meses)
			decimal := meses - float64(mesEntero)
			numDias := decimal * 30
			fechaFin = fechaInicio.AddDate(0, mesEntero, int(numDias))
			fmt.Println("Fecha Inicio Novedad:", fechaInicio)
			fmt.Println("Fecha Fin Novedad:", fechaFin)
		}
	}

	ArgoOtrosiPost := make(map[string]interface{})
	ArgoOtrosiPost = map[string]interface{}{
		"NumeroContrato":  numContrato,
		"Vigencia":        vigencia,
		"FechaRegistro":   time.Now(),
		"Contratista":     cesionario,
		"PlazoEjecucion":  tiempoprorroga,
		"FechaInicio":     fechaInicio,
		"FechaFin":        fechaFin,
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
		"FechaFin":       fechaFin,
		"NumeroContrato": numContrato,
		"Vigencia":       vigencia,
	}
	fmt.Println("TitanOtrosiPost:", TitanOtrosiPost)

	// var resultPost map[string]interface{}

	// url = "/novedad_postcontractual"
	// if err := SendJson(beego.AppConfig.String("AdministrativaAmazonService")+url, "POST", &resultPost, &ArgoCesionPost); err == nil {
	// 	url = "/novedadCPS/otrosi_contrato"
	// 	if err := SendJson(beego.AppConfig.String("TitanMidService")+url, "POST", &resultPost, &TitanCesionPost); err == nil {
	// 		fmt.Println("Registro en Titan exitoso!")
	// 	}
	// }

}
