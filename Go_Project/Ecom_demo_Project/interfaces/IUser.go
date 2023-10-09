package interfaces

import "github.com/vimalachandrashekhar/Ecom-demo-Project/entities"

type IUser interface {
	Register(u *entities.User) (*entities.UserResponse, error)
	Login(u *entities.User) (string, error)
	LogOut() string
}
