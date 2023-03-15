package request

type NyukaRequest struct {
	ITEMCD string `json:"product_cd"`
	CENTNO string `json:"center_no"`
}
