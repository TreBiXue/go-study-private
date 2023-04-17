Feature: login

Background:
  * url "https://ec-warehouse-et76x6ix4q-an.a.run.app"

Scenario: test for login(use id)
  Given path '/api/v1/login'
  And header Accept = 'application/json'
  And param id = 2
  When method get
  Then status 200
  And match response.Accounts.employee_code == '1234567'
  And match response.Accounts.employee_name == 'aaaa'
  # * assert responseStatus == 200
