package errorhandler

import (
	"errors"
	"fmt"
	"go-studying/api/request"
	"go-studying/repo"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			var err error
			var errorType ErrorType
			var errorMessage string

			for _, e := range c.Errors {
				if err == nil {
					err = e
				} else {
					err = fmt.Errorf("%v\n%v", err, e)
				}
			}

			switch {
			case errors.Is(err, ErrUnauthorized):
				errorType = ErrUnauthorized
				errorMessage = err.Error()
			case errors.Is(err, repo.ErrorNotFound):
				errorType = ErrNotFound
				errorMessage = err.Error()
			case errors.Is(err, repo.ErrorInternal):
				errorType = ErrNotFound
				errorMessage = err.Error()
			case errors.Is(err, request.ErrorBadRequest):
				errorType = ErrBadParams
				errorMessage = err.Error()
			//case errors.Is(err, ErrDatabase):
			//	if spanner.ErrCode(err) == codes.NotFound {
			//		errorType = ErrNotFound
			//		errorMessage = err.Error()
			//	} else if spanner.ErrCode(err) == codes.AlreadyExists {
			//		errorType = ErrUniqueViolation
			//		errorMessage = err.Error()
			//	}
			default:
				errorType = ErrUnknown
				errorMessage = err.Error()
			}

			c.JSON(GetHTTPStatus(errorType), gin.H{
				"code": errorType,
				"msg":  errorMessage,
			})
		}
	}
}
