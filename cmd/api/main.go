package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func init() {

}

func main() {

	r := chi.NewMux()

	r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World"))
		})
		r.Post("/sinal", nil)
	})

	server := &http.Server{
		Addr:        ":8080",
		Handler:     r,
		ReadTimeout: 3 * time.Minute,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
