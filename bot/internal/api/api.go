package api

import (
	"june/botmgr/internal/api/routes"

	"github.com/gin-gonic/gin"
)

type API struct {
	router *gin.Engine
}

func New() *API {
	return &API{
		router: gin.Default(),
	}
}

func (a *API) Start() error {
	a.router.Use(ZLogger())

	brg := a.router.Group("/bot")
	routes.RegisterBotRoutes(brg)

	a.router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return a.router.Run(":8080")
}
