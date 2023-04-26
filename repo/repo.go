package repo

import (
	"context"
	"errors"
	"go-studying/api/request"
	"go-studying/models"
	"time"
)

var (
	ErrorNotFound = errors.New("data not found")
	ErrorInternal = errors.New("internal error")
)

type IAccountRepo interface {
	GetByID(ctx context.Context, id string) (res models.Account, err error)
	UpdateLastAccessedByID(ctx context.Context, id string, lastAccessed *time.Time) error
}

type ICenterRepo interface {
	GetByCenterNo(ctx context.Context, centerNo string) (res models.Center, err error)
}

type IVenderRepo interface {
	GetByVenderNo(ctx context.Context, venderNo string) (res models.Vender, err error)
}

type IProductRepo interface {
	GetByProductCD(ctx context.Context, itemCD string) (res models.Product, err error)
}

type INyukaRepo interface {
	GetNyukaCount(ctx context.Context, nyukaInput *request.NyukaInputRequest) (res []models.NyukaInfo, err error)
	GetNyukaJANInfo(ctx context.Context, nyukaJAN *request.NyukaInputJANRequest) (res models.NyukaInfo, err error)
	UpdateNyukaInfo(ctx context.Context, nyukaJANJISU *request.NyukaInputJANNKAJISURequest) (err error)
}

type IAccessLogRepo interface {
	SaveAccessLog(ctx context.Context, record models.AccessLog) (err error)
}
