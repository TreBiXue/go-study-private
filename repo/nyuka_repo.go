package repo

import (
	"cloud.google.com/go/spanner"
	"context"
	"errors"
	"go-studying/api/request"
	"go-studying/models"
	"google.golang.org/api/iterator"
)

type spannerNyukaRepository struct {
	client *spanner.Client
}

func NewSpannerNyukaRepository(spanner *spanner.Client) INyukaRepo {
	return &spannerNyukaRepository{spanner}
}

func (sp spannerNyukaRepository) GetNyukaCount(ctx context.Context, nyukaInput *request.NyukaInputRequest) (res []models.NyukaInfo, err error) {

	flag := "0"
	sql := `select CENTNO,ORDNO,ENO,JAN,SIRENO,ORDERSU,NKAYOSU,NKAJISU,NKAYDATE,NKAUKEFLG,UP_DATE 
				from NYUKAUKE 
				where CENTNO = @centerNo and SIRENO = @sireNo and NKAYDATE >= @nyukaBegin and NKAYDATE <= @nyukaEnd  and NKAUKEFLG = @nyukaFlg `
	params := map[string]interface{}{
		"centerNo":   nyukaInput.CENTNO,
		"sireNo":     nyukaInput.SIRENO,
		"nyukaBegin": nyukaInput.NYUKABEGIN,
		"nyukaEnd":   nyukaInput.NYUKAEND,
		"nyukaFlg":   flag,
	}

	iter := sp.client.Single().Query(ctx, spanner.Statement{
		SQL:    sql,
		Params: params,
	})
	defer iter.Stop()

	res = []models.NyukaInfo{}

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var nyuka models.NyukaInfo
		if err := row.ToStruct(&nyuka); err != nil {
			return nil, err
		}
		res = append(res, nyuka)
	}
	return
}

func (sp spannerNyukaRepository) GetNyukaJANInfo(ctx context.Context, nyukaJAN *request.NyukaInputJANRequest) (res models.NyukaInfo, err error) {
	flag := "0"
	sql := `select CENTNO,ORDNO,ENO,JAN,SIRENO,ORDERSU,NKAYOSU,NKAJISU,NKAYDATE,NKAUKEFLG,UP_DATE 
				from NYUKAUKE 
				where CENTNO = @centerNo and ORDNO = @orderNo and JAN = @itemCd and NKAUKEFLG = @nyukaFlg `
	params := map[string]interface{}{
		"centerNo": nyukaJAN.CENTNO,
		"orderNo":  nyukaJAN.ORDERNO,
		"itemCd":   nyukaJAN.ITEMCD,
		"nyukaFlg": flag,
	}

	iter := sp.client.Single().Query(ctx, spanner.Statement{
		SQL:    sql,
		Params: params,
	})
	defer iter.Stop()

	row, err := iter.Next()

	if err != nil {
		return models.NyukaInfo{}, err
	}

	res = models.NyukaInfo{}
	row.ToStruct(&res)

	if err != nil {
		return models.NyukaInfo{}, err
	}

	return
}

func (sp spannerNyukaRepository) UpdateNyukaInfo(ctx context.Context, nyukaJANJISU *request.NyukaInputJANNKAJISURequest) (err error) {
	upDate := nyukaJANJISU.UP_DATE.Format("20060102")
	flag := "0"

	stmt := spanner.Statement{
		SQL: `UPDATE NYUKAUKE
				SET UP_DATE = @upDate,NKAJISU= @nkajisu,NKAUKEFLG= '1'
				where CENTNO = @centerNo and ORDNO = @orderNo and JAN = @itemCd and NKAUKEFLG = @nyukaFlg `,
		Params: map[string]interface{}{
			"upDate":   upDate,
			"nkajisu":  nyukaJANJISU.NYUKAJISU,
			"centerNo": nyukaJANJISU.CENTNO,
			"orderNo":  nyukaJANJISU.ORDERNO,
			"itemCd":   nyukaJANJISU.ITEMCD,
			"nyukaFlg": flag,
		},
	}

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
