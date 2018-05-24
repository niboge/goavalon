package handle

import (
	"avalon/plugin"
)

var session = plugin.NewRedis("")

type RoomLogic struct {
}
