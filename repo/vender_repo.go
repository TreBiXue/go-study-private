package repo

import (
	"context"
	"go-studying/models"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/spanner"
)

type spannerVenderRepository struct {
	client *spanner.Client
}

func NewSpannerVenderRepository(spanner *spanner.Client) IVenderRepo {
	return &spannerVenderRepository{spanner}
}

func (sp *spannerVenderRepository) GetByVenderNo(ctx context.Context, venderNo string) (res models.Vender, err error) {
	sql := `select SIRENO,SIRENM from SIRE_M where SIRENO = @venderNo`
	params := map[string]interface{}{"venderNo": venderNo}

	iter := sp.client.Single().Query(ctx, spanner.Statement{
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
	res = models.Vender{}
	row.ToStruct(&res)

	if err != nil {
		// todo：具体原因log记录
		err = ErrorInternal
	}

	return
}
