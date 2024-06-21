package handlers

import (
	"june/botmgr/internal/bot"
	"june/botmgr/internal/database"
	"june/botmgr/pkg/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BotHandler struct {
	db *database.Database
}

func NewBotHandler(db *database.Database) *BotHandler {
	return &BotHandler{db: db}
}

func (bh *BotHandler) Routes(r *gin.Engine) *gin.RouterGroup {
	br := r.Group("/bot")
	{
		br.POST("/create", bh.Create())
		br.GET("/status/:id", bh.Status())
	}

	return br
}

func (bh *BotHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.NewBotRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse body"})
			return
		}

		b := bot.New().SetName(req.Name)

		if _, err := bh.db.DB().Exec("insert into bot (name, domain) values (?, ?)", b.Name(), "test.com"); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, req)
	}
}

func (bh *BotHandler) Status() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		info := bh.db.DB().QueryRow("select status from bot where id = ?", id)

		var res int
		if err := info.Scan(&res); err != nil {
			ctx.JSON(http.StatusTeapot, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"bot_status": res})
	}
}
