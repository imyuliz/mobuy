package controllers

import (
	"github.com/astaxie/beego"
)

type AddressController struct{
    beego.Controller
}
func (li *AddressController) Get()  {
    li.TplName="person/address.html"
}