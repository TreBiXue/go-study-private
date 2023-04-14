package request

import "time"

type NyukaInputRequest struct {
	CENTNO     string `json:"center_no" example:"0093"`
	NYUKABEGIN string `json:"nyuka_begin" example:"20220101"`
	NYUKAEND   string `json:"nyuka_end" example:"20240101"`
	SIRENO     string `json:"sire_no" example:"4466"`
}

type NyukaInputJANRequest struct {
	CENTNO  string `json:"center_no"`
	ORDERNO string `json:"order_no"`
	ITEMCD  string `json:"product_cd"`
}

type NyukaInputJANNKAJISURequest struct {
	NyukaInputJANRequest
	NYUKAJISU int64     `json:"jiseki_su"`
	UP_DATE   time.Time `json:"up_date"`
}

func (r *NyukaInputJANNKAJISURequest) Validate() error {
	if r.UP_DATE.IsZero() {
		r.UP_DATE = time.Now()
	}
	return nil
}
