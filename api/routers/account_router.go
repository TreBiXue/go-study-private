package routers

import (
	"go-studying/api/handlers"
	repo "go-studying/repos"
	"go-studying/services"

	"cloud.google.com/go/spanner"
	"github.com/gin-gonic/gin"
)

func InitAccountRouter(db *spanner.Client, router *gin.Engine) {
	iAccountRepo := repo.NewSpannerAccountRepository(db)

	iAccountService := services.NewAccountService(iAccountRepo)
	accountHandler := handlers.NewAccountHandlers(iAccountService)

	groupRouter := router.Group("api/v1")

	groupRouter.GET("/login", accountHandler.LoginByID)

}
