package itest

import (
	"testing"
	"gopkg.in/h2non/baloo.v3"
)

var todofinderBalooServer = baloo.New("http://localhost:8089")

func TestWrongEndpoint_ShouldReturn404(t *testing.T) {
	todofinderBalooServer.
		Get("/test_wrong").
		Expect(t).
		Status(404).
		Type("application/json").
		JSON(map[string]interface{}{
		"code":    "NOT_FOUND",
		"message": "/test_wrong was not found.",
		"status":  404,
	}).
		Done()
}

func TestSearchEndpoint_BadParameters_ShouldReturn400(t *testing.T) {
	queryParam := map[string]string{
		"pattern": "TODO",
	}
	todofinderBalooServer.
		Get("/search").
		SetQueryParams(queryParam).
		Expect(t).
		Status(400).
		Type("application/json").
		JSON(map[string]interface{}{
		"code":    "BAD_PARAMETER",
		"message": "Parameter 'package' is not correct or is missing.",
		"status":  400,
	}).
		Done()
}

func TestSearchEndpoint_BadMethod_ShouldReturn405(t *testing.T) {
	todofinderBalooServer.
		Post("/search").
		Expect(t).
		Status(405).
		Type("application/json").
		JSON(map[string]interface{}{
		"code":    "METHOD_NOT_ALLOWED",
		"message": "Unexpected method : 'POST'.",
		"status":  405,
	}).
		Done()
}

func TestSearchEndpoint_HappyPath_ShouldReturn200(t *testing.T) {
	queryParam := map[string]string{
		"pattern": "TODO",
		"package": "github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder",
	}
	todofinderBalooServer.
		Get("/search").
		SetQueryParams(queryParam).
		Expect(t).
		Status(200).
		Type("application/json").
		Done()
}
