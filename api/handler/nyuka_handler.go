package handler

import (
	"fmt"
	"go-studying/api/request"
	"go-studying/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NyukaHandler struct {
	nyukaService service.INyukaService
}

func NewNyukaHandlers(service service.INyukaService) NyukaHandler {
	return NyukaHandler{nyukaService: service}
}

//	@Summary 获取未入荷条数
//
// @accept			application/json
//
// @Param   input body request.NyukaInputRequest true "test"
// @Success	200			{object}	response.NyukaInputResponse	"成功"
// @Router		/api/v1/nyuka/getcounts [post]
func (ch *NyukaHandler) GetNyukaCount(ctx *gin.Context) {

	request := request.NyukaInputRequest{}
	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		fmt.Printf("ShouldBindJSON error %v", err)
		return
	}

	response, err := ch.nyukaService.GetNyukaCount(ctx, &request)

	if err != nil {
		fmt.Printf("GetNyukaCount error %v", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"入荷件数": response})

}

//	@Summary JANによって検品情報取得
//
// @accept			application/json
//
// @Param   input body request.NyukaInputJANRequest true "jan info"
// @Success	200			{object}	response.NyukaInputJANResponse	"成功"
// @Router		/api/v1/nyuka/getjaninfo [post]
func (ch *NyukaHandler) GetNyukaJANInfo(ctx *gin.Context) {

	request := request.NyukaInputJANRequest{}
	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		fmt.Printf("error %v", err)
		return
	}

	response, err := ch.nyukaService.GetNyukaJANInfo(ctx, &request)

	if err != nil {
		fmt.Printf("error %v", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"商品情報明細": response})

}

//	@Summary 入荷数を更新する
//
// @accept			application/json
//
// @Param   input body request.NyukaInputJANNKAJISURequest true "jisekisu　input"
// @Success	200	{boolean} IsSuccess	"成功"
// @Router		/api/v1/nyuka/update [PATCH]
func (ch *NyukaHandler) UpdateNyukaInfo(ctx *gin.Context) {

	request := request.NyukaInputJANNKAJISURequest{}
	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		fmt.Printf("error %v", err)
		return
	}

	if err := request.Validate(); err != nil {
		fmt.Printf("error %v", err)
		return
	}

	response, err := ch.nyukaService.UpdateNyukaInfo(ctx, &request)

	if err != nil {
		fmt.Printf("error %v", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"実績数更新結果": response})

}
