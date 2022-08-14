package main

import (
	"tconn/router"
	"tconn/service/ws"
	"tconn/util/log"
)

func init() {
	log.Setup()
	ws.Setup()
}

func main() {
	router.InitRouter()
}
