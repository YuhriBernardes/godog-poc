Feature: Get random api
  As a developer
  I would like to request a random API
  In oder to use in my apps

  Scenario Outline:
    Given I request a random API
    Then the response status code is 200
    And the response body is not empty
    Examples:
