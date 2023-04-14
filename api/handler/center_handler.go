package handler

import (
	"go-studying/api/request"
	"go-studying/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CenterHandler struct {
	centerService service.ICenterService
}

func NewCenterHandlers(cs service.ICenterService) CenterHandler {
	return CenterHandler{centerService: cs}
}

// @Description Get center details by ID
// @Accept json
// @Produce json
// @Param	center_no	query	string	true	"center_no"
// @Success 200 {object} models.Center
// @Router /api/v1/center [get]
func (ch *CenterHandler) GetCenterInfo(ctx *gin.Context) {
	centerNo := ctx.Query("center_no")

	if centerNo == "" {
		_ = ctx.Error(request.ErrorBadRequest)
		return
	}

	res, err := ch.centerService.GetByCenterNo(ctx, centerNo)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"CenterInfo": res})
}
