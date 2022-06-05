package validator

import (
	"fmt"
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/entities/structure"
)

func (v *Validator) ValidateModelInfo(info *model.Info) error {
	if info == nil {
		return nil
	}

	return v.validateModelStructure(info.Structure())
}

func (v *Validator) validateModelStructure(info *structure.Info) error {
	if info == nil {
		return nil
	}

	if err := v.validateLayers(info); err != nil {
		return err
	}

	neuronIds, err := v.validateNeurons(info)
	if err != nil {
		return err
	}

	if err = v.validateNeuronOffsets(info, neuronIds); err != nil {
		return err
	}

	if err = v.validateNeuronLinks(info, neuronIds); err != nil {
		return err
	}

	return nil
}

func (v *Validator) validateLayers(info *structure.Info) error {
	if len(info.Layers()) > 0 && len(info.Neurons()) == 0 {
		return fmt.Errorf("validate layers: missing neuron info")
	}

	layerIds := make(map[int]struct{})
	for _, v := range info.Layers() {
		if _, found := layerIds[v.ID()]; found {
			return fmt.Errorf("duplicate layer info")
		}

		layerIds[v.ID()] = struct{}{}
	}

	fmt.Printf("FUCK: %#v | %#v\n", layerIds, info.Layers())
	for _, v := range info.Neurons() {
		if _, found := layerIds[v.LayerID()]; !found {
			return fmt.Errorf("missing layer info")
		}
	}

	return nil
}

func (v *Validator) validateNeurons(info *structure.Info) (map[int]struct{}, error) {
	neuronIds := make(map[int]struct{})
	for _, v := range info.Neurons() {
		if _, found := neuronIds[v.ID()]; found {
			return nil, fmt.Errorf("duplicate neuron info")
		}

		neuronIds[v.ID()] = struct{}{}
	}

	return neuronIds, nil
}

func (v *Validator) validateNeuronOffsets(info *structure.Info, neurons map[int]struct{}) (err error) {
	if neurons == nil {
		if neurons, err = v.validateNeurons(info); err != nil {
			return
		}
	}

	for _, w := range info.Weights() {
		for _, v := range w.Offsets() {
			if _, found := neurons[v.NeuronID()]; !found {
				return fmt.Errorf("validate offsets: missing neuron info")
			}
			if err = validateFloatValue(v.Offset()); err != nil {
				return fmt.Errorf("neuron offset: %w", err)
			}
		}
	}

	return
}

func (v *Validator) validateNeuronLinks(info *structure.Info, neurons map[int]struct{}) (err error) {
	if neurons == nil {
		if neurons, err = v.validateNeurons(info); err != nil {
			return
		}
	}

	linksIds := make(map[int]struct{})
	for _, v := range info.Links() {
		if _, found := linksIds[v.ID()]; found {
			return fmt.Errorf("duplicate link info")
		}

		linksIds[v.ID()] = struct{}{}

		if _, found := neurons[v.From()]; !found {
			return fmt.Errorf("validate links: missing neuron info")
		}
		if _, found := neurons[v.To()]; !found {
			return fmt.Errorf("validate to links: missing neuron info")
		}
	}

	for _, w := range info.Weights() {
		for _, v := range w.Weights() {
			if _, found := linksIds[v.LinkID()]; !found {
				return fmt.Errorf("validate weights: missing link info")
			}
			if err = validateFloatValue(v.Weight()); err != nil {
				return fmt.Errorf("neuron link weight: %w", err)
			}
		}
	}

	return nil
}
