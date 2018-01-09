package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "fmt"
	"reflect"
	"avalon/app/handle"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
    _ "github.com/go-sql-driver/mysql"
)

const PROJECT_NAME = 'avalon'

func init() {
	// config
    beego.AppConfigPath = "avalon/app/conf/db.conf"
    beego.ParseConfig()

    Printf("%V",beego.AppConfig)
	// orm
	orm.RegisterDataBase(PROJECT_NAME, "mysql", "root:xiaobaitu2@/time.geekbang.org?charset=utf8")
    orm.RunSyncdb(PROJECT_NAME, false, true)
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 加载模板 todo 一次不行么
	r.LoadHTMLGlob("app/templates/**/*.tpl")
	r.LoadHTMLGlob("app/templates/*.tpl")	

	// 静态文件
	r.StaticFile("/favicon.ico", "./public/favicon.ico")

	// 路由
	Router(r)

	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}


func Router(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello w Word")
	})

	router.GET("/index", handle.Index.Main)
}