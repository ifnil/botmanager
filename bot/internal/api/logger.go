package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ZLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.
			Info().
			Str("addr", c.RemoteIP()).
			Str("path", c.Request.URL.Path).
			Msg("new request")

		c.Next()
	}
}
