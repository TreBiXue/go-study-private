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
