package handle

import (
	"sync"

	"avalon/plugin"
	// "github.com/astaxie/beego/orm"
	"avalon/app/model"
	"net/http"

	"github.com/go-martini/martini"
)

type Chat struct {
	sync.Mutex
	Rooms map[string]*RoomSt
}

var chat *Chat
var once sync.Once

// NewChat ...
func GetChat() *Chat {
	once.Do(func() {
		chat = &Chat{sync.Mutex{}, map[string]*RoomSt{}}
	})
	return chat
}

// (chat *Chat) GetRoomByName ...
func (chat *Chat) GetRoomByName(roomName string) *RoomSt {
	chat.Lock()
	defer chat.Unlock()
	room, ok := chat.Rooms[roomName]
	if ok == true {
		return room
	} else {
		return nil
	}
}

// AddRoom ...
func (chat *Chat) AddRoom(roomName string, room *RoomSt) {
	chat.Lock()
	defer chat.Unlock()
	chat.Rooms[roomName] = room
}

// RemoveRoom
func (chat *Chat) RemoveChat(roomName string) {
	delete(chat.Rooms, roomName)
}

// HandleGameSocket ...
func ResistSocket(req *http.Request, params martini.Params, recevier <-chan *Message, sender chan<- *Message, done <-chan bool, disconnect chan<- int, err <-chan error) (int, string) {
	sessionKey := req.Header.Get("authSessionKey")
	roomName := params["name"]

	session, _ := plugin.NewRedis("")
	info, _ := session.Get("userInfo")
	if info == nil {
		userInfo := info.(model.UserSt)
		cli := Client{Name: sessionKey, UserInfo: userInfo, in: recevier, out: sender, done: done, diconnect: disconnect /*, err: <-err*/}
		room := chat.GetRoomByName(roomName)
		if room == nil {
			return 404, "{errorInfo:'can't find room'}"
		}
		room.AddClient(sessionKey, &cli)
		addMsg := &Message{From: sessionKey, EventName: "Join", Body: ""}
		addMsg.UserInfo.NickName = cli.UserInfo.NickName
		addMsg.UserInfo.Avatar = cli.UserInfo.Avatar
		room.BroadcastMessage(addMsg, &cli)
		for {
			select {
			case <-cli.err:
				// 处理错误消息
			case msg := <-cli.in:
				// 消息处理器
				ResistGameHandle(room, msg, &cli)
			case <-cli.done:
				// 处理掉线
				room.RemoveClient(cli.Name)
				msg := &Message{From: sessionKey, EventName: "Disconnect",
					Body: " "}
				msg.UserInfo.NickName = cli.UserInfo.NickName
				msg.UserInfo.Avatar = cli.UserInfo.Avatar
				room.BroadcastMessage(msg, &cli)
				return 200, "ok"
			}
		}

	} else {
		return 403, "{'errorInfo':'no login'}"
	}
}
