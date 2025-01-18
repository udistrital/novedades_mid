package controllers

import (
	"encoding/json"
	"strconv"

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
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
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

// GetOne ...
// @Title GetOne
// @Description get ContratoEstado by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Alert
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CambioEstadoContratoValidoController) GetOne() {

	var alerta models.Alert
	idStr := c.Ctx.Input.Param(":id")

	result, err := GetEstadoArgo(idStr)
	if result != nil {
		alerta.Type = "OK"
		alerta.Code = "200"
		alerta.Body = result
	} else {
		alerta.Type = "error"
		alerta.Code = "400"
		alerta.Body = err
	}

	c.Data["json"] = alerta
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Novedades
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	map[string]interface{}	true		"body for Novedades content"
// @Success 200 {object} models.Alert
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CambioEstadoContratoValidoController) Put() {
	var registroEstado map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroEstado); err == nil {

		result, err1 := ActualizarEstadoArgo(registroEstado)

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

// Delete ...
// @Title Delete
// @Description delete EstadoContrato by id
// @Param	body		body 	map[string]interface{}	true		"body for Novedades content"
// @Success 200 {object} models.Alert
// @Failure 403 :id is not int
// @router /:id [delete]
func (c *CambioEstadoContratoValidoController) Delete() {
	var alerta models.Alert
	idStr := c.Ctx.Input.Param(":id")
	result, err := DeleteEstadoArgo(idStr)
	if result != nil {
		alerta.Type = "OK"
		alerta.Code = "200"
		alerta.Body = result
	} else {
		alerta.Type = "error"
		alerta.Code = "400"
		alerta.Body = err
	}

	c.Data["json"] = alerta
	c.ServeJSON()
}

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

func GetEstadoArgo(idArgo string) (map[string]interface{}, map[string]interface{}) {
	var result map[string]interface{}
	query := "contrato_estado/" + idArgo
	if errArgo := models.GetJson(beego.AppConfig.String("AdministrativaAmazonService")+query, &result); errArgo == nil {
		return result, nil
	} else {
		outputError := map[string]interface{}{"funcion": "/GetEstadoArgo", "err": "No se pudo obtener el registro de la base de datos"}
		return nil, outputError
	}
}

func ActualizarEstadoArgo(registroEstado map[string]interface{}) (map[string]interface{}, map[string]interface{}) {
	var result string
	query := "contrato_estado/" + strconv.Itoa(int(registroEstado["Id"].(float64)))
	errRegDoc := request.SendJson(beego.AppConfig.String("AdministrativaAmazonService")+query, "PUT", &result, registroEstado)
	if errRegDoc == nil && result == "OK" {
		return map[string]interface{}{"funcion": "/ActualizarEstadoArgo", "result": result}, nil
	}
	errorMap := map[string]interface{}{"funcion": "/ActualizarEstadoArgo", "err": errRegDoc.Error()}
	return nil, errorMap
}

func DeleteEstadoArgo(idEstado string) (map[string]interface{}, map[string]interface{}) {
	var result string
	errRegDoc := request.SendJson(beego.AppConfig.String("AdministrativaAmazonService")+"contrato_estado/"+idEstado, "DELETE", &result, "")
	if errRegDoc == nil && result == "OK" {
		return map[string]interface{}{"result": result}, nil
	}
	errorMap := map[string]interface{}{"funcion": "/DeleteEstadoArgo", "err": "La función de borrado no hizo nada"}
	return nil, errorMap
}
