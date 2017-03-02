package models

import (
	"time"
)

// Comment 评价表
type Comment struct{
    //评价编号
    CommentId       int64    `orm:"pk;auto"`
    //笔记编号 外键
    NoteId          int64   `orm:"null"`
    //用户编号  外键
    UserId          int64   `orm:"null"`
    //评价内容
    ComContent      string  `orm:"null;;index"`
    //评价时间
    CommentTime     time.Time   `orm:"type(datetime);auto_now_add"`
    //信息准确度
    Precise         float64  `orm:"digits(2);decimals(1)"`  
    //总体评价
    Evaluation      float64  `orm:"digits(2);decimals(1)"`  
    //评价照片1
    ComPhoto1st     string   `orm:"null"`
    //评价照片2
    ComPhoto2       string   `orm:"null"`
    //评价照片3
    ComPhoto3       string   `orm:"null"`
}