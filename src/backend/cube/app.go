package cube

import (
	"os"

	"neural_storage/config/core/services/config"
	"neural_storage/cube/handlers"
	"neural_storage/pkg/logger"
)

type App struct {
	srv handlers.Server
}

func New() (*App, error) {
	filename := os.Getenv("CONFIG_PATH")
	if filename == "" {
		filename = "/tmp/config.yml"
	}
	config, err := config.New(filename)
	if err != nil {
		return nil, err
	}

	var lg = logger.New()

	return &App{srv: handlers.New(config, lg)}, nil
}

func (a *App) Run() error {
	return a.srv.Run()
}
