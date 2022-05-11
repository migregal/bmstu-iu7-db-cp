//go:build unit
// +build unit

package validator

import (
	"math"
	"testing"
)

func Test_validateFloatvalue(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ordinal float", args{10.0}, false},
		{"Nan float", args{math.NaN()}, true},
		{"Inf float", args{math.Inf(0)}, true},
		{"-Inf float", args{math.Inf(-1)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateFloatValue(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("validateFloatValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
