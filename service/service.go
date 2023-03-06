package service

import (
	"context"
	"go-studying/models"
	"time"
)

type IAccountService interface {
	Login(ctx context.Context, id string, lastAccessed *time.Time) (res []models.Account, err error)
}

type ICenterService interface {
	GetByCenterNo(ctx context.Context, centerNo string) (res models.Center, err error)
}
