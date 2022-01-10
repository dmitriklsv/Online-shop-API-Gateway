package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/arturzhamaliyev/Online-shop-API-Gateway/internal/auth/pb"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(r *http.Request, c pb.AuthServiceClient) (int64, int64, error) {
	var body RegisterRequestBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return 0, http.StatusBadRequest, err
	}
	defer r.Body.Close()

	registerResponse, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		return 0, http.StatusBadGateway, err
	}

	return registerResponse.GetId(), registerResponse.GetStatus(), nil
}
