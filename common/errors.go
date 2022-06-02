package common

import "errors"

var (
	InvalidRequestBody = errors.New("Invalid request body")
	SomethingWentWrong = errors.New("Something went wrong")
)
