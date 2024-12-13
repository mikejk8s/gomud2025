package services

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/go-fuego/fuego/examples/userstore/models"

	controller "gomud2025/controllers"
)

func NewInMemoryUsersService() *InMemoryUsersService {
	return &InMemoryUsersService{
		Users: []models.Users{},
		Incr:  new(int),
	}
}

type InMemoryUsersService struct {
	Users []models.Users
	Incr  *int
}

// FilterUsers implements controller.UsersService.
func (userService *InMemoryUsersService) FilterUsers(filter controller.UsersFilter) ([]models.Users, error) {
	users := []models.Users{}
	for _, p := range userService.Users {
		if filter.Name != "" && !strings.Contains(p.Name, filter.Name) {
			continue
		}
		if filter.YoungerThan != 0 && p.Age >= filter.YoungerThan {
			continue
		}

		users = append(users, p)
	}
	return users, nil
}

// GetUserByName implements controller.UsersService.
func (userService *InMemoryUsersService) GetUserByName(name string) (models.Users, error) {
	for _, p := range userService.Users {
		if p.Name == name {
			return p, nil
		}
	}
	return models.Users{}, errors.New("user not found")
}

// CreateUsers implements controller.UsersService.
func (userService *InMemoryUsersService) CreateUsers(c models.UsersCreate) (models.Users, error) {
	*userService.Incr++
	newUser := models.Users{
		ID:   fmt.Sprintf("user-%d", *userService.Incr),
		Name: c.Name,
		Age:  c.Age,
	}
	userService.Users = append(userService.Users, newUser)
	slog.Info("Created user", "id", newUser.ID)

	return newUser, nil
}

// DeleteUsers implements controller.UsersService.
func (userService *InMemoryUsersService) DeleteUsers(id string) (any, error) {
	for i, p := range userService.Users {
		if p.ID == id {
			userService.Users = append(userService.Users[:i], userService.Users[i+1:]...)
			return nil, nil
		}
	}
	return nil, errors.New("user not found")
}

// GetAllUsers implements controller.UsersService.
func (userService *InMemoryUsersService) GetAllUsers() ([]models.Users, error) {
	return userService.Users, nil
}

// GetAllUsersByAge implements controller.UsersService.
func (userService *InMemoryUsersService) GetAllUsersByAge() ([][]models.Users, error) {
	maxAge := 0
	for _, p := range userService.Users {
		if maxAge < p.Age {
			maxAge = p.Age
		}
	}
	users := make([][]models.Users, maxAge+1)
	for _, p := range userService.Users {
		users[p.Age] = append(users[p.Age], p)
	}
	return users, nil
}

// GetUsers implements controller.UsersService.
func (userService *InMemoryUsersService) GetUsers(id string) (models.Users, error) {
	for _, p := range userService.Users {
		if p.ID == id {
			return p, nil
		}
	}
	return models.Users{}, errors.New("user not found")
}

// UpdateUsers implements controller.UsersService.
func (userService *InMemoryUsersService) UpdateUsers(id string, input models.UsersUpdate) (models.Users, error) {
	for i, p := range userService.Users {
		if p.ID == id {
			if input.Name != "" {
				p.Name = input.Name
			}
			if input.Age != 0 {
				p.Age = input.Age
			}
			userService.Users[i] = p
			return p, nil
		}
	}
	return models.Users{}, errors.New("user not found")
}

var _ controller.UsersService = &InMemoryUsersService{}
