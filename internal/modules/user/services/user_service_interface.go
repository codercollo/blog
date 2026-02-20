package services

import (
	"github.com/codercollo/blog/internal/modules/user/requests/auth"
	UserResponse "github.com/codercollo/blog/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponse.User, error)
	CheckUserExists(email string) bool
	HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
}
