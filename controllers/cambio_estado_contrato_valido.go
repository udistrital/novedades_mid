package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/models"

	//. "github.com/udistrital/golog"
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

	errRegDoc := request.SendJson(beego.AppConfig.String("AdminMidApi")+"/validarCambioEstado", "POST", &resEstado, &estados)

	if errRegDoc == nil {
		return resEstado, nil
	}
	return "", errRegDoc
}
