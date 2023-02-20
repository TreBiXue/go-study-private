package services

import (
	"go-studying/models"
	repo "go-studying/repos"

	"github.com/gin-gonic/gin"
)

type accountService struct {
	accountRepo repo.IAccountRepo
}

func NewAccountService(a repo.IAccountRepo) IAccountService {
	return &accountService{accountRepo: a}
}

func (a *accountService) GetByID(c *gin.Context, id string) (res []models.Account, err error) {
	res, err = a.accountRepo.GetByID(c, id)
	if err != nil {
		return nil, err
	}
	return
}
