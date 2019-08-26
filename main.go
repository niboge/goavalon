package main

import (
	"avalon/app/handle"
	"github.com/gin-gonic/gin"
	"net/http"

	"avalon/app/model"

//	. "fmt"
)

const PROJECT_NAME = "avalon"

func init() {
	model.Register()
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	Route(r)

	r.Run("shikii.cc:80") // listen and serve on 0.0.0.0:8080
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
	group := router.Group("/index", middelware(handle.Index))
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
		group.POST("/InceptionSpace", handle.Room.InitRoomGame)
	}
}

func middelware(cont handle.BaseI) gin.HandlerFunc {
	return func(c *gin.Context) {
		cont.BeforeHandle(c)
	}
}
