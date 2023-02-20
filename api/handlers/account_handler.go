package handlers

import (
	"fmt"
	"go-studying/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	AccountService services.IAccountService
}

func NewAccountHandlers(as services.IAccountService) AccountHandler {
	return AccountHandler{AccountService: as}
}

func (a *AccountHandler) LoginByID(c *gin.Context) {
	id := c.Query("id")
	listAc, err := a.AccountService.GetByID(c, id)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Accounts": listAc})
}
