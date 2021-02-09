package godog_poc

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/go-resty/resty/v2"
)

type(
	Entry struct {
		API         string `json:"API"`
		Description string `json:"Description"`
		Auth        string `json:"Auth"`
		HTTPS       bool   `json:"HTTPS"`
		Cors        string `json:"Cors"`
		Link        string `json:"Link"`
		Category    string `json:"Category"`
	}

	randomApiResponse struct {
		Count   int `json:"count"`
		Entries []Entry `json:"entries"`
	})

type randomApiScenario struct {
	endpoint string
	response *resty.Response
}

func (s *randomApiScenario) prepareScenario(*godog.Scenario){
	s.endpoint = "https://api.publicapis.org/random"
}

func (s *randomApiScenario) iRequestARandomApi() error {
	var err error = nil
	s.response, err = resty.New().NewRequest().SetError(&err).Get(s.endpoint)

	return err
}

func (s *randomApiScenario) theResponseStatusCodeIs(code int) error{
	actualCode :=  s.response.StatusCode()
	if code != actualCode{
		return fmt.Errorf("Expected status code %d, but received %d", code, actualCode)
	}
	return nil
}

func (s *randomApiScenario) theResponseBodyIsNotEmpty() error{
	responseBody := s.response.Body()
	if responseBody == nil || len(responseBody) <=0 {
		return fmt.Errorf("Expected body to not be empty, but got %v", responseBody)
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext){
	scenario := new(randomApiScenario)

	ctx.BeforeScenario(scenario.prepareScenario)

	ctx.Step(`^I request a random API`, scenario.iRequestARandomApi)
	ctx.Step(`^the response status code is (\d+)$`, scenario.theResponseStatusCodeIs)
	ctx.Step(`^the response body is not empty`, scenario.theResponseBodyIsNotEmpty)

}
