package models

// Collection 收藏表
type Collection struct{
    //收藏编号
    CollectionId    int64   `orm:"pk;auto"`
    //笔记编号 外键 暂空
    NoteId          int64   `orm:"null"`
    //用户编号 外键 暂空
    UserId          int64   `orm:"null"`
}