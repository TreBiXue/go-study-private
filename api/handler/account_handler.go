package handler

import (
	"fmt"
	"go-studying/api/request"
	"go-studying/api/response"
	"go-studying/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	AccountService service.IAccountService
}

func NewAccountHandlers(as service.IAccountService) AccountHandler {
	return AccountHandler{AccountService: as}
}

func (a *AccountHandler) LoginByID(c *gin.Context) {
	req := &request.LoginByIdRequest{}
	req.ID = c.Query("id")
	// if err := c.ShouldBind(&req); err != nil {
	// 	fmt.Printf("error %v", err)
	// 	return
	// }
	if err := req.Validate(); err != nil {
		fmt.Printf("error %v", err)
		return
	}

	// id := c.Query("id")
	// lastAccessed := time.Now()
	listAc, err := a.AccountService.Login(c, req.ID, &req.AccessTime)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}
	resp := &response.LoginByIDResponse{
		EmployeeCode: listAc[0].EmployeeCode,
		EmployeeName: listAc[0].EmployeeName,
	}
	c.JSON(http.StatusOK, gin.H{"Accounts": resp})
}
