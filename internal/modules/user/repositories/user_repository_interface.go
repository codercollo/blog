package repositories

import (
	userModel "github.com/codercollo/blog/internal/modules/user/models"
)

type UserRepositoryInterface interface {
	Create(user userModel.User) userModel.User
}
