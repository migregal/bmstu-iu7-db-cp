package structure

import (
	"neural_storage/cube/handlers/http/v1/entities/neuron"
	"neural_storage/cube/handlers/http/v1/entities/neuron/link"
	"neural_storage/cube/handlers/http/v1/entities/structure/layer"
	"neural_storage/cube/handlers/http/v1/entities/structure/weights"
)

type Info struct {
	Name    string         `json:"title"`
	Neurons []neuron.Info  `json:"neurons"`
	Layers  []layer.Info   `json:"layers"`
	Links   []link.Info    `json:"links"`
	Weights []weights.Info `json:"weights"`
}
