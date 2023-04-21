package middleware

import (
	"errors"
	"fmt"
	"go-studying/pkg/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Print("aaaa")
		header := ctx.GetHeader("Authorization")
		headerList := strings.Split(header, " ")
		if len(headerList) != 2 {
			err := errors.New("authorization 失敗")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return

		}
		t := headerList[0]
		content := headerList[1]

		if t != "Bearer" {
			err := errors.New("認証失敗 Bearerのみ")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		if _, err := token.Verify([]byte(content)); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
