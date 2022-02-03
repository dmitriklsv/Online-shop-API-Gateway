package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/arturzhamaliyev/Online-shop-API-Gateway/internal/auth/pb"
	"github.com/arturzhamaliyev/Online-shop-API-Gateway/pkg/utils"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			utils.SendJson(w, http.StatusUnauthorized, utils.HTTPError{Message: "unauthorized"})
			return
		}

		token := strings.Split(auth, "Bearer ")
		if len(token) < 2 {
			utils.SendJson(w, http.StatusUnauthorized, utils.HTTPError{Message: "unauthorized"})
			return
		}

		res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
			Token: token[1],
		})
		if err != nil || res.Status != http.StatusOK {
			utils.SendJson(w, http.StatusUnauthorized, utils.HTTPError{Message: "unauthorized"})
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), utils.UserID, res.UserId)))
	})
}
