package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/helpers"
	"github.com/udistrital/novedades_mid/models"
	"github.com/udistrital/utils_oas/requestresponse"
)

func AnularNovedadPorID(id string) requestresponse.APIResponse {
	url := beego.AppConfig.String("NovedadesCrudService") + "/novedades_poscontractuales/" + id

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return helpers.ErrEmiter(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return helpers.ErrEmiter(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return helpers.ErrEmiter(
			fmt.Errorf("fallo al obtener la novedad: código %d", resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return helpers.ErrEmiter(err)
	}

	var novedad models.Novedad
	if err := json.Unmarshal(body, &novedad); err != nil {
		return helpers.ErrEmiter(fmt.Errorf("error al convertir novedad con id %s", id))
	}

	layoutEntrada := "2006-01-02 15:04:05 -0700 -0700"
	layoutSalida := "2006-01-02 15:04:05"

	if parsed, err := time.Parse(layoutEntrada, novedad.FechaCreacion); err == nil {
		novedad.FechaCreacion = parsed.Format(layoutSalida)
	}
	if parsed, err := time.Parse(layoutEntrada, novedad.FechaModificacion); err == nil {
		novedad.FechaModificacion = parsed.Format(layoutSalida)
	}

	novedad.Activo = false

	jsonBody, err := json.Marshal(novedad)
	if err != nil {
		return helpers.ErrEmiter(err)
	}

	putReq, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return helpers.ErrEmiter(err)
	}
	putReq.Header.Set("Content-Type", "application/json")

	putResp, err := client.Do(putReq)
	if err != nil {
		return helpers.ErrEmiter(err)
	}
	defer putResp.Body.Close()

	putBody, err := io.ReadAll(putResp.Body)
	if err != nil {
		return helpers.ErrEmiter(err)
	}

	return requestresponse.APIResponseDTO(true, putResp.StatusCode, json.RawMessage(putBody))
}
