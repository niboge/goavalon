package model

import (
	"errors"
	"fmt"

	// "github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Room *RoomSt
var RoomArr []RoomSt

func init() {
	Room = new(RoomSt)
	Room.tableName = "room"
}

type RoomSt struct {
	Base
	//Id          int
	Name        string
	Password    string
	Avatar      string
	Salt        string
	Status      int
	IsLocked    int
	CreatedTime string
	UpdatedTime string
	Owner       int
	Notice      string
	Type        int
	// Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

func (this *RoomSt) TableName() string {
	return "room"
}

func (this *RoomSt) FindFirst(cond interface{}) (*RoomSt, bool) {
	var res RoomSt

	err := this.Base.FindFirst(cond).Find(&res).Error

	return &res, err == nil
}

func (this *RoomSt) Find(cond interface{}) ([]RoomSt, bool) {
	res := make([]RoomSt, 1)

	err := this.Base.FindAll(cond).Find(&res)

	return res, err != nil
}

func (this *RoomSt) Insert(data *RoomSt) (int, error) {
	ok := this.Create(&data)
	if !ok {
		return 0, errors.New("Room.Save fail")
	}
	fmt.Printf("Room.Save: %+v \n", data)

	return data.Id, nil
}

func (this *RoomSt) Update(cond interface{}, data map[string]interface{}) bool {
	return this.Base.Modify(cond, data)
}