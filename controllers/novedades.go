package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/helpers"
	"github.com/udistrital/novedades_mid/models"
	"github.com/udistrital/novedades_mid/services"
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
	c.Mapping("Patch", c.Patch)
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
	error := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/?query=contrato_id:"+idStr+",vigencia:"+vigencia+",activo:true&limit=0&sortby=FechaCreacion&order=asc", &novedades)

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
// @router /:id/:vigencia [put]
func (c *NovedadesController) Put() {

	idStr := c.Ctx.Input.Param(":id")
	var registroNovedad map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroNovedad); err == nil {

		result, err1 := ActualizarNovedad(idStr)

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

// Delete ...
// @Title Delete
// @Description delete the Novedades
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *NovedadesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	// var registroNovedad map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	result, err1 := EliminarNovedad(idStr)
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

	c.Data["json"] = alertErr
	c.ServeJSON()
}

// Patch ...
// @Title Inactivar novedad por ID
// @Description Inactiva una novedad específica por su ID
// @Param	id		path	string	true	"ID de la novedad"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @router /:id [patch]
func (c *NovedadesController) Patch() {
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		c.Data["json"] = helpers.ErrEmiter(nil, "id vacío")
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	usuario := "CC" + c.Ctx.Input.Header("X-User")
	if usuario == "CC" {
		usuario = "MID"
	}

	result := services.AnularNovedadYRevertirEstado(id, usuario)

	if !result.Success {
		c.Data["json"] = result
		c.Ctx.Output.SetStatus(result.Status)
	} else {
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// RegistrarNovedadMongo Función para registrar la novedad en postgresql
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
		fmt.Println("2: ", NovedadPoscontractualPost)
	case "NP_PRO":
		// prórroga
		fmt.Println("Novedad de prorroga")
		NovedadPoscontractualPost = models.ConstruirNovedadProrrogaPost(registroNovedadPost)
	case "NP_ADPRO":
		// adicion/prorroga
		fmt.Println("Novedad de adicion/prorroga")
		NovedadPoscontractualPost = models.ConstruirNovedadAdProrrogaPost(registroNovedadPost)
	}
	// codTerminada := ""
	codEntramite := ""
	// codEnejecuion := ""
	// var estadoTerminada map[string]interface{}
	// error3 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro?query=TipoParametroId.CodigoAbreviacion:ENOV,CodigoAbreviacion:TERM", &estadoTerminada)
	// if error3 == nil {
	// 	if len(estadoTerminada) != 0 {
	// 		inter := estadoTerminada["Data"].([]interface{})
	// 		data := inter[0].(map[string]interface{})
	// 		idEstado, _ := data["Id"].(float64)
	// 		codTerminada = strconv.FormatFloat(idEstado, 'f', -1, 64)
	// 	}
	// }
	var estadoEntramite map[string]interface{}
	error4 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro?query=TipoParametroId.CodigoAbreviacion:ENOV,CodigoAbreviacion:ENTR", &estadoEntramite)
	if error4 == nil {
		if len(estadoEntramite) != 0 {
			inter := estadoEntramite["Data"].([]interface{})
			data := inter[0].(map[string]interface{})
			idEstado, _ := data["Id"].(float64)
			codEntramite = strconv.FormatFloat(idEstado, 'f', -1, 64)
		}
	}
	// var estadoEnejecucion map[string]interface{}
	// error5 := request.GetJson(beego.AppConfig.String("ParametrosCrudService")+"/parametro?query=TipoParametroId.CodigoAbreviacion:ENOV,CodigoAbreviacion:ENEJ", &estadoEnejecucion)
	// if error5 == nil {
	// 	if len(estadoEnejecucion) != 0 {
	// 		inter := estadoEnejecucion["Data"].([]interface{})
	// 		data := inter[0].(map[string]interface{})
	// 		idEstado, _ := data["Id"].(float64)
	// 		codEnejecuion = strconv.FormatFloat(idEstado, 'f', -1, 64)
	// 	}
	// }

	if registroNovedadPost["tiponovedad"] == "NP_CES" {
		novedad := NovedadPoscontractualPost["NovedadPoscontractual"].(map[string]interface{})
		if novedad["Estado"] == codEntramite {
			errRegNovedad = request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/trNovedad", "POST", &resultadoRegistro, NovedadPoscontractualPost)
		} else {
			errRegNovedad = request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/trNovedad/trnovedadpoliza", "POST", &resultadoRegistro, NovedadPoscontractualPost)
		}
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

// Función que duplicará los datos de registro de novedades de adición y cesión
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

func ActualizarNovedad(id string) (status interface{}, outputError interface{}) {

	var novedad map[string]interface{}
	var resultadoRegistro map[string]interface{}
	err := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/"+id, &novedad)
	if err == nil {
		novedad["Estado"] = "TERMINADA"
		errRegNovedad := request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/"+id, "PUT", &resultadoRegistro, novedad)
		if errRegNovedad == nil {
			fmt.Println("Novedad actualizada!!")
		}
	}
	return nil, nil
}

func EliminarNovedad(id string) (status map[string]interface{}, outputError interface{}) {
	var result map[string]interface{}
	var resultadoNov map[string]interface{}
	var resultPropiedades []map[string]interface{}
	var resultFechas []map[string]interface{}
	var resultPolizas []map[string]interface{}
	var delFechas bool
	var delPropiedad bool
	var delPoliza bool
	if err := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/"+id, &resultadoNov); err == nil {
		idTipo, _ := resultadoNov["TipoNovedad"].(float64)
		codTipo := strconv.FormatFloat(idTipo, 'f', -1, 64)
		err1 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/propiedad?query=IdNovedadesPoscontractuales.Id:"+id, &resultPropiedades)
		if err1 == nil {
			err2 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/fechas?query=IdNovedadesPoscontractuales.Id:"+id, &resultFechas)
			if err2 == nil {
				if len(resultFechas[0]) > 0 {
					for _, fecha := range resultFechas {
						delFechas = false
						idFecha, _ := fecha["Id"].(float64)
						idstr := strconv.FormatFloat(idFecha, 'f', -1, 64)
						err3 := request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/fechas/"+idstr, "DELETE", &result, nil)
						if err3 == nil {
							delFechas = true
							// fmt.Println("Registro de Fecha Eliminado")
						} else {
							return nil, result
						}
					}
				} else {
					delFechas = true
				}
				if len(resultPropiedades[0]) > 0 {
					for _, propiedad := range resultPropiedades {
						delPropiedad = false
						idPropiedad := propiedad["Id"].(float64)
						idstr := strconv.FormatFloat(idPropiedad, 'f', -1, 64)
						err3 := request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/propiedad/"+idstr, "DELETE", &result, nil)
						if err3 == nil {
							delPropiedad = true
							// fmt.Println("Registro de Propiedad Eliminado")
						} else {
							return nil, result
						}
					}
				} else {
					delPropiedad = true
				}
				if codTipo == "2" {
					err3 := request.GetJson(beego.AppConfig.String("NovedadesCrudService")+"/poliza?query=IdNovedadesPoscontractuales.Id:"+id, &resultPolizas)
					if err3 == nil && len(resultPolizas[0]) > 0 {
						for _, poliza := range resultPolizas {
							delPoliza = false
							idPoliza, _ := poliza["Id"].(float64)
							idstr := strconv.FormatFloat(idPoliza, 'f', -1, 64)
							err4 := request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/poliza/"+idstr, "DELETE", &result, nil)
							if err4 == nil {
								delPoliza = true
							} else {
								return nil, result
							}
						}
					} else {
						delPoliza = true
					}
				} else {
					delPoliza = true
				}
				if delFechas && delPropiedad && delPoliza {
					err4 := request.SendJson(beego.AppConfig.String("NovedadesCrudService")+"/novedades_poscontractuales/"+id, "DELETE", &result, nil)
					if err4 == nil {
						fmt.Println("Registro de novedad eliminado, Id: ", id)
						return result, nil
					}
				} else {
					fmt.Println("Error al eliminar")
					return nil, result
				}
			} else {
				return nil, resultFechas
			}
		} else {
			return nil, resultPropiedades
		}
	} else {
		return nil, resultadoNov
	}
	errorMap := map[string]interface{}{"funcion": "/EliminarNovedad", "err": "La función de borrado no hizo nada (else sin return)"}
	return nil, errorMap
}
