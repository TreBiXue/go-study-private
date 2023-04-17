package handler

import (
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

// @Description ユーザー登録
// @Accept json
// @Produce json
// @Param	id	query	string	true	"id"
// @Success 200 {object} []models.Account
// @Router /api/v1/login [get]
func (a *AccountHandler) LoginByID(ctx *gin.Context) {
	req := &request.LoginByIdRequest{}
	req.ID = ctx.Query("id")

	if req.ID == "" {
		_ = ctx.Error(request.ErrorBadRequest)
		return
	}

	res, err := a.AccountService.Login(ctx, req.ID, &req.AccessTime)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp := &response.LoginByIDResponse{
		EmployeeCode: res.EmployeeCode,
		EmployeeName: res.EmployeeName,
	}
	ctx.JSON(http.StatusOK, gin.H{"Accounts": resp})
}
