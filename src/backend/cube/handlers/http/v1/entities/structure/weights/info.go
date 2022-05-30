package weights

import (
	"neural_storage/cube/handlers/http/v1/entities/neuron/link/weight"
	"neural_storage/cube/handlers/http/v1/entities/neuron/offset"
)

type Info struct {
	ID      string        `json:"id"`
	Name    string        `json:"name"`
	Weights []weight.Info `json:"weights"`
	Offsets []offset.Info `json:"offsets"`
}
