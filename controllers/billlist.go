package controllers

import (
	"github.com/astaxie/beego"
)

type BillListContorller struct{
    beego.Controller
}

func (li *BillListContorller) Get()  {
    li.TplName="person/billlist.html"
}