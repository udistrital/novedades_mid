package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/novedades_mid/models"
	"github.com/udistrital/utils_oas/request"
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

	//fmt.Println(registroNovedad, alertErr, horaRegistro)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroNovedad); err == nil {

		//fmt.Println(registroNovedad)

		result, err1 := RegistrarNovedadMongo(registroNovedad)

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

//RegistrarNovedadMongo Función para registrar la novedad en mongodb
func RegistrarNovedadMongo(novedad map[string]interface{}) (status interface{}, outputError interface{}) {
	horaRegistro := time_bogota.Tiempo_bogota()

	registroNovedadPost := make(map[string]interface{})

	registroNovedadPost = novedad

	var resultadoRegistroMongo map[string]interface{}
	//var identificacion map[string]interface{}

	fmt.Println("Ingresa a la función de post \n", registroNovedadPost, horaRegistro)

	errRegNovedadMongo := request.SendJson("http://"+beego.AppConfig.String("NovedadesApiMongoService")+"/v1/novedad", "POST", &resultadoRegistroMongo, registroNovedadPost)

	fmt.Println(resultadoRegistroMongo, errRegNovedadMongo, beego.AppConfig.String("NovedadesApiMongoService"))

	result := resultadoRegistroMongo["Body"]

	//fmt.Println("Aqui tambien es \n", identificacion)
	//resultado, _ := identificacion.([]map[string]interface{})

	//Resultado2 := resultado[0]
	//identificacion = map[string]interface{}{"res": resultadoRegistroMongo["Body"].(map[string]interface{})["Response"]}
	//logs.Info(produccionAcademicaPost)
	//fmt.Println("Aqui es \n", Resultado2)
	if errRegNovedadMongo != nil {
		//return nil, identificacion

		return nil, result

	} else {
		return result, nil
	}

	//fmt.Println("Ingresa a la función de post \n", registroNovedadPost, horaRegistro)

	//return status, outputError
}
