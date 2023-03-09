package service

import (
	"context"
	"go-studying/models"
	"go-studying/repo"
)

type venderService struct {
	venderRepo repo.IVenderRepo
}

func NewVenderService(vr repo.IVenderRepo) IVenderService {
	return &venderService{venderRepo: vr}
}

func (vr *venderService) GetByVenderNo(ctx context.Context, venderNo string) (res models.Vender, err error) {
	res, err = vr.venderRepo.GetByVenderNo(ctx, venderNo)
	return
}
