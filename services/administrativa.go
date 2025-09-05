package services

import (
	"fmt"
	neturl "net/url"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/helpers"
	"github.com/udistrital/utils_oas/request"
)

func getNumeroContratoID(numeroSuscrito int, vigencia int) (int, error) {
	base := beego.AppConfig.String("AdministrativaAmazonService")
	if base == "" {
		return 0, fmt.Errorf("AdministrativaAmazonService no configurado")
	}

	var r1 []map[string]interface{}
	url1 := fmt.Sprintf("%s/contrato_suscrito?query=NumeroContratoSuscrito:%d,Vigencia:%d&limit=1", base, numeroSuscrito, vigencia)
	if err := request.GetJson(url1, &r1); err == nil && len(r1) > 0 {
		if nc, ok := r1[0]["NumeroContrato"].(map[string]interface{}); ok {
			switch v := nc["Id"].(type) {
			case float64:
				return int(v), nil
			case string:
				if iv, err := strconv.Atoi(v); err == nil {
					return iv, nil
				}
			}
		}
	}

	var r2 []map[string]interface{}
	url2 := fmt.Sprintf("%s/contrato_suscrito?query=numero_contrato_suscrito:%d,vigencia:%d&limit=1", base, numeroSuscrito, vigencia)
	if err := request.GetJson(url2, &r2); err == nil && len(r2) > 0 {
		if nc, ok := r2[0]["numero_contrato"].(map[string]interface{}); ok {
			switch v := nc["Id"].(type) {
			case float64:
				return int(v), nil
			case string:
				if iv, err := strconv.Atoi(v); err == nil {
					return iv, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("NumeroContrato.Id no encontrado para NumeroContratoSuscrito:%d Vigencia:%d", numeroSuscrito, vigencia)
}

func getNumeroContratoSuscritoStr(numeroSuscrito int, vigencia int) (string, error) {
	base := beego.AppConfig.String("AdministrativaAmazonService")
	if base == "" {
		return "", fmt.Errorf("AdministrativaAmazonService no configurado")
	}

	{ // mayúsculas
		var r1 []map[string]interface{}
		url1 := fmt.Sprintf("%s/contrato_suscrito?query=NumeroContratoSuscrito:%d,Vigencia:%d&limit=1", base, numeroSuscrito, vigencia)
		if err := request.GetJson(url1, &r1); err == nil && len(r1) > 0 {
			if v, ok := r1[0]["NumeroContratoSuscrito"]; ok {
				if s, ok2 := v.(string); ok2 && strings.TrimSpace(s) != "" {
					return s, nil
				}
			}
		}
	}
	{ // minúsculas
		var r2 []map[string]interface{}
		url2 := fmt.Sprintf("%s/contrato_suscrito?query=numero_contrato_suscrito:%d,vigencia:%d&limit=1", base, numeroSuscrito, vigencia)
		if err := request.GetJson(url2, &r2); err == nil && len(r2) > 0 {
			if v, ok := r2[0]["numero_contrato_suscrito"]; ok {
				if s, ok2 := v.(string); ok2 && strings.TrimSpace(s) != "" {
					return s, nil
				}
			}
		}
	}
	return "", fmt.Errorf("NumeroContratoSuscrito string no encontrado")
}

func getTopAndPrevNovedadAmazon(contratoId, vigencia int) (map[string]interface{}, map[string]interface{}, error) {
	base := beego.AppConfig.String("AdministrativaAmazonService")
	if base == "" {
		return nil, nil, fmt.Errorf("AdministrativaAmazonService no configurado")
	}
	var arr []map[string]interface{}
	url := fmt.Sprintf("%s/novedad_postcontractual/?query=NumeroContrato:%d,Vigencia:%d&sortby=FechaInicio&order=desc&limit=2",
		base, contratoId, vigencia)
	if err := request.GetJson(url, &arr); err != nil {
		return nil, nil, err
	}
	if len(arr) == 0 {
		return nil, nil, nil
	}
	if len(arr) == 1 {
		return arr[0], nil, nil
	}
	return arr[0], arr[1], nil
}

func getFechasActaInicioPorIdVigencia(numeroContratoID int, vigencia int) (time.Time, time.Time, error) {
	base := beego.AppConfig.String("AdministrativaAmazonService")
	if base == "" {
		return time.Time{}, time.Time{}, fmt.Errorf("AdministrativaAmazonService no configurado")
	}

	q := neturl.QueryEscape(fmt.Sprintf("NumeroContrato:%d,Vigencia:%d", numeroContratoID, vigencia))
	url := fmt.Sprintf("%s/acta_inicio?query=%s&sortby=FechaInicio&order=desc&limit=1", base, q)
	list, err := helpers.GetDataListFromURL(url)
	if err != nil || len(list) == 0 {
		q2 := neturl.QueryEscape(fmt.Sprintf("numero_contrato:%d,vigencia:%d", numeroContratoID, vigencia))
		url2 := fmt.Sprintf("%s/acta_inicio?query=%s&sortby=fecha_inicio&order=desc&limit=1", base, q2)
		list, err = helpers.GetDataListFromURL(url2)
		if err != nil || len(list) == 0 {
			return time.Time{}, time.Time{}, fmt.Errorf("acta_inicio sin resultados para NumeroContrato(Id)=%d Vigencia=%d", numeroContratoID, vigencia)
		}
	}

	m := list[0]
	fi := helpers.ParseTimeAny(m["FechaInicio"])
	ff := helpers.ParseTimeAny(m["FechaFin"])
	if fi.IsZero() {
		fi = helpers.ParseTimeAny(m["fecha_inicio"])
	}
	if ff.IsZero() {
		ff = helpers.ParseTimeAny(m["fecha_fin"])
	}
	if fi.IsZero() && ff.IsZero() {
		return time.Time{}, time.Time{}, fmt.Errorf("acta_inicio sin fechas válidas para NumeroContrato(Id)=%d Vigencia=%d", numeroContratoID, vigencia)
	}
	if ff.IsZero() {
		ff = fi
	}
	return fi, ff, nil
}

func postContratoEstado(estadoID int, numeroContratoID int, vigencia int, usuario string, fechaISO string) (interface{}, error) {
	base := beego.AppConfig.String("AdministrativaAmazonService")
	if base == "" {
		return nil, fmt.Errorf("AdministrativaAmazonService no configurado")
	}

	payload := map[string]interface{}{
		"Estado":         map[string]interface{}{"Id": estadoID},
		"FechaRegistro":  fechaISO,
		"NumeroContrato": strconv.Itoa(numeroContratoID),
		"Usuario":        usuario,
		"Vigencia":       vigencia,
	}

	var resp interface{}
	if err := request.SendJson(base+"/contrato_estado", "POST", &resp, payload); err != nil {
		return nil, err
	}
	return resp, nil
}
