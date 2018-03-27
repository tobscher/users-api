package handlers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/satori/go.uuid"

	"github.com/stretchr/testify/assert"
	"github.com/tobscher/users-api/core"
)

type stubUsersService struct {
	users []*core.User
	err   error
}

func (s *stubUsersService) All() ([]*core.User, error) {
	return s.users, s.err
}

func TestUsersIndexOnErrorReturnsInternalServerError(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	handler := &Users{
		service: &stubUsersService{
			err: errors.New("An unexpected error occurred"),
		},
	}
	handler.index(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, http.StatusInternalServerError, "expected response status code to be 501")
	assert.Equal(t, string(body), "{\"message\":\"An unexpected error occurred\"}\n", "expected body to include error")
}

func TestUsersIndex(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	id, _ := uuid.FromString("961dfc7b-f02e-4db7-b543-3e904a3a830b")
	id2, _ := uuid.FromString("b9fbdbe0-0b55-41c3-aecb-8f18103fa8da")
	handler := &Users{
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
	handler.index(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, http.StatusOK, "expected response status code to be 200")
	assert.Equal(t, string(body), "[{\"id\":\"961dfc7b-f02e-4db7-b543-3e904a3a830b\",\"name\":\"Foo\"},{\"id\":\"b9fbdbe0-0b55-41c3-aecb-8f18103fa8da\",\"name\":\"Bar\"}]\n", "expected body to return users")
}
