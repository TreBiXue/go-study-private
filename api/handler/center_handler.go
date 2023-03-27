package handler

import (
	"fmt"
	"go-studying/service"

	"github.com/gin-gonic/gin"
)

// @title APIドキュメントのタイトル
// @version バージョン(1.0)
// @description 仕様書に関する内容説明
// @termsOfService 仕様書使用する際の注意事項

// @host localhost:8080
// @BasePath /

type CenterHandler struct {
	centerService service.ICenterService
}

func NewCenterHandlers(cs service.ICenterService) CenterHandler {
	return CenterHandler{centerService: cs}
}

// @description テスト用APIの詳細
// @version 1.0
// @accept application/x-json-stream
// @param none query string false "必須ではありません。"
// @Success 200 {string} string    "ok"
// @router /api/v1/center [get]
func (ch *CenterHandler) GetCenterInfo(ctx *gin.Context) {
	centerNo := ctx.Query("center_no")

	//res, err := ch.centerService.GetByCenterNo(ctx, centerNo)
	_, err := ch.centerService.GetByCenterNo(ctx, centerNo)

	if err != nil {
		fmt.Printf("error %v", err)
		return
	}

	// ctx.JSON(http.StatusOK, gin.H{"CenterInfo": res})
	ctx.JSON(200, "ok")

}

// // @description テスト用APIの詳細
// // @version 1.0
// // @accept application/x-json-stream
// // @param none query string false "必須ではありません。"
// // @Success 200 {string} string    "ok"
// // @router /test/ [get]
// func test(c *gin.Context) {
// 	c.JSON(200, "ok")
// }
