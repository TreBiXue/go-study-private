package models

type NyukaInfo struct {
	CENTNO    string `json:"center_no"`
	ORDNO     string `json:"order_no"`
	ENO       string `json:"meisai_no"`
	JAN       string `json:"product_cd"`
	SIRENO    string `json:"sire_no"`
	ORDERSU   int64  `json:"order_su"`
	NKAYOSU   int64  `json:"nyuka_yotei"`
	NKAJISU   int64  `json:"nyuka_jiseiki"`
	NKAYDATE  string `json:"nyuka_date"`
	NKAUKEFLG string `json:"nyuka_flg"`
	UP_DATE   string `json:"nyuka_update"`
}