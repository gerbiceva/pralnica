package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

// http-swagger middleware

const PORT = 1234

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("user"))
	})

	fmt.Printf("Serving on: %d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), r)
}
