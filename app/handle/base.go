package handle

import (
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "fmt"
	. "avalon/plugin/selftype"
	_ "log"
)

type BaseSt struct {
	Data Object
}

func (this *BaseSt) setRetCode(code int) {
	this.Data["code"] = code
}

func (this *BaseSt) Fail(msg interface{}) Object {
	this.Data["data"] = Object{}
	this.Data["msg"] = msg
	this.Data["code"] = -1

	return this.Data
}

func (this *BaseSt) Succ(data Object) (Data Object) {
	this.Data["data"] = data
	this.Data["msg"] = ""
	this.Data["code"] = 0
	return this.Data
}

func (this *BaseSt) IsAjax() bool {
	// return input.Header("X-Requested-With") == "XMLHttpRequest"
	return true
}

// func init(){

// }
