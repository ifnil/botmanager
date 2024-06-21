package common

import "github.com/gin-gonic/gin"

type Handler interface {
	Create() gin.HandlerFunc
}
