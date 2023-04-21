package app

import (
	"go-studying/middleware"
	"go-studying/middleware/errorhandler"

	"github.com/gin-gonic/gin"
)

func NewGinEngine() *gin.Engine {
	r := gin.Default()
	r.Use(errorhandler.ErrorHandler())
	r.Use(middleware.RecordUaAndTime)
	return r
}
