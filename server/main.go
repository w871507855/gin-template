package main

import (
	"github.com/w871507855/gin-template/server/global"
	"github.com/w871507855/gin-template/server/initialize"
)

func main() {
	r := initialize.Routers()
	initialize.Viper()
	panic(r.Run(global.CONF.System.Addr))
}
