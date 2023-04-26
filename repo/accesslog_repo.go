package repo

import (
	"context"
	"errors"
	"go-studying/models"

	"cloud.google.com/go/spanner"
	"github.com/google/uuid"
)

type spannerAccessLogRepository struct {
	client *spanner.Client
}

func NewSpannerAccesslogRepository(spanner *spanner.Client) IAccessLogRepo {
	return &spannerAccessLogRepository{spanner}
}

func (sp *spannerAccessLogRepository) SaveAccessLog(ctx context.Context, record models.AccessLog) (err error) {
	sql := `insert AccessLog (Id,Url,Path,Ua,Status,Elapsed) VALUES(@id,@url,@path,@ua,@status,@elapsed)`
	params := map[string]interface{}{"id": uuid.New().String(),
		"url":     record.URL,
		"path":    record.Path,
		"ua":      record.Ua,
		"status":  record.Status,
		"elapsed": record.Elapsed}

	stmt := spanner.Statement{SQL: sql, Params: params}
	_, err = sp.client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		row, err := txn.Update(ctx, stmt)

		if err != nil {
			return err
		}
		if row == 0 {
			err = errors.New("更新失敗！")
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
