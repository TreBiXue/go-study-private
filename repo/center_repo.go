package repo

import (
	"cloud.google.com/go/spanner"
	"context"
	"go-studying/models"
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

	if err != nil {
		return models.Center{}, err
	}

	res = models.Center{}
	row.ToStruct(&res)

	if err != nil {
		return models.Center{}, err
	}

	return
}
