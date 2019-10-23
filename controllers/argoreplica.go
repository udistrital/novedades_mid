package controllers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

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
	var resultjbpm interface{}
	var resultjbpmjson models.JbpmReplica
	idStr := c.Ctx.Input.Param(":id")
	var alerta models.Alert
	//error := request.GetJson(beego.AppConfig.String("jbpmService")+"/services/bodega_temporal.HTTPEndpoint/novedad/"+idStr, &resultjbpm)
	error := getXml(beego.AppConfig.String("jbpmService")+"/services/bodega_temporal.HTTPEndpoint/novedad/"+idStr, &resultjbpm)

	fmt.Println(beego.AppConfig.String("jbpmService")+"/services/bodega_temporal.HTTPEndpoint/novedad/"+idStr+"\n", resultjbpm)

	if error == nil {

		json_jbpmreplica, error_json := json.Marshal(resultjbpm)

		if error_json == nil {

			if err := json.Unmarshal(json_jbpmreplica, &resultjbpmjson); err == nil {

				fmt.Println("entro al true \n", resultjbpmjson)
				alerta.Type = "OK"
				alerta.Code = "200"
				alerta.Body = resultjbpmjson
				c.Data["json"] = alerta
				c.Ctx.Output.SetStatus(200)

			} else {
				alerta.Type = "error"
				alerta.Code = "400"
				alerta.Body = "falló el unmarshall"
				c.Data["json"] = alerta
				//c.Abort("400")
				c.Ctx.Output.SetStatus(400)

			}

		} else {
			alerta.Type = "error"
			alerta.Code = "400"
			alerta.Body = "No se pudo formatear a JSON"
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

func getXml(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return xml.NewDecoder(r.Body).Decode(target)
}
