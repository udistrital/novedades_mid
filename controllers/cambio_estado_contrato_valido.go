package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	//"github.com/udistrital/novedades_mid/helpers"
	. "github.com/udistrital/golog"
	"github.com/udistrital/novedades_mid/models"
)

// CambioEstadoContratoValidoController operations for CambioEstadoContratoValido
type CambioEstadoContratoValidoController struct {
	beego.Controller
}

// URLMapping ...
func (c *CambioEstadoContratoValidoController) URLMapping() {
	c.Mapping("ValidarCambioEstado", c.ValidarCambioEstado)
}

// ValidarCambiosEstado ...
// @Title ValidarCambiosEstado
// @Description create ValidarCambiosEstado
// @Success 201 {int} models.EstadoContrato
// @Failure 403 body is empty
// @router / [post]
func (c *CambioEstadoContratoValidoController) ValidarCambioEstado() {

	var estados []models.EstadoContrato //0: actual y 1:siguiente

	reglasbase, err := CargarReglasBase("AdministrativaContratacion")
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}

	m := NewMachine().Consult(reglasbase)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &estados); err == nil {

		if m.CanProve(`estado(` + strings.ToLower(estados[0].NombreEstado) + `,` + strings.ToLower(estados[1].NombreEstado) + `).`) {
			c.Data["json"] = "true"
		} else {
			c.Data["json"] = "false"
		}

	} else {
		c.Data["json"] = err.Error()
		fmt.Println("error1: ", err)
	}

	c.ServeJSON()

}
