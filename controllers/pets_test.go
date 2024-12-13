package controller_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-fuego/fuego/examples/userstore/lib"
)

func TestGetAllUsers(t *testing.T) {
	t.Run("can get all users", func(t *testing.T) {
		s := lib.NewMudServer()

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/users/all?per_page=5", nil)

		s.Mux.ServeHTTP(w, r)

		require.Equal(t, http.StatusOK, w.Code)
	})
}

func TestFilterUsers(t *testing.T) {
	t.Run("can filter users", func(t *testing.T) {
		s := lib.NewMudServer()

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/users/?name=kit&per_page=5", nil)
		s.Mux.ServeHTTP(w, r)
		t.Log(w.Body.String())
		require.Equal(t, http.StatusOK, w.Code)

		w = httptest.NewRecorder()
		r = httptest.NewRequest("GET", "/users/?name=kit&younger_than=1&per_page=5", nil)
		s.Mux.ServeHTTP(w, r)
		t.Log(w.Body.String())
		require.Equal(t, http.StatusOK, w.Code)
	})
}

func TestPostUsers(t *testing.T) {
	t.Run("can create a user", func(t *testing.T) {
		s := lib.NewMudServer()

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/users/", strings.NewReader(`{"name": "kitkat"}`))

		s.Mux.ServeHTTP(w, r)

		require.Equal(t, http.StatusCreated, w.Code)
		userId := w.Body.String()
		t.Log(userId)
		require.NotEmpty(t, userId)
	})
}

func TestGetUsers(t *testing.T) {
	t.Run("can get a user by id", func(t *testing.T) {
		s := lib.NewMudServer()

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/users/", strings.NewReader(`{"name": "kitkat"}`))
		s.Mux.ServeHTTP(w, r)
		t.Log(w.Body.String())
		require.Equal(t, http.StatusCreated, w.Code)

		w = httptest.NewRecorder()
		r = httptest.NewRequest("GET", "/users/user-1", nil)

		s.Mux.ServeHTTP(w, r)

		t.Log(w.Body.String())
		require.Equal(t, http.StatusOK, w.Code)
	})
}

func TestGetAllPestByAge(t *testing.T) {
	t.Run("can get a user by id", func(t *testing.T) {
		s := lib.NewMudServer()

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/users/", strings.NewReader(`{"name": "kitkat"}`))
		s.Mux.ServeHTTP(w, r)
		t.Log(w.Body.String())
		require.Equal(t, http.StatusCreated, w.Code)

		w = httptest.NewRecorder()
		r = httptest.NewRequest("GET", "/users/by-age", nil)

		s.Mux.ServeHTTP(w, r)

		t.Log(w.Body.String())
		require.Equal(t, http.StatusOK, w.Code)
	})
}

func TestGetUsersByName(t *testing.T) {
	t.Run("can get a user by name", func(t *testing.T) {
		s := lib.NewMudServer()

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/users/", strings.NewReader(`{"name": "kitkat"}`))
		s.Mux.ServeHTTP(w, r)
		t.Log(w.Body.String())
		require.Equal(t, http.StatusCreated, w.Code)

		w = httptest.NewRecorder()
		r = httptest.NewRequest("GET", "/users/by-name/kitkat", nil)

		s.Mux.ServeHTTP(w, r)

		t.Log(w.Body.String())
		require.Equal(t, http.StatusOK, w.Code)
	})
}

func TestPutUsers(t *testing.T) {
	t.Run("can update a user", func(t *testing.T) {
		s := lib.NewMudServer()

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/users/", strings.NewReader(`{"name": "kitkat"}`))
		s.Mux.ServeHTTP(w, r)
		t.Log(w.Body.String())
		require.Equal(t, http.StatusCreated, w.Code)

		w = httptest.NewRecorder()
		r = httptest.NewRequest("PUT", "/users/user-1", strings.NewReader(`{"name": "snickers"}`))

		s.Mux.ServeHTTP(w, r)

		t.Log(w.Body.String())
		require.Equal(t, http.StatusOK, w.Code)
	})
}

func TestDeleteUsers(t *testing.T) {
	t.Run("can delete a user", func(t *testing.T) {
		s := lib.NewMudServer()

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/users/", strings.NewReader(`{"name": "kitkat"}`))
		s.Mux.ServeHTTP(w, r)
		t.Log(w.Body.String())
		require.Equal(t, http.StatusCreated, w.Code)

		w = httptest.NewRecorder()
		r = httptest.NewRequest("DELETE", "/users/user-1", nil)

		s.Mux.ServeHTTP(w, r)

		t.Log(w.Body.String())
		require.Equal(t, http.StatusOK, w.Code)
	})
}
