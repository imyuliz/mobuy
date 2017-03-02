package models

import (
	"time"
)
// Activity 活动
type Activity struct{
    // 活动编号
    ActivityId      int64       `orm:"pk;auto"`
    // 开始时间
    ActivityTime    time.Time    `orm:"auto_now_add;type(datetime);index"`
    // 截止时间
    StopTime        time.Time    `orm:"auto_now;type(datetime);null"`
    // 题库编号
    RewardId        int64        `orm:"null"`
    // 笔记编号
    NoteId          int64       `orm:"null"`
    // 活动标语
    ActivityContent string      `orm:"size(50);null"`
}
