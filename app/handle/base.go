package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	. "avalon/plugin/selftype"
	_ "log"
)

var secrets = Object{
	"foo":    Object{"email": "foo@bar.com", "phone": "123433"},
	"austin": Object{"email": "austin@example.com", "phone": "666"},
	"lena":   Object{"email": "lena@guapa.com", "phone": "523443"},
}


type BaseSt struct{
	Data Object
	c *gin.Context
	auth bool
}

type BaseI interface{
	BeforeHandle(*gin.Context)
}

func (this *BaseSt) BeforeHandle(c *gin.Context) {
	// context
	if this.auth != false {
		gin.BasicAuth(gin.Accounts{
			"foo":    "bar",
			"austin": "1234",
			"lena":   "hello2",
			"manu":   "4321",
		})

		user,_ := c.GetQuery(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, Object{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, Object{"user": user, "secret": "NO SECRET :("})
		}
	}

	this.c = c
}

func (this *BaseSt) setRetCode(code int) *BaseSt{
	this.Data["code"] = code
	return this
}

func (this *BaseSt) fail(msg interface{}) {
	this.Data["data"] = Object{}
	this.Data["msg"] = msg.(string)
	this.Data["code"] = -1

	if this.isAjax() {
		this.c.JSON(http.StatusBadRequest, this.Data)		
	}else{
		this.c.String(http.StatusInternalServerError, msg.(string))
	}
}

func (this *BaseSt) succ(data Object, tpl_name string) {
	this.Data["data"] = data
	this.Data["msg"] = ""

	fmt.Printf("\n%V\n",this.Data["code"])

	this.Data["code"] = 0

	this.c.HTML(http.StatusOK, tpl_name, this.Data)
}

func (this *BaseSt) isAjax() bool {
	// return input.Header("X-Requested-With") == "XMLHttpRequest"
	return true
}

// func init(){

// }
