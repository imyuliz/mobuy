package models

// FeedbackInfo 信息反馈表
type FeedbackInfo   struct{
    // 反馈编号
    FeedbackId      int64       `orm:"pk;auto"`
    // 用户编号 外键
    UserId          int64       `orm:"null"`
    // 管理员编号
    ManagerId       int64       `orm:"null"`
    // 反馈信息
    FeedbackInfo    string
    // 反馈回答结果
    feedbackResu    string      `orm:"null"`
}