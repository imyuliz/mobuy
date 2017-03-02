package models

// Reply 回复表
type Reply  struct{
    //回复编号
    ReplyId         int64   `orm:"pk;auto"`
    //评论编号 外键 暂空
    CommentId       int64   `orm:"null"`
    //回复内容
    ReplyContent    string  `orm:"null"`
}