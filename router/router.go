package router

import (
	"tconn/middleware"
	"tconn/router/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	router.Use(middleware.Cors())

	g1 := router.Group("")
	{
		g1.GET("/ws", api.WS)
	}

	router.Run(":8080")
}
