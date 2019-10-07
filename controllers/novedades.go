package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// NovedadesController operations for Novedades
type NovedadesController struct {
	beego.Controller
}

// URLMapping ...
func (c *NovedadesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Novedades
// @Param	body		body 	models.Novedades	true		"body for Novedades content"
// @Success 201 {object} models.Novedades
// @Failure 403 body is empty
// @router / [post]
func (c *NovedadesController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Novedades by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Novedades
// @Failure 403 :id is empty
// @router /:id [get]
func (c *NovedadesController) GetOne() {
	var resultado map[string]interface{}
	var Result interface{}
	var alerta models.Alert
	alertas := append([]interface{}{"error"})
	idStr := c.Ctx.Input.Param(":id")

	//errResultado := request.GetJson("http://"+beego.AppConfig.String("EventoService")+"/calendario_evento/"+idStr, &resultado)

}

// GetAll ...
// @Title GetAll
// @Description get Novedades
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Novedades
// @Failure 403
// @router / [get]
func (c *NovedadesController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Novedades
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Novedades	true		"body for Novedades content"
// @Success 200 {object} models.Novedades
// @Failure 403 :id is not int
// @router /:id [put]
func (c *NovedadesController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Novedades
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *NovedadesController) Delete() {

}

//Función que construirá la novedad a ser consultada.
func ConstruirNovedad(idNovedad string) (novedad map[string]interface{}, outputError interface{}) {

	novedadid := idNovedad

	result, err1 := ConsultarNovedadID(novedadid)
	fmt.Println(result, err1)

	return nil, nil
}

//Funcion que consulta la novedad por id
func ConsultarNovedadID(idNovedad string) (novedad map[string]interface{}, outputError interface{}) {

	novedadid := idNovedad
	var getnovedad []map[string]interface{}
	errnovedad := request.GetJson("http://"+beego.AppConfig.String("NovedadesCrudService")+"/v1/novedades_poscontractuales/?query=ContratoId:"+novedadid, &getnovedad)

	return nil, errnovedad
}
