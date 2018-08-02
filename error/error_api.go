package error

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strings"
	"fmt"
)

// Error is a wrapper representing an error that can be sent by the application.
type Error struct {
	// ErrorCode is a custom error code of the error thrown by the application
	ErrorCode string
	// Param is used to replace placeholders in an error template with the corresponding values
	Param Params
	// Error contains the original error
	Error error
}

// Constant defining all the different error code the API can throw.
// Each error code is mapped with a specific http code and http message
// defined in errors.yaml
const (
	NOT_FOUND             = "NOT_FOUND"
	METHOD_NOT_ALLOWED    = "METHOD_NOT_ALLOWED"
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	UNAUTHORIZED          = "UNAUTHORIZED"
	BAD_PARAMETER         = "BAD_PARAMETER"
	PACKAGE_NOT_FOUND     = "PACKAGE_NOT_FOUND"
	NO_SOURCE             = "NO_SOURCE"
	SOURCE_NOT_READABLE   = "SOURCE_NOT_READABLE"
)

type (
	// Params is used to replace placeholders in an error template with the corresponding values.
	Params map[string]interface{}

	errorTemplate struct {
		//ErrorMessage is the message to display to the end-users.
		ErrorMessage string `yaml:"message"`
		//HttpStatus is the http code to return.
		HttpStatus int `yaml:"status"`
	}
)

// Templates is a map between an api error code and errorTemplate.
var Templates map[string]errorTemplate

// LoadMessages reads a YAML file containing error templates.
func LoadMessages(file string) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	Templates = map[string]errorTemplate{}
	return yaml.Unmarshal(bytes, &Templates)
}

func replacePlaceholders(message string, params Params) string {
	if len(message) == 0 {
		return ""
	}
	for key, value := range params {
		message = strings.Replace(message, "{"+key+"}", fmt.Sprint(value), -1)
	}
	return message
}

// GetMessage return a verbose message corresponding to the api error.
func (error *Error) GetMessage() string {
	return replacePlaceholders(Templates[error.ErrorCode].ErrorMessage, error.Param)
}
