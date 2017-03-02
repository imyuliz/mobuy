package models

import (
	"time"
)
// Notice 公告表
type Notice struct{
    // 公告编号
    NoticeId    int64           `orm:"pk;auto;"`
    // 公告类型 外键
    NoticeType  int64           `orm:"null;"`
    // 公告名称
    NoticeName  string          `orm:"size(50);null"`
    // 发布时间
    ReleaTime   time.Time       `orm:"auto_now_add;type(datetime);index"`
    // 发布的管理员编号
    ManagerId   int             `orm:"null"`
    //截止时间
    StopTime    time.Time       `orm:"auto_now;type(datetime)"`
}

// NoticeType 公告类型
type NoticeType struct{
    // 公告类型编号
    NoticeTypeId    int         `orm:"pk;auto"`
    // 公告类型名称
    NoticeName      string      `orm:"size(50);null"`
}