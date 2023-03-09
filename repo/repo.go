package repo

import (
	"context"
	"go-studying/models"
	"time"
)

type IAccountRepo interface {
	GetByID(ctx context.Context, id string) (res []models.Account, err error)
	UpdateLastAccessedByID(ctx context.Context, id string, lastAccessed *time.Time) error
}

type ICenterRepo interface {
	GetByCenterNo(ctx context.Context, centerNo string) (res models.Center, err error)
}

type IVenderRepo interface {
	GetByVenderNo(ctx context.Context, venderNo string) (res models.Vender, err error)
}
