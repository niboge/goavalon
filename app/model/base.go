package model

import (
    "github.com/astaxie/beego"
    Orm "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    
    // . "avalon/plugin/selftype"
    . "fmt"
)

var orm Orm.Ormer
var db Orm.QueryBuilder

type Base struct{
    tableName string
}

// 构建model查询
type ModelCond struct{
    Where string
    Bind interface{}
    Column string
    Order string
    Group string
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

func (this *Base)FindFirst(cond interface{}) Orm.RawSeter{
    db, _ = Orm.NewQueryBuilder("mysql")

    var res Orm.RawSeter
    var sql string
    
    switch cond.(type) {
    case bool:
        // res = db.Select("*").From("room").String()
    case int:
        sql = db.Select("*").From(this.tableName).Where(Sprintf("id=%d", cond)).String()

        res = orm.Raw(sql)
    case string:
        sql = db.Select("*").From(this.tableName).Where(cond.(string)).String()
        res = orm.Raw(sql)
    
    case ModelCond:
        conds := cond.(ModelCond)
        if conds.Column == "" {
            conds.Column = "*"
        }
        if conds.Where == "" {
            conds.Where = "TURE"
        }

        sql = db.Select(conds.Column).From(this.tableName).Where(conds.Where).String()
        res = orm.Raw(sql, conds.Bind)
    }

    // Printf("[DB-FindFirst]: %v \n",sql)
    
    return res
}

func (this *Base)Find(cond interface{}) Orm.RawSeter{
    db,_ = Orm.NewQueryBuilder("mysql")

    var res Orm.RawSeter
    
    switch cond.(type) {
    case bool:
        res = orm.Raw( db.Select("*").From("room").String())
    case int:
        res = orm.Raw( db.Select("*").From(this.tableName).Where(Sprintf("id=%d", cond)).String())
    case string:
        db.Select("*").From(this.tableName).Where(cond.(string))
    case ModelCond:
        conds := cond.(ModelCond)
        db.Select("*").From(this.tableName).Where(conds.Where)
    default:
        conds := Orm.Params{}
        where := conds["where"]
        limit := conds["limit"]
        Printf("%V %V \n",where,limit)
        // res = db.Select("name").From("_room").Where("sdf").String()
    }

    Printf("[DB-Find]: %v \n",res)
    
    return res
}