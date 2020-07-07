package main

import (
	"blog/pkg/router"
	"blog/pkg/util"
)

func main() {
	util.SystemRun()
	engine := router.InitRouter()
	engine.Run()
}
