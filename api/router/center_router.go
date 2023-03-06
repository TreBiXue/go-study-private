package router

import (
	"github.com/gin-gonic/gin"
	"go-studying/api/handler"
)

// import (
// 	"go-studying/api/handler"
// 	"go-studying/repo"
// 	"go-studying/service"

// 	"cloud.google.com/go/spanner"
// 	"github.com/gin-gonic/gin"
// )

// func InitAccountRouter(db *spanner.Client, router *gin.Engine) {
// 	iAccountRepo := repo.NewSpannerAccountRepository(db)

// 	iAccountService := service.NewAccountService(iAccountRepo)
// 	accountHandler := handler.NewAccountHandlers(iAccountService)

// 	groupRouter := router.Group("api/v1")

// 	groupRouter.GET("/login", accountHandler.LoginByID)

// }

type CenterRouter struct {
	centerHandler handler.CenterHandler
}

func (r *CenterRouter) Center(router *gin.Engine) {
	groupRouter := router.Group("api/v1")
	{
		groupRouter.GET("/center", r.centerHandler.GetCenterInfo)
	}

}

func NewCenterRouter(centerHandler handler.CenterHandler) *CenterRouter {
	router := &CenterRouter{centerHandler: centerHandler}
	return router
}
