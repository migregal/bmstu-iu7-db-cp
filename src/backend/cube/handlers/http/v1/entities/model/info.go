package model

import (
	"neural_storage/cube/handlers/http/v1/entities/structure"
)

type Info struct {
	ID        string         `json:"id"`
	Name      string         `json:"title"`
	OwnerID   string         `json:"owner_id"`
	Structure *structure.Info `json:"structure,omitempty"`
}
