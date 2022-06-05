//go:build unit
// +build unit

package validator

import (
	"neural_storage/config/adapters/validator/mock"
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/entities/neuron"
	"neural_storage/cube/core/entities/neuron/link"
	"neural_storage/cube/core/entities/neuron/link/weight"
	"neural_storage/cube/core/entities/neuron/offset"
	"neural_storage/cube/core/entities/structure"
	"neural_storage/cube/core/entities/structure/layer"
	"neural_storage/cube/core/entities/structure/weights"
	"neural_storage/cube/core/ports/config"
	"reflect"
	"testing"
)

func TestValidator_ValidateModelInfo(t *testing.T) {
	type fields struct {
		conf config.ValidatorConfig
	}
	type args struct {
		info *model.Info
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"no info",
			fields{&mock.ValidatorConfig{}},
			args{model.NewInfo("", "", "", nil)},
			false,
		},
		{
			"empty struct",
			fields{&mock.ValidatorConfig{}},
			args{
				model.NewInfo(
					"",
					"",
					"",
					structure.NewInfo(
						"",
						"",
						nil,
						nil,
						nil,
						nil,
					),
				),
			},
			false,
		},
		{
			"ordinal stuct",
			fields{&mock.ValidatorConfig{}},
			args{
				model.NewInfo(
					"id",
					"",
					"",
					structure.NewInfo(
						"id",
						"",
						[]*neuron.Info{
							neuron.NewInfo(0, 0),
						},
						[]*layer.Info{
							layer.NewInfo(0, "func 1", "func 2"),
						},
						[]*link.Info{
							link.NewInfo(0, 0, 0),
						},
						[]*weights.Info{
							weights.NewInfo(
								"",
								"",
								[]*weight.Info{
									weight.NewInfo(0, 0, 10.0),
								},
								[]*offset.Info{
									offset.NewInfo(0, 0, -2.0),
								},
							),
						},
					),
				),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validator{
				conf: tt.fields.conf,
			}
			if err := v.ValidateModelInfo(tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("ValidateModelInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidator_validateLayers(t *testing.T) {
	type fields struct {
		conf config.ValidatorConfig
	}
	type args struct {
		info *structure.Info
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validator{
				conf: tt.fields.conf,
			}
			if err := v.validateLayers(tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("validateLayers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidator_validateModelStructure(t *testing.T) {
	type fields struct {
		conf config.ValidatorConfig
	}
	type args struct {
		info *structure.Info
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validator{
				conf: tt.fields.conf,
			}
			if err := v.validateModelStructure(tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("validateModelStructure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidator_validatelinks(t *testing.T) {
	type fields struct {
		conf config.ValidatorConfig
	}
	type args struct {
		info    *structure.Info
		neurons map[int]struct{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validator{
				conf: tt.fields.conf,
			}
			if err := v.validateNeuronLinks(tt.args.info, tt.args.neurons); (err != nil) != tt.wantErr {
				t.Errorf("validatelinks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidator_validateNeuronOffsets(t *testing.T) {
	type fields struct {
		conf config.ValidatorConfig
	}
	type args struct {
		info    *structure.Info
		neurons map[int]struct{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validator{
				conf: tt.fields.conf,
			}
			if err := v.validateNeuronOffsets(tt.args.info, tt.args.neurons); (err != nil) != tt.wantErr {
				t.Errorf("validateNeuronOffsets() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidator_validateNeurons(t *testing.T) {
	type fields struct {
		conf config.ValidatorConfig
	}
	type args struct {
		info *structure.Info
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]struct{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validator{
				conf: tt.fields.conf,
			}
			got, err := v.validateNeurons(tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateNeurons() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateNeurons() got = %v, want %v", got, tt.want)
			}
		})
	}
}
