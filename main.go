package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "fmt"
	"reflect"
	"avalon/app/handle"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 加载模板
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
		// c.JSON(200, gin.H{
		// 	"message": "hi",
		// })
		Printf("%v %T \n", gin.H{}, gin.H{})
		c.String(http.StatusOK, "Hello w Word")
	})

	router.GET("/index", handle.Index.Main)
}



/*
import (
	"net/http"

	"avalon/conf"
	"avalon/db"
	"avalon/handle"
	"avalon/middleware"

	"github.com/beatrichartz/martini-sockets"
	"github.com/go-martini/martini"
)

func main() {

	m := martini.Classic()
	config := conf.CreateConfig("./config/config.yaml")
	ConfigMartini(m, config)
	RouterConfig(m)
	m.Run()
}

func ConfigMartini(m *martini.ClassicMartini, config *conf.Config) *martini.ClassicMartini {
	orm := db.SetEngine(config.DataBase.DbPath)
	// 初始化用户表	
	orm.Sync(new(db.User))
	sessionManager := middleware.GetSessionManager(7200)
	// 配置DATABASES
	m.Map(orm)
	// 全局配置信息
	m.Map(config)
	// 全局Wxssion管理器
	m.Map(sessionManager)
	// handle.GetChat()

	return m
}

func RouterConfig(m *martini.ClassicMartini) {
	m.Get("/", func() string {
		return "hello,word"
	})
	// m.Get("/login", handle.LoginWechatUser)
	m.Get("/login", func (req *http.Request) (int, string) {
		return 200, "hello, word"
		})
	m.Post("/registerUser", handle.RegisterWechatUser)
	m.Get("/game/room/:name", sockets.JSON(handle.Message{}), handle.ResistSocket)
}
*/