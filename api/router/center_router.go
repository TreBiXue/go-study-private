package router

import (
	"go-studying/api/handler"

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
		groupRouter.GET("/center", r.centerHandler.GetCenterInfo)
		// groupRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

}

func NewCenterRouter(centerHandler handler.CenterHandler) *CenterRouter {
	centerrouter := &CenterRouter{centerHandler: centerHandler}
	return centerrouter
}
