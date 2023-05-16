package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/w871507855/gin-template/server/api"
)

type PingRouter struct {
}

func (*PingRouter) InitPing(r *gin.Engine) {
	group := r.Group("/ping")
	apiGroup := api.ApiGroupApp.PingApiGroup
	group.GET("/", apiGroup.Ping)
}
