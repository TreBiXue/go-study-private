package handler

import (
	"fmt"
	"go-studying/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	AccountService service.IAccountService
}

func NewAccountHandlers(as service.IAccountService) AccountHandler {
	return AccountHandler{AccountService: as}
}

func (a *AccountHandler) LoginByID(c *gin.Context) {
	id := c.Query("id")
	lastAccessed := time.Now()
	listAc, err := a.AccountService.Login(c, id, &lastAccessed)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Accounts": listAc})
}
