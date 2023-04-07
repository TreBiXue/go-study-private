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

/*
GetCenterInfo @Summary Get user by ID
@Description Get user details by ID
@Tags users
@ID getUserById
@Accept json
@Produce json
@Param	center_no	query	string	false	"必須ではありません。"
@Success 200 {object} User
@Router /api/v1/center [get]
*/
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
