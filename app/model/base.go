package model

import (
	. "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"reflect"
	"strings"
)

type Base struct {
	tableName string
	Id        int
}

// 构建model查询
type ModelCond struct {
	Where  interface{}
	Bind   []string
	Column string
	Order  string
	Group  string
}

var Engine *gorm.DB

func Init() *gorm.DB {
	var err error

	Engine, err = gorm.Open("mysql", beego.AppConfig.String("db_mysql"))
	if err != nil {
		panic("failed to connect database!")
	}

	Engine.DB().SetConnMaxLifetime(3)
	Engine.DB().SetMaxIdleConns(10)
	Engine.DB().SetMaxOpenConns(100)
	Engine.LogMode(true)

	return Engine
}

func (this *Base) BeforeUpdate() {}
func (this *Base) AfterUpdate() {}


func (this *Base) FindFirst(cond interface{}) *gorm.DB {
	return this.condition(cond).Limit(1)
}

func (this *Base) FindAll(cond interface{}) *gorm.DB {
	//rows :=
	//if err != nil {
	//	logs.Warning("DB.Find fail",err)
	//}

	return this.condition(cond)
}

func (this *Base) Create(model interface{}) (ok bool) {
	db := Engine.Table(this.tableName).Create(model)

	if db.Error != nil {
		logs.Warning("DB.Insert fail", db.Error)
	}else {
		ok = true
	}

	return
}

func (this *Base) Modify(cond interface{}, data interface{}) (ok bool) {
	var data_map map[string]interface{}

	if data != nil {
		data_map = data.(map[string]interface{})
	}else{
		//todo op 不推荐用,把所有字段都更新了,待优化
		data_map = this.getModelMap()
		if cond == nil {
			cond = this.Id
		}
	}

	db := this.condition(cond).Updates(data_map)
	if db.Error != nil {
		logs.Warning("DB.Insert fail", db.Error)
	}else {
		ok = true
	}

	return
}

func (this *Base) Exec(sql string,args ...interface{}) *gorm.DB {
	return Engine.Exec(sql, args)
}

func (this *Base) condition(cond interface{}) *gorm.DB{
	var db *gorm.DB
	switch cond.(type) {
	case bool:
		db = Engine.Select("*").Table(this.tableName).Where("true")
	case []string:
		db = Engine.Select("*").Table(this.tableName).Where("id IN (?)", cond.([]string))
	case []int:
		db = Engine.Select("*").Table(this.tableName).Where("id IN (?)", cond.([]int))
	case int:
		db = Engine.Select("*").Table(this.tableName).Where(Sprintf("id=%d", cond.(int)))
	case string:
		db = Engine.Select("*").Table(this.tableName).Where(cond.(string))
	case ModelCond:
		condComplex := cond.(ModelCond)
		column := "*"
		if condComplex.Column != "" {
			column = condComplex.Column
		}

		db = Engine.Select(column).
			Table(this.tableName).
			Where(condComplex.Where, append(make([]interface{},0), condComplex.Bind)...)

		if condComplex.Order != "" {
			db = db.Order(condComplex.Order)
		}
	}

	return db
}

func (this *Base) getModelMap() map[string]interface{} {
	ret := make(map[string]interface{},10)

	t := reflect.TypeOf(*this)
	v := reflect.ValueOf(*this)

	for i := 0; i < t.NumField(); i++ {
		ret[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}

	return ret
}