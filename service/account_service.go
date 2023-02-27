package service

import (
	"context"
	"go-studying/models"
	"go-studying/repo"
)

type accountService struct {
	accountRepo repo.IAccountRepo
}

func NewAccountService(a repo.IAccountRepo) IAccountService {
	return &accountService{accountRepo: a}
}

func (a *accountService) GetByID(c context.Context, id string) (res []models.Account, err error) {
	res, err = a.accountRepo.GetByID(c, id)
	if err != nil {
		return nil, err
	}
	return
}
