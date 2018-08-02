package error

import (
	"encoding/json"
)

// HttpError represent the error response return by the application in server mode.
type HttpError struct {
	//ErrorCode is unique code identifying the error
	ErrorCode string `json:"code"`
	//Message is the message that can be displayed to end users
	Message string `json:"message"`
	//HttpStatus is http status code
	HttpStatus int `json:"status"`
}

// ToJson convert an error object to JSON string.
func (errorResponse *HttpError) ToJson() (string, *Error) {
	errorJson, err := json.Marshal(errorResponse)
	if err != nil {
		return "", &Error{INTERNAL_SERVER_ERROR, nil, err}
	}
	return string(errorJson), nil
}