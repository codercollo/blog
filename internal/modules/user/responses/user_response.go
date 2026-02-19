package responses

import (
	"fmt"

	userModel "github.com/codercollo/blog/internal/modules/user/models"
)

type User struct {
	ID    uint
	Image string
	Name  string
	Email string
}

type Users struct {
	Data []User
}

func ToUser(user userModel.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: fmt.Sprintf("http://ui-avatars.com/api/?name=%s", user.Name),
	}
}
