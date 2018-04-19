package main

import (
	"avalon/app/handle"
	"avalon/app/model"
	"avalon/util"

	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/astaxie/beego/session"
	"github.com/gorilla/websocket"

	"container/list"
	"encoding/json"
	"sync"

	. "fmt"
	"log"
	"os"
)

func init() {
	CreteRoom("tes", 12)
	go manager.start()
}

var Upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ClientManager struct {
	clients    map[string]map[string]*Client
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	user     model.UserSt
	id       string
	roomName string
	socket   *websocket.Conn
	send     chan []byte
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    map[string]map[string]*Client{},
}

func CreteRoom(roomName string, roomSize int) *handle.RoomSt {
	dismissVote := util.NewVote() // 反对票仓
	agreeVote := util.NewVote()   // 同意票仓
	turnTalkList := list.New()    // 轮流发言链表
	room := handle.RoomSt{
		Mutex:         sync.Mutex{},
		Name:          roomName,
		DisVote:       dismissVote,
		AgrVote:       agreeVote,
		Clients:       make(map[string]*handle.Client, 25),
		RoomSize:      roomSize,
		TurnsTalkList: turnTalkList,
		Captains:      []string{}}
	return &room
}

func WebsocketLoop(w http.ResponseWriter, r *http.Request) {
	conn, error := Upgrader.Upgrade(w, r, nil)
	if error != nil {
		log.Fatal("Connect Websocket Fail!", error)
		http.NotFound(w, r)
		return
	}

	session, err := globalSession.SessionStart(w, r)
	if err != nil {
		panic("[Error] !! 503 session fail!")
	}

	user := session.Get("UserAuth")
	if user == nil {
		jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected.no auth"})
		conn.WriteMessage(websocket.CloseMessage, jsonMessage)
		conn.Close()
		return
	}

	// go room
	r.ParseForm()
	client := &Client{id: Sprintf("%s", user.(model.UserSt).Id), user: user.(model.UserSt), roomName: r.Form.Get("roomName"), socket: conn, send: make(chan []byte)}
	manager.register <- client

	go client.read()
	go client.write()
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.addClient(conn.roomName, conn)
			jsonMessage, _ := json.Marshal(&Message{Content: "[" + conn.user.NickName + "]进入了房间!"})
			manager.sendRoom(jsonMessage, conn.roomName)
		case conn := <-manager.unregister:
			if a, ok := manager.clients[conn.roomName]; ok {
				close(conn.send)
				// delete(manager.clients[conn.roomName], conn.roomName)
				jsonMessage, _ := json.Marshal(&Message{Content: "[" + conn.user.NickName + "]离开了房间!"})
				manager.send(jsonMessage, conn)
			}
		case message := <-manager.broadcast:
			for _, conn := range manager.clients["圣杯战争"] {
				Printf("[!TEST!] %V \n", conn)
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients[conn.roomName], conn.id)
				}
			}
		}
	}
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for _, conn := range manager.clients["圣杯战争"] {
		// if conn != ignore {
		conn.send <- message
		// }
	}
}

func (manager *ClientManager) sendRoom(message []byte, room string) {
	for _, conn := range manager.clients[room] {
		conn.send <- message
	}
}

func (c *Client) read() {
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		manager.broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (this *ClientManager) addClient(room string, c *Client) {
	if _, ok := this.clients[room]; !ok {
		this.clients[room] = make(map[string]*Client)
	}

	this.clients[room][c.id] = c
}
