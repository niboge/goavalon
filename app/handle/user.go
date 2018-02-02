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
       User = UserSt{ BaseSt{Object{},nil,true} }
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

   user := model.User.FindFirst(2)
   if user == false{
           panic(" no this user")
   }

   var data = Object{
           "account":  "阿瓦隆",
           "nick":    "玩家",
           "score":       "game",
           "win": 10,
           "lose": 3,
           "winrate": 56,
   }

   this.succ(data, "user/personal.tpl")
}
