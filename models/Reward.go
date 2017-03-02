package models

import (
	"time"
)


// Reward 智库悬赏表
type Reward struct{
    // 悬赏编号
    RewardId        int64   `orm:"pk;auto"`

    // 回答用户 外键
    UserId          int64   `orm:"null"`
    // 悬赏科目 外键
    RewardSub       int     `orm:"null"`
    // 难度系数 外键 
    RewardDiff      int     `orm:"null"`
    // 适用人群 外键
    SuitPeo         int     `orm:"null"`
    // 悬赏状态 外键
    RewardStus      int     `orm:"null"`
    // 悬赏问题
    RewardQue       string  `orm:"null;index"`
    // 悬赏价格
    RewardPrice     float64 `orm:"digits(10);decimals(1)"`  
    // 发布时间
    ReleaseTime     time.Time    `orm:"auto_now_add;type(datetime);index"`
    // 积分值
    MemberPoint     int64
    // 截止时间
    CloseTime       time.Time    `orm:"auto_now;type(datetime)"`
    // 题目来源
    ReferanceBook   string      `orm:"size(50);null"`
    // 点赞数
    RewardValue     int64       `orm:"null"`
    // 问题图片路径
    PicturAdress    string      `orm:"null"`
    // 回答图片路径
    AnswerAdess     string      `orm:"null"`
}
// RewardRel 发布用户和智库关系表
type RewardRel  struct{
    // 关系编号
    RewardRelId     int64       `orm:"pk;auto"`
    // 用户编号 外键 
    UserId          int64       `orm:"null"`
    // 悬赏编号 外键
    RewardId        int64       `orm:"null"`
}   

// DiffDegree 难度系数表
type DiffDegree struct{
    // 难度编号
    DiffId          int         `orm:"pk;auto"`
    // 难度级别
    DiffLevel       int        `orm:"index"`  
}
// RewardStus 悬赏状态表
type RewardStus struct{
    // 状态编号
    RewardStusId    int         `orm:"pk;auto"`
    // 状态名称
    StusName        string      `orm:"size(10)"`
}