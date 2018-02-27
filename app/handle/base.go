package handle

import (
	. "avalon/plugin/selftype"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/session"

	_ "log"
	"net/http"
)

var secrets = Object{
	"haibo":    Object{"email": "haibo@qq.com", "phone": "13720044402", "pwd":"haibo"},
	"qiqi": Object{"email": "qiqi@qq.com", "phone": "11111","pwd":"qiqi"},
	"z":   Object{"email": "lena@guapa.com", "phone": "523443"},
}

type BaseSt struct {
	Data Object
	c    *gin.Context
	auth bool
	
	sessionManager *session.Manager
	session session.Store
}

type BaseI interface {
	BeforeHandle(*gin.Context, *session.Manager)
}

func (this *BaseSt) BeforeHandle(c *gin.Context, sessionM *session.Manager) {

	// context
	this.c = c

	// auth
	if this.auth != false {
		user, _ := this.c.GetQuery(gin.AuthUserKey)
		if secret, ok := secrets[user]; ok {
			if secret.(Object)["pwd"] != this.c.Query("pwd") {
				this.failPage("valid auth pwd")
			}
		} else {
			this.failPage("valid auth")
		}
	}

	// session
	this.sessionManager = sessionM
}

func (this *BaseSt) setRetCode(code int) *BaseSt {
	this.Data["code"] = code
	return this
}

func (this *BaseSt) fail(msg interface{}) {
	defer this.sessionRelease()

	this.Data["data"] = Object{}
	this.Data["msg"] = msg.(string)
	this.Data["code"] = -1
	
	this.c.JSON(http.StatusBadRequest, this.Data)
}

func (this *BaseSt) failPage(msg interface{}) {
	defer this.sessionRelease()

	this.Data["data"] = Object{}
	this.Data["msg"] = msg.(string)
	this.Data["code"] = -1

	this.c.HTML(http.StatusInternalServerError, "error.tpl", this.Data)
	this.c.Abort()
}

func (this *BaseSt) succ(data interface{}, tpl_name string) {
	defer this.sessionRelease()

	this.Data["data"] = data
	this.Data["msg"] = ""

	fmt.Printf("\n INFO: %V %v \n", data)

	this.Data["code"] = 0

	this.c.HTML(http.StatusOK, tpl_name, this.Data)
}

func (this *BaseSt) isAjax() bool {
	// return input.Header("X-Requested-With") == "XMLHttpRequest"
	return true
}


func (this *BaseSt) SessionStart() session.Store {
	if this.session == nil {
		obj, err := this.sessionManager.SessionStart(this.c.Writer, this.c.Request)
		if err != nil {
			// logs.Error(err)
			panic("503 session fail!")
		}

		this.session = obj
	}

	return this.session
}

func (this *BaseSt) SetSession(name interface{}, value interface{}) {
	if this.session == nil {
		this.SessionStart()
	}
	this.session.Set(name, value)
}

func (this *BaseSt) GetSession(name interface{}) interface{} {
	if this.session == nil {
		this.SessionStart()
	}
	return this.session.Get(name)
}

func (this *BaseSt) DelSession(name interface{}) {
	if this.session == nil {
		this.SessionStart()
	}
	this.session.Delete(name)
}

// SessionRegenerateID regenerates session id for this session.
// the session data have no changes.
func (this *BaseSt) SessionRegenerateID() {
	if this.session != nil {
		this.session.SessionRelease(this.c.Writer)
	}
	this.session = this.sessionManager.SessionRegenerateID(this.c.Writer, this.c.Request)
}

// DestroySession cleans session data and session cookie.
func (this *BaseSt) DestroySession() {
	this.session.Flush()
	this.session = nil
	this.sessionManager.SessionDestroy(this.c.Writer, this.c.Request)
}

func (this *BaseSt) sessionRelease() {
	if this.session != nil {
		this.session.SessionRelease(this.c.Writer)
	}
}
// func init(){

// }
