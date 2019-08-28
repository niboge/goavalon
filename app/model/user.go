package model

import (
	// "github.com/astaxie/beego"

	"errors"
	. "fmt"
)

var User *UserSt

func init() {
	User = new(UserSt)
	User.tableName = "user"
}

type UserSt struct {
	Base
	//Id        int    `json:"Id"`
	Account   string `json:"Account"`
	NickName  string `json:"NickName"`
	Avatar    string `json:"Avatar"`
	Pwd       string `json:"Pwd"`
	Mobile    string `json:"Mobile"`
	Score     int    `json:"Score"`
	Win       int    `json:"Win"`
	Lose      int    `json:"Lose"`
	WinRate   string `json:"WinRate"`
	LoginTime int    `json:"LoginTime"`
	// Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

func (this *UserSt) TableName() string {
	return "user"
}

func (this *UserSt) FindFirst(cond interface{}) (*UserSt, bool) {
	var res UserSt

	err := this.Base.FindFirst(cond).Find(&res).Error

	return &res, err == nil
}

func (this *UserSt) Insert(data *UserSt) (int, error) {
	ok := this.Create(&data)
	if !ok {
		return 0, errors.New("user.Save fail")
	}
	Printf("user.Save: %+v \n", data)

	return data.Id, nil
}

func (this *UserSt) Update(cond interface{}, data map[string]interface{}) bool {
	return this.Modify(cond, data)
}