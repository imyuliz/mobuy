package controllers

import (
	"github.com/astaxie/beego"
	"mobuy/models"
	
	
)

type LoginController struct{
    beego.Controller
}

func (li *LoginController) Get()  {
    beego.ReadFromRequest(&li.Controller)
    //如果点击了退出登录，情况cookie
    isExit:=li.Input().Get("exit")=="true"
    if isExit {
        //清空账号  -1是清除f
        li.Ctx.SetCookie("useraccount","",-1,"/")
        li.Ctx.SetCookie("userpwd","",-1,"/")
        li.Redirect("/",302)
        return
    }
    li.TplName="html/login.html"
}

func (li *LoginController)  Post(){
    useraccount := li.Input().Get("useraccount")
    userpwd := li.Input().Get("userpwd")
    userpwd=models.EncodeMessageMd5(userpwd)
    autoLogin:=li.Input().Get("autologin")=="on"
    flash := beego.NewFlash()
    if flag,_:= models.CheckAccount(useraccount,userpwd);flag {
        //接下来是设置cookie
		//设置cookie的时间
		maxAge := 0 //默认是浏览器生命周期
		//判断是否自动登录，如果自动登录设置cookie的时间
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		//设置cookie
		//参数列表      1、名     2、值 3、最大时间  4、路径
		li.Ctx.SetCookie("useraccount", useraccount, maxAge, "/") //账号
		li.Ctx.SetCookie("userpwd", userpwd, maxAge, "/")     //密码

		//4、重定向
		li.Redirect("/", 301)
		return
		//为什么要return呢？因为避免再次请求节省资源
    } else{
        flash.Error("账号或者密码错误，请重新输入")
        flash.Store(&li.Controller)
        li.Redirect("/login",302)
         return
    }
   

}