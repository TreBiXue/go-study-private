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

func (ns nyukaService) GetNyukaInfo(ctx context.Context, nyukaInfo *request.NyukaRequest) (res response.NyukaResponse, err error) {

	nyukai, _ := ns.nyukaRepo.GetNyukaInfo(ctx, nyukaInfo)

	productInfo, _ := ns.productRepo.GetByProductCD(ctx, nyukai.ITEMCD)

	res = response.NyukaResponse{ITEMCD: nyukaInfo.ITEMCD,
		ITEMNM: productInfo.ITEMNM,
		CENTNO: nyukai.CENTNO,
	}
	return
}

func (ns nyukaService) SetNyukaInfo(ctx context.Context, nyukaInfo *request.NyukaRequest) (res response.NyukaResponse, err error) {
	return
}
