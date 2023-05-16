package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/w871507855/gin-template/server/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	pingGroup := router.RouterGroupApp.PingRouterGroup
	pingGroup.InitPing(r)
	return r
}
