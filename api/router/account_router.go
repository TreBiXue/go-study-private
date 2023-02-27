package router

import (
	"go-studying/api/handler"
	"go-studying/repo"
	"go-studying/service"

	"cloud.google.com/go/spanner"
	"github.com/gin-gonic/gin"
)

func InitAccountRouter(db *spanner.Client, router *gin.Engine) {
	iAccountRepo := repo.NewSpannerAccountRepository(db)

	iAccountService := service.NewAccountService(iAccountRepo)
	accountHandler := handler.NewAccountHandlers(iAccountService)

	groupRouter := router.Group("api/v1")

	groupRouter.GET("/login", accountHandler.LoginByID)

}
