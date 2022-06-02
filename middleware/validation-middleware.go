package middleware

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
	"github.com/stkr89/modelsvc/types"
	"strings"
)

func ValidateDeleteInput() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*types.DeleteRequest)
			err := validator.New().Struct(req)
			err = validateUtil(err)
			if err != nil {
				return nil, err
			}

			return next(ctx, req)
		}
	}
}

func ValidateUpdateInput() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*types.UpdateRequest)
			err := validator.New().Struct(req)
			err = validateUtil(err)
			if err != nil {
				return nil, err
			}

			return next(ctx, req)
		}
	}
}

func ValidateGetInput() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*types.GetRequest)
			err := validator.New().Struct(req)
			err = validateUtil(err)
			if err != nil {
				return nil, err
			}

			return next(ctx, req)
		}
	}
}

func ValidateCreateInput() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*types.CreateRequest)
			err := validator.New().Struct(req)
			err = validateUtil(err)
			if err != nil {
				return nil, err
			}

			return next(ctx, req)
		}
	}
}

func validateUtil(err error) error {
	var validationErrors validator.ValidationErrors
	if err != nil {
		validationErrors = err.(validator.ValidationErrors)
	}

	if len(validationErrors) > 0 {
		var allErrs []string
		for _, e := range validationErrors {
			allErrs = append(allErrs, e.Error())
		}

		return errors.New(strings.Join(allErrs, ",\n"))
	}

	return nil
}
