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
	topNov, prevNov, _ := getTopAndPrevNovedadAmazon(n.ContratoId, n.Vigencia)
	var topIni, topFin time.Time
	if topNov != nil {
		if v, ok := topNov["FechaInicio"]; ok {
			topIni = helpers.ParseTimeAny(v)
		}
		if v, ok := topNov["FechaFin"]; ok {
			topFin = helpers.ParseTimeAny(v)
		}
		if topFin.IsZero() && !topIni.IsZero() {
			topFin = topIni
		}
	}

	var bodyIni, bodyFin time.Time
	if prevNov != nil {
		if v, ok := prevNov["FechaInicio"]; ok {
			bodyIni = helpers.ParseTimeAny(v)
		}
		if v, ok := prevNov["FechaFin"]; ok {
			bodyFin = helpers.ParseTimeAny(v)
		}
	}
	if bodyIni.IsZero() {
		if aiIni, aiFin, err := getFechasActaInicioPorIdVigencia(numeroContratoID, n.Vigencia); err == nil {
			bodyIni, bodyFin = aiIni, aiFin
		} else {
			beego.Warn("[AnularNovedad] No se pudo derivar fechas de acta_inicio: ", err)
		}
	}
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
	var amazonEliminada interface{} = nil
	if topNov != nil {
		if idAmazon, ok := topNov["Id"].(float64); ok {
			urlDelete := fmt.Sprintf("%s/novedad_postcontractual/%d",
				beego.AppConfig.String("AdministrativaAmazonService"),
				int(idAmazon),
			)
			var deleteResp interface{}
			if err := request.SendJson(urlDelete, "DELETE", &deleteResp, nil); err == nil {
				amazonEliminada = deleteResp
			} else {
				amazonEliminada = map[string]interface{}{"error": err.Error()}
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

	initial := helpers.PickInitialRow(rows)
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

	mesesAnulados := helpers.MesesEntreSafe(topIni, topFin)
	mesesRestituir := helpers.MesesEntreSafe(bodyIni, bodyFin)

	replicaMeses := ReplicafechaAnterior(contratoFull, mesesAnulados, mesesRestituir, bodyIni, bodyFin, contratoIds)

	combined := map[string]interface{}{
		"novedad_anulada":  respPUTNov,
		"amazon_eliminada": amazonEliminada,
		"cambios_estado":   cambiosEstado,
		"titan": map[string]interface{}{
			"numero_contrato_utilizado": numStr,

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
