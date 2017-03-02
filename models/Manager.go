package models

// Manager 管理员表
type Manager struct{
    //管理员编号
    ManagerId       int     `orm:"pk;auto"`
    //管理员编号
    ManagerAcc      string  `orm:"size(20)"`
    //管理员名字
    ManagerName     string  `orm:"size(20)"`
    //管理员权限
    ManagerPower    string  `orm:"size(20);null"`
}
// ManaUseriden 管理员和用户审核关系表
type ManaUseriden   struct{
    //关系编号
    ManuserRelId    int64    `orm:"pk;auto"`
    // 审核编号 外键 暂空
    DeathId         int64    `orm:"null"`
    // 管理员编号 外键 暂空
    ManagerId       int      `orm:"null"`
}
// ManagerWork 管理员事务表
type ManagerWork    struct{
    //事务编号
    WorkId          int64    `orm:"auto"`
    //管理员编号
    ManagerId       int      `orm:"null"`
    // 管理员处理的编号  投诉等，不包括审核用户的身份 是投诉表的外键
    ComplainId      int64    `orm:"null"`
}