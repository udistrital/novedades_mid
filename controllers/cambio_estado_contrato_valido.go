package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/models"
	//. "github.com/udistrital/golog"
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

// DeleteEstadoContrato ...
// @Title DeleteEstadoContrato
// @Description delete EstadoContrato by id
// @Param	id		path 	string	true		"id de estado a eliminar"
// @Success 200 {string} OK!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CambioEstadoContratoValidoController) DeleteEstadoContrato() {
	id := c.Ctx.Input.Param(":id")
	if id != "" {
		err := DeleteEstadoContrato(id)
		if err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(400)
		}
	} else {
		c.Data["json"] = "id is empty"
		c.Ctx.Output.SetStatus(403)
	}
	c.ServeJSON()
}

func DeleteEstadoContrato(idEstado string) error {
	var resEstado string
	errRegDoc := models.SendJson(beego.AppConfig.String("AdministrativaAmazonService")+"contrato_estado/"+idEstado, "DELETE", &resEstado, nil)
	if errRegDoc == nil {
		return nil
	}
	return errRegDoc
}
