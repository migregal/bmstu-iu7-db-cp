package model

import (
	"neural_storage/cube/core/entities/structure"
)

type Info struct {
	ownerID   string
	id        string
	name      string
	structure *structure.Info
}

func NewInfo(ownerID, name string, structure *structure.Info) *Info {
	return &Info{ownerID: ownerID, name: name, structure: structure}
}
