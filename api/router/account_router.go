package router

import (
	"go-studying/api/handler"

	"github.com/gin-gonic/gin"
)

// func InitAccountRouter(db *spanner.Client, router *gin.Engine) {
// 	iAccountRepo := repo.NewSpannerAccountRepository(db)

// 	iAccountService := service.NewAccountService(iAccountRepo)
// 	accountHandler := handler.NewAccountHandlers(iAccountService)

// 	groupRouter := router.Group("api/v1")
// 	{
// 		groupRouter.GET("/login", accountHandler.LoginByID)
// 	}

// }

type AccountRouter struct {
	accountHandler handler.AccountHandler
}

func (r *AccountRouter) Login(router *gin.Engine) {
	groupRouter := router.Group("api/v1")
	{
		groupRouter.GET("/login", r.accountHandler.LoginByID)
	}

}

func NewAccountRouter(accountHandler handler.AccountHandler) *AccountRouter {
	accountrouter := &AccountRouter{accountHandler: accountHandler}
	return accountrouter
}
