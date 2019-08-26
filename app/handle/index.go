package handle

import (
	. "fmt"
	"github.com/gin-gonic/gin"
	"log"
	// "net/http"
	Str "strings"

	"avalon/app/model"
	"avalon/plugin"
	"encoding/json"
	"time"
)

var Index *IndexSt

func init() {
	Index = &IndexSt{"[圣杯战争] [刺客练刀房] [莫甘娜演技培训班] [神民] [狼美人叫嚣] [毒奶]", BaseSt{Data: Object{}, c: nil}}
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
			this.fail(Sprintf("%s", err))
		}
	}()

	// suc, user := model.User.FindFirst(2)
	// if suc == false {
	// 	panic(" no this user")
	// }

	var data = Object{
		"game_name":  "阿瓦隆",
		"game_slogn": Str.Split(this.Info, " "),
		"players":    "玩家",
		"game":       "game",
		// "info":       user,
	}

	this.succ(data, "index.tpl")
}

func (this *IndexSt) Login(context *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Print("[Error] ", err)
			this.fail(Sprintf("%s", err))
		}
	}()

	// -登录
	if this.c.Request.Method == "POST" {
		uname := this.c.PostForm("username")
		pwd := this.c.PostForm("password")

		ok, user := model.User.FindFirst(model.ModelCond{Where: "account=?", Bind: uname})
		if ok && user.Pwd == pwd {
			this.sendTicket(&user)
			this.c.Redirect(302, "/user")
		} else {
			this.setRetMsg("error pwd or account").succ(nil, "login.tpl")
		}

		return
	}

	// -登录页
	if this.user == nil {
		this.succ(Object{}, "login.tpl")
	} else {
		if this.user.Lose+this.user.Win != 0 {
			this.user.WinRate = Sprintf("%.2f", float32(this.user.Win)/float32(this.user.Win+this.user.Lose)*100)
		}
		this.succ(this.user, "personal.tpl")
	}

}

func (this *IndexSt) sendTicket(user *model.UserSt) {
	user.LoginTime = int(time.Now().Unix())
	userencode, _ := json.Marshal(user)
	this.session.Set(Sprintf("UserAuth:%d", user.Id), userencode)

	aes, _ := plugin.AesEncrypt(Sprintf("%d-%d", user.Id, user.LoginTime))
	this.c.SetCookie("ticket", string(user.Id)+","+string(aes), 86400, "", "", false, false)
}

func (this *IndexSt) Self() string {
	return this.Info
}

func (this *IndexSt) Test() (str string) {
	str = this.Info
	this.Info = this.Info + " 测试test"
	return
}
