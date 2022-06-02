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
}

type UpdateRequest struct {
}

type UpdateResponse struct {
}

type DeleteRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}
