package model

import "neural_storage/cube/core/entities/structure"

func (i *Info) ID() string {
	return i.id
}

func (i *Info) OwnerID() string {
	return i.ownerID
}

func (i *Info) Name() string {
	return i.name
}

func (i *Info) SetName(name string) {
	i.name = name
}

func (i *Info) Structure() *structure.Info {
	return i.structure
}

func (i *Info) SetStructure(structure *structure.Info) {
	i.structure = structure
}
