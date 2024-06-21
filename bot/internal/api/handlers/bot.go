package handlers

import (
	"fmt"
	"june/botmgr/pkg/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewBot() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.ContentType() != "application/json" {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Incorrect Content-Type"})
		}

		var newBotRequest request.NewBotRequest
		if err := c.BindJSON(&newBotRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse body"})
		}
	}
}

func BotStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("bot %s doing fine", id)})
	}
}
