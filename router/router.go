package router

import (
	

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	
)

func Router() chi.Router {

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"LINK"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	

return router


}
