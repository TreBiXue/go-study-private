package repo

import (
	"cloud.google.com/go/spanner"
	"context"
	"go-studying/models"
	"google.golang.org/api/iterator"
)

type spannerCenterRepository struct {
	client *spanner.Client
}

func NewSpannerCenterRepository(spanner *spanner.Client) ICenterRepo {
	return &spannerCenterRepository{spanner}
}

func (sp spannerCenterRepository) GetByCenterNo(ctx context.Context, centerNo string) (res models.Center, err error) {
	sql := `select CENTNO,CENTNM from CENTER_M where CENTNO = @centerNo`
	params := map[string]interface{}{"centerNo": centerNo}

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

	res = models.Center{}
	row.ToStruct(&res)
	if err != nil {
		// todo：具体原因log记录
		err = ErrorInternal
	}

	return
}
