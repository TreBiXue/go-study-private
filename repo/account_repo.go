package repo

import (
	"context"
	"errors"
	"go-studying/models"
	"google.golang.org/api/iterator"
	"time"

	"cloud.google.com/go/spanner"
)

type spannerAccountRepository struct {
	spanner_client *spanner.Client
}

func NewSpannerAccountRepository(spanner *spanner.Client) IAccountRepo {
	return &spannerAccountRepository{spanner}
}

func (sp *spannerAccountRepository) GetByID(ctx context.Context, id string) (res models.Account, err error) {
	sql := `select ID,EmployeeCode,Token,EmployeeName,LastAccessed from Operators where id = @id`
	params := map[string]interface{}{"id": id}

	iter := sp.spanner_client.Single().Query(ctx, spanner.Statement{
		SQL:    sql,
		Params: params,
	})
	defer iter.Stop()
	row, err := iter.Next()

	if err == iterator.Done {
		err = ErrorNotFound
		return
	} else if err != nil {
		// todo：具体原因log记录
		err = ErrorInternal
		return
	}

	res = models.Account{}
	row.ToStruct(&res)
	if err != nil {
		// todo：具体原因log记录
		err = ErrorInternal
	}

	return
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
