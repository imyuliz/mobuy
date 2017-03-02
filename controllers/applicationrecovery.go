package controllers

import (
	"github.com/astaxie/beego"
)
// ApplicationRecoveryController 回收申请
type ApplicationRecoveryController struct{
    beego.Controller
}

func (li *ApplicationRecoveryController) Get()  {
    li.TplName="person/applicationrecovery.html"
}