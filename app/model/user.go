package model


import (
    // "github.com/astaxie/beego"
    Orm "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"

    . "fmt"

)

var User *UserSt

func init() {
    User = new(UserSt)
    Orm.RegisterModel(User)
    User.tableName = "user"
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
    WinRate string
    // Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

func (this *UserSt) TableName() string {
    return "user"
}

func (this *UserSt) FindFirst(cond interface{}) (bool,UserSt){
    var res UserSt
    
    err := this.Base.FindFirst(cond).QueryRow(&res)

    return err == nil, res
}

func (this *UserSt) Save(cond interface{}) (int , error) {
    user := new(UserSt)
    user.NickName = "海波"
    id, err := orm.Insert(user)
    Printf("%V %V \n", id , err)

    return int(id), err
}