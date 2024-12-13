package lib

import (
	"github.com/go-fuego/fuego"
	controller "github.com/go-fuego/fuego/examples/userstore/controllers"
	"github.com/go-fuego/fuego/examples/userstore/services"
)

func NewMudServer(options ...func(*fuego.Server)) *fuego.Server {
	s := fuego.NewServer(options...)

	usersResources := controller.UsersResources{
		UsersService: services.NewInMemoryUsersService(), // Dependency injection: we can pass a service here (for example a database service)
	}
	usersResources.Routes(s)

	return s
}
