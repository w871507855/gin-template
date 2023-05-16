package router

import "github.com/w871507855/gin-template/server/router/ping"

type RouterGroup struct {
	PingRouterGroup ping.PingRouter
}

var RouterGroupApp = new(RouterGroup)
