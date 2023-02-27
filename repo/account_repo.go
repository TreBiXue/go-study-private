package repo

import (
	"context"
	"errors"
	"go-studying/models"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

type spannerAccountRepository struct {
	spanner_client *spanner.Client
}

func NewSpannerAccountRepository(spanner *spanner.Client) IAccountRepo {
	return &spannerAccountRepository{spanner}
}

func (sp *spannerAccountRepository) GetByID(ctx context.Context, id string) (res []models.Account, err error) {
	sql := `select ID,EmployeeCode,Token,EmployeeName,LastAccessed from Operators where id = @id`
	params := map[string]interface{}{"id": id}

	iter := sp.spanner_client.Single().Query(ctx, spanner.Statement{
		SQL:    sql,
		Params: params,
	})
	defer iter.Stop()

	account_list := []models.Account{}

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var account models.Account
		if err := row.ToStruct(&account); err != nil {
			return nil, err
		}
		account_list = append(account_list, account)
	}
	return account_list, nil
}

func (sp *spannerAccountRepository) UpdateLastAccessedByID(ctx context.Context, id string, lastAccessed *time.Time) error {
	stmt := spanner.Statement{
		SQL: `UPDATE Operators
				SET LastAccessed = @lastAccessed
				WHERE ID = @id`,
		Params: map[string]interface{}{
			"id":           id,
			"lastAccessed": lastAccessed,
		},
	}

	_, err := sp.spanner_client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
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
