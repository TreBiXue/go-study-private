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

	ctx.JSON(http.StatusOK, gin.H{"CountInfo": response})

}

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

	ctx.JSON(http.StatusOK, gin.H{"CountInfo": response})

}

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

	ctx.JSON(http.StatusOK, gin.H{"CountInfo": response})

}
