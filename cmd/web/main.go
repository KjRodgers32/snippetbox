package main

import (
	"log/slog"
	"net/http"
	"os"

	env "github.com/KjRodgers32/snippetbox/internal"
)

type config struct {
	addr      string
	staticDir string
}

type application struct {
	logger *slog.Logger
	config config
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	addr, err := env.GetString("ADDR")
	if err != nil {
		logger.Error("error getting port address")
	}
	staticDir, err := env.GetString("STATIC_ADDRESS")
	if err != nil {
		logger.Error("error getting static address")
	}

	config := config{
		addr:      addr,
		staticDir: staticDir,
	}

	app := &application{
		logger: logger,
		config: config,
	}

	logger.Info("starting server on", "addr", app.config.addr)

	err = http.ListenAndServe(app.config.addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
