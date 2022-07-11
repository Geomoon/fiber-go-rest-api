package api

import (
	"accounts-api/pkg/user"
)

func toUser(request *SignupRequest) *user.User {
	return &user.User{
		Name:     request.Name,
		LastName: request.LastName,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
}

func toResponse(user *user.User) *UserResponse {
	return &UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		LastName: user.LastName,
		Username: user.Username,
		Email:    user.Email,
	}
}
