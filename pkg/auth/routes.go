package auth

import (
	"net/http"

	"github.com/arturzhamaliyev/Online-shop-API-Gateway/pkg/auth/routes"
	"github.com/arturzhamaliyev/Online-shop-API-Gateway/pkg/config"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", svc.Register)
	})

	return svc
}

func (svc *ServiceClient) Register(w http.ResponseWriter, r *http.Request) {
	// TODO: check for content-type JSON

	routes.Register(r, svc.Client)
	defer r.Body.Close()
}
