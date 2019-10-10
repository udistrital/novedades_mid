package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

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

	var registroNovedad map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	//fmt.Println(registroNovedad, alertErr, horaRegistro)
	fmt.Println("Ingresa a la función del controlador para post \n")
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroNovedad); err == nil {

		//fmt.Println(registroNovedad)

		result, err1 := RegistrarNovedad(registroNovedad)

		//fmt.Println(registroNovedadPost, horaRegistro)
		if err == nil {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = result
		} else {
			alertErr.Type = "error"
			alertErr.Code = "400"
			alertas = append(alertas, err1)
			alertErr.Body = alertas
		}

	} else {
		alertErr.Type = "error"
		alertErr.Code = "400"
		alertas = append(alertas, err.Error())
		alertErr.Body = alertas
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
// @router /:id [get]
func (c *NovedadesController) GetOne() {
	// var resultado map[string]interface{}
	// var Result interface{}
	// var alerta models.Alert
	// alertas := append([]interface{}{"error"})
	// idStr := c.Ctx.Input.Param(":id")

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

//RegistrarNovedadMongo Función para registrar la novedad en mongodb
func RegistrarNovedad(novedad map[string]interface{}) (status interface{}, outputError interface{}) {

	registroNovedadPost := make(map[string]interface{})
	registroNovedadPost = novedad

	var NovedadPoscontractualPost map[string]interface{}
	var resultadoRegistro map[string]interface{}
	//var resultadoRegistroMongo map[string]interface{}

	switch registroNovedadPost["tiponovedad"] {
	case "59d7965e867ee188e42d8c72":
		//suspensión
		fmt.Println("Novedad de suspensión")
	case "59d79683867ee188e42d8c97":
		//cesión
		fmt.Println("Novedad de cesión")
	case "59d796ac867ee188e42d8cbf":
		//reinicio
		fmt.Println("Novedad de reinicio")
	case "59d797aa867ee188e42d8db6":
		//liquidación
		fmt.Println("Novedad de liquidación")
	case "59d79809867ee188e42d8e0d":
		//terminacion anticipada
		fmt.Println("Novedad de terminación anticipada")
	case "59d7985e867ee188e42d8e64":
		//adición
		fmt.Println("Novedad de adición")
		NovedadPoscontractualPost = ConstruirNovedadAdicionPost(registroNovedadPost)
		fmt.Println(NovedadPoscontractualPost)
	case "59d79894867ee188e42d8e9b":
		//prórroga
		fmt.Println("Novedad de prorroga")
		NovedadPoscontractualPost = ConstruirNovedadProrrogaPost(registroNovedadPost)
		fmt.Println(NovedadPoscontractualPost)
	case "59d79904867ee188e42d8f02":
		//adicion/prorroga
		fmt.Println("Novedad de adicion/prorroga")
	}

	errRegNovedad := request.SendJson("http://"+beego.AppConfig.String("NovedadesCrudService")+"/v1/trNovedad", "POST", &resultadoRegistro, NovedadPoscontractualPost)

	if errRegNovedad != nil {
		fmt.Println("\n entro al error \n")
		return nil, NovedadPoscontractualPost

	} else {
		fmt.Println("\n entro al true \n")
		return NovedadPoscontractualPost, nil

	}

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

func ConstruirNovedadAdicionPost(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadAdicion := make(map[string]interface{})
	NovedadAdicion = novedad

	NovedadAdicionPost := make(map[string]interface{})
	contratoid, _ := strconv.ParseInt(NovedadAdicion["contrato"].(string), 10, 32)
	numerocdpid, _ := strconv.ParseInt(NovedadAdicion["numerocdp"].(string), 10, 32)
	numerosolicitudentero := NovedadAdicion["numerosolicitud"].(float64)
	numerosolicitud := strconv.FormatFloat(numerosolicitudentero, 'f', -1, 64)
	vigencia, _ := strconv.ParseInt(NovedadAdicion["vigencia"].(string), 10, 32)
	vigenciacdp, _ := strconv.ParseInt(NovedadAdicion["vigencia"].(string), 10, 32)

	fmt.Println(NovedadAdicion["contrato"], NovedadAdicion["numerocdp"], NovedadAdicion["numerosolicitud"], NovedadAdicion["vigencia"], NovedadAdicion["vigencia"])
	fmt.Println("\n", contratoid, numerocdpid, numerosolicitud, vigencia, vigenciacdp, "\n")

	NovedadAdicionPost["NovedadPoscontractual"] = map[string]interface{}{
		"Aclaracion":        nil,
		"Activo":            true,
		"ContratoId":        contratoid,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"Motivo":            NovedadAdicion["motivo"],
		"NumeroCdpId":       numerocdpid,
		"NumeroSolicitud":   numerosolicitud,
		"Observacion":       nil,
		"TipoNovedad":       6,
		"Vigencia":          vigencia,
		"VigenciaCdp":       vigenciacdp,
	}

	fechas := make([]map[string]interface{}, 0)

	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadAdicion["fechasolicitud"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 7,
		},
	})
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadAdicion["fechaadicion"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 1,
		},
	})

	NovedadAdicionPost["Fechas"] = fechas

	propiedades := make([]map[string]interface{}, 0)
	propiedades = append(propiedades, map[string]interface{}{
		"Activo":            true,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoPropiedad": map[string]interface{}{
			"Id": 6,
		},
		"propiedad": NovedadAdicion["valoradicion"],
	})

	propiedades = append(propiedades, map[string]interface{}{
		"Activo":            true,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoPropiedad": map[string]interface{}{
			"Id": 2,
		},
		"propiedad": NovedadAdicion["cesionario"],
	})

	NovedadAdicionPost["Propiedad"] = propiedades

	fmt.Println(NovedadAdicionPost)

	return NovedadAdicionPost
}

func ConstruirNovedadProrrogaPost(novedad map[string]interface{}) (novedadformatted map[string]interface{}) {
	NovedadProrroga := make(map[string]interface{})
	NovedadProrroga = novedad

	NovedadProrrogaPost := make(map[string]interface{})
	contratoid, _ := strconv.ParseInt(NovedadProrroga["contrato"].(string), 10, 32)
	numerocdpid, _ := strconv.ParseInt(NovedadProrroga["numerocdp"].(string), 10, 32)
	numerosolicitudentero := NovedadProrroga["numerosolicitud"].(float64)
	numerosolicitud := strconv.FormatFloat(numerosolicitudentero, 'f', -1, 64)
	vigencia, _ := strconv.ParseInt(NovedadProrroga["vigencia"].(string), 10, 32)
	vigenciacdp, _ := strconv.ParseInt(NovedadProrroga["vigencia"].(string), 10, 32)

	NovedadProrrogaPost["NovedadPoscontractual"] = map[string]interface{}{
		"Aclaracion":        nil,
		"Activo":            true,
		"ContratoId":        contratoid,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"Motivo":            NovedadProrroga["motivo"],
		"NumeroCdpId":       numerocdpid,
		"NumeroSolicitud":   numerosolicitud,
		"Observacion":       nil,
		"TipoNovedad":       7,
		"Vigencia":          vigencia,
		"VigenciaCdp":       vigenciacdp,
	}

	fechas := make([]map[string]interface{}, 0)

	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadProrroga["fechasolicitud"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 7,
		},
	})
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             "2019-10-08T15:43:51.710Z",
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 1,
		},
	})
	fechas = append(fechas, map[string]interface{}{
		"Activo":            true,
		"Fecha":             NovedadProrroga["fechaprorroga"],
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoFecha": map[string]interface{}{
			"Id": 4,
		},
	})

	NovedadProrrogaPost["Fechas"] = fechas

	propiedades := make([]map[string]interface{}, 0)
	propiedades = append(propiedades, map[string]interface{}{
		"Activo":            true,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoPropiedad": map[string]interface{}{
			"Id": 6,
		},
		"propiedad": NovedadProrroga["valoradicion"],
	})

	propiedades = append(propiedades, map[string]interface{}{
		"Activo":            true,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoPropiedad": map[string]interface{}{
			"Id": 5,
		},
		"propiedad": NovedadProrroga["tiempoprorroga"],
	})

	propiedades = append(propiedades, map[string]interface{}{
		"Activo":            true,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Id":                0,
		"IdNovedadesPoscontractuales": map[string]interface{}{
			"Id": nil,
		},
		"IdTipoPropiedad": map[string]interface{}{
			"Id": 2,
		},
		"propiedad": NovedadProrroga["cesionario"],
	})

	NovedadProrrogaPost["Propiedad"] = propiedades

	fmt.Println(NovedadProrrogaPost)

	return NovedadProrrogaPost
}
