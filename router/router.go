package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/to4to/go-rss-feed/handler"
)


func Router(){
	router := chi.NewRouter()

	router.Use(
		cors.Handler(cors.Options{

			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
	)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handler.HandlerReadiness)
	v1Router.Get("/err", handler.HandlerErr)
	v1Router.Post("/users",apiCfg.HandlerCreateUser)

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
}