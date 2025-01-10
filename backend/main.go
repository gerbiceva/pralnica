package main

import (
	"fmt"
	"net/http"
	dbclient "rso/m/dbClient"
	"rso/m/docs"
	"rso/m/paths"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	httpSwagger "github.com/swaggo/http-swagger"

	chiprometheus "github.com/766b/chi-prometheus"
)

const PORT = 1234

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	m := chiprometheus.NewMiddleware("serviceName")
	r.Use(m)
	r.Handle("/metrics", promhttp.Handler())

	r.Mount("/swagger", httpSwagger.WrapHandler)

	r.Get("/reservations/user/{uuid}", paths.GetUserReservations)
	r.Get("/reservations/month/{month}", paths.GetReservationsByMonth)
	r.Post("/reservations", paths.AddReservation)
	r.Delete("/reservations/{id}", paths.DeleteReservation)

	r.Get("/users/search", paths.SearchUsers)
	r.Get("/users", paths.GetUsers)
	r.Get("/users/self", paths.GetSelf)
	r.Get("/users/{uuid}", paths.GetUser)
	r.Post("/users/{uuid}", paths.EditUser)
	r.Post("/users/confirm/{uuid}", paths.ConfirmUser)
	r.Post("/users/ban/{uuid}", paths.BanUser)
	r.Post("/users/{uuid}", paths.EditUser)

	docgen.PrintRoutes(r)
	docs.WriteDocsMarkdown(r)
	docs.WriteDocsJSON(r)

	fmt.Printf("Serving on: %d\n", PORT)
	dbclient.InitDbClient()

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), r)
	for {
		time.Sleep(100 * time.Second)
	}
}
