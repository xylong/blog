package main

import "blog/pkg/router"

func main() {
	engine := router.InitRouter()
	engine.Run()
}
