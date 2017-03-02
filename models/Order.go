package models

import (
	"time"
)

// Oder 订单表
type Oder struct{
    //订单编号
    OrderId     int64       `orm:"pk;auto"`
    //笔记编号 外键 暂空
    NoteId      int64       `orm:"null"`
    // 用户编号 外键 暂空
    UserId      int64       `orm:"null"`
    //订单状态
    OrderStatus string      `orm:"size(20);index"`
    //订单价格
    OrderMoney  float64     `orm:"digits(10);decimals(2)"`
    //支付方式
    OrderPay    string      `orm:"size(20)"`
    //订单生成时间
    CreateTime  time.Time  `orm:"auto_now_add;type(datetime);index"`
    //订单完成时间
    FinishTime  time.Time   `orm:"auto_now;type(datetime)"`
    //订单取消时间
    CancelTime  time.Time   `orm:"auto_now;type(datetime)"`
}
// Refund 退款表
type Refund struct{
    // 退款编号
    RefundId    int64       `orm:"pk;auto"`
    //退款原因
    Reason      string      `orm:"null"`
    //退款状态
    ReStatus    string      `orm:"size(20)"`
}
// OrderRefund 订单表和退款表的关系表 即订单退款表
type OrderRefund struct{
    //订单退款编号
    OrderReId   int64       `orm:"pk;auto"`
    //订单编号 外键 暂空
    OrderId     int64       `orm:"null"`
    //退款编号  外键 暂空
    RefundId    int64       `orm:"null"`
}