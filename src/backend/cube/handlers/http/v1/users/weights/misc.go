package weights

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"neural_storage/cube/core/entities/neuron/link/weight"
	"neural_storage/cube/core/entities/neuron/offset"
	"neural_storage/cube/core/entities/structure/weights"
	httpweight "neural_storage/cube/handlers/http/v1/entities/neuron/link/weight"
	httpoffset "neural_storage/cube/handlers/http/v1/entities/neuron/offset"
	httpweights "neural_storage/cube/handlers/http/v1/entities/structure/weights"
)

func weightToBL(info httpweights.Info) weights.Info {
	linkWeights := []*weight.Info{}
	for _, lw := range info.Weights {
		linkWeights = append(linkWeights, weight.NewInfo(lw.ID, lw.LinkID, lw.Weight))
	}

	offsets := []*offset.Info{}
	for _, o := range info.Offsets {
		offsets = append(offsets, offset.NewInfo(o.ID, o.NeuronID, o.Offset))
	}

	return *weights.NewInfo(info.ID, info.Name, linkWeights, offsets)
}

func weightFromBL(info weights.Info) httpweights.Info {
	linkWeights := []httpweight.Info{}
	for _, lw := range info.Weights() {
		linkWeights = append(linkWeights, httpweight.Info{ID: lw.ID(), LinkID: lw.LinkID(), Weight: lw.Weight()})
	}

	offsets := []httpoffset.Info{}
	for _, o := range info.Offsets() {
		offsets = append(offsets, httpoffset.Info{ID: o.ID(), NeuronID: o.NeuronID(), Offset: o.Offset()})
	}

	return httpweights.Info{ID: info.ID(), Name: info.Name(), Weights: linkWeights, Offsets: offsets}
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
