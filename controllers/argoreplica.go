package controllers

import (
	"fmt"
	"strconv"

	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/models"
)

// ArgoReplicaController operations for ArgoReplica
type ArgoReplicaController struct {
	beego.Controller
}

// URLMapping ...
func (c *ArgoReplicaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create ArgoReplica
// @Param	body		body 	models.ArgoReplica	true		"body for ArgoReplica content"
// @Success 201 {object} models.ArgoReplica
// @Failure 403 body is empty
// @router / [post]
func (c *ArgoReplicaController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get ArgoReplica by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ArgoReplica
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ArgoReplicaController) GetOne() {
	var resultjbpm map[string]interface{}
	var novedad_replicaCollection map[string]interface{}
	var novedad []interface{}
	var NovedadGET []map[string]interface{}
	var datos_replica map[string]interface{}
	var novedadid float64

	var novedadformated map[string]interface{}
	//var resultjbpmjson models.JbpmReplica
	idStr := c.Ctx.Input.Param(":id")
	var alerta models.Alert
	error := models.GetJsonWSO2(beego.AppConfig.String("jbpmService")+"/services/bodega_temporal.HTTPEndpoint/novedad/"+idStr, &resultjbpm)
	formatdata.JsonPrint(resultjbpm)

	if error == nil {
		novedad_replicaCollection = resultjbpm["novedad_replicaCollection"].(map[string]interface{})
		if novedad_replicaCollection["novedad"] != nil {

			novedad = novedad_replicaCollection["novedad"].([]interface{})
			datos_replica = novedad[0].(map[string]interface{})
			novedadid, _ = strconv.ParseFloat(datos_replica["novedad_id"].(string), 32)
			idStr1 := strconv.FormatFloat(novedadid, 'f', -1, 64)
			error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/?query=id:"+idStr1+"&limit=0", &NovedadGET)
			if error == nil {

				NovedadGETtrad := NovedadGET[0]
				fmt.Println("\n", novedadid, "\n", error)

				switch NovedadGETtrad["TipoNovedad"].(float64) {
				case 8:
					novedadformated = models.GetNovedadAdicion(NovedadGETtrad)
				case 2:
					novedadformated = models.GetNovedadCesion(NovedadGETtrad)
				}

				formatdata.JsonPrint(NovedadGETtrad)
				formatdata.JsonPrint(novedadformated)
				alerta.Type = "OK"
				alerta.Code = "200"
				alerta.Body = novedadformated
				c.Data["json"] = alerta
				//c.Abort("400")
				c.Ctx.Output.SetStatus(200)

			} else {
				alerta.Type = "error"
				alerta.Code = "400"
				alerta.Body = "No se pudo realizar el get con el id específicado"
				c.Data["json"] = alerta
				//c.Abort("400")
				c.Ctx.Output.SetStatus(400)
			}

		} else {
			alerta.Type = "error"
			alerta.Code = "400"
			alerta.Body = "No se encontraron datos asociados a ese id"
			c.Data["json"] = alerta
			//c.Abort("400")
			c.Ctx.Output.SetStatus(400)
		}

	} else {
		fmt.Println("entro al error\n", error)
		alerta.Type = "error"
		alerta.Code = "400"
		alerta.Body = "No se ha podido realizar la petición GET"
		c.Data["json"] = alerta
		//c.Abort("400")
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get ArgoReplica
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ArgoReplica
// @Failure 403
// @router / [get]
func (c *ArgoReplicaController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the ArgoReplica
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ArgoReplica	true		"body for ArgoReplica content"
// @Success 200 {object} models.ArgoReplica
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ArgoReplicaController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the ArgoReplica
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ArgoReplicaController) Delete() {

}
