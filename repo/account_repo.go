package repo

import (
	"context"
	"go-studying/models"

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
