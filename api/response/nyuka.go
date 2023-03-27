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
	ORDERSU   string `json:"order_su"`
	NYUKAYOSU string `json:"yotei_su"`
	NYUKAJISU string `json:"jiseki_su"`
}
