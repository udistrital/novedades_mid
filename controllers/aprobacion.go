package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/models"
	"github.com/udistrital/utils_oas/request"
)

type AprobacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *AprobacionController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
}

// Get ...
// @Title Get
// @Description get Novedades by id
// @Param	rol		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Novedades
// @Failure 403 :rol is empty
// @router /:rol [get]
func (c *AprobacionController) Get() {

	var alerta models.Alert
	rol := c.Ctx.Input.Param(":rol")

	resultTabla, err := models.ConsultaTablaAprobacion(rol)

	if err == nil {
		alerta.Type = "OK"
		alerta.Code = "200"
		alerta.Body = resultTabla
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
// @Param	body		body 	models.Novedades	true		"body for Novedades content"
// @Success 200 {object} models.Novedades
// @Failure 403 :id is not int
// @router /:id/:estado [put]
func (c *AprobacionController) Put() {

	idStr := c.Ctx.Input.Param(":id")
	estado := c.Ctx.Input.Param(":estado")
	var registroNovedad map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroNovedad); err == nil {

		result, err1 := ActualizarEstadoNovedad(idStr, estado)

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

func ActualizarEstadoNovedad(id string, estado string) (status interface{}, outputError interface{}) {

	var novedad map[string]interface{}
	var resultadoRegistro map[string]interface{}
	err := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/nvedades_poscontractuales/"+id, &novedad)
	if err == nil {
		novedad["Estado"] = estado
		errRegNovedad := request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/"+id, "PUT", &resultadoRegistro, novedad)
		if errRegNovedad == nil {
			fmt.Println("Novedad actualizada!!")
			return resultadoRegistro, nil
		} else {
			outputError = map[string]interface{}{"funcion": "/ActualizarEstadoNovedad", "err": errRegNovedad}
			return nil, outputError
		}
	} else {
		outputError = map[string]interface{}{"funcion": "/ActualizarEstadoNovedad", "err": err}
		return nil, outputError
	}
}
