Feature: go-study API Test

Background:
  # todo local
  * url "http://localhost:8080/api/v1"
  # todo Dev
  # * url "https://ec-warehouse-et76x6ix4q-an.a.run.app/api/v1"

  * def csv = read('go-studying.csv')

Scenario: Login-正常登陆
  Given path "/login"
  And param id = 1
  When method get
  Then status 200
  #判空#
  ####And match response.Accounts != null
  ####And match response contains {"Accounts":"#notnull"}
  #匹配某个值#
  ####And match response.Accounts.employee_code == "1"
  #匹配ALL#
  ####And match response == {"Accounts":{"employee_code":"1","employee_name":"test"}}
  ####And match response contains  {"Accounts":{"employee_code":"1","employee_name":"test"}}
  #匹配+忽略某值#
  ####And match response contains deep  {"Accounts":{"employee_code":"1"}}
  #匹配+忽略某值#
  And match response contains deep  {"Accounts":{"employee_code":"1", "employee_name":"#notnull"}}

Scenario: Login-Error-用户不存在
  Given path "/login"
  And param id = 999
  When method get
  Then status 404
  And match response == {"code":"not_found","msg":"data not found"}

#### Scenario: Login-Error-时间更新失败

# 匹配table类型 + 循环的写法
Scenario Outline: Center-取得-正常取得
  Given path "/center"

  And param  center_no = <center_no>
  When method get
  Then status <code>
  And match response == <center_response>

  Examples:
    | center_no | center_response                                                          | code |
    | "0093"    | {"CenterInfo":{"center_no":"0093","center_name":"白鳥常温センター"}} | 200 |
    | "9999"    | {"code":"not_found","msg":"data not found"}                    | 404 |

# 匹配CSV类型的写法
Scenario Outline: Vender-取得-正常取得
    Given path "/vender"
    And param vender_no = <vender_no>
    When method get
    Then status <code>
    And match response == <vender_response>

    Examples:
      |csv|


