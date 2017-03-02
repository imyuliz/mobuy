package controllers

import (
	"github.com/astaxie/beego"
    "mobuy/models"
)

type  RegisterController struct{
    beego.Controller
}

func (li *RegisterController) Get()  {
    beego.ReadFromRequest(&li.Controller)
    li.TplName="html/register.html"
}
// user register
func (li *RegisterController) Post(){
    flash := beego.NewFlash()
    useraccount := li.Input().Get("useraccount")
    userpwd := li.Input().Get("userpwd")
    surepwd := li.Input().Get("surepwd")
    //判断是否勾选
    agree:=li.Input().Get("agree")=="on"
    var err error
    if flag,_ := models.FindUserAccByUserAccount(useraccount);flag==false {
        if (len(useraccount)!=0)&&(userpwd==surepwd)&&(len(userpwd)!=0)&&(len(surepwd)!=0) {
            if agree {
                userpwd=models.EncodeMessageMd5(userpwd)
                err= models.AddUser(useraccount,userpwd)
                if err!=nil {
                    beego.Error(err)
                }
                //设置浏览器的最大生命周期和cookie
                maxAge := 1<<31 - 1 //默认浏览器的生命周期为： 0
                //参数列表      1、名     2、值 3、最大时间  4、路径
                li.Ctx.SetCookie("useraccount",useraccount,maxAge,"/") //账号
                li.Ctx.SetCookie("userpwd",userpwd,maxAge,"/") //账号
                li.Redirect("/",301)
            } else{
                flash.Error("请阅读协议并确认")
                flash.Store(&li.Controller)
                li.Redirect("/register",302)
            }
           
        } else{
               flash.Error("用户名或者密码错误")
                flash.Store(&li.Controller)
                li.Redirect("/register",302)
        }     
    } else{
             flash.Error("用户已存在")
             flash.Store(&li.Controller)
              li.Redirect("/register",302)
    }
    return
}
