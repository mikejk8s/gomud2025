package lib

import (
	"github.com/go-fuego/fuego"

	controller "gomud2025/controllers"
	services "gomud2025/services"
)

func NewMudServer(options ...func(*fuego.Server)) *fuego.Server {
	s := fuego.NewServer(options...)

	usersResources := controller.UsersResources{
		UsersService: services.NewInMemoryUsersService(), // Dependency injection: we can pass a service here (for example a database service)
	}
	usersResources.Routes(s)

	return s
}
