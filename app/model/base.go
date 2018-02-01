package model

import (
    "github.com/astaxie/beego"
    Orm "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"

    // . "fmt"
)

var orm Orm.Ormer

type Base struct{
}

func Register() {
    // db init
    Orm.Debug, _ = beego.AppConfig.Bool("debug")
    Orm.RegisterDriver("mysql", Orm.DRMySQL)
    err := Orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("db_mysql"))
    orm = Orm.NewOrm()
    orm.Using("default")
}

func (this *Base)baseFindFirst(cond interface{}) interface{} {
    // return this.FindFirst()
    return cond
}