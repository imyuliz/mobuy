package models

import (
	"time"
)

// AdInfo 广告信息表
type AdInfo     struct{
    // 广告编号
    AdId            int64       `orm:"pk;auto"`
    // 广告名字
    AdName          string      `orm:"size(30)"`
    // 广告所属商家
    AdSeller        string      `orm:"size(30)"`
    // 发布时间
    ReleaseTime     time.Time   `orm:"auto_now_add;type(datetime);index"`
    // 截止时间
    StopTime        time.Time   `orm:"auto_now;type(datetime);null"`
    // 广告图片地址一
    AdImgAdrone     string      `orm:"null"`
    // 广告图片地址二
    AdImgAdrtwo     string      `orm:"null"`
    // 广告图片地址三
    AdImgAdrthre     string     `orm:"null"`
}