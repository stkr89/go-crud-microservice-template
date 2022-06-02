package endpoints

import (
	"context"
	"github.com/stkr89/modelsvc/types"

	"github.com/go-kit/kit/endpoint"
	"github.com/stkr89/modelsvc/service"
)

type Endpoints struct {
	Add endpoint.Endpoint
}

func MakeEndpoints(s service.ModelService) Endpoints {
	return Endpoints{
		Add: makeAddEndpoint(s),
	}
}

func makeAddEndpoint(s service.ModelService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*types.MathRequest)
		return s.Add(req)
	}
}
