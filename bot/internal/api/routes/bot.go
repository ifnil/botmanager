package routes

import (
	"june/botmgr/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterBotRoutes(rg *gin.RouterGroup) {
	rg.POST("/new", handlers.NewBot())
	rg.GET("/status/:id", handlers.BotStatus())
}
