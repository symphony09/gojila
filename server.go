package main

import (
	"github.com/gin-gonic/gin"
	"github.com/symphony09/gojila/global"
	"github.com/symphony09/gojila/router"
)

func main() {
	defer func() {
		global.CloseDB()
	}()

	gin.ForceConsoleColor()
	g := gin.Default()
	router.Load(g)
	_ = g.Run(global.Config.GetString("address"))
}
