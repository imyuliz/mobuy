package main

import (
	_ "mobuy/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"mobuy/models"
)

func main() {
	orm.RegisterDataBase("default","mysql",beego.AppConfig.String("username")+":"+beego.AppConfig.String("password")+ "@/mobuy?charset=utf8&loc=Asia%2FShanghai",30)
	orm.RegisterModel(
		new(models.Activity),
		new(models.AdInfo),
		new(models.Appeal),
		new(models.UserApeal),
		new(models.Collection),
		new(models.Comment),
		new(models.School),
		new(models.BookVer),
		new(models.Grade),
		new(models.NoteSubject),
		new(models.Rank),
		new(models.Label),
		new(models.SuitPeo),
		new(models.FeedbackInfo),
		new(models.Location),
		new(models.Province),
		new(models.City),
		new(models.County),
		new(models.Manager),
		new(models.ManaUseriden),
		new(models.ManagerWork),
		new(models.Message),
		new(models.Note),
		new(models.NoteType),
		new(models.CertiLevel),
		new(models.SaleStatus),
		new(models.NoteImage),
		new(models.NoteToLable),
		new(models.Notice),
		new(models.NoticeType),
		new(models.Oder),
		new(models.Refund),
		new(models.OrderRefund),
		new(models.Recovery),
		new(models.RecoveryType),
		new(models.RecoveryStatus),
		new(models.Reply),
		new(models.Reward),
		new(models.RewardRel),
		new(models.DiffDegree),
		new(models.RewardStus),
		new(models.User),
		new(models.UserIdentity),
		new(models.Death),
		new(models.NoteCategory),
		new(models.UserToDeath))
		orm.RunSyncdb("default",false,true)
	beego.Run()
}

