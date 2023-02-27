package repo

import (
	"context"
	"go-studying/models"
)

type IAccountRepo interface {
	GetByID(ctx context.Context, ID string) (res []models.Account, err error)
}
