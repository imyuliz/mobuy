package models

// Appeal 投诉表
type Appeal struct{
    // 投诉编号
    AppealId        int64   `orm:"pk;auto"`
    // 投诉内容
    AppealContent   string  
    // 处理结语
    AppealAppreal   string  `orm:"null"`
    // 是否处理
    AppealDeal      string  `orm:"size(20);null"`
}
// UserApeal 用户投诉关系
type UserApeal  struct{
    // 关系编号
    UserApealId     int64   `orm:"pk;auto"`
    // 投诉编号 外键 暂空
    AppealId        int64   `orm:"null"`
    // 用户编号 外键 暂空
    UserId          int64   `orm:"null"`
}