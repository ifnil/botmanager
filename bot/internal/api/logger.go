package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ZLogger(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		log.
			Info().
			Str("addr", c.RemoteIP()).
			Str("path", c.Request.URL.Path).
			Msg("new request")

		cc := c.Copy()

		stmt, err := db.Prepare("insert into http_log (client_ip, remote_ip, method, url, host, path, body) values (?,?,?,?,?,?,?)")
		if err != nil {
			return
		}

		go func(stmt *sql.Stmt) {
			_, err = stmt.Exec(cc.ClientIP(),
				cc.RemoteIP(),
				cc.Request.Method,
				cc.Request.URL.String(),
				cc.Request.URL.Host,
				cc.FullPath(),
				"")
			if err != nil {
				log.Error().Err(err).Msg("error inserting log")
			}
		}(stmt)
	}
}
