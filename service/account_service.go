package service

import (
	"context"
	"go-studying/models"
	"go-studying/repo"
	"time"
)

type accountService struct {
	accountRepo repo.IAccountRepo
}

func NewAccountService(a repo.IAccountRepo) IAccountService {
	return &accountService{accountRepo: a}
}

func (a *accountService) Login(ctx context.Context, id string, lastAccessed *time.Time) (res []models.Account, err error) {
	res, err = a.accountRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = a.accountRepo.UpdateLastAccessedByID(ctx, id, lastAccessed)
	if err != nil {
		return nil, err
	}

	return
}
