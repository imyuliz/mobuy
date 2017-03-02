package models

import (
	"time"
)

// Message 留言表
type Message struct{
    //留言编号
    MessageId   int64       `orm:"pk;auto"`
    //留言内容
    MessContent string      `orm:null`
    // 留言时间
    MessDate    time.Time   `orm:"auto_now_add;type(datetime);index"`
    //用户编号  外键  暂可空
    UserId      int64       `orm:null`
    //留言所属笔记
    NoteId      int64       `orm:null`
}