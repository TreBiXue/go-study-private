package response

type NyukaInputResponse struct {
	COUNT       int
	ORDERNOLIST []string `json:"orderNo_List"`
}

type NyukaInputJANResponse struct {
	ORDNO     string `json:"order_no"`
	ENO       string `json:"eda_no"`
	ITEMCD    string `json:"product_cd"`
	ITEMNM    string `json:"product_name"`
	SIRENO    string `json:"sire_no"`
	SIRENM    string `json:"sire_name"`
	ORDERSU   int64  `json:"order_su"`
	NYUKAYOSU int64  `json:"yotei_su"`
	NYUKAJISU int64  `json:"jiseki_su"`
}
