package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/w871507855/gin-template/server/response"
)

type PingApi struct {
}

func (*PingApi) Ping(c *gin.Context) {
	response.SuccessWithMessage(c, "pong")
}
