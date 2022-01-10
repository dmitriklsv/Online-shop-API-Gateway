package auth

import (
	"net/http"

	"github.com/arturzhamaliyev/Online-shop-API-Gateway/internal/auth/routes"
	"github.com/arturzhamaliyev/Online-shop-API-Gateway/internal/config"
	"github.com/arturzhamaliyev/Online-shop-API-Gateway/pkg/utils"
	"github.com/go-chi/chi/v5"
)

type PostBody struct {
	Id int64
}

func RegisterRoutes(r chi.Router, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", svc.Register)
	})

	return svc
}

// @Summary		Create user
// @Description	register user
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			input			body		routes.RegisterRequestBody	true	"registering new user with sent credentials"
// @Success		201				{object}		PostBody
// @Failure		400				{object}	utils.HTTPError
// @Failure		404				{object}	utils.HTTPError
// @Failure		409				{object}	utils.HTTPError
// @Failure		500				{object}	utils.HTTPError
// @Router			/auth/register	[post]
func (svc *ServiceClient) Register(w http.ResponseWriter, r *http.Request) {
	// TODO: check for content-type JSON

	id, status, err := routes.Register(r, svc.Client)
	if err != nil {
		utils.SendJson(w, status, utils.HTTPError{Message: err.Error()})
		return
	}

	utils.SendJson(w, status, &PostBody{id})
}
