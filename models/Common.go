package models

// School 学校表
type School struct{
    //学校编号
    SchoolId    int         `orm:"pk;auto"`
    //学校名字
    SchoolName  string      `orm:"size(30);index"`
}

// BookVer 教材版本表
type BookVer struct{
    //教材编号
    BookId      int64       `orm:"pk;auto"`
    //教材版本
    BookName    string      `orm:"size(30);index"`    
}

// Grade 年级表
type Grade  struct{
    //年级编号
    GradeId     int         `orm:"pk;auto"`
    //年级名称
    GradeName   string      `orm:"size(20);index"`
}

// NoteSubject 所属学科表
type NoteSubject struct{
    //学科主键
    SubjectId   int         `orm:"pk;auto"`
    //学科名称
    SubjectName string      `orm:"size(20);index"`
}


// Rank 笔记评价等级表
type Rank struct{
    // 等级编号
    RankId      int64       `orm:"pk;auto"`
    // 等级名称
    RankName    string      `orm:"size(20)"`
    //笔记编号 外键 暂可为空
    NoteId      int64       `orm:"null"`
}

// Label   标签表
type Label struct{
    //标签主键  
    LableId     int64       `orm:"pk;auto"`
    //标签名字
    LableName   string      `orm:"size(20);index"`
}

// 合适人群
type SuitPeo struct{
    // 人群编号
    SuitId          int64       `orm:"pk;auto"`
    // 人群名字
    SuitName        string      `orm:"null;size(20);index"`
}