package service

import (
	"context"
	"go-studying/models"
	"time"
)

type IAccountService interface {
	Login(ctx context.Context, id string, lastAccessed *time.Time) (res []models.Account, err error)
}
