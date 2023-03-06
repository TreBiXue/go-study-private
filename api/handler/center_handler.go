package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-studying/service"
	"net/http"
)

type CenterHandler struct {
	centerService service.ICenterService
}

func NewCenterHandlers(cs service.ICenterService) CenterHandler {
	return CenterHandler{centerService: cs}
}

func (ch *CenterHandler) GetCenterInfo(ctx *gin.Context) {
	centerNo := ctx.Query("center_no")

	res, err := ch.centerService.GetByCenterNo(ctx, centerNo)

	if err != nil {
		fmt.Printf("error %v", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"CenterInfo": res})

}
