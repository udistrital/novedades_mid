package services

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/helpers"
	"github.com/udistrital/novedades_mid/models"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/requestresponse"
)

func AnularNovedadPorID(id string) requestresponse.APIResponse {
	var novedades []models.Novedad

	urlGet := beego.AppConfig.String("NovedadesCrudService") + "/novedades_poscontractuales/?query=id:" + id + "&limit=0"
	err := request.GetJson(urlGet, &novedades)
	if err != nil || len(novedades) == 0 {
		return helpers.ErrEmiter(fmt.Errorf("error obteniendo novedad con id %s: %v", id, err))
	}

	novedad := novedades[0]
	novedad.Activo = false
	novedad.FechaModificacion = time.Now().Format("2006-01-02 15:04:05")
	novedad.FechaCreacion = time.Now().Format("2006-01-02 15:04:05")

	urlPut := beego.AppConfig.String("NovedadesCrudService") + "/novedades_poscontractuales/" + id
	var respuesta map[string]interface{}
	err = request.SendJson(urlPut, "PUT", &respuesta, novedad)
	if err != nil {
		return helpers.ErrEmiter(fmt.Errorf("error al enviar PUT para novedad con id %s: %v", id, err))
	}

	var amazonEliminada interface{} = nil
	var amazonNovedades []map[string]interface{}
	urlAmazonGet := fmt.Sprintf(
		"%s/novedad_postcontractual/?query=NumeroContrato:%s,Vigencia:%d&sortby=FechaInicio&order=desc&limit=-1",
		beego.AppConfig.String("AdministrativaAmazonService"),
		strconv.Itoa(novedad.ContratoId),
		novedad.Vigencia,
	)

	if err := request.GetJson(urlAmazonGet, &amazonNovedades); err == nil && len(amazonNovedades) > 0 {
		if idAmazon, ok := amazonNovedades[0]["Id"].(float64); ok {
			urlDelete := fmt.Sprintf("%s/novedad_postcontractual/%d",
				beego.AppConfig.String("AdministrativaAmazonService"),
				int(idAmazon),
			)
			var deleteResp map[string]interface{}
			if err := request.SendJson(urlDelete, "DELETE", &deleteResp, nil); err == nil {
				amazonEliminada = deleteResp
			}
		}
	}

	combined := map[string]interface{}{
		"novedad_anulada":  respuesta,
		"amazon_eliminada": amazonEliminada,
	}
	raw, _ := json.Marshal(combined)
	return requestresponse.APIResponseDTO(true, 200, json.RawMessage(raw))
}
