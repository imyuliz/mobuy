package routers

import (
	"mobuy/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&controllers.HomeController{})
	beego.Router("/login",&controllers.LoginController{})
	// beego.Router("/login/:exit",&controllers.LoginController{})
	beego.Router("/register",&controllers.RegisterController{})
	beego.Router("/introducation",&controllers.IntroducationController{})
	//个人页面路由
	beego.Router("/person",&controllers.PersonController{})
	beego.Router("/person/:id([0-9]+)",&controllers.PersonController{},"GET:View")
	//用户地址路由
	// beego.Router("/useradress",&controllers.AddressController{})
	// beego.Router("/useradress/:uid([0-9]+)",&controllers.AddressController{},"GET:View")
	//用户安全路由
	// beego.Router("/safety",&controllers.SafetyController{})
	//用户信息路由
	// beego.Router("/userinformation",&controllers.UserInformationController{})
	beego.Router("/userinformation/:uid([0-9]+)",&controllers.UserInformationController{},"GET:View")
	
	//用户路由
	beego.Router("/userinformation/modify",&controllers.UserInformationController{},"GET:Modify")
	beego.Router("/userinformation/updateimag",&controllers.UserInformationController{},"POST:UpdateImag")
	beego.Router("/userinformation/updateuser",&controllers.UserInformationController{},"POST:UpdateUser")
	beego.Router("/userinformation/updatepwd",&controllers.UserInformationController{},"POST:UpdatePwd")

	// 用户安全
	beego.Router("/userinformation/usersafety/:uid([0-9]+)",&controllers.UserInformationController{},"GET:UserSafety")

	//用户地址路由
	beego.Router("/userinformation/address",&controllers.UserInformationController{},"GET:Address")
	beego.Router("/userinformation/addaddress",&controllers.UserInformationController{},"POST:AddAddress")


	// beego.Router("/userinformation/modify/:uid([0-9]+)",&controllers.UserInformationController{},"GET:Modify")
	beego.Router("/applicationrecovery",&controllers.ApplicationRecoveryController{})
	beego.Router("/bill",&controllers.BillController{})
	beego.Router("/billlist",&controllers.BillListContorller{})
	
	
	
}
