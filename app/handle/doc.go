package handle

import (

	"github.com/gin-gonic/gin"
)

// main page
func Rule(context *gin.Context) {
	context.HTML(200, "static.tpl", Object{"data":"rule"})
}

func Ss(context *gin.Context) {
	context.HTML(200, "static.tpl", Object{"data":"ss"})
}