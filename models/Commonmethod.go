package models

import (

	"crypto/md5"
    
	"encoding/base64"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	
	"github.com/astaxie/beego/orm"
	"strconv"
)

type UserLocation struct{
	//收件人
	Consignee   	string 
	//收件人电话
	ConsigneeTel    string
	// 省
	ProName			string	
	// 市区		
	CityName		string	
	// 县		
	CountyName		string		
	// 详细地址	
	LocationDes		string			
}

type TempLocation struct{
	
	//省编号
	ProId	int
	// 市编号
	CityId	int
	// 县编号
	CountyId	int
}

// EncodeMessageMd5 把字符串加密后返回一个字符串
func EncodeMessageMd5(msg string) string {
	h := md5.New()
	coding := base64.NewEncoding(beego.AppConfig.String("base64key"))
	h.Write([]byte(msg)) // 需要加密的字符串为 123456
	key := []byte(beego.AppConfig.String("md5key"))
	md5string := h.Sum([]byte(key))

	return coding.EncodeToString(md5string)
}
// CheckAccountbyCookie 通过cookie的账号密码判断是否查询用户
func CheckAccountbyCookie(li *context.Context) (bool,User){
	var zuser User 
	
	//获取cookie的账户
	ck,err:=li.Request.Cookie("useraccount")
	if err!=nil {
		return false,zuser
	}
	//获取cookie的密码
	useraccount:=ck.Value
	ck , err=li.Request.Cookie("userpwd")
	if err!=nil {
		return false,zuser
	}
	userpwd:=ck.Value

	flag ,user := CheckAccount(useraccount,userpwd)
	return flag,user

}

// CheckAccount 用用户账号和密码查出user
func  CheckAccount(useraccount,userpwd string) (bool,User)  {
    o:=orm.NewOrm()
    var user User
    err := o.QueryTable(user).Filter("UserAccount",useraccount).Filter("UserPwd",userpwd).One(&user)
    return err!=orm.ErrNoRows ,user
}
// FindUserByUserId 通过用户编号查出用户信息
func FindUserByUserId(userid string) (error,*User) {
	err,num := StringToInt(userid)
	if err!=nil {
		return err,nil
	}else{
		o:=orm.NewOrm()
		user :=new(User)
		err=o.QueryTable(user).Filter("UserId",num).One(user)
		if err!=nil {
		return err,nil
		}else{
			return nil,user
		}
	}
	
}
// StringToInt 把string装换为int
func StringToInt(args string) (error,int64)  {
	argss,err:= strconv.ParseInt(args,10,64)
	if err!=nil {
		return err,0
	}
		return nil,argss
}
// FindAddressByUserId 通过用户编号查询用户地址
func FindAddressByUserId(userid string) (error,[]*UserLocation) {
	err,num := StringToInt(userid)
	if err!=nil {
		return err,nil
	}else{
		o:=orm.NewOrm()
		userlocation:=make([]*UserLocation, 0)
	// var userlocation []*UserLocation
	// s:=reflect.TypeOf(UserLocation).Elem()
		// userlocation,err:=o.Raw("SELECT location_des,pro_name,city_name,county_name FROM location AS l,province AS p,city AS c,county AS co WHERE l.user_prov=p.pro_id AND l.user_city=c.city_id AND l.user_area=co.county_id AND l.user_id=?",num).Exec()
		o.Raw("SELECT l.location_des,l.consignee,l.consignee_tel,p.pro_name,c.city_name,co.county_name "+
			"FROM location AS l,province AS p,city AS c,county AS co "+
			"WHERE l.user_prov=p.pro_id AND l.user_city=c.city_id AND"+
			" l.user_area=co.county_id AND l.user_id=?",num).QueryRows(&userlocation)
		return nil,userlocation
		}
}
// UpdateUser 更新用户
func UpdateUser(user *User)  {
	o := orm.NewOrm()
	o.Update(user)
}
// FindProIdByProname 通过省的名字查询到其编号
func FindProIdByProname(proname string) Province {
	o:=orm.NewOrm()
	var  province Province
	o.Raw("SELECT pro_id FROM province WHERE pro_name=?",proname).QueryRow(&province)
	return province
	// var proid	sql.Result
	// proid:=o.Raw("SELECT pro_id FROM province WHERE pro_name=?",proname)
	// return proid
}
// FindCityIdByCityName 通过市的名字查询到编号
func FindCityIdByCityName(cityname string)  City {
	o:=orm.NewOrm()
	var city City
	// var cityId sql.Result
	o.Raw("SELECT city_id FROM city WHERE city_name=?",cityname).QueryRow(&city)
	return city
}
// FindCountyIdByCountyname 通过县的名字查询县的编号
func FindCountyIdByCountyname(countyName	string) County {
	o:=orm.NewOrm()
	var county County
	// var countyid sql.Result
	o.Raw("SELECT county_id FROM county WHERE county_name=?",countyName).QueryRow(&county)
	return county
}
// InsertAddress  添加收件人地址
func InsertAddress(userid int64,proid,cityid,countyid int,location,consignee,consigneetel string) error {
	o:=orm.NewOrm()
	relocation:=&Location{
		UserId:userid,
		UserProv:proid,
		UserCity:cityid,
		UserArea:countyid,
		LocationDes:location,
		Consignee:consignee,
		ConsigneeTel:consigneetel,
	}
	_,err:=o.Insert(relocation)
	return err
}

// func FindThreeIdByThreeName(proname,cityname,countyName string) *TempLocation {
// 	o:=orm.NewOrm()
// 	// var  TempLocation
// 	templocation:= new(TempLocation)
// 	r:=o.Raw("SELECT p.pro_id,c.city_id,co.county_id "+
// 	"FROM province AS p,city AS c,county AS co "+
// 	"WHERE p.pro_name=? AND c.city_name=? AND co.county_name=?",proname,cityname,countyName).RowsToStruct(templocation,"key","value")
// 	// templocations:=r.Exec()
// 	// return templocation
// }


func Int64ToInt(num int64)  int {
	
	// strconv.FormatInt(num,10)
	tonum,_:= strconv.Atoi(strconv.FormatInt(num,10))

	return tonum
}