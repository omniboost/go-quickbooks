package quickbooks

import "fmt"

// ErrorObject quickbooks error object
type ErrorObject struct {
	Fault fault `json:"Fault"`
	// @TODO: change this to a Time/Timestamp type
	// can be 13234232 or "2018-11-29T02:53:10.369-08:00"
	Time string `json:"time"`
}

type fault struct {
	Error []faultError `json:"Error"`
	Type  string       `json:"type"`
}

type faultError struct {
	Message string `json:"Message"`
	Detail  string `json:"Detail"`
	Code    string `json:"code"`
	Element string `json:"element"`
}

// Error() ErrorObject error interface method
func (e ErrorObject) Error() string {
	// errStr := fmt.Sprintf("Type: %s Code: %s Message: %s", e.Fault.Type, e.Fault.Error[0].Code, e.Fault.Error[0].Message)
	errStr := fmt.Sprintf("Type: %s Code: %s Detail: %s", e.Fault.Type, e.Fault.Error[0].Code, e.Fault.Error[0].Detail)
	return errStr
}

// SDKError customer error object
type SDKError struct {
	Type    string
	Code    string
	Message string
}

// Error() SDKError error interface method
func (e SDKError) Error() string {
	errStr := fmt.Sprintf("Type: %s Code: %s Message: %s", e.Type, e.Code, e.Message)
	return errStr
}

// New creates a new SDKError object
func (e SDKError) New(errorType string, errorCode string, errorMessage string) SDKError {
	return SDKError{errorType, errorCode, errorMessage}
}
