package types

import "github.com/google/uuid"

type CreateRequest struct {
}

type CreateResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type GetResponse struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type ListRequest struct {
}

type ListResponse struct {
	Data []*ListResponseData `json:"data"`
}

type ListResponseData struct {
}

type UpdateRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type UpdateResponse struct {
}

type DeleteRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}
