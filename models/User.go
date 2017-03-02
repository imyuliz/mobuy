package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	
	
)


// User 用户表
type User struct{
    //用户编号
    UserId          int     `orm:"pk;auto"`
    //用户昵称
    NickName    string      `orm:"size(30);index"`
    //用户账户
    UserAccount string      `orm:"size(30);index"`
    //用户头像
    UserImage   string      `orm:"null"`
    //用户密码
    UserPwd     string     
    //用户邮箱
    UserEmail   string      `orm:"null"`
    // 用户支付账户
    PayAccount  string      `orm:"null;index"`
    //用户自我介绍
    SelfTro     string      `orm:"type(text);null"`
    //用户电话
    UserTel     string       `orm:"null"`
    //用户积分
    UserScore   int64       `orm:"default(0)"`
    //用户所在学校
    UserSchool  string      `orm:"null"`
    //用户性别
    UserSex     string      `orm:"size(5);null"`
    //用户生日
    UserBirth   string      `orm:"size(30);null"`
    //用户核实身份照片
    UserPicture string      `orm:"null"`
    //用户创建时间
    RegisterTime time.Time  `orm:"auto_now_add;type(datetime);null"`
}

//UserIdentity 用户核实身份图片表
type UserIdentity struct{
    //主键
    Id          int64       `orm:"pk;auto"`
    //身份照片一
    IdenPircone string      `orm:"null"`
    //身份照片二
    IdenPirctwo string      `orm:"null"`
}

//Death 用户身份审核表
type Death struct{
    //审核表编号
    DeathId     int64       `orm:"pk;auto"`
    //管理员审核评价
    DeathDes    string      `orm:"null"`
    //管理员审核结果
    DeathResult string      `orm:"size(10)"`
}
//UserToDeath 审核表和用户表的中间表即 用户身份审核表
type UserToDeath struct{
    //关系编号
    RelId       int64       `orm:"pk;auto"`
    //用户编号  暂时允许为空 
    UserId      int64       `orm:"null"`
    //审核编号  暂时允许为空
    DeathId     int64       `orm:"null"`
}
// FindUserAccByUserAccount  通过useraccount查找User
func FindUserAccByUserAccount(useraccount string) (bool , User)  {
    // o:=orm.NewOrm()
    // user := new(User)
    // qs:=o.QueryTable("user")
    // err:= qs.Filter("user_account",).One(user)
    // userAccount:=user.UserAccount
    // if err!=nil {
    //     return ""
    // }
    // return userAccount
    o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("user_account", useraccount).One(&user)
	return err != orm.ErrNoRows, user

}
// AddUser 添加用户
func AddUser(useraccount ,userpwd string)  error {
    o:=orm.NewOrm()
    user:=&User{
        UserAccount:useraccount,
        UserPwd:userpwd,
    }
    _,err := o.Insert(user)
        return err
}


