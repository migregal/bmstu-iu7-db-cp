package structure

import (
	"neural_storage/cube/handlers/http/v1/entities/neuron"
	"neural_storage/cube/handlers/http/v1/entities/neuron/link"
	"neural_storage/cube/handlers/http/v1/entities/structure/layer"
	"neural_storage/cube/handlers/http/v1/entities/structure/weights"
)

type Info struct {
	ID      string         `json:"id,omitempty"`
	Name    string         `json:"title,omitempty"`
	Neurons []neuron.Info  `json:"neurons,omitempty"`
	Layers  []layer.Info   `json:"layers,omitempty"`
	Links   []link.Info    `json:"links,omitempty"`
	Weights []weights.Info `json:"weights,omitempty"`
}
