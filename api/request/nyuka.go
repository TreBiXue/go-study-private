package request

import "time"

type NyukaInputRequest struct {
	CENTNO     string `json:"center_no"`
	NYUKABEGIN string `json:"nyuka_begin"`
	NYUKAEND   string `json:"nyuka_end"`
	SIRENO     string `json:"sire_no"`
}

type NyukaInputJANRequest struct {
	CENTNO  string `json:"center_no"`
	ORDERNO string `json:"order_no"`
	ITEMCD  string `json:"product_cd"`
}

type NyukaInputJANNKAJISURequest struct {
	NyukaInputJANRequest
	NYUKAJISU string    `json:"jiseki_su"`
	UP_DATE   time.Time `json:"up_date"`
}

func (r *NyukaInputJANNKAJISURequest) Validate() error {
	if r.UP_DATE.IsZero() {
		r.UP_DATE = time.Now()
	}
	return nil
}
