package main

import (
	"log"
	"net/http"

	"github.com/arturzhamaliyev/Online-shop-API-Gateway/pkg/auth"
	"github.com/arturzhamaliyev/Online-shop-API-Gateway/pkg/config"
	"github.com/go-chi/chi/v5"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := chi.NewRouter()

	authSvc := *auth.RegisterRoutes(r, c)

	err = http.ListenAndServe(c.Port, r)
	if err != nil {
		log.Fatalln(err)
	}
}
