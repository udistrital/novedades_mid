package services

import (
	"fmt"
	neturl "net/url"
	"strconv"
	"time"

	"github.com/udistrital/novedades_mid/helpers"
	"github.com/udistrital/utils_oas/request"
)

func queryTitanContratoRowsPorNumero(numeroContrato string, vigencia int) ([]map[string]interface{}, error) {
	base := titanCrudBase()
	if base == "" {
		return nil, fmt.Errorf("TitanCrudService no configurado")
	}
	q := neturl.QueryEscape(fmt.Sprintf("NumeroContrato:%s,Vigencia:%d", numeroContrato, vigencia))
	urlGet := fmt.Sprintf("%s/contrato?limit=-1&query=%s", base, q)
	return helpers.GetDataListFromURL(urlGet)
}

func listTitanContratoRowsByAny(contratoIdNovedad, numeroContratoIdArgo, vigencia int) (rows []map[string]interface{}, usado string, tried []string, err error) {
	tried = []string{}

	if sus, e := getNumeroContratoSuscritoStr(contratoIdNovedad, vigencia); e == nil && sus != "" {
		tried = append(tried, sus)
		if r, e2 := queryTitanContratoRowsPorNumero(sus, vigencia); e2 == nil && len(r) > 0 {
			return r, sus, tried, nil
		}
	}

	ci := strconv.Itoa(contratoIdNovedad)
	tried = append(tried, ci)
	if r, e := queryTitanContratoRowsPorNumero(ci, vigencia); e == nil && len(r) > 0 {
		return r, ci, tried, nil
	}

	argo := strconv.Itoa(numeroContratoIdArgo)
	tried = append(tried, argo)
	if r, e := queryTitanContratoRowsPorNumero(argo, vigencia); e == nil && len(r) > 0 {
		return r, argo, tried, nil
	}
	return nil, "", tried, fmt.Errorf("No se hallaron filas en Titan para candidatos %v", tried)
}

func putContratoRow(row map[string]interface{}) (map[string]interface{}, error) {
	base := titanCrudBase()
	if base == "" {
		return nil, fmt.Errorf("TitanCrudService no configurado")
	}
	id := helpers.GetRowId(row)
	if id == 0 {
		return nil, fmt.Errorf("Contrato.Id inválido")
	}
	var resp map[string]interface{}
	if err := request.SendJson(fmt.Sprintf("%s/contrato/%d", base, id), "PUT", &resp, row); err != nil {
		return nil, err
	}
	return resp, nil
}

func getTitanContratoById(id int) (map[string]interface{}, error) {
	base := titanCrudBase()
	if base == "" {
		return nil, fmt.Errorf("TitanCrudService no configurado")
	}
	url := fmt.Sprintf("%s/contrato/%d", base, id)
	obj, err := helpers.GetObjectFromURL(url)
	if err != nil || obj == nil {
		return nil, fmt.Errorf("No se encontró contrato Titan %d: %v", id, err)
	}
	return obj, nil
}

func deleteTitanContratoById(id int) error {
	base := titanCrudBase()
	if base == "" {
		return fmt.Errorf("TitanCrudService no configurado")
	}
	var resp map[string]interface{}
	return request.SendJson(fmt.Sprintf("%s/contrato/%d", base, id), "DELETE", &resp, nil)
}

func getContratoPreliqRowsByMonth(contratoId, ano, mes int) ([]map[string]interface{}, error) {
	base := titanCrudBase()
	if base == "" {
		return nil, fmt.Errorf("TitanCrudService no configurado")
	}
	q := neturl.QueryEscape(fmt.Sprintf("ContratoId.Id:%d,PreliquidacionId.Ano:%d,PreliquidacionId.Mes:%d,Activo:true", contratoId, ano, mes))
	urlGet := fmt.Sprintf("%s/contrato_preliquidacion?limit=-1&query=%s", base, q)
	return helpers.GetDataListFromURL(urlGet)
}

func deleteCPAndDetailsForMonth(ano, mes int, contratoIds []int) []map[string]interface{} {
	base := titanCrudBase()
	var out []map[string]interface{}
	for _, contratoId := range contratoIds {
		cps, err := getContratoPreliqRowsByMonth(contratoId, ano, mes)
		if err != nil {
			out = append(out, map[string]interface{}{"contrato_id": contratoId, "ok": false, "error": err.Error()})
			continue
		}
		if len(cps) == 0 {
			out = append(out, map[string]interface{}{"contrato_id": contratoId, "cp_id": 0, "ok": true})
			continue
		}
		for _, cp := range cps {
			cpId := helpers.GetRowId(cp)
			// borrar detalles
			qDet := neturl.QueryEscape(fmt.Sprintf("ContratoPreliquidacionId.Id:%d,Activo:true", cpId))
			urlDet := fmt.Sprintf("%s/detalle_preliquidacion?limit=-1&query=%s", base, qDet)
			if dets, err := helpers.GetDataListFromURL(urlDet); err == nil {
				for _, d := range dets {
					did := helpers.GetRowId(d)
					if did != 0 {
						var delResp map[string]interface{}
						_ = request.SendJson(fmt.Sprintf("%s/detalle_preliquidacion/%d", base, did), "DELETE", &delResp, nil)
					}
				}
			}
			// borrar cabecera
			var delCP map[string]interface{}
			if err := request.SendJson(fmt.Sprintf("%s/contrato_preliquidacion/%d", base, cpId), "DELETE", &delCP, nil); err != nil {
				out = append(out, map[string]interface{}{"contrato_id": contratoId, "cp_id": cpId, "ok": false, "error": err.Error()})
			} else {
				out = append(out, map[string]interface{}{"contrato_id": contratoId, "cp_id": cpId, "ok": true})
			}
		}
	}
	return out
}

func deleteAllCPAndDetailsForContrato(ano int, contratoId int) []map[string]interface{} {
	var resumen []map[string]interface{}
	for m := 1; m <= 12; m++ {
		part := deleteCPAndDetailsForMonth(ano, m, []int{contratoId})
		resumen = append(resumen, part...)
	}
	return resumen
}

func deleteCPAndDetailsForContratoRange(contratoId int, fi, ff time.Time) []map[string]interface{} {
	var out []map[string]interface{}
	if fi.IsZero() || ff.IsZero() {
		return out
	}
	if ff.Before(fi) {
		fi, ff = ff, fi
	}
	start := time.Date(fi.Year(), fi.Month(), 1, 0, 0, 0, 0, time.UTC)
	last := time.Date(ff.Year(), ff.Month(), 1, 0, 0, 0, 0, time.UTC)
	for t := start; !t.After(last); t = t.AddDate(0, 1, 0) {
		part := deleteCPAndDetailsForMonth(t.Year(), int(t.Month()), []int{contratoId})
		out = append(out, part...)
	}
	return out
}
