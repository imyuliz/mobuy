package models

import (
	"time"
)


// Recovery 回收表
type Recovery   struct{
    // 回收编号
    RecoveryId      int64       `orm:"pk;auto"`
    // 回收标题
    RecoveryTitle   string      `orm:"size(20);null"`
    // 用户编号
    UserId          int64       `orm:"null"`
    // 详细描述
    RecoveryDes     string      `orm:"null"`
    // 所属学科
    BookSubject     int         `orm:null`
    // 所属学校
    BookSchool      int         `orm:"null"`
    // 教材版本
    BookVer         int         `orm:"null"`
    // 回收类型
    RecoveryType    int         `orm:"null"`
    // 申请单价
    AppliPrice      float64     `orm:"digits(5);decimals(2)"`
    // 适合人群
    SuitPeo         int         `orm:"null"`
    // 回收状态
    RecoveryStatus  int         `orm:"null"`
    // 时间
    TimeId          int         `orm:null`
    // 申请时间
    AppliTime       time.Time   `orm:"auto_now_add;type(datetime);index"`
    // 审核管理员编号
    ManagId         int         `orm:"null"`
    // 管理员审核评语
    OfficComment    string      `orm:"size(50);null"`
    // 照片地址一
    LocationAdres1  string      `orm:"null"`
    // 照片地址二
    LocationAdres2  string      `orm:"null"`    
    // 照片地址三
    LocationAdres3  string      `orm:"null"`
}
// RecoveryType 回收类型表 
type RecoveryType   struct{
    // 回收类型编号
    RecoveryTypeId  int         `orm:"pk;auto"`
    // 回收类型名
    RecoveryName    string      `orm:"size(11)"`
}
// RecoveryStatus 回收状态表
type RecoveryStatus  struct{
    // 状态编号
    RecovStatusId  int         `orm:"pk;auto"`
    //回收状态名称
    RecovStaName    string      `orm:"size(11)"`
}