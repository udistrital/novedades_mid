package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/time_bogota"
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

	var registroNovedad map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroNovedad); err == nil {

		result, err1 := RegistrarNovedad(registroNovedad)

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

// GetOne ...
// @Title GetOne
// @Description get Novedades by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Novedades
// @Failure 403 :id is empty
// @router /:id/:vigencia [get]
func (c *NovedadesController) GetOne() {
	var novedades []map[string]interface{}
	var novedadformated map[string]interface{}
	novedadesformated := make([]map[string]interface{}, 0)
	// var vacio map[string]interface{}
	//var Result interface{}
	var alerta models.Alert
	//alertas := append([]interface{}{"error"})
	idStr := c.Ctx.Input.Param(":id")
	vigencia := c.Ctx.Input.Param(":vigencia")

	error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/?query=contrato_id:"+idStr+",vigencia:"+vigencia+"&limit=0&sortby=FechaCreacion&order=asc", &novedades)

	if len(novedades) != 0 {
		if novedades[0]["TipoNovedad"] != nil {
			for _, novedad := range novedades {

				idTipoNovedad := novedad["TipoNovedad"].(float64)

				switch idTipoNovedad {
				case 1:
					fmt.Println("Novedad suspensión")
					novedadformated = models.GetNovedadSuspension(novedad)
					novedadesformated = append(novedadesformated, novedadformated)
				case 2:
					fmt.Println("Novedad Cesión")
					novedadformated = models.GetNovedadCesion(novedad)
					novedadesformated = append(novedadesformated, novedadformated)
				case 3:
					fmt.Println("Novedad Reinicio")
					novedadformated = models.GetNovedadReinicio(novedad)
					novedadesformated = append(novedadesformated, novedadformated)
				case 5:
					fmt.Println("Novedad Terminación Anticipada")
					novedadformated = models.GetNovedadTAnticipada(novedad)
					novedadesformated = append(novedadesformated, novedadformated)
				case 6:
					fmt.Println("Novedad Adición")
					novedadformated = models.GetNovedadAdicion(novedad)
					novedadesformated = append(novedadesformated, novedadformated)
				case 7:
					fmt.Println("Novedad Prórroga")
					novedadformated = models.GetNovedadProrroga(novedad)
					novedadesformated = append(novedadesformated, novedadformated)
				case 8:
					fmt.Println("Novedad Adición/prorroga")
					novedadformated = models.GetNovedadAdProrroga(novedad)
					novedadesformated = append(novedadesformated, novedadformated)
				}

			}
		} else {
			novedadesformated = []map[string]interface{}{}
		}
	}

	fmt.Println(error)
	if novedades != nil {
		alerta.Type = "OK"
		alerta.Code = "200"
		alerta.Body = novedadesformated
	} else {
		alerta.Type = "error"
		alerta.Code = "400"
		alerta.Body = "No se ha podido realizar la petición GET"
	}

	c.Data["json"] = alerta
	c.ServeJSON()

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
	var resultado interface{}
	//var Result interface{}
	var alerta models.Alert
	//resultado = models.GetNovedades("/v1/novedades_poscontractuales")
	//fmt.Println(errResultado, resultado)
	alerta.Type = "OK"
	alerta.Code = "200"
	alerta.Body = resultado
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
// @router /:id [put]
func (c *NovedadesController) Put() {
	var reinicio map[string]interface{} //[]models.NovedadSuspensionPut
	var alertErr models.Alert
	var result map[string]interface{}
	alertas := append([]interface{}{"Response:"})
	idStr := c.Ctx.Input.Param(":id")
	url := "/novedad_postcontractual/" + idStr

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &reinicio); err == nil {
		if err := models.SendJson(beego.AppConfig.String("AdministrativaAmazonService")+url, "PUT", &result, &reinicio); err == nil {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = result

		} else {
			alertErr.Type = "error"
			alertErr.Code = "400"
			alertas = append(alertas, err)
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

// Delete ...
// @Title Delete
// @Description delete the Novedades
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *NovedadesController) Delete() {

}

//RegistrarNovedadMongo Función para registrar la novedad en postgresql
func RegistrarNovedad(novedad map[string]interface{}) (status interface{}, outputError interface{}) {

	registroNovedadPost := make(map[string]interface{})
	registroNovedadPost = novedad

	var NovedadPoscontractualPost map[string]interface{}
	var resultadoRegistro map[string]interface{}
	var errRegNovedad interface{}

	switch registroNovedadPost["tiponovedad"] {
	case "NP_SUS":
		// suspensión
		fmt.Println("Novedad de suspensión")
		NovedadPoscontractualPost = models.ConstruirNovedadSuspension(registroNovedadPost)
	case "NP_CES":
		// cesión
		fmt.Println("Novedad de cesión")
		NovedadPoscontractualPost = models.ConstruirNovedadCesion(registroNovedadPost)
	case "NP_REI":
		// reinicio
		fmt.Println("Novedad de reinicio")
		NovedadPoscontractualPost = models.ConstruirNovedadReinicio(registroNovedadPost)
	case "NP_TER":
		// terminacion anticipada
		fmt.Println("Novedad de terminación anticipada")
		NovedadPoscontractualPost = models.ConstruirNovedadTAnticipada(registroNovedadPost)
	case "NP_ADI":
		// adición
		fmt.Println("Novedad de adición")
		NovedadPoscontractualPost = models.ConstruirNovedadAdicionPost(registroNovedadPost)
	case "NP_PRO":
		// prórroga
		fmt.Println("Novedad de prorroga")
		NovedadPoscontractualPost = models.ConstruirNovedadProrrogaPost(registroNovedadPost)
	case "NP_ADPRO":
		// adicion/prorroga
		fmt.Println("Novedad de adicion/prorroga")
		NovedadPoscontractualPost = models.ConstruirNovedadAdProrrogaPost(registroNovedadPost)
	}

	if registroNovedadPost["tiponovedad"] == "NP_CES" {
		errRegNovedad = request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/trNovedad/trnovedadpoliza", "POST", &resultadoRegistro, NovedadPoscontractualPost)
	} else {
		errRegNovedad = request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/trNovedad", "POST", &resultadoRegistro, NovedadPoscontractualPost)
	}

	if resultadoRegistro["Status"] == "400" || errRegNovedad != nil {
		fmt.Println(errRegNovedad)
		return nil, resultadoRegistro

	} else {
		formatdata.JsonPrint(resultadoRegistro)
		return resultadoRegistro, nil

	}

}

//Función que duplicará los datos de registro de novedades de adición y cesión
func RegistroAdministrativaAmazon(Novedad map[string]interface{}) (idRegistroAdmAmazon int, outputError interface{}) {
	NovedadAmazon := Novedad
	var NovedadGET []map[string]interface{}
	var NovedadAdmAmazonFormatted map[string]interface{}
	var resultadoregistroadmamazon map[string]interface{}
	var resultadoregistrojbpm map[string]interface{}
	var errRegNovedad error

	NovedadMap := NovedadAmazon["NovedadPoscontractual"].(map[string]interface{})
	idStrf64 := NovedadMap["Id"].(float64)
	idStr := strconv.FormatFloat(idStrf64, 'f', -1, 64)

	// ID de la novedad que se acaba de guardar

	error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/?query=id:"+idStr+"&limit=0", &NovedadGET)

	//Para novedad de adición prorroga

	NovedadAdmAmazonFormatted = models.FormatAdmAmazonNovedadCesion(NovedadGET)
	urladm := beego.AppConfig.String("AdministrativaAmazonService") + "/novedad_postcontractual"
	errRegNovedad = request.SendJson(urladm, "POST", &resultadoregistroadmamazon, NovedadAdmAmazonFormatted)
	formatdata.JsonPrint(resultadoregistroadmamazon)

	fmt.Println("Aqui se muestra la traducción de la novedad para replica en AdmAmazon \n", NovedadAdmAmazonFormatted, error)
	formatdata.JsonPrint(NovedadAdmAmazonFormatted)

	if errRegNovedad == nil && resultadoregistroadmamazon["Id"] != nil {
		idResultRegistroAdmAmazon := resultadoregistroadmamazon["Id"]

		registrojbpm := map[string]interface{}{
			"_post_novedad": map[string]interface{}{
				"argonovedad_id":     idResultRegistroAdmAmazon.(float64),
				"novedad_id":         idStrf64,
				"fecha_creacion":     time_bogota.TiempoBogotaFormato(),
				"fecha_modificacion": time_bogota.TiempoBogotaFormato(),
				"activo":             true,
			},
		}

		formatdata.JsonPrint(registrojbpm)

		errRegNovedad = models.SendJson(beego.AppConfig.String("jbpmService")+"/services/bodega_temporal.HTTPEndpoint/novedad", "POST", &resultadoregistrojbpm, registrojbpm)

		return 0, nil

	} else {
		errorRegistro := "No se pudo guardar en Administrativa amazon"
		return 0, errorRegistro
	}

}
