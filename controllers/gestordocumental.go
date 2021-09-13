package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// GestorDocumentalController operations for Nuxeo
type GestorDocumentalController struct {
	beego.Controller
}

// URLMapping ...
func (c *GestorDocumentalController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetOne", c.GetOne)
}

// GetGestorDocumental ...
// @Title GetGestorDocumental
// @Description obtener documento por enlace
// @Param	enlace		path 	string	true		"The key for staticblock"
// @Success 200 {}
// @Failure 403 :enlace is empty
// @router /:enlace [get]
func (c *GestorDocumentalController) GetOne() {
	var novedad map[string]interface{}
	idStr := c.Ctx.Input.Param(":enlace")
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	error := request.GetJson(beego.AppConfig.String("GestorDocumentalMid")+"/document/"+idStr, &novedad)

	if error == nil {
		if len(novedad) == 2 && novedad["Status"].(string) == "500" {
			alertErr.Type = "error"
			alertErr.Code = "500"
			alertErr.Body = novedad
			c.Ctx.Output.SetStatus(500)
		} else {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = novedad
		}

	} else {
		alertErr.Type = "error"
		alertErr.Code = "400"
		alertas = append(alertas, error)
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(400)
	}
	c.Data["json"] = alertErr
	c.ServeJSON()
}

// PostGestorDocumental ...
// @Title PostGestorDocumental
// @Description Crear documento en Nuxeo
// @Param   body        body    {}  true        "Crear documento en Nuxeo"
// @Success 200 {}
// @Failure 403 body is empty
// @router / [post]
func (c *GestorDocumentalController) Post() {

	var registroDoc []map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroDoc); err == nil {

		result, err1 := RegistrarDoc(registroDoc)

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

func RegistrarDoc(documento []map[string]interface{}) (status interface{}, outputError interface{}) {

	var resultadoRegistro map[string]interface{}
	var errRegDoc interface{}

	errRegDoc = models.SendJson(beego.AppConfig.String("GestorDocumentalMid")+"/upload", "POST", &resultadoRegistro, documento)

	if resultadoRegistro["Status"].(string) == "200" && errRegDoc == nil {

		jsonString, _ := json.Marshal(resultadoRegistro["res"])
		return jsonString, nil

	} else {
		return nil, resultadoRegistro["Error"].(string)
	}

}
