package controller

import (
	"log/slog"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/examples/userstore/models"
	"github.com/go-fuego/fuego/option"
	"github.com/go-fuego/fuego/param"
)

// default pagination options
var optionPagination = option.Group(
	option.QueryInt("per_page", "Number of items per page", param.Required()),
	option.QueryInt("page", "Page number", param.Default(1), param.Example("1st page", 1), param.Example("42nd page", 42), param.Example("100th page", 100)),
	option.ResponseHeader("Content-Range", "Total number of users", param.StatusCodes(200, 206), param.Example("42 users", "0-10/42")),
)

type UsersResources struct {
	UsersService UsersService
}

type UsersError struct {
	Err     error  `json:"-" xml:"-"`
	Message string `json:"message" xml:"message"`
}

var _ error = UsersError{}

func (e UsersError) Error() string { return e.Err.Error() }

func (rs UsersResources) Routes(s *fuego.Server) {
	usersGroup := fuego.Group(s, "/users", option.Header("X-Header", "header description"))

	fuego.Get(usersGroup, "/", rs.filterUsers,
		optionPagination,
		option.Query("name", "Filter by name", param.Example("cat name", "felix"), param.Nullable()),
		option.QueryInt("younger_than", "Only get users younger than given age in years", param.Default(3)),
		option.Description("Filter users"),
	)

	fuego.Get(usersGroup, "/all", rs.getAllUsers,
		optionPagination,
		option.Tags("my-tag"),
		option.Description("Get all users"),
	)

	fuego.Get(usersGroup, "/by-age", rs.getAllUsersByAge, option.Description("Returns an array of users grouped by age"))
	fuego.Post(usersGroup, "/", rs.postUsers,
		option.DefaultStatusCode(201),
		option.AddError(409, "Conflict: User with the same name already exists", UsersError{}),
	)

	fuego.Get(usersGroup, "/{id}", rs.getUsers,
		option.Path("id", "User ID", param.Example("example", "123")),
	)
	fuego.Get(usersGroup, "/by-name/{name...}", rs.getUserByName)
	fuego.Put(usersGroup, "/{id}", rs.putUsers)
	fuego.Put(usersGroup, "/{id}/json", rs.putUsers,
		option.Summary("Update a user with JSON-only body"),
		option.RequestContentType("application/json"),
	)
	fuego.Delete(usersGroup, "/{id}", rs.deleteUsers)
}

func (rs UsersResources) getAllUsers(c fuego.ContextNoBody) ([]models.Users, error) {
	page := c.QueryParamInt("page")
	pageWithTypo := c.QueryParamInt("page-with-typo") // this shows a warning in the logs because "page-with-typo" is not a declared query param
	slog.Info("query params", "page", page, "page-with-typo", pageWithTypo)
	return rs.UsersService.GetAllUsers()
}

func (rs UsersResources) filterUsers(c fuego.ContextNoBody) ([]models.Users, error) {
	return rs.UsersService.FilterUsers(UsersFilter{
		Name:        c.QueryParam("name"),
		YoungerThan: c.QueryParamInt("younger_than"),
	})
}

func (rs UsersResources) getAllUsersByAge(c fuego.ContextNoBody) ([][]models.Users, error) {
	return rs.UsersService.GetAllUsersByAge()
}

func (rs UsersResources) postUsers(c *fuego.ContextWithBody[models.UsersCreate]) (models.Users, error) {
	body, err := c.Body()
	if err != nil {
		return models.Users{}, err
	}

	return rs.UsersService.CreateUsers(body)
}

func (rs UsersResources) getUsers(c fuego.ContextNoBody) (models.Users, error) {
	id := c.PathParam("id")

	return rs.UsersService.GetUsers(id)
}

func (rs UsersResources) getUserByName(c fuego.ContextNoBody) (models.Users, error) {
	name := c.PathParam("name")

	return rs.UsersService.GetUserByName(name)
}

func (rs UsersResources) putUsers(c *fuego.ContextWithBody[models.UsersUpdate]) (models.Users, error) {
	id := c.PathParam("id")

	body, err := c.Body()
	if err != nil {
		return models.Users{}, err
	}

	return rs.UsersService.UpdateUsers(id, body)
}

func (rs UsersResources) deleteUsers(c *fuego.ContextNoBody) (any, error) {
	return rs.UsersService.DeleteUsers(c.PathParam("id"))
}

type UsersFilter struct {
	Name        string
	YoungerThan int
}

type UsersService interface {
	GetUsers(id string) (models.Users, error)
	GetUserByName(name string) (models.Users, error)
	CreateUsers(models.UsersCreate) (models.Users, error)
	GetAllUsers() ([]models.Users, error)
	FilterUsers(UsersFilter) ([]models.Users, error)
	GetAllUsersByAge() ([][]models.Users, error)
	UpdateUsers(id string, input models.UsersUpdate) (models.Users, error)
	DeleteUsers(id string) (any, error)
}
