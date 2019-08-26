package handle

import (
	"avalon/app/model"
	"avalon/plugin"
	"fmt"
	"github.com/gin-gonic/gin"

	_ "log"
	"net/http"
)

var secrets = Object{
	"haibo": Object{"email": "haibo@qq.com", "phone": "13720044402", "pwd": "haibo"},
	"qiqi":  Object{"email": "qiqi@qq.com", "phone": "11111", "pwd": "qiqi"},
	"z":     Object{"email": "lena@guapa.com", "phone": "523443"},
}

type BaseSt struct {
	Data    Object
	c       *gin.Context
	auth    bool
	user    *model.UserSt
	session *plugin.Redis
}

type Object map[string]interface{}

type BaseI interface {
	BeforeHandle(*gin.Context)
}

func (this *BaseSt) BeforeHandle(c *gin.Context) {
	this.c = c
	this.session, _ = plugin.NewRedis("")

	// response
	this.Data["msg"] = ""
	this.Data["code"] = 0

	// auth cookie
	if this.auth != false {
		ticket, err := this.c.Cookie("ticket")
		if err != nil {
			this.setRetCode(-401).failPage("auth valid!")
			return
		}

		user := plugin.UserAuth(ticket)
		if user == nil {
			this.setRetCode(-401).failPage("auth valid!")
			return
		}

		this.user = user
	}

}

func (this *BaseSt) setRetCode(code int) *BaseSt {
	this.Data["code"] = code
	return this
}

func (this *BaseSt) setRetMsg(msg string) *BaseSt {
	this.Data["msg"] = msg
	return this
}

func (this *BaseSt) fail(msg interface{}) {

	this.Data["data"] = Object{}
	this.Data["msg"] = msg.(string)

	if this.Data["code"] == 0 {
		this.Data["code"] = -1
	}

	this.c.JSON(http.StatusBadRequest, this.Data)
	this.c.Abort()
}

func (this *BaseSt) failPage(msg interface{}) {

	this.Data["data"] = Object{}
	this.Data["msg"] = msg.(string)
	if this.Data["code"] == 0 {
		this.Data["code"] = -1
	}

	this.c.HTML(http.StatusInternalServerError, "error.tpl", this.Data)
	this.c.Abort()
}

func (this *BaseSt) succ(data interface{}, tpl_name string) {

	this.Data["data"] = data

	fmt.Printf("\n respons INFO: %+v \n", data)

	if tpl_name == "" {
		this.c.JSON(http.StatusOK, this.Data)
		this.c.Abort()
	} else {
		this.c.HTML(http.StatusOK, tpl_name, this.Data)
		this.c.Abort()
	}
}

func (this *BaseSt) isAjax() bool {
	return this.c.GetHeader("X-Requested-With") == "XMLHttpRequest"
}

// func (this *BaseSt) SessionStart(sessionM *session.Manager) session.Store {
// 	if sessionM != nil {
// 		this.sessionManager = sessionM
// 		obj, err := this.sessionManager.SessionStart(this.c.Writer, this.c.Request)
// 		if err != nil {
// 			// logs.Error(err)
// 			panic("503 session fail!")
// 		}

// 		this.session = obj
// 	} else {
// 		if this.session == nil {
// 			obj, err := this.sessionManager.SessionStart(this.c.Writer, this.c.Request)
// 			if err != nil {
// 				// logs.Error(err)
// 				panic("503 session fail!")
// 			}

// 			this.session = obj
// 		}
// 	}

// 	return this.session
// }

// func (this *BaseSt) SetSession(name interface{}, value interface{}) {
// 	if this.session == nil {
// 		this.SessionStart(nil)
// 	}
// 	this.session.Set(name, value)
// }

// func (this *BaseSt) GetSession(name interface{}) interface{} {
// 	if this.session == nil {
// 		this.SessionStart(nil)
// 	}
// 	return this.session.Get(name)
// }

// func (this *BaseSt) DelSession(name interface{}) {
// 	if this.session == nil {
// 		this.SessionStart(nil)
// 	}
// 	this.session.Delete(name)
// }

// // SessionRegenerateID regenerates session id for this session.
// // the session data have no changes.
// func (this *BaseSt) SessionRegenerateID() {
// 	if this.session != nil {
// 		this.session.SessionRelease(this.c.Writer)
// 	}
// 	this.session = this.sessionManager.SessionRegenerateID(this.c.Writer, this.c.Request)
// }

// // DestroySession cleans session data and session cookie.
// func (this *BaseSt) DestroySession() {
// 	this.session.Flush()
// 	this.session = nil
// 	this.sessionManager.SessionDestroy(this.c.Writer, this.c.Request)
// }

// func (this *BaseSt) sessionRelease() {
// 	if this.session != nil {
// 		this.session.SessionRelease(this.c.Writer)
// 	}
// }

// func init(){

// }
