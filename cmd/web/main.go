package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/KjRodgers32/snippetbox/internal/models"

	env "github.com/KjRodgers32/snippetbox/internal"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	addr      string
	staticDir string
}

type application struct {
	logger   *slog.Logger
	config   config
	snippets *models.SnippetModel
}

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
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
	dsn, err := env.GetString("DSN")
	if err != nil {
		logger.Error("error getting dsn")
	}

	db, err := OpenDB(dsn)
	if err != nil {
		logger.Error(err.Error())
	}

	defer db.Close()

	config := config{
		addr:      addr,
		staticDir: staticDir,
	}

	app := &application{
		logger:   logger,
		config:   config,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("starting server on", "addr", app.config.addr)

	err = http.ListenAndServe(app.config.addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
