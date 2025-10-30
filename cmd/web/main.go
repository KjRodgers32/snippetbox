package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	env "github.com/KjRodgers32/snippetbox/internal"
)

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

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(staticDir))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)

	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	logger.Info("starting server on", addr)

	err = http.ListenAndServe(addr, mux)
	log.Fatal(err)
}
