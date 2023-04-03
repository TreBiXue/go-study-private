package repo

import (
	"cloud.google.com/go/spanner"
	"context"
	"go-studying/models"
)

type spannerProductRepository struct {
	client *spanner.Client
}

func NewSpannerProductRepository(spanner *spanner.Client) IProductRepo {
	return &spannerProductRepository{spanner}
}

func (sp spannerProductRepository) GetByProductCD(ctx context.Context, itemCD string) (res models.Product, err error) {
	sql := `select ITEMCD,ITEMNM from ITEM_M where ITEMCD = @itemCD`
	params := map[string]interface{}{"itemCD": itemCD}

	iter := sp.client.Single().Query(ctx, spanner.Statement{
		SQL:    sql,
		Params: params,
	})
	defer iter.Stop()

	row, err := iter.Next()

	if err != nil {
		return models.Product{}, err
	}

	res = models.Product{}
	row.ToStruct(&res)

	if err != nil {
		return models.Product{}, err
	}

	return
}
