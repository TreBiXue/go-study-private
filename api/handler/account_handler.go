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

// @Description Get account details by ID
// @Accept json
// @Produce json
// @Param	id	query	string	true	"id"
// @Success 200 {object} []models.Account
// @Router /api/v1/login [get]
func (a *AccountHandler) LoginByID(c *gin.Context) {
	req := &request.LoginByIdRequest{}
	req.ID = c.Query("id")

	if err := req.Validate(); err != nil {
		fmt.Printf("error %v", err)
		return
	}

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
