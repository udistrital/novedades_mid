package services

import (
	"fmt"
	"time"

	"github.com/udistrital/novedades_mid/helpers"
	"github.com/udistrital/novedades_mid/models"
	"github.com/udistrital/utils_oas/request"
)

func reprocesarPreliquidacionConBody(body interface{}) (map[string]interface{}, error) {
	base := titanMidBase()
	if base == "" {
		return nil, fmt.Errorf("TitanMidService no configurado")
	}
	var resp map[string]interface{}
	if err := request.SendJson(base+"/preliquidacion/", "POST", &resp, body); err != nil {
		return nil, err
	}
	return resp, nil
}

func PostReplica_Titan(contrato map[string]interface{}, fi, ff time.Time, ano, mes int) map[string]interface{} {
	resp := map[string]interface{}{
		"funcion": "/PostReplica_Titan",
		"ano":     ano,
		"mes":     mes,
	}
	payload := models.NewPreliquidacionReplicaFromContrato(contrato, fi, ff)
	pre, err := reprocesarPreliquidacionConBody(payload)
	resp["payload"] = payload // queda serializable como struct
	resp["resp"] = pre
	resp["error"] = helpers.ErrorString(err)
	return resp
}
func ReplicafechaAnterior(contrato map[string]interface{},
	mesesEliminar [][2]int, mesesCrear [][2]int,
	globalIni, globalFin time.Time, contratoIds []int) []map[string]interface{} {

	var salida []map[string]interface{}

	for _, par := range mesesEliminar {
		m := par[0]
		a := par[1]
		del := deleteCPAndDetailsForMonth(a, m, contratoIds)
		salida = append(salida, map[string]interface{}{
			"eliminacion_mes": []int{m, a},
			"eliminacion":     del,
		})
	}

	for _, par := range mesesCrear {
		m := par[0]
		a := par[1]

		delPrev := deleteCPAndDetailsForMonth(a, m, contratoIds)

		fiMes, ffMes := helpers.MonthBoundsClipped(a, m, globalIni, globalFin)
		if fiMes.After(ffMes) {
			salida = append(salida, map[string]interface{}{
				"post_mes":               []int{m, a},
				"eliminacion_previa_mes": delPrev,
				"skip":                   "FechaInicio > FechaFin (mes fuera del rango restituido)",
			})
			continue
		}
		post := PostReplica_Titan(contrato, fiMes, ffMes, a, m)
		salida = append(salida, map[string]interface{}{
			"post_mes":               []int{m, a},
			"eliminacion_previa_mes": delPrev,
			"post":                   post,
		})
	}

	return salida
}
