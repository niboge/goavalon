package handle

import (
	// "github.com/astaxie/beego/orm"
	"avalon/app/model"
)

type Client struct {
	Name      string // 客户端链接的名字，这里一般为3rdsession的字符串
	UserInfo  model.UserSt
	in        <-chan *Message // 传进来的消息管道
	out       chan<- *Message // 发出去的消息管道
	done      <-chan bool     // 结束的bool
	err       <-chan error    // 错误管道
	diconnect chan<- int      // 断开链接的管道
}

type RoomUserInfo struct {
	Name      string `json:"name"`
	NickName  string `json:"nickName"`
	Avatar string `json:"Avatar"`
}

type Message struct {
	From      string `json:"from"`
	EventName string `json:"eventName"`
	Body      string `json:"body"`
	RoleInfo  struct {
		Role    string `json:"role"`
		Captain string `json:"captain"`
	} `json:"roleInfo"`
	UserInfo struct {
		NickName  string `json:"nickName"`
		Avatar string `json:"Avatar"`
	} `json:"userInfo"`
	UserList []RoomUserInfo `json:"userList"`
	TeamSize int            `json:"teamSize"`
	TeamList []string       `json:"teamList"`
}
