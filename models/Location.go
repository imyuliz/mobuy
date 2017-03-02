package models

// location 地址表
type Location struct{
    //地址主键
    LocationId  int64       `orm:"pk;auto"`
    //用户编号   暂时允许为空 外键
    UserId      int64       `orm:"null"`
    //省份或者直辖市    暂时允许为空     外键
    UserProv    int         `orm:"null"`
    //市  暂时允许为空     外键
    UserCity    int         `orm:"null"`
    //县、区    暂时允许为空     外键
    UserArea    int         `orm:"null"`
    //详细地址   暂时允许为空     外键
    LocationDes string      `orm:"null"`
    //收件人
    Consignee   string      `orm:"size(11);null"`    
    //收件人电话
    ConsigneeTel    string  `orm:"size(11);null"`

}

// Province 省及直辖市表
type Province struct{
    //省会/直辖市编号
    ProId       int         `orm:"pk;auto"`
    //省名
    ProName     string      `orm:"null;size(30);index"`
}
// City 市级表
type City struct{
    //市级编号
    CityId      int         `orm:"pk;auto"`
    // 市级名字
    CityName    string      `orm:"null;size(30);index"`
    //市属省会，直辖市编号 -外键  暂时允许为空
    ProId       int         `orm:"null"`
}
// County 县级表
type County struct{
    //县编号
    CountyId    int         `orm:"pk;auto"`
    //县名字
    CountyName  string      `orm:"null;size(30);index"`
    //县所属市  外键 暂可为空
    CityId      int         `orm:"null"`
}
