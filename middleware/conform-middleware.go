package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/leebenson/conform"
	"github.com/stkr89/modelsvc/types"
)

func ConformDeleteInput() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*types.DeleteRequest)
			err := conform.Strings(req)
			if err != nil {
				return nil, err
			}
			return next(ctx, req)
		}
	}
}

func ConformUpdateInput() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*types.UpdateRequest)
			err := conform.Strings(req)
			if err != nil {
				return nil, err
			}
			return next(ctx, req)
		}
	}
}

func ConformGetInput() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*types.GetRequest)
			err := conform.Strings(req)
			if err != nil {
				return nil, err
			}
			return next(ctx, req)
		}
	}
}

func ConformCreateInput() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*types.CreateRequest)
			err := conform.Strings(req)
			if err != nil {
				return nil, err
			}
			return next(ctx, req)
		}
	}
}
