package v1

import (
	"github.com/go-chi/chi"
	"github.com/goandfootball/test-api/internal/data"
	"github.com/goandfootball/test-api/internal/server/v1/user"
	"net/http"
)

func New() http.Handler {
	r := chi.NewRouter()

	ur := &user.URouter{
		Repository: &data.UserRepository{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.UPaths())

	return r
}
