package models

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/entities/neuron"
	"neural_storage/cube/core/entities/neuron/link"
	"neural_storage/cube/core/entities/neuron/link/weight"
	"neural_storage/cube/core/entities/neuron/offset"
	"neural_storage/cube/core/entities/structure"
	"neural_storage/cube/core/entities/structure/layer"
	"neural_storage/cube/core/entities/structure/weights"
	httpmodel "neural_storage/cube/handlers/http/v1/entities/model"
	httpneuron "neural_storage/cube/handlers/http/v1/entities/neuron"
	httplink "neural_storage/cube/handlers/http/v1/entities/neuron/link"
	httpweight "neural_storage/cube/handlers/http/v1/entities/neuron/link/weight"
	httpoffset "neural_storage/cube/handlers/http/v1/entities/neuron/offset"
	httpstructure "neural_storage/cube/handlers/http/v1/entities/structure"
	httplayer "neural_storage/cube/handlers/http/v1/entities/structure/layer"
	httpweights "neural_storage/cube/handlers/http/v1/entities/structure/weights"
)

func structToBL(info httpstructure.Info) *structure.Info {
	neurons := []*neuron.Info{}
	for _, v := range info.Neurons {
		neurons = append(neurons, neuron.NewInfo(v.ID, v.LayerID))
	}

	layers := []*layer.Info{}
	for _, v := range info.Layers {
		layers = append(layers, layer.NewInfo(v.ID, v.LimitFunc, v.ActivationFunc))
	}

	links := []*link.Info{}
	for _, v := range info.Links {
		links = append(links, link.NewInfo(v.ID, v.From, v.To))
	}

	weights := []*weights.Info{}
	for _, w := range info.Weights {
		weights = append(weights, weightToBL(w))
	}
	return structure.NewInfo("", info.Name, neurons, layers, links, weights)
}

func weightToBL(info httpweights.Info) *weights.Info {
	linkWeights := []*weight.Info{}
	for _, lw := range info.Weights {
		linkWeights = append(linkWeights, weight.NewInfo(lw.ID, lw.LinkID, lw.Weight))
	}

	offsets := []*offset.Info{}
	for _, o := range info.Offsets {
		offsets = append(offsets, offset.NewInfo(o.ID, o.NeuronID, o.Offset))
	}

	return weights.NewInfo(info.ID, info.Name, linkWeights, offsets)
}

func modelFromBL(info *model.Info) httpmodel.Info {
	return httpmodel.Info{
		ID:        info.ID(),
		OwnerID:   info.OwnerID(),
		Name:      info.Name(),
		Structure: structFromBL(info.Structure()),
	}
}

func structFromBL(info *structure.Info) *httpstructure.Info {
	if info == nil {
		return nil
	}

	layers := []httplayer.Info{}
	for _, v := range info.Layers() {
		layers = append(layers, httplayer.Info{ID: v.ID(), ActivationFunc: v.ActivationFunc(), LimitFunc: v.LimitFunc()})
	}

	neurons := []httpneuron.Info{}
	for _, v := range info.Neurons() {
		neurons = append(neurons, httpneuron.Info{ID: v.ID(), LayerID: v.LayerID()})
	}

	links := []httplink.Info{}
	for _, v := range info.Links() {
		links = append(links, httplink.Info{ID: v.ID(), From: v.From(), To: v.To()})
	}

	return &httpstructure.Info{
		ID:      info.ID(),
		Name:    info.Name(),
		Layers:  layers,
		Neurons: neurons,
		Links:   links,
		Weights: weightFromBL(info.Weights()),
	}
}

func weightFromBL(info []*weights.Info) []httpweights.Info {
	weights := []httpweights.Info{}
	for _, i := range info {
		linkWeights := []httpweight.Info{}
		for _, lw := range i.Weights() {
			linkWeights = append(linkWeights, httpweight.Info{ID: lw.ID(), LinkID: lw.LinkID(), Weight: lw.Weight()})
		}

		offsets := []httpoffset.Info{}
		for _, o := range i.Offsets() {
			offsets = append(offsets, httpoffset.Info{ID: o.ID(), NeuronID: o.NeuronID(), Offset: o.Offset()})
		}

		weights = append(weights, httpweights.Info{ID: i.ID(), Name: i.Name(), Weights: linkWeights, Offsets: offsets})
	}

	return weights
}

func jsonGzip(data interface{}) ([]byte, error) {
	resp, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to form response: %v", err)
	}

	buf := new(bytes.Buffer)
	gz := gzip.NewWriter(buf)
	if _, err := gz.Write(resp); err != nil {
		gz.Close()
		return nil, fmt.Errorf("failed to gzip response: %v", err)
	}

	gz.Close()
	return buf.Bytes(), nil
}

func unGzip(data []byte) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to read gzipped cache: %v", err)
	}
	defer gz.Close()

	s, err := ioutil.ReadAll(gz)
	if err != nil {
		return nil, fmt.Errorf("failed to decode gzipped cache: %v", err)
	}

	return s, nil
}
