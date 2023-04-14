package app

import (
	"github.com/gin-gonic/gin"
	"go-studying/middleware/errorhandler"
)

func NewGinEngine() *gin.Engine {
	r := gin.Default()
	r.Use(errorhandler.ErrorHandler())
	return r
}
