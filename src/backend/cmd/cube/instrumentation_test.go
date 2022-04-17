//go:build instrumentation

package main

import (
	"testing"

	"neural_storage/pkg/instrumentation"
)

func init() {
	instrumentation.PreInit()
}

func TestInstumentation(t *testing.T) {
	main()
}
