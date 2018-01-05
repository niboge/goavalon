package handle

import (
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "fmt"
    _"log"
    . "avalon/plugin/selftype"
)


type BaseSt struct {
    // return http-client
    Data Object
}

func (this *BaseSt) Fail(msg string, code int) {
    this.Data["data"] = Object{}
    this.Data["msg"] = msg
    this.Data["code"] = code
}

func (this *BaseSt) IsAjax() bool {
    // return input.Header("X-Requested-With") == "XMLHttpRequest"
    return true
}

// func init(){

// }

