package handlers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"

	"github.com/stretchr/testify/assert"
	"github.com/tobscher/users-api/core"
)

func createJwtToken(access []string, key string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access": access,
	})
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

type stubUsersService struct {
	users []*core.User
	user  *core.User
	err   error
}

func (s *stubUsersService) All() ([]*core.User, error) {
	return s.users, s.err
}

func (s *stubUsersService) Show(id uuid.UUID) (*core.User, error) {
	return s.user, s.err
}

func TestUsersIndexOnErrorReturnsInternalServerError(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/users", nil)
	w := httptest.NewRecorder()

	handler := &users{
		service: &stubUsersService{
			err: errors.New("An unexpected error occurred"),
		},
	}
	handler.index().ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "expected response status code to be 500")
	assert.Equal(t, "{\"message\":\"An unexpected error occurred\"}\n", string(body), "expected body to include error")
}

func TestUsersIndex(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/users", nil)
	w := httptest.NewRecorder()

	id, _ := uuid.FromString("961dfc7b-f02e-4db7-b543-3e904a3a830b")
	id2, _ := uuid.FromString("b9fbdbe0-0b55-41c3-aecb-8f18103fa8da")
	handler := &users{
		service: &stubUsersService{
			users: []*core.User{
				&core.User{
					ID:   id,
					Name: "Foo",
				},
				&core.User{
					ID:   id2,
					Name: "Bar",
				},
			},
		},
	}
	handler.index().ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "expected response status code to be 200")
	assert.Equal(t, "[{\"id\":\"961dfc7b-f02e-4db7-b543-3e904a3a830b\",\"name\":\"Foo\"},{\"id\":\"b9fbdbe0-0b55-41c3-aecb-8f18103fa8da\",\"name\":\"Bar\"}]\n", string(body), "expected body to return users")
}

func TestUsersShowInvalidID(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/users/z", nil)
	validToken := createJwtToken([]string{"user"}, "valid")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", validToken))
	w := httptest.NewRecorder()

	handler := &users{
		service: &stubUsersService{},
	}

	key := []byte("valid")
	NewRouter(handler, key).ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "expected response status code to be 400")
	assert.Equal(t, "{\"message\":\"uuid: incorrect UUID length: z\"}\n", string(body), "expected body to return error")
}

func TestUsersShowOnErrorReturnsInternalServerError(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/users/2723b889-f546-4030-bebd-6a880889f82e", nil)
	validToken := createJwtToken([]string{"user"}, "valid")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", validToken))
	w := httptest.NewRecorder()

	handler := &users{
		service: &stubUsersService{
			err: errors.New("error"),
		},
	}

	key := []byte("valid")
	NewRouter(handler, key).ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "expected response status code to be 500")
	assert.Equal(t, "{\"message\":\"error\"}\n", string(body), "expected body to return error")
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"), "expected content type to be json")
}

func TestUsersShow(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/users/2723b889-f546-4030-bebd-6a880889f82e", nil)
	validToken := createJwtToken([]string{"user"}, "valid")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", validToken))
	w := httptest.NewRecorder()

	id, _ := uuid.FromString("2723b889-f546-4030-bebd-6a880889f82e")
	handler := &users{
		service: &stubUsersService{
			user: &core.User{
				ID:   id,
				Name: "Test",
			},
		},
	}

	key := []byte("valid")
	NewRouter(handler, key).ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "expected response status code to be 200")
	assert.Equal(t, "{\"id\":\"2723b889-f546-4030-bebd-6a880889f82e\",\"name\":\"Test\"}\n", string(body), "expected body to return error")
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"), "expected content type to be json")
}
