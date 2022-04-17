package main

import (
	"log"
	"neural_storage/cube"
	"neural_storage/pkg/instrumentation"
)

func main() {
	app, err := cube.New()
	if err != nil {
		log.Fatal(err)
	}

	errs := make(chan error, 1)
	go func() { errs <- app.Run() }()

	select {
	case err := <-errs:
		log.Fatal(err)
	case <-instrumentation.Done():
	}
}
