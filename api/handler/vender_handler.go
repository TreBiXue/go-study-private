package handler

import (
	"fmt"
	"go-studying/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VenderHandler struct {
	venderService service.IVenderService
}

func NewVenderHandlers(vs service.IVenderService) VenderHandler {
	return VenderHandler{venderService: vs}
}

// @Description Get vender details by ID
// @Accept json
// @Produce json
// @Param	vender_no	query	string	true	"vender_no"
// @Success 200 {object} models.Vender
// @Router /api/v1/vender [get]
func (vh *VenderHandler) GetVenderInfo(ctx *gin.Context) {
	venderNo := ctx.Query("vender_no")

	res, err := vh.venderService.GetByVenderNo(ctx, venderNo)

	if err != nil {
		fmt.Printf("error %v", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"venderInfo": res})

}
