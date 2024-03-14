package main

import (
	_ "github.com/nazudis/disbursement/src/config"

	_ "github.com/nazudis/disbursement/src/model/migration"

	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/nazudis/disbursement/engine/rest/routes"
	"github.com/nazudis/disbursement/src/helper"
	"github.com/nazudis/netper"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.NoCache)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res := netper.PlugResponse(w)
		respData := helper.ReplySuccess("server is running", nil)
		_ = res.ReplyCustom(respData.Status, respData)
	})

	r.Route("/v1", routes.AppRoutes)

	port := "3000"
	fmt.Printf("listen on port :%s \n", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%s", port), r)
}
