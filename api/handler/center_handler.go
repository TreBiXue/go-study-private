package handler

import (
	"fmt"
	"go-studying/service"

	"github.com/gin-gonic/gin"
)

type CenterHandler struct {
	centerService service.ICenterService
}

func NewCenterHandlers(cs service.ICenterService) CenterHandler {
	return CenterHandler{centerService: cs}
}

// @description aaaaaa
// @version 1.0
// @accept application/x-json-stream
// @param none query string false "必須ではありません。"
// @Success 200 {string} string    "ok"
// @router /GetCenterInfo/ [get]
func (ch *CenterHandler) GetCenterInfo(ctx *gin.Context) {
	centerNo := ctx.Query("center_no")

	// res, err := ch.centerService.GetByCenterNo(ctx, centerNo)
	_, err := ch.centerService.GetByCenterNo(ctx, centerNo)

	if err != nil {
		fmt.Printf("error %v", err)
		return
	}

	// ctx.JSON(http.StatusOK, gin.H{"CenterInfo": res})
	ctx.JSON(200, "ok")

}
