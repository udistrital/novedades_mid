package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// CambioEstadoContratoValidoController operations for CambioEstadoContratoValido
type CambioEstadoContratoValidoController struct {
	beego.Controller
}

// URLMapping ...
func (c *CambioEstadoContratoValidoController) URLMapping() {
	c.Mapping("ValidarCambioEstado", c.ValidarCambioEstado)
}

// ValidarCambiosEstado ...
// @Title ValidarCambiosEstado
// @Description create ValidarCambiosEstado
// @Param   body        body    {}  true        "Crear documento en Nuxeo"
// @Success 201 {object} models.Alert
// @Failure 400 body is empty
// @router / [post]
func (c *CambioEstadoContratoValidoController) ValidarCambioEstado() {

	var estados []models.EstadoContrato //0: actual y 1:siguiente
	var cambioEstado map[string]interface{}
	var alertErr models.Alert
	alertas := []interface{}{"Response:"}

	//result, err1 := consultarCambioEstado(estados)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &estados); err == nil {
		result, err1 := consultaEstado(estados)
		if err1 == nil {
			alertErr.Type = "OK"
			alertErr.Code = "201"
			alertErr.Body = result

		} else {
			alertErr.Type = "error"
			alertErr.Code = "400"
			alertas = append(alertas, err1.Error())
			alertErr.Body = alertas
			c.Ctx.Output.SetStatus(400)
		}

	} else {
		if err2 := json.Unmarshal(c.Ctx.Input.RequestBody, &cambioEstado); err2 == nil {
			if result, err1 := ActivarContrato(cambioEstado); err1 == nil {
				alertErr.Type = "OK"
				alertErr.Code = "200"
				alertas = append(alertas, result)
				alertErr.Body = alertas
				c.Ctx.Output.SetStatus(200)
			} else {
				alertErr.Type = "error"
				alertErr.Code = "400"
				alertas = append(alertas, err1)
				alertErr.Body = alertas
				c.Ctx.Output.SetStatus(400)
			}
		} else {
			alertErr.Type = "error"
			alertErr.Code = "400"
			alertas = append(alertas, err.Error())
			alertErr.Body = alertas
			c.Ctx.Output.SetStatus(400)
		}
	}

	c.Data["json"] = alertErr
	c.ServeJSON()

}

func consultaEstado(estados []models.EstadoContrato) (status string, err error) {
	var resEstado string

	errRegDoc := models.SendJson(beego.AppConfig.String("AdminMidApi")+"/validarCambioEstado", "POST", &resEstado, &estados)

	if errRegDoc == nil {
		return resEstado, nil
	}
	return "", errRegDoc
}

func ActivarContrato(cambioEstado map[string]interface{}) (map[string]interface{}, map[string]interface{}) {
	var result map[string]interface{}
	estadoStruct := models.CambioEstado{}
	estadoStruct.Estado.Id = int(cambioEstado["Estado"].(map[string]interface{})["Id"].(float64))
	estadoStruct.FechaRegistro = cambioEstado["FechaRegistro"].(string)
	estadoStruct.NumeroContrato = cambioEstado["NumeroContrato"].(string)
	estadoStruct.Usuario = cambioEstado["Usuario"].(string)
	estadoStruct.Vigencia = int(cambioEstado["Vigencia"].(float64))
	errRegDoc := request.SendJson(beego.AppConfig.String("AdministrativaAmazonService")+"/contrato_estado", "POST", &result, estadoStruct)
	if errRegDoc != nil {
		fmt.Println("Error sending JSON:", errRegDoc)
		errorMap := map[string]interface{}{"funcion": "/ActivarContrato", "err": errRegDoc.Error()}
		return nil, errorMap
	}
	return result, nil
}
