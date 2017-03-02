package controllers

import (
	"github.com/astaxie/beego"
)
// BillController 个人账单
type BillController struct{
    beego.Controller
}
func (li *BillController) Get()  {
    li.TplName="person/bill.html"
}