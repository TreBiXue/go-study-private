Feature: getcounts

Background:
  * url "https://ec-warehouse-et76x6ix4q-an.a.run.app"

Scenario: test for getcounts
  Given path '/api/v1/nyuka/getcounts'
  And header Accept = 'application/json'
  And request { center_no: '0093', nyuka_begin: '20230205', nyuka_end: '20230210', sire_no: '4466' }
  When method post
  Then status 200
  And match response.入荷件数.COUNT == 3
  And match response.入荷件数.orderNo_List contains '092628664'
  And match response.入荷件数.orderNo_List contains '092629268'
  # * assert responseStatus == 200