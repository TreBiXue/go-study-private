package router

import (
	"go-studying/api/handler"
	"go-studying/middleware"

	"github.com/gin-gonic/gin"
	// _ "openapi/docs"
)

// func InitCenterRouter(db *spanner.Client, router *gin.Engine) {
// 	iCenterRepo := repo.NewSpannerCenterRepository(db)

// 	iCenterService := service.NewCenterService(iCenterRepo)
// 	centerHandler := handler.NewCenterHandlers(iCenterService)

// 	groupRouter := router.Group("api/v1")
// 	{
// 		groupRouter.GET("/center", centerHandler.GetCenterInfo)
// 	}

// }

type CenterRouter struct {
	centerHandler handler.CenterHandler
}

func (r *CenterRouter) Center(router *gin.Engine) {
	groupRouter := router.Group("api/v1")
	{
		groupRouter.Use(middleware.AuthJWT())
		groupRouter.GET("/center", r.centerHandler.GetCenterInfo)
	}

}

func NewCenterRouter(centerHandler handler.CenterHandler) *CenterRouter {
	centerrouter := &CenterRouter{centerHandler: centerHandler}
	return centerrouter
}
