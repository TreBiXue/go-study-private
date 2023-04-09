Feature: getjaninfo

Background:
  * url "https://ec-warehouse-et76x6ix4q-an.a.run.app"
  * def janinfo = read('getjaninfo.csv')
  * print janinfo

# Scenario: test for getcounts
#   Given path '/api/v1/nyuka/getjaninfo'
#   And header Accept = 'application/json'
#   And request { center_no: janinfo[0]['in_center_no'], order_no: janinfo[0]['in_order_no'], product_cd: janinfo[0]['in_product_cd']}
#   When method post
#   Then status 200
#   And match response.商品情報明細.order_no == janinfo[0]['out_order_no']
#   And match response.商品情報明細.eda_no == janinfo[0]['out_eda_no']
#   And match response.商品情報明細.product_cd == janinfo[0]['out_product_cd']
#   And match response.商品情報明細.product_name == janinfo[0]['out_product_name']
#   And match response.商品情報明細.sire_no == janinfo[0]['out_sire_no']
#   And match response.商品情報明細.sire_name == janinfo[0]['out_sire_name']
#   And match response.商品情報明細.order_su == janinfo[0]['out_order_su']
#   And match response.商品情報明細.yotei_su == janinfo[0]['out_yotei_su']
#   And match response.商品情報明細.jiseki_su == janinfo[0]['out_jiseki_su']


#   Examples:
#     |janinfo|