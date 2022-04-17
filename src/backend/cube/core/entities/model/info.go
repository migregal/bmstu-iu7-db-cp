package model

import "neural_storage/cube/core/entities/structure"

type Info struct {
	id        string
	name      string
	structure *structure.Info
}

func NewInfo(name string, structure *structure.Info) *Info {
	return &Info{name: name, structure: structure}
}
