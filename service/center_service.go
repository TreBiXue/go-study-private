package service

import (
	"context"
	"go-studying/models"
	"go-studying/repo"
)

type centerService struct {
	centerRepo repo.ICenterRepo
}

func NewCenterService(a repo.ICenterRepo) ICenterService {
	return &centerService{centerRepo: a}
}

func (c centerService) GetByCenterNo(ctx context.Context, centerNo string) (res models.Center, err error) {
	res, err = c.centerRepo.GetByCenterNo(ctx, centerNo)
	return
}

