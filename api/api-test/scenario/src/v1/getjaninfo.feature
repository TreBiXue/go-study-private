Feature: getjaninfo

Background:
* url 'https://ec-warehouse-et76x6ix4q-an.a.run.app'
* def csv = read('../../data/v1/getjaninfo.csv')


Scenario Outline: get janinfo by jan

* def input = { center_no: '<in_center_no>', order_no: '<in_order_no>', product_cd: '<in_product_cd>' }
* def expected = { order_no: '<out_order_no>', eda_no: '<out_eda_no>', product_cd: '<out_product_cd>', product_name: '<out_product_name>', sire_no: '<out_sire_no>', sire_name: '<out_sire_name>', order_su: '<out_order_su>', yotei_su: '<out_yotei_su>', jiseki_su: '<out_jiseki_su>' }
Given path '/api/v1/nyuka/getjaninfo'
And header Accept = 'application/json'
And request input
When method post
Then status 200
And match response.商品情報明細.order_no == expected.order_no
And match response.商品情報明細.eda_no == expected.eda_no
And match response.商品情報明細.product_cd == expected.product_cd
And match response.商品情報明細.product_name == expected.product_name
And match response.商品情報明細.sire_no == expected.sire_no
And match response.商品情報明細.sire_name == expected.sire_name
And match response.商品情報明細.order_su == Number(expected.order_su)
And match response.商品情報明細.yotei_su == Number(expected.yotei_su)
And match response.商品情報明細.jiseki_su == Number(expected.jiseki_su)

Examples:
|csv|