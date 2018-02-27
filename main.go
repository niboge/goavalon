package main

import (
	"avalon/app/handle"
	"github.com/gin-gonic/gin"
	"net/http"

	"avalon/app/model"
	"github.com/astaxie/beego/session"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"

	. "fmt"
)

const PROJECT_NAME = "avalon"

func init() {
	model.Register()
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	Session()

	Route(r)

	r.Run("127.0.0.1:80") // listen and serve on 0.0.0.0:8080
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
	// 加载模板
	router.LoadHTMLGlob("app/templates/**/*.tpl")

	// 静态文件
	router.Static("/public", "./public")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")

	router.GET("/room/saber", func(c *gin.Context) { WebsocketLoop(c.Writer, c.Request) })

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
	}

	// user
	group = router.Group("/user", middelware(&handle.User))
	{
		group.GET("", handle.User.Info)
	}

	// room
	group = router.Group("/room", middelware(&handle.Room))
	{
		group.GET("", handle.Room.Main)
	}
}

func middelware(cont handle.BaseI) gin.HandlerFunc {

	return func(c *gin.Context) {
		cont.BeforeHandle(c, globalSession)
	}
}

func WebsocketLoop(w http.ResponseWriter, r *http.Request) {

	_, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	// for {
	//     t, msg, err := conn.ReadMessage()
	//     if err != nil {
	//     	conn.WriteMessage(t, []byte(Sprintf("%s", err)))
	//         break
	//     }
	//     conn.WriteMessage(t, msg)
	// }

	// 	var cmd []string = make([]string, 3)
	// 	var to_id int

	// 	client := G.AddClient(ws)

	// 	if client == nil {
	// 		log.Printf("Too many connection. Addr=%s", client.Addr)
	// 		return
	// 	}

	// 	//	defer func() {
	// 	//		if x := recover(); x != nil {
	// 	//			log.Printf("run time panic: %v", x)
	// 	//		}
	// 	//	}()

	// Loop:
	// 	for {
	// 		var reply string
	// 		if err := websocket.Message.Receive(ws, &reply); err != nil {
	// 			log.Printf("Receive Error: %s", err)
	// 			break
	// 		}
	// 		if etc.Debug {
	// 			log.Printf("[#%d] %s", client.Id, reply)
	// 		}
	// 		if EqualFold("", TrimSpace(reply)) {
	// 			fmt.Println("empty")
	// 			break
	// 		}
	// 		cmd = SplitN(reply, " ", 3)
	// 		//至少有2个参数
	// 		if len(cmd) < 3 {
	// 			goto unknow
	// 		}
	// 		switch cmd[0] {
	// 		case "set":
	// 			switch cmd[1] {
	// 			//修改名称
	// 			case "name":
	// 				info := SplitN(cmd[2], "|", 2)
	// 				client.SetName(info[0], info[1])
	// 				continue
	// 				break
	// 			default:
	// 				goto unknow
	// 			}
	// 			break
	// 		case "get":
	// 			switch cmd[1] {
	// 			case "name":
	// 				client.Write("names " + G.GetClients())
	// 				break
	// 			default:
	// 				goto unknow
	// 			}
	// 			break

	// 		//向某个人发送信息
	// 		case "sendto":
	// 			to_id, _ = strconv.Atoi(cmd[1])
	// 			G.Sendto(client.Id, to_id, cmd[2])
	// 			break

	// 		//向某个组或所有人发送信息
	// 		case "sendm":
	// 			to_id, _ = strconv.Atoi(cmd[1])
	// 			G.Broadcast(fmt.Sprintf("msgall %d %s", client.Id, cmd[2]), client.Id)
	// 			break
	// 		case "logout":
	// 			G.RemoveClient(client.Id)
	// 			break Loop

	// 		default:
	// 			goto unknow
	// 		}
	// 		continue
	// 	unknow:
	// 		if err := websocket.Message.Send(ws, "cmd unknow"); err != nil {
	// 			log.Printf("Can't send.%s", err)
	// 			panic(100)
	// 		}
	// 	}
	// 	G.RemoveClient(client.Id)
}
