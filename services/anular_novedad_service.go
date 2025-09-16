package services

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/helpers"
	"github.com/udistrital/novedades_mid/models"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/requestresponse"
)

func AnularNovedadYRevertirEstado(id string, usuario string) requestresponse.APIResponse {
	var novedades []models.Novedad
	urlGet := beego.AppConfig.String("NovedadesCrudService") + "/novedades_poscontractuales/?query=id:" + id + "&limit=0"
	if err := request.GetJson(urlGet, &novedades); err != nil || len(novedades) == 0 {
		return helpers.ErrEmiter(fmt.Errorf("error obteniendo novedad %s", id))
	}
	n := novedades[0]

	numeroContratoID, errNC := getNumeroContratoID(n.ContratoId, n.Vigencia)
	if errNC != nil || numeroContratoID == 0 {
		return helpers.ErrEmiter(fmt.Errorf("no se pudo resolver NumeroContrato.Id: %v", errNC))
	}
	numStr, _ := getNumeroContratoSuscritoStr(n.ContratoId, n.Vigencia)
	if strings.TrimSpace(numStr) == "" {
		numStr = strconv.Itoa(n.ContratoId)
	}

	amazonList, errAL := fetchAmazonNovedades(numStr, n.Vigencia)
	if errAL != nil {
		return helpers.ErrEmiter(fmt.Errorf("error consultando novedades Amazon: %v", errAL))
	}

	var bodyIni, bodyFin time.Time
	var idsToDelete []int
	var keepID int

	if len(amazonList) <= 1 {
		if aiIni, aiFin, err := getFechasActaInicioPorIdVigencia(numeroContratoID, n.Vigencia); err == nil {
			bodyIni, bodyFin = aiIni, aiFin
		}
		if len(amazonList) == 1 {
			if idAmazon := helpers.GetRowId(amazonList[0]); idAmazon > 0 {
				idsToDelete = append(idsToDelete, idAmazon)
			}
		}
	} else {
		adpro220 := findLatestTipo(amazonList, 220)
		if adpro220 == nil {
			raw, _ := json.Marshal(map[string]interface{}{
				"error": "No se puede anular: existen múltiples novedades en Amazon pero ninguna es de tipo 220.",
			})
			return requestresponse.APIResponseDTO(false, 400, json.RawMessage(raw))
		}
		bodyIni, bodyFin = fechasDeAmazon(adpro220)
		keepID = helpers.GetRowId(adpro220)
		for _, rec := range amazonList {
			idAmazon := helpers.GetRowId(rec)
			if idAmazon != 0 && idAmazon != keepID {
				idsToDelete = append(idsToDelete, idAmazon)
			}
		}
	}

	if bodyIni.IsZero() || bodyFin.IsZero() {
		if aiIni, aiFin, err := getFechasActaInicioPorIdVigencia(numeroContratoID, n.Vigencia); err == nil {
			bodyIni, bodyFin = aiIni, aiFin
		}
	}

	fechaISO := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	cambiosEstado := map[string]interface{}{"estado_10": nil, "estado_4": nil}
	resp10, err10 := postContratoEstado(10, numeroContratoID, n.Vigencia, usuario, fechaISO)
	if err10 != nil {
		cambiosEstado["estado_10"] = map[string]interface{}{"error": err10.Error()}
		raw, _ := json.Marshal(map[string]interface{}{"cambios_estado": cambiosEstado})
		return requestresponse.APIResponseDTO(false, 400, json.RawMessage(raw))
	}
	cambiosEstado["estado_10"] = resp10
	resp4, err4 := postContratoEstado(4, numeroContratoID, n.Vigencia, usuario, fechaISO)
	if err4 != nil {
		cambiosEstado["estado_4"] = map[string]interface{}{"error": err4.Error()}
		raw, _ := json.Marshal(map[string]interface{}{"cambios_estado": cambiosEstado})
		return requestresponse.APIResponseDTO(false, 400, json.RawMessage(raw))
	}
	cambiosEstado["estado_4"] = resp4

	n.Activo = false
	now := time.Now().Format("2006-01-02 15:04:05")
	n.FechaModificacion = now
	n.FechaCreacion = now
	urlPutN := beego.AppConfig.String("NovedadesCrudService") + "/novedades_poscontractuales/" + id
	var respPUTNov map[string]interface{}
	if err := request.SendJson(urlPutN, "PUT", &respPUTNov, n); err != nil {
		raw, _ := json.Marshal(map[string]interface{}{
			"cambios_estado":  cambiosEstado,
			"novedad_anulada": map[string]interface{}{"error": err.Error()},
		})
		return requestresponse.APIResponseDTO(false, 400, json.RawMessage(raw))
	}

	var amazonEliminadas []map[string]interface{}
	if len(idsToDelete) > 0 {
		for _, delID := range idsToDelete {
			urlDelete := fmt.Sprintf("%s/novedad_postcontractual/%d", beego.AppConfig.String("AdministrativaAmazonService"), delID)
			var deleteResp interface{}
			if err := request.SendJson(urlDelete, "DELETE", &deleteResp, nil); err == nil {
				amazonEliminadas = append(amazonEliminadas, map[string]interface{}{"id": delID, "ok": true})
			} else {
				amazonEliminadas = append(amazonEliminadas, map[string]interface{}{"id": delID, "ok": false, "error": err.Error()})
			}
		}
	}

	rows, usado, tried, errRows := listTitanContratoRowsByAny(n.ContratoId, numeroContratoID, n.Vigencia)
	if errRows != nil || len(rows) == 0 {
		raw, _ := json.Marshal(map[string]interface{}{
			"cambios_estado":  cambiosEstado,
			"novedad_anulada": respPUTNov,
			"error":           "No se hallaron filas de contrato en Titan para NumeroContrato/Vigencia",
			"candidatos":      tried,
		})
		return requestresponse.APIResponseDTO(false, 400, json.RawMessage(raw))
	}
	_ = usado

	var initial map[string]interface{}
	if helpers.IsAdicionOProrroga(n.TipoNovedad) {
		initial = helpers.PickInitialRow(rows)
	} else {
		initial = helpers.PickRowPreferExtraRP(rows)
	}
	if initial == nil {
		initial = helpers.PickInitialRow(rows)
	}
	initialId := helpers.GetRowId(initial)

	helpers.AjustarValorContratoParaAnulacion(initial, rows, n.TipoNovedad)

	if bodyIni.IsZero() || bodyFin.IsZero() {
		fi0, ff0 := helpers.FechasDeRow(initial)
		if bodyIni.IsZero() {
			bodyIni = fi0
		}
		if bodyFin.IsZero() {
			bodyFin = ff0
		}
	}
	if bodyIni.IsZero() || bodyFin.IsZero() {
		raw, _ := json.Marshal(map[string]interface{}{
			"cambios_estado":  cambiosEstado,
			"novedad_anulada": respPUTNov,
			"error":           "No fue posible determinar fechas para regenerar preliquidación.",
		})
		return requestresponse.APIResponseDTO(false, 400, json.RawMessage(raw))
	}

	var updates []map[string]interface{}
	var contratosEliminados []map[string]interface{}
	var contratosExtraCPEliminados []map[string]interface{}

	if _, err := putContratoRow(helpers.UpdateRowFechasActivo(initial, bodyIni, bodyFin, true)); err == nil {
		updates = append(updates, map[string]interface{}{"id": initialId, "ok": true})
	} else {
		updates = append(updates, map[string]interface{}{"id": initialId, "ok": false, "error": err.Error()})
	}

	for _, r := range rows {
		rid := helpers.GetRowId(r)
		if rid == initialId {
			continue
		}
		if n.TipoNovedad == 8 || n.TipoNovedad == 1 || n.TipoNovedad == 2 {
			fiExtra, ffExtra := helpers.FechasDeRow(r)
			if !fiExtra.IsZero() || !ffExtra.IsZero() {
				contratosExtraCPEliminados = append(contratosExtraCPEliminados, deleteCPAndDetailsForContratoRange(rid, fiExtra, ffExtra)...)
			}
			contratosExtraCPEliminados = append(contratosExtraCPEliminados, deleteAllCPAndDetailsForContrato(n.Vigencia, rid)...)
			if err := deleteTitanContratoById(rid); err != nil {
				r["Activo"] = false
				r["FechaModificacion"] = time.Now().Format("2006-01-02 15:04:05")
				if _, err2 := putContratoRow(r); err2 != nil {
					contratosEliminados = append(contratosEliminados, map[string]interface{}{"id": rid, "ok": false, "error": "DELETE:" + err.Error() + " | PUT Activo=false:" + err2.Error()})
				} else {
					contratosEliminados = append(contratosEliminados, map[string]interface{}{"id": rid, "ok": true, "soft_deleted": true})
				}
			} else {
				if obj, errGet := getTitanContratoById(rid); errGet == nil && helpers.GetBool(obj, "Activo") {
					obj["Activo"] = false
					obj["FechaModificacion"] = time.Now().Format("2006-01-02 15:04:05")
					_, _ = putContratoRow(obj)
					contratosEliminados = append(contratosEliminados, map[string]interface{}{"id": rid, "ok": true, "soft_deleted": true})
				} else {
					contratosEliminados = append(contratosEliminados, map[string]interface{}{"id": rid, "ok": true})
				}
			}
		} else {
			if helpers.GetBool(r, "Activo") {
				r["Activo"] = false
				r["FechaModificacion"] = time.Now().Format("2006-01-02 15:04:05")
				if _, err := putContratoRow(r); err != nil {
					updates = append(updates, map[string]interface{}{"id": rid, "ok": false, "error": err.Error()})
				} else {
					updates = append(updates, map[string]interface{}{"id": rid, "ok": true})
				}
			}
		}
	}

	contratoFull, errK := getTitanContratoById(initialId)
	if errK != nil {
		raw, _ := json.Marshal(map[string]interface{}{
			"cambios_estado":  cambiosEstado,
			"novedad_anulada": respPUTNov,
			"error":           fmt.Sprintf("No se pudo recargar el contrato Titan %d: %v", initialId, errK),
		})
		return requestresponse.APIResponseDTO(false, 400, json.RawMessage(raw))
	}

	var contratoIds []int
	for _, r := range rows {
		if idc := helpers.GetRowId(r); idc != 0 {
			contratoIds = append(contratoIds, idc)
		}
	}

	var elimStart, elimEnd time.Time
	for _, delID := range idsToDelete {
		if rec := findByID(amazonList, delID); rec != nil {
			fi, ff := fechasDeAmazon(rec)
			if !fi.IsZero() && (elimStart.IsZero() || fi.Before(elimStart)) {
				elimStart = fi
			}
			if !ff.IsZero() && (elimEnd.IsZero() || ff.After(elimEnd)) {
				elimEnd = ff
			}
		}
	}
	mesesAnulados := helpers.MesesEntreSafe(elimStart, elimEnd)
	mesesRestituir := helpers.MesesEntreSafe(bodyIni, bodyFin)

	replicaMeses := ReplicafechaAnterior(contratoFull, mesesAnulados, mesesRestituir, bodyIni, bodyFin, contratoIds)

	combined := map[string]interface{}{
		"novedad_anulada":   respPUTNov,
		"amazon_eliminadas": amazonEliminadas,
		"cambios_estado":    cambiosEstado,
		"titan": map[string]interface{}{
			"numero_contrato_utilizado":     numStr,
			"contrato_inicial_id":           initialId,
			"contrato_updates":              updates,
			"contratos_eliminados":          contratosEliminados,
			"contratos_extra_cp_eliminados": contratosExtraCPEliminados,
			"replica": map[string]interface{}{
				"funcion":            "/ReplicafechaAnterior",
				"periodo_restituido": []string{bodyIni.Format(time.RFC3339), bodyFin.Format(time.RFC3339)},
				"meses_eliminados":   mesesAnulados,
				"meses_recreados":    mesesRestituir,
				"detalle":            replicaMeses,
			},
		},
	}
	raw, _ := json.Marshal(combined)
	return requestresponse.APIResponseDTO(true, 200, json.RawMessage(raw))
}

func fetchAmazonNovedades(numContrato string, vigencia int) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/novedad_postcontractual?query=NumeroContrato:%s,Vigencia:%d", beego.AppConfig.String("AdministrativaAmazonService"), numContrato, vigencia)
	var data []map[string]interface{}
	if err := request.GetJson(url, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func fechasDeAmazon(rec map[string]interface{}) (time.Time, time.Time) {
	var fi, ff time.Time
	if v, ok := rec["FechaInicio"]; ok {
		fi = helpers.ParseTimeAny(v)
	}
	if v, ok := rec["FechaFin"]; ok {
		ff = helpers.ParseTimeAny(v)
	}
	if ff.IsZero() && !fi.IsZero() {
		ff = fi
	}
	return fi, ff
}

func findLatestTipo(list []map[string]interface{}, tipo int) map[string]interface{} {
	var best map[string]interface{}
	var bestId int
	for _, rec := range list {
		if t, ok := rec["TipoNovedad"]; ok {
			tv := intFromAny(t)
			if tv == tipo {
				rid := helpers.GetRowId(rec)
				if rid > bestId {
					best = rec
					bestId = rid
				}
			}
		}
	}
	return best
}

func findByID(list []map[string]interface{}, id int) map[string]interface{} {
	for _, rec := range list {
		if helpers.GetRowId(rec) == id {
			return rec
		}
	}
	return nil
}

func intFromAny(v interface{}) int {
	switch t := v.(type) {
	case int:
		return t
	case int64:
		return int(t)
	case float64:
		return int(t)
	case string:
		i, _ := strconv.Atoi(strings.TrimSpace(t))
		return i
	default:
		return 0
	}
}
