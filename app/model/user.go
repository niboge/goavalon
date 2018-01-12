package model


import (
    // "github.com/astaxie/beego"
    Orm "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"

    . "fmt"

    . "avalon/plugin/selftype"
)

var User *UserSt

func init() {
    User = new(UserSt)
    Orm.RegisterModel(User)
}

type UserSt struct {
    Base
    Id      int
    Account string
    NickName string
    Avatar string
    Pwd string
    Mobile string
    Score int
    Win int
    Lose int
    // Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

func (this *UserSt) TableName() string {
    return "user"
}

func (this *UserSt) FindFirst(cond interface{}) interface{}{
    res := this.baseFindFirst(cond)

    var obj UserSt
    switch t := cond.(type) {
    case int:
        obj = UserSt{Id:cond.(int)}
        if err := orm.Read(&obj, "Id"); err != nil {
            return false
        }
    case string:
        where := cond.(string)
        
    default:
        conds := cond.(Object)
        where := conds["where"]
        limit := conds["limit"]
        Printf("%V %V \n",where,limit)
    }
    
    return obj

}

func (this *UserSt) Save(cond interface{}) (int , error) {
    user := new(UserSt)
    user.NickName = "海波"
    id, err := orm.Insert(user)
    Printf("%V %V \n", id , err)

    return int(id), err
}