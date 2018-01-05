package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "fmt"
    Str "strings"
    "log"
    
    . "avalon/plugin/selftype"
)

var Index IndexSt

func init(){
    Index = IndexSt{Info:"[圣杯战争] [刺客练刀房] [莫甘娜演技培训班] [神民] [叫嚣全场狼美人] [叫嚣的忠臣奥博伦]"}
}

/**
    Hanlde: Index Class
**/
type IndexSt struct { 
    Info string
    BaseSt;
}



// main page
func (this *IndexSt) Main(context *gin.Context){
    panic(" 测试啊 ")
    defer func() { 
        if err := recover(); err != nil {log.Print("[Warning] ", err)}
    }()
    
    var data = Object{
        "game_name" : "阿瓦隆",
        "game_slogn" : Str.Split(this.Info, " "),
        "players" : "玩家",
        "game" : "game",
    }
    

    context.HTML(http.StatusOK, "index.tpl", gin.H{
        "msg": "成功",
        "data": data,
        "code": 0,
    })
}

func (this *IndexSt) Self() string{
    return this.Info
}

func (this *IndexSt) Test() (str string){
    str = this.Info
    this.Info = this.Info + " 测试test"
    return 
}



