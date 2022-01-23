package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/arturzhamaliyev/Online-shop-API-Gateway/internal/auth/pb"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(r *http.Request, c pb.AuthServiceClient) (int64, string, error) {
	var body LoginRequestBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return http.StatusBadRequest, "", err
	}
	defer r.Body.Close()

	loginResponse, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		return http.StatusBadGateway, "", err
	}

	return loginResponse.Status, loginResponse.Token, nil
}
