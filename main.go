package main

import (
	"github.com/astaxie/beego"

	"avalon/app/handle"
	. "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"

	"avalon/app/model"
)

const PROJECT_NAME = "avalon"

func init() {
	// init orm
	model.Register()
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
