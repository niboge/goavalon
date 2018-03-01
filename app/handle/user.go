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

  var user model.UserSt

	auth := this.GetSession("UserAuth")
	if auth != nil {
    user = auth.(model.UserSt)
    if user.Lose + user.Win != 0 {
      user.WinRate = Sprintf("%.2f", float32(user.Win)/float32(user.Win+user.Lose)*100)
    }
  }

	this.succ(user, "personal.tpl")
}


