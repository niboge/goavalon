package handle

import (
	. "fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	Str "strings"

	"avalon/app/model"

	. "avalon/plugin/selftype"
)

var Index IndexSt

func init() {
	Index = IndexSt{"[圣杯战争] [刺客练刀房] [莫甘娜演技培训班] [神民] [叫嚣全场狼美人] [忠臣假跳炸奥博伦]", BaseSt{Object{}}}
}

/**
    Hanlde: Index Class
**/
type IndexSt struct {
	Info string
	BaseSt
}

// main page
func (this *IndexSt) Main(context *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Print("[Error] ", err)
			context.HTML(http.StatusInternalServerError, "index.tpl", this.Fail(err))
			return
		}
	}()

	user := model.User.FindFirst(2)
	if user == false{
		panic(" no this user")
	}

	var data = Object{
		"game_name":  "阿瓦隆",
		"game_slogn": Str.Split(this.Info, " "),
		"players":    "玩家",
		"game":       "game",
		"info": user,
	}

	context.HTML(http.StatusOK, "index.tpl", this.Succ(data))
}

func (this *IndexSt) Self() string {
	return this.Info
}

func (this *IndexSt) Test() (str string) {
	str = this.Info
	this.Info = this.Info + " 测试test"
	return
}
