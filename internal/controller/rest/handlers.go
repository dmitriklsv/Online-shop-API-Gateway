package rest

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/arturzhamaliyev/Online-shop-API-Gateway/docs"
	"github.com/arturzhamaliyev/Online-shop-API-Gateway/internal/auth"
	"github.com/arturzhamaliyev/Online-shop-API-Gateway/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Run(c *config.Config) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost%s/swagger/doc.json", c.Port)),
	))

	auth.RegisterRoutes(r, c)

	return r
}
