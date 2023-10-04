package models

import "github.com/google/uuid"

type NewOrderEventDto struct {
	Id uuid.UUID `json:"id"`
}
