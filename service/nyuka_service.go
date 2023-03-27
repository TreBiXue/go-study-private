package service

import (
	"context"
	"go-studying/api/request"
	"go-studying/api/response"
	"go-studying/repo"
)

type nyukaService struct {
	productRepo repo.IProductRepo
	nyukaRepo   repo.INyukaRepo
}

func NewNyukaService(prod repo.IProductRepo, nyuka repo.INyukaRepo) INyukaService {
	return &nyukaService{productRepo: prod,
		nyukaRepo: nyuka}

}

func (ns nyukaService) GetNyukaCount(ctx context.Context, nyukaInput *request.NyukaInputRequest) (res response.NyukaInputResponse, err error) {
	res, err = ns.nyukaRepo.GetNyukaCount(ctx, nyukaInput)
	// ORDERNO重複対応
	return
}

func (ns nyukaService) GetNyukaJANInfo(ctx context.Context, nyukaJAN *request.NyukaInputJANRequest) (res response.NyukaInputJANResponse, err error) {
	res, err = ns.nyukaRepo.GetNyukaJANInfo(ctx, nyukaJAN)
	// JANによる商品名取得
	// SIRENOによる仕入れ名取得
	return
}

func (ns nyukaService) UpdateNyukaInfo(ctx context.Context, nyukaJANJISU *request.NyukaInputJANNKAJISURequest) (IsSuccess bool, err error) {
	IsSuccess = false
	err = ns.nyukaRepo.UpdateNyukaInfo(ctx, nyukaJANJISU)
	if err != nil {
		return
	}
	IsSuccess = true
	return

}
