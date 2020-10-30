package v1

import (
	"github.com/go-chi/chi"
	"github.com/goandfootball/test-api/internal/data"
	"net/http"
)

func New() http.Handler {
	r := chi.NewRouter()

	ur := &URouter{
		Repository: &data.UserRepository{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.UPaths())

	return r
}
