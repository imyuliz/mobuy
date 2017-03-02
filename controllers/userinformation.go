package controllers

import (
	"github.com/astaxie/beego"
	"mobuy/models"
    "regexp"
	
	"strings"
	"strconv"
	
)

type UserInformationController struct{
    beego.Controller
}

func (li *UserInformationController) View()  {
        islogin,user:=models.CheckAccountbyCookie(li.Ctx)
        li.Data["IsLogin"] = islogin
        //  fmt.Println(islogin)
    if !islogin {
        li.Redirect("/login",302)
        
    }else{
       
         li.Data["User"]=user
        userid:=li.Ctx.Input.Param(":uid")
         
    //    err,user:= models.FindUserByUserId(userid)
        err,address:=models.FindAddressByUserId(userid)
       if err!=nil {
           beego.Error(err)
       }else{
           li.Data["Adress"]=address
               
           
       }
    }
    li.TplName="person/information.html"
}

func (li *UserInformationController) Modify(){
    beego.ReadFromRequest(&li.Controller)
     islogin,user:=models.CheckAccountbyCookie(li.Ctx)
        li.Data["IsLogin"] = islogin
        // userid:=li.Input().Get("userid")
        //    err,user:= models.FindUserByUserId(userid)
         usersex:=user.UserSex
           li.Data["User"]=user
           

             if (usersex)=="男" {
       li.Data["IsMan"]=true
   }else if (usersex)=="女" {
        
       li.Data["IsWoman"]=true
   }else{
       li.Data["BaoMi"]=true
   }
        li.TplName="person/editinformation.html"
}
// UpdateImag 修改头像
func (li *UserInformationController) UpdateImag()  {
    flash := beego.NewFlash()
    f,h,err:=li.GetFile("userimage")
    defer f.Close()
    if err!=nil {
        flash.Error("上传失败")
        flash.Store(&li.Controller)
        li.Redirect("/userinformation/modify",302)
    }else{
        li.SaveToFile("userimage", "static/upload/img/" + h.Filename)
        _,user:=models.CheckAccountbyCookie(li.Ctx)
        user.UserImage="/static/upload/img/" + h.Filename
        models.UpdateUser(&user)
        flash.Success("上传成功")
        flash.Store(&li.Controller)
       li.Redirect("/userinformation/modify",302)
    }

}
// UpdateUser 修改用户的基本信息
func (li *UserInformationController) UpdateUser()  {
    flash := beego.NewFlash()
    _,user:=models.CheckAccountbyCookie(li.Ctx)
    nickname:=li.Input().Get("nickname")
    usersex:=li.Input().Get("usersex")
   userbirth:= li.Input().Get("userbirth")
   usertel:= li.Input().Get("usertel")
   useremail:= li.Input().Get("useremail")
   if len(useremail)>0 {
       ok, _ := regexp.MatchString("^([a-z0-9A-Z]+[-|_|\\.]?)+[a-z0-9A-Z]@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\.)+[a-zA-Z]{2,}$", useremail)
       if !ok {
           flash.Error("请输入正确的邮箱地址")
           flash.Store(&li.Controller)
           li.Redirect("/userinformation/modify",302)
           return
       }else{
        user.UserEmail=useremail
       }
    }
    if len(usertel)>0 {
        if len(usertel)!=11 {
           flash.Error("请输入正确的电话号码")
           flash.Store(&li.Controller)
           li.Redirect("/userinformation/modify",302)
           return
    }else{
        user.UserTel=usertel
    }
    }
    if len(nickname)>0 {
        if len(nickname)>7&&len(nickname)<3 {
           flash.Error("请输入正确的昵称")
           flash.Store(&li.Controller)
           li.Redirect("/userinformation/modify",302)
           return
    }else{
        user.NickName=nickname
    }
}
    if  userbirth!=""||len(userbirth)>0 {
        user.UserBirth=userbirth
    }
    if len(usersex)>0 {
        user.UserSex=usersex
    }
    models.UpdateUser(&user)
    flash.Success("基本信息更新成功")
    flash.Store(&li.Controller)
    li.Redirect("/userinformation/modify",302)
}

func (li *UserInformationController) UpdatePwd()  {
    flash := beego.NewFlash()
    _,user:=models.CheckAccountbyCookie(li.Ctx)
    userpwd:=li.Input().Get("userpwd")
    surepwd:=li.Input().Get("surepwd")
    if len(userpwd)>0 &&len(surepwd)>0 {
        if strings.EqualFold(userpwd,surepwd) {
            userpwd=models.EncodeMessageMd5(userpwd)
            user.UserPwd=userpwd
            models.UpdateUser(&user)
            flash.Success("密码修改成功")
            flash.Store(&li.Controller)
            li.Redirect("/userinformation/modify",302)
        }else{
            flash.Error("密码更新失败，请重新输入...")
            flash.Store(&li.Controller)
            li.Redirect("/userinformation/modify",302)
        }
    }

}

func (li *UserInformationController) UserSafety(){
    islogin,user:=models.CheckAccountbyCookie(li.Ctx)
        li.Data["IsLogin"] = islogin
        //  fmt.Println(islogin)
    if !islogin {
        li.Redirect("/login",302)
        
    }else{       
         li.Data["User"]=user
        // userid:=li.Ctx.Input.Param(":uid")

}
        li.TplName="person/safety.html"
}

func (li *UserInformationController) Address()  {
    beego.ReadFromRequest(&li.Controller)
     islogin,user:=models.CheckAccountbyCookie(li.Ctx)
        li.Data["IsLogin"] = islogin
        //  fmt.Println(islogin)
    if !islogin {
        li.Redirect("/login",302)
        
    }else{       
        userid:=user.UserId
        userids:=strconv.Itoa(userid)
        // userid:=li.Ctx.Input.Param(":uid")
        err,location:=models.FindAddressByUserId(userids)
        if err!=nil {
            beego.Error()
        }
        li.Data["UserLocation"]=location
         li.Data["User"]=user
        // userid:=li.Ctx.Input.Param(":uid")
}
    li.TplName="person/address.html"
}

func (li *UserInformationController) AddAddress()  {
     _,user:=models.CheckAccountbyCookie(li.Ctx)
     userid:=user.UserId
     userids,_:=strconv.ParseInt(strconv.Itoa(userid),10,64)
    flash := beego.NewFlash()
   consignee := li.Input().Get("consignee")
   consigneetel := li.Input().Get("consigneetel")
   province:=li.Input().Get("province")
    city:=li.Input().Get("city")
     county:=li.Input().Get("county")
     locationdeslocationdes:= li.Input().Get("locationdes")
     if len(consignee)==0||len(consigneetel)==0||len(province)==0||len(city)==0||len(county)==0||len(locationdeslocationdes)==0 {
         flash.Error("请输入完整的信息。")
        
     }else if len(consigneetel)!=11 {
         flash.Error("请输入正确的电话号码")
        
     }else{

        Province    :=models.FindProIdByProname(province)
        City        :=models.FindCityIdByCityName(city)
        County    :=models.FindCountyIdByCountyname(county)
        proId       :=Province.ProId
        cityId      :=City.CityId
        countyId:=County.CountyId
        // proid       :=models.Int64ToInt(proId)
        // cityid      := models.Int64ToInt(cityId)
        // countyid    :=models.Int64ToInt(countyId)


        // alllocationid := models.FindThreeIdByThreeName(province,city,county)
        // proId:= alllocationid.ProId
        // fmt.Println(proId)
        // cityId:=alllocationid.CityId
        // countyId:=alllocationid.CountyId
        err:=models.InsertAddress(userids,proId,cityId, countyId,locationdeslocationdes,consignee,consigneetel)
        if err!=nil {
            flash.Error("添加失败")
        }else{
            flash.Success("添加成功")
        }

    } 
        flash.Store(&li.Controller)
        li.Redirect("/userinformation/address",302)
     
   
}