package repo

import (
	"context"
	"go-studying/models"

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

	if err != nil {
		return models.Vender{}, err
	}

	res = models.Vender{}
	row.ToStruct(&res)

	if err != nil {
		return models.Vender{}, err
	}

	return
}
