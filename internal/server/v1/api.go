package v1

import (
	"github.com/go-chi/chi"
	"github.com/goandfootball/test-api/internal/data"
	user2 "github.com/goandfootball/test-api/internal/data/user"
	"github.com/goandfootball/test-api/internal/server/v1/user"
	"net/http"
)

func New() http.Handler {
	r := chi.NewRouter()

	ur := &user.URouter{
		Repository: &user2.DbUser{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.UPaths())

	return r
}
