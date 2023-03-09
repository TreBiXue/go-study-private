package router

import (
	"go-studying/api/handler"

	"github.com/gin-gonic/gin"
)

type VenderRouter struct {
	venderHandler handler.VenderHandler
}

func (vr *VenderRouter) Vender(router *gin.Engine) {
	groupRouter := router.Group("api/v1")
	{
		groupRouter.GET("/vender", vr.venderHandler.GetVenderInfo)
	}

}

func NewVenderRouter(venderHandler handler.VenderHandler) *VenderRouter {
	venderRouter := &VenderRouter{venderHandler: venderHandler}
	return venderRouter
}
