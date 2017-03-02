package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)


// Note 笔记表
type Note struct{
    //笔记编号 主键
    NoteId      int64        `orm:"pk;auto"`
    //学校编号 外键   暂可为空
    SchoolId    int          `orm:"null"`
    //年级编号 外键   暂可为空
    GradeId     int          `orm:"null"`
    //笔记属于教材版本 外键    暂可为空
    BookVer     int64        `orm:null"`
    //销售状态 外键    暂可为空
    AlesStatus  string       `orm:"null"`
     //分类编号外键 暂时可以为空
    CategoriesId    int64    `orm:"null"`
    //笔记所属证书等级
    CerLankId   int          `orm:"null"`
    //时间编号
    TimeId      int          `orm:"null"`
    // 笔记标题 
    NoteTitle   string   
    //笔记的记录时间   
    NoteTime    time.Time    `orm:"type(datetime);null"`
    //笔记详情描述
    NoteDesc    string       `orm:"null"`
    // 笔记价格
    NotePrice   float64     `orm:"digits(10);decimals(2)"`
    //笔记的文档地址
    NoteAdress  string      `orm:"null"`
    //发布时间
    CarrTime    time.Time   `orm:"auto_now;type(datetime);index"`
    //下架时间
    ShelfTime   time.Time   `orm:"type(datetime)"`
    //笔记的库存数量
    NoteNum     int         `orm:"null"`
    // 笔记购买数量
    PurchNum    int         `orm:"null"`
    //笔记的评论数量
    CommentNum  int         `orm:"null;index"`
    //笔记的好评率
    HightComme  float64     `orm:"digits(4);decimals(2);index"`
    // 笔记的浏览次数
    ScanTimes   int64       `orm:"null;index;index"`
    // 笔记是否为实体还是虚拟商品
    NoteEntity  int         `orm:"null;index"`
    // 笔记的浏览总数
    NoteMess    int         `orm:"null"`
    // 笔记的官方评价
    NoteGrade   float64     `orm:"digits(2);decimals(1)"`
    // 笔记的官方评价
    NoteAppra   string      `orm:"null"`
    // 优惠折扣
    NoteSale    float64     `orm:"digits(2);decimals(1);index"`
}
// 笔记是否为实体
type NoteType struct{
    TypeId      int         `orm:"pk;auto"`
    TypeName    string      `orm:"sie(20)"`
}
// CertiLevel 笔记所属证书类型
type CertiLevel   struct{
    CerLankId   int         `orm:"pk;auto"`
    CerLankName string      `orm:"size(20)"`
}


// SaleStatus 销售状态表
type SaleStatus   struct{
    //销售状态编号
    SalesId     int         `orm:"pk;auto"`
    //销售状态名称
    SaleName    string      `orm:"size(20);index"`
}
// NoteImage 笔记展示图片表
type NoteImage  struct{
    //图片编号
    NoteImaId   int64       `orm:"pk;auto"`
    //笔记编号 外键 暂可为空
    NoteId      int64       `orm:"null"`
    //图片路径一
    PathOne     string      `orm:"null"`
    //图片路径二
    PathTwo     string      `orm:"null"`
    //图片路径三
    PathThree   string      `orm:"null"`
}

// NoteToLable 笔记标签关系表
type NoteToLable struct{
    //笔记标签关系主键
    ToLableId   int64       `orm:"pk;auto"`
    //笔记编号 外键
    NoteId      int64       `orm:"null"`
    // 标签编号 外键
    LableId     int64        `orm:"null"`
}
// NoteCaties  笔记分类表
type NoteCategory struct{
    //分类编号
    CategoriesId    int64     `orm:"pk;auto"`
    //分类名字
    CategoriesName  string     `orm:"size(20);index"`

}
// FindNoteCategory 获取所有笔记分类
func FindNoteCategory()  ([]*NoteCategory,error) {
    o:= orm.NewOrm()
    cates:=make([]*NoteCategory, 0)
    qs:=o.QueryTable("note_category")
    _,err:=qs.All(&cates)
    return cates,err
}

