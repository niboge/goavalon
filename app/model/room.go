package model


import (
    // "github.com/astaxie/beego"
    Orm "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"

    . "fmt"
)

var Room *RoomSt
var RoomArr []RoomSt

func init() {
    Room = new(RoomSt)
    Orm.RegisterModel(Room)
    Room.tableName = "room"
}

type RoomSt struct {
    Base
    Id      int
    Name string
    Password string
    Avatar string
    Salt string
    Status int
    IsLocked int
    CreatedTime string
    UpdatedTime string
    Owner string
    Notice string
    Type int
    // Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

func (this *RoomSt) TableName() string {
    return "room"
}

func (this *RoomSt) FindFirst(cond interface{}) RoomSt{
    var res RoomSt
    
    this.Base.FindFirst(cond).QueryRow(&res)

    return res
}

func (this *RoomSt) Find(cond interface{}) interface{}{
    raw := this.Base.Find(cond)
    raw.QueryRows(&RoomArr)
    return RoomArr
}

func (this *RoomSt) Save(cond interface{}) (int , error) {
    user := new(RoomSt)
    user.Name = "Saber圣杯战争"
    id, err := orm.Insert(user)
    Printf("%V %V \n", id , err)

    return int(id), err
}