package controllers

import (
	"github.com/astaxie/beego"
)

type IntroducationController struct{
    beego.Controller
}
func (li *IntroducationController) Get()  {
    li.TplName="html/introducation.html"
}