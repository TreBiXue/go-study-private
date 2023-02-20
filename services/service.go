package services

import (
	"go-studying/models"

	"github.com/gin-gonic/gin"
)

type IAccountService interface {
	GetByID(c *gin.Context, id string) (res []models.Account, err error)
}
