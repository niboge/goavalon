package handle

import (
	. "fmt"
	"github.com/gin-gonic/gin"
	"log"
	// "net/http"
	// Str "strings"

	"avalon/app/model"

	. "avalon/plugin/selftype"
)

var User UserSt

func init() {
	User = UserSt{BaseSt{Data:Object{}, c:nil, auth:true}}
}

/**
    Hanlde: Index Class
**/
type UserSt struct {
	BaseSt
}

// main page
func (this *UserSt) Info(context *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Print("[Error] ", err)
			this.fail(Sprintf("%s", err))
		}
	}()

  this.SetSession("name","value")
  
  var data = Object{
    "account": "阿瓦隆Demo",
    "nick":    "玩家Demo",
    "score":   "100",
    "win":     10.0,
    "lose":    3.0,
    "winrate": 56.0,
  }

	succ, user := model.User.FindFirst(2)
	if succ {
    data["account"] = user.Account
    data["nick"] = user.NickName
    data["avatar"] = this.c
    data["score"] = user.Score
    data["win"] = user.Win
    data["lose"] = user.Lose
    if user.Lose + user.Win != 0 {
      data["winrate"] = Sprintf("%.2f", float32(user.Win)/float32(user.Win+user.Lose)*100)
    }else {
      data["winrate"] = '-'
    }
  }

	this.succ(data, "personal.tpl")
}
