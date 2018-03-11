package main

import (
	"avalon/app/handle"
	"github.com/gin-gonic/gin"
	"net/http"

	"avalon/app/model"
	"github.com/astaxie/beego/session"
	// _ "github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"

	. "fmt"
)

const PROJECT_NAME = "avalon"

func init() {
	model.Register()
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	Session()

	Route(r)

	r.Run("xdd.cn:80") // listen and serve on 0.0.0.0:8080
}

var globalSession *session.Manager

func Session() {
	beego.BConfig.WebConfig.Session.SessionOn = true

	sessionConfig := &session.ManagerConfig{
		CookieName:      "go." + PROJECT_NAME,
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	globalSession, _ = session.NewManager("memory", sessionConfig)
	go globalSession.GC()
}

func Route(router *gin.Engine) {
	//
	router.LoadHTMLGlob("app/templates/**/*.tpl")
	router.Static("/public", "./public")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")

	// websocket server
	router.GET("/websocket", func(c *gin.Context) { WebsocketLoop(c.Writer, c.Request) })

	// doc
	router.GET("/doc/rule", handle.Rule)
	router.GET("/doc/ss", handle.Ss)

	// index
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index/main")
	})
	group := router.Group("/index", middelware(&handle.Index))
	{
		group.GET("/main", handle.Index.Main)
		group.GET("/login", handle.Index.Login)
		group.POST("/login", handle.Index.Login)
	}

	// user
	group = router.Group("/user", middelware(&handle.User))
	{
		group.GET("", handle.User.Info)
	}

	// room
	group = router.Group("/room", middelware(&handle.Room))
	{
		group.GET("", handle.Room.List)
		group.GET("/in:roomName", handle.Room.Game)
	}
}

func middelware(cont handle.BaseI) gin.HandlerFunc {
	return func(c *gin.Context) {
		cont.BeforeHandle(c, globalSession)
	}
}

