package services

import (
	"testing"

	"github.com/stretchr/testify/require"

	controller "github.com/go-fuego/fuego/examples/userstore/controllers"
	"github.com/go-fuego/fuego/examples/userstore/models"
)

func TestInMemoryUsers(t *testing.T) {
	service := NewInMemoryUsersService()

	t.Run("can create a user", func(t *testing.T) {
		newUser, err := service.CreateUsers(models.UsersCreate{Name: "kitkat", Age: 1})
		require.NoError(t, err)
		newUser2, err := service.CreateUsers(models.UsersCreate{Name: "payday", Age: 3})
		require.NoError(t, err)
		require.Equal(t, "user-1", newUser.ID)
		require.Equal(t, "user-2", newUser2.ID)
	})

	t.Run("can get a user by name", func(t *testing.T) {
		newUser, err := service.GetUserByName("kitkat")
		require.NoError(t, err)
		require.Equal(t, "kitkat", newUser.Name)
		require.Equal(t, 1, newUser.Age)
	})

	t.Run("cannot get a user by name if it doesn't exists", func(t *testing.T) {
		_, err := service.GetUserByName("snickers")
		require.Error(t, err)
	})

	t.Run("can get a user by id", func(t *testing.T) {
		newUser, err := service.GetUsers("user-1")
		require.NoError(t, err)
		require.Equal(t, "kitkat", newUser.Name)
		require.Equal(t, 1, newUser.Age)
	})

	t.Run("can get all users", func(t *testing.T) {
		users, err := service.GetAllUsers()
		require.NoError(t, err)
		require.Len(t, users, 2)
	})

	t.Run("can filter users", func(t *testing.T) {
		users, err := service.FilterUsers(controller.UsersFilter{Name: "kit", YoungerThan: 5})
		require.NoError(t, err)
		require.Len(t, users, 1)
		require.Equal(t, "kitkat", users[0].Name)

		users, err = service.FilterUsers(controller.UsersFilter{Name: "kit", YoungerThan: 1})
		require.NoError(t, err)
		require.Len(t, users, 0)
	})

	t.Run("can get all users by age", func(t *testing.T) {
		users, err := service.GetAllUsersByAge()
		require.NoError(t, err)
		require.Len(t, users, 4)
		require.Equal(t, "kitkat", users[1][0].Name)
		require.Equal(t, "payday", users[3][0].Name)
	})

	t.Run("can update a user", func(t *testing.T) {
		updatedUser, err := service.UpdateUsers("user-1", models.UsersUpdate{Name: "snickers", Age: 2})
		require.NoError(t, err)
		require.Equal(t, "snickers", updatedUser.Name)
		require.Equal(t, 2, updatedUser.Age)
	})

	t.Run("can delete a user", func(t *testing.T) {
		_, err := service.DeleteUsers("user-1")
		require.NoError(t, err)

		_, err = service.GetUsers("user-1")
		require.Error(t, err)
	})

	t.Run("cannot get a user that does not exist", func(t *testing.T) {
		_, err := service.GetUsers("user-1")
		require.Error(t, err)
	})

	t.Run("cannot update a user that does not exist", func(t *testing.T) {
		_, err := service.UpdateUsers("user-1", models.UsersUpdate{Name: "snickers", Age: 2})
		require.Error(t, err)
	})

	t.Run("cannot delete a user that does not exist", func(t *testing.T) {
		_, err := service.DeleteUsers("user-1")
		require.Error(t, err)
	})
}
