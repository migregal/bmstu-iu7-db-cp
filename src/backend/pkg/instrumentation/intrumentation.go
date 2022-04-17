package instrumentation

import (
	"os"
	"os/signal"
	"syscall"
)

var enabled = false

func PreInit() {
	enabled = true
}

func Done() <-chan struct{} {
	if !enabled {
		return nil
	}

	done := make(chan struct{}, 1)

	sigHandler := make(chan os.Signal, 1)
	signal.Notify(sigHandler, syscall.SIGUSR2)

	go func() {
		<-sigHandler
		done <- struct{}{}
	}()

	return done
}
