package handle

import (
	"container/list"
	"math/rand"
	"sync"

	"log"
	"fmt"
	"avalon/app/model"
	"avalon/util"
	"github.com/gin-gonic/gin"
	// "log"
	// . "fmt"
)

var Room RoomSt

func init() {
	Room = RoomSt{BaseSt: BaseSt{Data: Object{}, c: nil}, Mutex: sync.Mutex{}}
}

type RoomSt struct {
	sync.Mutex                       // 互斥锁，保证线程安全
	Stash         string             //状态
	RoomSize      int                // 房间人数
	Name          string             // 创建房间时的名字，创建时为uuid，并分享时候将该uuid带上
	GameNum       int                // 房间当前局数
	GoodManWins   int                // 抵抗组织成员获胜局数
	BadGuysWins   int                // 间谍成员获胜局数
	TurnsTalkList *list.List 		 // 轮流发言链表
	DisVote       *util.VoteSet           // 反对票仓
	AgrVote       *util.VoteSet           // 赞成票仓
	Captains      []string           // 队长链表
	Clients       map[string]*Client // 客户端管理池
	BaseSt
}

func (this *RoomSt) List(context *gin.Context) {
	rooms := model.Room.Find(true)
	if rooms == nil {
		panic("no  room")
	}

	this.succ(rooms, "room.tpl")
}

func (this *RoomSt) Game(context *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Print("[Error] ", err)
			this.fail(fmt.Sprintf("%s", err))
		}
	}()

	name := this.c.Param("roomName")
	room := model.Room.FindFirst( model.ModelCond{Where:"name=?", Bind:name})

	if room.Id == 0 {
		panic("no this room")
	}

	this.succ(room, "room/main.tpl")
}

// func CreteRoom(roomName string, roomSize int) *RoomSt {
// 	rooms := model.Room.Find(true)
// 	if rooms == nil {
// 		panic("no  room")
// 	}

// 	// for _, room := range rooms {
		
// 	// }


// 	dismissVote := NewVote()   // 反对票仓
// 	agreeVote := NewVote()     // 同意票仓
// 	turnTalkList := list.New() // 轮流发言链表
// 	room := RoomSt{
// 		Mutex:         sync.Mutex{},
// 		Name:          roomName,
// 		DisVote:       dismissVote,
// 		AgrVote:       agreeVote,
// 		Clients:       map[string]*Client{},
// 		RoomSize:      roomSize,
// 		TurnsTalkList: turnTalkList,
// 		Captains:      []string{}}
// 	return &room
// }

// 初始化游戏信息
func (room *RoomSt) InitRoomGame() {
	// 随机选择队长
	captaignsName, _ := room.TakeRandCaptains()

	badManList := []string{}
	// 分配坏蛋
	clientList := room.ClientNameList()
	// 初始化信息
	msg := &Message{From: "SYSTEM", EventName: "INIT"}
	for i := 0; i <= 2; i++ {
		point := rand.Intn(len(clientList))
		badManList = append(badManList, clientList[point])
		clientList = append(clientList[:point], clientList[point+1:]...)
	}
	// 根据分配选择给各个客户端返回信息
	for cliName, cli := range room.Clients {
		for badmanpoint := range badManList {
			if cliName == badManList[badmanpoint] {
				msg.RoleInfo.Role = "BADMAN"
				msg.RoleInfo.Captain = captaignsName
				msg.TeamList = badManList
				cli.out <- msg
			} else {
				msg.RoleInfo.Role = "GOODMAN"
				msg.RoleInfo.Captain = captaignsName
				cli.out <- msg
			}
		}
	}
}

// 添加房间客户端
func (room *RoomSt) AddClient(clientName string, client *Client) bool {
	room.Lock()
	defer room.Unlock()
	// 通知其他用户发送
	joinMsg := &Message{From: "SYSTEM", EventName: "JOIN"}
	joinMsg.UserInfo.NickName = client.UserInfo.NickName
	joinMsg.UserInfo.Avatar = client.UserInfo.Avatar
	room.BroadcastMessage(joinMsg, client)
	if len(room.ClientNameList()) < room.RoomSize-1 {
		room.Clients[clientName] = client // 加入房间的客户端池
		room.Captains = append(room.Captains, clientName)
		// 将clientName加入到发言队列中去
		room.TurnsTalkList.PushBack(clientName)
		return true
	} else if len(room.ClientNameList()) == room.RoomSize-1 {
		room.Clients[clientName] = client
		room.Captains = append(room.Captains, clientName)
		room.TurnsTalkList.PushBack(clientName)
		// 增加一个发言队列的标记
		room.TurnsTalkList.PushBack("END")
		// 发送一个标记告诉客户端人满了
		for _, cli := range room.Clients {
			readMsg := &Message{From: "System", EventName: "READY"}
			cli.out <- readMsg
		}
		// 初始化第一局游戏的信息
		room.InitRoomGame()
		return true
	} else if len(room.ClientNameList()) == room.RoomSize {
		startMsg := &Message{From: "SYSTEM", EventName: "Start", Body: ""}
		for _, cli := range room.Clients {
			cli.out <- startMsg
		}
		return true
	} else {
		return false
	}
}

// 删除房间的客户端
func (room *RoomSt) RemoveClient(clientName string) {
	room.Lock()
	defer room.Unlock()
	delete(room.Clients, clientName)
	var turnsTalkNext *list.Element
	// 删除队长备选
	for i := range room.Captains {
		if room.Captains[i] == clientName {
			room.Captains = append(room.Captains[:i], room.Captains[i+1:]...)
		}
	}
	// 将该玩家移除发言队列
	for te := room.TurnsTalkList.Front(); te != nil; {
		if te.Value.(string) == clientName {
			turnsTalkNext = te.Next()
			room.TurnsTalkList.Remove(te)
			te = turnsTalkNext
		} else {
			te = te.Next()
		}
	}
}

// 清空房间票仓
func (room *RoomSt) ClearVoteSet() {
	room.Lock()
	defer room.Unlock()
	room.DisVote.Clear()
	room.AgrVote.Clear()
}

// 投同意票
func (room *RoomSt) VoteAgreeVote(clientName string) {
	room.Lock()
	defer room.Unlock()
	room.AgrVote.Add(clientName)
}

// 投反对票
func (room *RoomSt) VoteDisVote(clientName string) {
	room.Lock()
	defer room.Unlock()
	room.DisVote.Add(clientName)
}

// 统计投票数，参数为模式，mission(任务执行模式)|team(组队模式) 返回如果味true则同意多，如果为fale则反对多,
func (room *RoomSt) CountVote(modle string) (bool, bool) {
	room.Lock()
	defer room.Unlock()
	agrvotes := room.AgrVote.Len()
	disvotes := room.DisVote.Len()
	if modle == "mission" {
		if disvotes >= 1 {
			return false, true
		} else {
			return true, true
		}
	} else if modle == "team" {
		if agrvotes > disvotes {
			return true, true
		} else {
			return false, true
		}
	} else {
		return false, false
	}
}

// 获取所有投票数，第一个为同意第二个为反对
func (room *RoomSt) GetVotes() (int, int) {
	room.Lock()
	defer room.Unlock()
	agrvotes := room.AgrVote.Len()
	disvotes := room.DisVote.Len()
	return agrvotes, disvotes
}

func (room *RoomSt) GetMissionConfig() int {
	room.Lock()
	defer room.Unlock()
	switch room.RoomSize {
	case 5:
		missionNum := [5]int{2, 3, 2, 3, 3}
		return missionNum[room.GameNum]
	case 6:
		missionNum := [5]int{2, 3, 4, 3, 4}
		return missionNum[room.GameNum]
	case 7:
		missionNum := [5]int{2, 3, 3, 4, 4}
		return missionNum[room.GameNum]
	default:
		return 2
	}
}

// 增加局数
func (room *RoomSt) AddGameNum() bool {
	room.Lock()
	defer room.Unlock()
	if room.GameNum < 5 {
		room.GameNum++
		return true
	} else {
		return false
	}
}

// 获取轮流发言客户端名字
func (room *RoomSt) TakeTurnsClientName() string {
	room.Lock()
	defer room.Unlock()
	cliNmae := room.TurnsTalkList.Front()
	room.TurnsTalkList.MoveToBack(cliNmae)
	return cliNmae.Value.(string)
}

// 随机获取队长
func (room *RoomSt) TakeRandCaptains() (string, bool) {
	room.Lock()
	defer room.Unlock()
	if len(room.Captains) != 0 {
		captainPoint := rand.Intn(len(room.Captains))
		captainName := room.Captains[captainPoint]
		room.Captains = append(room.Captains[:captainPoint], room.Captains[captainPoint+1:]...)
		return captainName, true
	} else {
		return "", false

	}

}

// 增加好人获胜局数
func (room *RoomSt) AddGoodManWins() bool {
	room.Lock()
	defer room.Unlock()
	if room.GoodManWins < 3 && room.GoodManWins+room.BadGuysWins < 5 {
		room.GameNum++
		return true
	} else {
		return false
	}
}

// 增加坏人获胜局数
func (room *RoomSt) AddBadGuysWins() bool {
	room.Lock()
	defer room.Unlock()
	if room.BadGuysWins < 3 && room.GoodManWins+room.BadGuysWins < 5 {
		room.BadGuysWins++
		return true
	} else {
		return false
	}
}

// 客户端名字列表
func (room *RoomSt) ClientNameList() []string {
	room.Lock()
	defer room.Unlock()
	list := []string{}
	for clientName := range room.Clients {
		list = append(list, clientName)
	}
	return list
}

// 给房间所有人发送消息
func (room *RoomSt) BroadcastMessage(msg *Message, client *Client) {
	room.Lock()
	defer room.Unlock()
	for _, cli := range room.Clients {
		if cli.Name != client.Name {
			client.out <- msg
		}
	}
}

func (room *RoomSt) BroadcastAll(msg *Message) {
	room.Lock()
	defer room.Unlock()
	for _, cli := range room.Clients {
		cli.out <- msg
	}

}

// (room *RoomSt) SendMessage ...
func (room *RoomSt) SendMessage(msg *Message, clientName string) {
	room.Lock()
	defer room.Unlock()
	msg.From = clientName
	client := room.Clients[clientName]
	client.out <- msg
}

// ChangeRoomStash ...
func (room *RoomSt) ChangeRoomStash(stash string) {
	room.Lock()
	defer room.Unlock()
	room.Stash = stash
}

func (room *RoomSt) GetClientByName(name string) *Client {
	room.Lock()
	defer room.Unlock()
	cli := room.Clients[name]
	return cli
}
