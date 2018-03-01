package model


import (
    // "github.com/astaxie/beego"
    Orm "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"

    . "fmt"

    . "avalon/plugin/selftype"
)

var Room *RoomSt
var RoomArr []RoomSt

func init() {
    Room = new(RoomSt)
    Orm.RegisterModel(Room)
    Room.tableName = "room"
    // Room.before()
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

func (this *RoomSt) FindFirst(cond interface{}) interface{}{
    this.Base.FindFirst(cond)

    var obj RoomSt
    switch cond.(type) {
    case int:
        obj = RoomSt{Id:cond.(int)}
        if err := orm.Read(&obj, "Id"); err != nil {
            return false
        }
    case string:
        // where := cond.(string)
        
    default:
        conds := cond.(Object)
        where := conds["where"]
        limit := conds["limit"]
        Printf("%V %V \n",where,limit)
    }
    
    return obj
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