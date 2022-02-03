package rest

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/arturzhamaliyev/Online-shop-API-Gateway/docs"
	"github.com/arturzhamaliyev/Online-shop-API-Gateway/internal/auth"
	"github.com/arturzhamaliyev/Online-shop-API-Gateway/internal/config"
	"github.com/arturzhamaliyev/Online-shop-API-Gateway/pkg/utils"
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

	authSvc := auth.RegisterRoutes(r, c)

	r.Route("/order", func(order chi.Router) {
		authMiddleware := auth.InitAuthMiddleware(authSvc)

		order.Use(authMiddleware.AuthRequired)

		order.Get("/", func(w http.ResponseWriter, r *http.Request) {
			utils.SendJson(w, 200, "hi")
		})
	})

	return r
}
