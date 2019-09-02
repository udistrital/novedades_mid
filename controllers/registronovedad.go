package controllers

import (
	"github.com/astaxie/beego"
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

}
