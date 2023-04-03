package router

import (
	"go-studying/api/handler"

	"github.com/gin-gonic/gin"
)

type NyukaRouter struct {
	nyukaHandler handler.NyukaHandler
}

func (vr *NyukaRouter) Nyuka(router *gin.Engine) {
	groupRouter := router.Group("api/v1")
	{
		groupRouter.POST("/nyuka/getcounts", vr.nyukaHandler.GetNyukaCount)
		groupRouter.POST("/nyuka/getjaninfo", vr.nyukaHandler.GetNyukaJANInfo)
		groupRouter.PATCH("/nyuka/update", vr.nyukaHandler.UpdateNyukaInfo)
	}

}

func NewNyukaRouter(handler handler.NyukaHandler) *NyukaRouter {
	nyukarRouter := &NyukaRouter{nyukaHandler: handler}
	return nyukarRouter
}
