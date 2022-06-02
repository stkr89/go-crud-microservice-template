package test

import (
	"github.com/joho/godotenv"
	"github.com/stkr89/modelsvc/cmd/server"
	"github.com/stkr89/modelsvc/common"
	"github.com/stkr89/modelsvc/endpoints"
	"github.com/stkr89/modelsvc/service"
	transport "github.com/stkr89/modelsvc/transports"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type HTTPIntegrationTestSuite struct {
	suite.Suite
	handler http.Handler
}

func (suite *HTTPIntegrationTestSuite) SetupSuite() {
	err := godotenv.Load("../.env")
	suite.NoError(err)

	e := endpoints.MakeEndpoints(service.NewModelServiceImpl())
	server.StartServer(common.NewLogger(), e, false, true)

	suite.handler = transport.NewHTTPHandler(e)
}

// new test cases

func TestHTTPIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(HTTPIntegrationTestSuite))
}
