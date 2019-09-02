package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/models"
	"github.com/udistrital/utils_oas/time_bogota"
)

// RegistroNovedadController operations for RegistroNovedad
type RegistroNovedadController struct {
	beego.Controller
}

// URLMapping ...
func (c *RegistroNovedadController) URLMapping() {
	c.Mapping("PostRegistroNovedad", c.PostRegistroNovedad)
	//c.Mapping("GetOneRegistroNovedad", c.GetOneRegistroNovedad)
	//c.Mapping("GetAllRegistroNovedad", c.GetAllRegistroNovedad)
	//c.Mapping("PutRegistroNovedad", c.PutRegistroNovedad)
	//c.Mapping("DeleteRegistroNovedad", c.DeleteRegistroNovedad)
}

// PostRegistroNovedad ...
// @Title PostRegistroNovedad
// @Description Agregar RegistroNovedad
// @Param   body        body    {}  true        "body Agregar RegistroNovedad content"
// @Success 200 {}
// @Failure 403 body is empty
// @router / [post]
func (c *RegistroNovedadController) PostRegistroNovedad() {
	var registroNovedad map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	horaRegistro := time_bogota.Tiempo_bogota()

	//fmt.Println(registroNovedad, alertErr, horaRegistro)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroNovedad); err == nil {

		registroNovedadPost := make(map[string]interface{})
		registroNovedadPost["registroNovedad"] = map[string]interface{}{
			"Novedad":       registroNovedad["Novedad"],
			"TipoNovedad":   registroNovedad["TipoNovedad"],
			"Contrato":      registroNovedad["Contrato"],
			"FechaRegistro": horaRegistro,
		}

		fmt.Println(registroNovedadPost, horaRegistro)
		alertErr.Type = "OK"
		alertErr.Code = "200"
		alertas = append(alertas, registroNovedadPost)
	} else {
		alertErr.Type = "error"
		alertErr.Code = "400"
		alertas = append(alertas, err.Error())
	}

	alertErr.Body = alertas
	c.Data["json"] = alertErr
	c.ServeJSON()

}
