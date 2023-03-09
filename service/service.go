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

type IVenderService interface {
	GetByVenderNo(ctx context.Context, venderNo string) (res models.Vender, err error)
}
