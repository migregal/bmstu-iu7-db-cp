package validator

import (
	"fmt"
	"math"
)

func validateFloatValue(value float64) error {
	if math.IsNaN(value) {
		return fmt.Errorf("NaN value")
	}
	if math.IsInf(value, 1) {
		return fmt.Errorf("+Inf value")
	}
	if math.IsInf(value, -1) {
		return fmt.Errorf("-Inf value")
	}

	return nil
}
