package response

type NyukaResponse struct {
	ITEMCD string `json:"product_cd"`
	ITEMNM string `json:"product_name"`
	CENTNO string `json:"center_no"`
}
