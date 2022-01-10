package rest

import (
	"fmt"
	"net/http"

	_ "github.com/arturzhamaliyev/Online-shop-API-Gateway/docs"
	"github.com/arturzhamaliyev/Online-shop-API-Gateway/internal/auth"
	"github.com/arturzhamaliyev/Online-shop-API-Gateway/internal/config"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Run(c *config.Config) http.Handler {
	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost%s/swagger/doc.json", c.Port)),
	))

	auth.RegisterRoutes(r, c)

	return r
}
