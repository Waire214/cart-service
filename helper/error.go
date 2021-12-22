package helper

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	//ResourceCreationFailed = "resource creation faild"
	NoResourceFound = "this resource does not exist"
	NoRecordFound   = "sorry. no record found"
	NoErrorsFound   = "no errors at the moment"
)

func (err ErrorResponse) Error() string {
	var errorBody ErrorBody
	return fmt.Sprintf("%v", errorBody)
}

func ErrorArrayToError(fieldErrors []error) error {
	var errorResponse ErrorResponse
	errorResponse.TimeStamp = time.Now().Format(time.RFC3339)
	errorResponse.ErrorReference = uuid.New()

	for _, value := range fieldErrors {
		body := ErrorBody{Code: "400 {validation error}", Source: "cart-service", Message: value.Error()}
		errorResponse.Errors = append(errorResponse.Errors, body)
	}
	return errorResponse
}

func PrintErrorMessage(code string, message string) error {
	var errorResponse ErrorResponse
	errorResponse.TimeStamp = time.Now().Format(time.RFC3339)
	errorResponse.ErrorReference = uuid.New()
	errorResponse.Errors = append(errorResponse.Errors, ErrorBody{Code: code, Source: "cart-service", Message: message})

	return errorResponse
}
