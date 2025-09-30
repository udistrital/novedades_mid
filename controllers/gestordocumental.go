package controllers

import (
	"encoding/json"
	"fmt"

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
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
}

// GetOne ...
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
	alertas := []interface{}{"Response:"}

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

// Post ...
// @Title PostGestorDocumental
// @Description Crear documento en Nuxeo
// @Param   body        body    {}  true        "Crear documento en Nuxeo"
// @Success 200 {}
// @Failure 403 body is empty
// @router / [post]
func (c *GestorDocumentalController) Post() {

	var registroDoc []map[string]interface{}
	var alertErr models.Alert
	alertas := []interface{}{"Response:"}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroDoc); err == nil {

		result, err1 := RegistrarDoc(registroDoc, "upload")

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

func RegistrarDoc(documento []map[string]interface{}, url string) (status interface{}, outputError interface{}) {

	var resultadoRegistro map[string]interface{}
	var errRegDoc interface{}
	errRegDoc = models.SendJson(beego.AppConfig.String("GestorDocumentalMid")+"/document/"+url, "POST", &resultadoRegistro, documento)
	if resultadoRegistro != nil {
		return resultadoRegistro["res"], nil
	} else {
		return nil, errRegDoc
	}
	// if resultadoRegistro["Status"].(string) == "200" && errRegDoc == nil {

	// 	// jsonString, _ := json.Marshal(resultadoRegistro["res"])

	// } else {

	// }
}

// Post ...
// @Title PostGestorDocumental
// @Description Crear documento en Nuxeo
// @Param   body        body    {}  true        "Crear documento en Nuxeo"
// @Success 200 {}
// @Failure 403 body is empty
// @router /:url [put]
func (c *GestorDocumentalController) Put() {

	var registroDoc []map[string]interface{}
	var alertErr models.Alert
	alertas := []interface{}{"Response:"}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroDoc); err == nil {
		result, err1 := RegistrarDoc(registroDoc, "firma_electronica")

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

func ConsultarEstadoNovedad(idRegistro string) (estado string) {
	var response map[string]interface{}
	var estadoNovedad string
	err := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro/"+idRegistro, &response)
	if err == nil {
		var data map[string]interface{}
		data = response["Data"].(map[string]interface{})
		estadoNovedad = data["Nombre"].(string)
	}
	return estadoNovedad
}

func ActualizarEstadoDocNovedad(documento string, estado string) (err error) {

	var resultadoRegistro map[string]interface{}
	var errRegDoc interface{}

	errRegDoc = models.SendJson(beego.AppConfig.String("GestorDocumentalMid")+"/document/"+documento+"/metadata", "POST", &resultadoRegistro, documento)
	fmt.Println(errRegDoc)
	// var estructura []interface{}
	// {
	// 	"properties":{
	// 		"dc:description": "ejemplo",
	// 		"dc:source":"prueba metadatos 2021",
	// 		"dc:publisher": "cristian alape",
	// 		"dc:rights": "Universidad Distrital Francisco José de Caldas",
	// 		"dc:title": "prueba_core_2021_3",
	// 		"dc:language": "Español",
	// 		"nxtag:tags": [
	// 				{
	// 					"label": "etiqueta_1",
	// 					"username": "cristian alape"
	// 				},
	// 								{
	// 					"label": "etiqueta_2",
	// 					"username": "cristian alape"
	// 				}
	// 			]
	// 	}
	// }
	return nil
}
