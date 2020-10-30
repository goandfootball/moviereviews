package user

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (ur *URouter) UPaths() http.Handler {
	r := chi.NewRouter()

	r.Get("/", ur.GetAll)

	return r
}
