# go-crud-microservice-template

## Steps

- Rename module in `go.mod`
- Rename model in `models/models.go` and add fields
- Populate fields for request and response objects in `types/types.go`
- Populate fields for request and response objects in `pb/model.proto`
- Generate gRPC client and server code by running the following:
```shell
$ cd pb
$ make generate
```
- Add fields to crud methods in `service/model_service.go`
- Add following environment variables:
  - `DB_USERNAME`
  - `DB_PASSWORD`
  - `DB_HOST`
  - `DB_PORT`
  - `DB_NAME`
  - `DB_CLUSTER`
  - `GRPC_PORT`
  - `HTTP_PORT`