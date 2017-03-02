package controllers

import (
	"github.com/astaxie/beego"
    
	"mobuy/models"
	
)

type HomeController struct{
    beego.Controller
}
func (li *HomeController) Get()  {
    
    islogin,user:= models.CheckAccountbyCookie(li.Ctx)
    li.Data["IsLogin"] = islogin
    if islogin {
        li.Data["User"]=user
    }
    categories,err:=models.FindNoteCategory()
    if err!=nil {
        beego.Error(err)
    }
    li.Data["Categories"]=categories


    li.TplName="html/home.html"
}