package service

import (
	"context"
	"go-studying/api/request"
	"go-studying/api/response"
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

type INyukaService interface {
	GetNyukaCount(ctx context.Context, nyukaInput *request.NyukaInputRequest) (res response.NyukaInputResponse, err error)
	GetNyukaJANInfo(ctx context.Context, nyukaJAN *request.NyukaInputJANRequest) (res response.NyukaInputJANResponse, err error)
	UpdateNyukaInfo(ctx context.Context, nyukaJANJISU *request.NyukaInputJANNKAJISURequest) (IsSuccess bool, err error)
}
