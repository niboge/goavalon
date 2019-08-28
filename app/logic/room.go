package logic

import (
	"avalon/app/model"
	"avalon/plugin"
	"fmt"
	//	"fmt"
	"strconv"
)

var session, _ = plugin.NewRedis("")

type RoomCfg struct {
	gametype int    // 1狼人杀 2阿瓦隆
	notice   string // 公告&&主题

	// wolf
	wolf        int
	wolf_white  int
	wolf_beauty int

	//farmer
	farmer    int
	prophet  bool
	witch    bool
	hunter   bool
	guard    bool
	idiot    bool
	magician bool

	// skill adjust
	self_rescue int //0 可自救 1第一天自救 2不可
}

type RoomLogic struct {
	RoomCfg
	id string
}

func NewRoom(roomid string) (logic *RoomLogic) {
	logic = new(RoomLogic)
	logic.id = roomid
	return logic
}

func (this *RoomLogic) AlterCfg(param map[string]string, uid int) bool {
	roomid := param["id"]

	farmer, _ := strconv.Atoi(param["farmer"])
	wolf, _ := strconv.Atoi(param["wolf"])
	self_rescue, _ := strconv.Atoi(param["self_rescue"])

	if farmer < 1 || farmer > 5 {
		return false
	}
	if wolf < 1 || wolf > 5 {
		return false
	}
	if self_rescue < 1 || self_rescue > 3 {
		return false
	}

	if _, ok := param["notice"]; ok {
		this.notice = param["notice"]
	}

	this.gametype = 1 //狼人杀
	this.gametype = 2 //阿瓦隆

	this.self_rescue = self_rescue
	this.wolf = wolf
	this.wolf_white, _ = strconv.Atoi(param["wolf_white"])
	this.wolf_beauty, _ = strconv.Atoi(param["wolf_beauty"])

	this.farmer = farmer
	this.prophet, _ = strconv.ParseBool(param["prophet"])
	this.witch, _ = strconv.ParseBool(param["witch"])
	this.hunter, _ = strconv.ParseBool(param["hunter"])
	this.guard, _ = strconv.ParseBool(param["guard"])
	this.idiot, _ = strconv.ParseBool(param["idiot"])
	this.magician, _ = strconv.ParseBool(param["magician"])

	// to redis
	session.Set("Room:"+this.id, this.RoomCfg)

	return model.Room.Modify(fmt.Sprintf("id=%d AND owner=%d",roomid,uid), param)
}

func _formatParam(param map[string]string) {

}