package repo

import (
	"context"
	"go-studying/api/request"
	"go-studying/api/response"
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

type IProductRepo interface {
	GetByProductCD(ctx context.Context, productCD string) (res models.Product, err error)
}

type INyukaRepo interface {
	GetNyukaInfo(ctx context.Context, nyukaInfo *request.NyukaRequest) (res response.NyukaResponse, err error)
	UpdateNyukaInfo(ctx context.Context, nyukaInfo *request.NyukaRequest) (res response.NyukaResponse, err error)
}
