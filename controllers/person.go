package controllers

import (
	"github.com/astaxie/beego"
    
	"mobuy/models"
)

type PersonController struct{
    beego.Controller
}

func (li *PersonController) Get()  {
   
        li.Redirect("/login",302)
        return
    
}

func (li *PersonController) View()  {
    islogin,user:=models.CheckAccountbyCookie(li.Ctx)
        li.Data["IsLogin"] = islogin

    if !islogin {
        li.Redirect("/login",302)
        return
    }else{
        // userid:=li.Input().Get("uid")
    //    err,user:= models.FindUserByUserId(userid)
       
           li.Data["User"]=user
       
    }
    li.TplName="person/index.html"
}