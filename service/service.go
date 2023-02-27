package service

import (
	"context"
	"go-studying/models"
)

type IAccountService interface {
	GetByID(c context.Context, id string) (res []models.Account, err error)
}
