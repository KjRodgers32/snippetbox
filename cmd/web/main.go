package main

import (
	"log"
	"net/http"

	env "github.com/KjRodgers32/snippetbox/internal"
)

func main() {
	addr, err := env.GetString("ADDR")
	if err != nil {
		log.Fatal(err)
	}
	staticDir, err := env.GetString("STATIC_ADDRESS")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(staticDir))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)

	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	err = http.ListenAndServe(addr, mux)
	log.Fatal(err)
}
