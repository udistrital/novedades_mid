package controllers

import (
	"encoding/json"

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
	c.Mapping("DeleteEstadoContrato", c.DeleteEstadoContrato)
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
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

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
			alertas = append(alertas, err.Error())
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

// ValidarCambiosEstado ...
// @Title ValidarCambiosEstado
// @Description create ValidarCambiosEstado
// @Param   body        body    {}  true        "Crear documento en Nuxeo"
// @Success 201 {object} models.Alert
// @Failure 400 body is empty
// @router / [post]

// DeleteEstadoContrato ...
// @Title DeleteEstadoContrato
// @Description delete EstadoContrato by id
// @Param	id		path 	string	true		"id de estado a eliminar"
// @Success 200 {string} OK!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CambioEstadoContratoValidoController) DeleteEstadoContrato() {

	idStr := c.Ctx.Input.Param(":id")
	// var registroNovedad map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	result, err1 := DeleteEstadoArgo(idStr)
	if err1 == nil {
		alertErr.Type = "OK"
		alertErr.Code = "200"
		alertErr.Body = result
	} else {
		alertErr.Type = "error"
		alertErr.Code = "400"
		alertas = append(alertas, err1)
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(400)
	}

	c.Data["json"] = alertErr
	c.ServeJSON()
}

func DeleteEstadoArgo(idEstado string) (status map[string]interface{}, outputError interface{}) {
	var result map[string]interface{}
	errRegDoc := request.SendJson(beego.AppConfig.String("AdministrativaAmazonService")+"contrato_estado/"+idEstado, "DELETE", &result, nil)
	if errRegDoc == nil {
		return nil, result
	}
	errorMap := map[string]interface{}{"funcion": "/DeleteEstadoArgo", "err": "La función de borrado no hizo nada"}
	return nil, errorMap
}
