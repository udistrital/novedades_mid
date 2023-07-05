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
}

// Get ...
// @Title Get
// @Description get Novedades by id
// @Param	rol		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Novedades
// @Failure 403 :rol is empty
// @router /:rol [get]
func (c *AprobacionController) Get() {

	rol := c.Ctx.Input.Param(":rol")

	var novedades []map[string]interface{}
	var response []map[string]interface{}
	var response1 []map[string]interface{}
	var response2 []map[string]interface{}

	var alerta models.Alert

	switch rol {
	case "SUPERVISOR":
		if errSup := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:4536", &response); errSup == nil {
			novedades = append(novedades, response...)
		}
		if errSup1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:4537", &response1); errSup1 == nil {
			novedades = append(novedades, response1...)
		}
		if errSup2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:4538", &response2); errSup2 == nil {
			novedades = append(novedades, response2...)
		}
	case "ORDENADOR_DEL_GASTO":
		if err1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:4539", &response); err1 == nil {
			novedades = append(novedades, response...)
		}
		if err2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:4540", &response1); err2 == nil {
			novedades = append(novedades, response1...)
		}
	case "ASISTENTE_JURIDICA":
		if err1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:4541", &response); err1 == nil {
			novedades = append(novedades, response...)
		}
		if err2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"novedades_poscontractuales?limit=0&sortby=FechaCreacion&order=asc&query=estado:4542", &response1); err2 == nil {
			novedades = append(novedades, response1...)
		}
	default:
		fmt.Println("El rol no coincide con alguno registrado!")
	}

	if novedades != nil {
		alerta.Type = "OK"
		alerta.Code = "200"
		alerta.Body = novedades
	} else {
		alerta.Type = "error"
		alerta.Code = "400"
		alerta.Body = "No se ha podido realizar la petición GET"
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
	err := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/"+id, &novedad)
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
