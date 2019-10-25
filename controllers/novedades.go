package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
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

	var registroNovedad map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	//fmt.Println(registroNovedad, alertErr, horaRegistro)
	fmt.Println("Ingresa a la función del controlador para post \n")
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroNovedad); err == nil {

		fmt.Println("Aqui se muestra JSON de entrada \n", registroNovedad)

		result, err1 := RegistrarNovedad(registroNovedad)

		//fmt.Println(registroNovedadPost, horaRegistro)
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
	var vacio map[string]interface{}
	//var Result interface{}
	var alerta models.Alert
	//alertas := append([]interface{}{"error"})
	idStr := c.Ctx.Input.Param(":id")
	vigencia := c.Ctx.Input.Param(":vigencia")
	fmt.Println(idStr)

	error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/?query=contrato_id:"+idStr+",vigencia:"+vigencia+"&limit=0", &novedades)
	fmt.Println(beego.AppConfig.String("NovedadesCrudService") + "/novedades_poscontractuales/?query=contrato_id:" + idStr + ",vigencia:" + vigencia + "&limit=0")

	fmt.Println("posicion 1 del vector ", novedades[0], vacio)
	if novedades[0]["TipoNovedad"] != nil {

		fmt.Println("No está vacío", len(novedades))
		for _, novedad := range novedades {

			fmt.Println(novedad["TipoNovedad"])
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
			//fmt.Println(novedadformated)

		}
	} else {
		novedadesformated = []map[string]interface{}{}
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

	//fmt.Println(novedades)

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
	// fmt.Println(errResultado, resultado)
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
	//var resultadoRegistroMongo map[string]interface{}

	switch registroNovedadPost["tiponovedad"] {
	case "59d7965e867ee188e42d8c72":
		//suspensión
		fmt.Println("Novedad de suspensión")
		NovedadPoscontractualPost = models.ConstruirNovedadSuspension(registroNovedadPost)
		fmt.Println(NovedadPoscontractualPost)
	case "59d79683867ee188e42d8c97":
		//cesión
		fmt.Println("Novedad de cesión")
		NovedadPoscontractualPost = models.ConstruirNovedadCesion(registroNovedadPost)
		// fmt.Println(NovedadPoscontractualPost)
	case "59d796ac867ee188e42d8cbf":
		//reinicio
		fmt.Println("Novedad de reinicio")
		NovedadPoscontractualPost = models.ConstruirNovedadReinicio(registroNovedadPost)
		fmt.Println(NovedadPoscontractualPost)
	case "59d79809867ee188e42d8e0d":
		//terminacion anticipada
		fmt.Println("Novedad de terminación anticipada")
		NovedadPoscontractualPost = models.ConstruirNovedadTAnticipada(registroNovedadPost)
		fmt.Println(NovedadPoscontractualPost)
	case "59d7985e867ee188e42d8e64":
		//adición
		fmt.Println("Novedad de adición")
		NovedadPoscontractualPost = models.ConstruirNovedadAdicionPost(registroNovedadPost)
		fmt.Println(NovedadPoscontractualPost)
	case "59d79894867ee188e42d8e9b":
		//prórroga
		fmt.Println("Novedad de prorroga")
		NovedadPoscontractualPost = models.ConstruirNovedadProrrogaPost(registroNovedadPost)
		fmt.Println(NovedadPoscontractualPost)
	case "59d79904867ee188e42d8f02":
		//adicion/prorroga
		fmt.Println("Novedad de adicion/prorroga")
		NovedadPoscontractualPost = models.ConstruirNovedadAdProrrogaPost(registroNovedadPost)
		fmt.Println(NovedadPoscontractualPost)
	}

	if registroNovedadPost["tiponovedad"] == "59d79683867ee188e42d8c97" {
		errRegNovedad = request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/trNovedad/trnovedadpoliza", "POST", &resultadoRegistro, NovedadPoscontractualPost)
	} else {
		errRegNovedad = request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/trNovedad", "POST", &resultadoRegistro, NovedadPoscontractualPost)
	}

	if resultadoRegistro["Status"] == "400" || errRegNovedad != nil {
		fmt.Println("\n entro al error \n")
		fmt.Println(errRegNovedad)
		//fmt.Println(resultadoRegistro["Status"])
		return nil, resultadoRegistro

	} else {
		fmt.Println("\n entro al true \n", resultadoRegistro)
		fmt.Println()

		idRegistroAdmAmazon, error_registroamazon := RegistroAdministrativaAmazon(resultadoRegistro)

		if error_registroamazon == nil {

		}

		fmt.Println(idRegistroAdmAmazon)

		return resultadoRegistro, nil

	}

}

//Función que duplicará los datos de registro de novedades de adición y cesión
func RegistroAdministrativaAmazon(Novedad map[string]interface{}) (idRegistroAdmAmazon int, outputError interface{}) {
	NovedadAmazon := Novedad
	var NovedadGET []map[string]interface{}
	var NovedadAdmAmazonFormatted map[string]interface{}
	var resultadoregistroadmamazon map[string]interface{}
	var errRegNovedad error

	NovedadMap := NovedadAmazon["NovedadPoscontractual"].(map[string]interface{})
	idStrf64 := NovedadMap["Id"].(float64)
	idStr := strconv.FormatFloat(idStrf64, 'f', -1, 64)

	fmt.Println("\n aqui se muestra el ID de la novedad que se acaba de guardar \n", idStr)

	error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/?query=id:"+idStr+"&limit=0", &NovedadGET)
	fmt.Println("Aqui muestro la novedad obtenida mediante GET \n", NovedadGET)

	//Para novedad de adición prorroga
	if NovedadGET[0]["TipoNovedad"].(float64) == 8 {
		NovedadAdmAmazonFormatted = models.FormatAdmAmazonNovedadAdProrroga(NovedadGET)
		urladm := beego.AppConfig.String("AdministrativaAmazonService") + "/novedad_postcontractual"
		errRegNovedad = request.SendJson(urladm, "POST", &resultadoregistroadmamazon, NovedadAdmAmazonFormatted)
		fmt.Println(beego.AppConfig.String("AdministrativaAmazonService") + "/novedad_postcontractual")
		formatdata.JsonPrint(resultadoregistroadmamazon)
		fmt.Println("Aquí se muestra el resultado del post a AdmAzamon \n", errRegNovedad, resultadoregistroadmamazon)
	}
	//errRegNovedad = request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/trNovedad/trnovedadpoliza", "POST", &resultadoRegistro, NovedadPoscontractualPost)

	fmt.Println("Aqui se muestra la traducción de la novedad para replica en AdmAmazon \n", NovedadAdmAmazonFormatted, error)
	formatdata.JsonPrint(NovedadAdmAmazonFormatted)
	return 0, nil
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
