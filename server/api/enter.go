package api

import "github.com/w871507855/gin-template/server/api/v1/ping"

type ApiGroup struct {
	PingApiGroup ping.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
