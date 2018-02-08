package model

import (
    "github.com/astaxie/beego"
    Orm "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"

    . "fmt"
)

var orm Orm.Ormer
var db Orm.QueryBuilder

type Base struct{
    tableName string
}

func Register() {
    // db init
    Orm.Debug, _ = beego.AppConfig.Bool("debug")
    Orm.RegisterDriver("mysql", Orm.DRMySQL)
    Orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("db_mysql"))
    orm = Orm.NewOrm()
    orm.Using("default")
}

// func (this *Base) before () {
    // this.TableName()
// }

func (this *Base)baseFindFirst(cond interface{}) interface{} {
    // return this.FindFirst()
    return cond
}

func (this *Base)Find(cond interface{}) string{
    db,_ = Orm.NewQueryBuilder("mysql")

    var res string
    
    switch cond.(type) {
    case bool:
        res = db.Select("*").From("room").String()
    case int:
        res = db.Select("*").
        From(this.tableName).
        Where("id=" + "1").String()
    case string:
        db.Select("*").From(this.tableName).Where(cond.(string))
    default:
        conds := Orm.Params{}
        where := conds["where"]
        limit := conds["limit"]
        Printf("%V %V \n",where,limit)
        // res = db.Select("name").From("_room").Where("sdf").String()
    }

    Printf("[DB-INFO]: %v \n",res)
    
    return res
}