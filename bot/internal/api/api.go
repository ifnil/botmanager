package api

import (
	"context"
	"june/botmgr/internal/api/handlers"
	"june/botmgr/internal/database"

	"github.com/gin-gonic/gin"
)

type API struct {
	router *gin.Engine
	ctx    context.Context
	cancel context.CancelFunc
}

func New(ctx context.Context, cancel context.CancelFunc) *API {
	return &API{
		router: gin.Default(),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (a *API) Start() error {
	db := database.New()
	if err := db.Init(); err != nil {
		return err
	}

	a.router.Use(ZLogger(db.DB()))

	handlers.NewBotHandler(db).Routes(a.router)
	a.router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return a.router.Run(":8080")
}
