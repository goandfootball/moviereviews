package user

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (ur *URouter) UPaths() http.Handler {
	r := chi.NewRouter()

	r.Get("/", ur.GetAll)
	r.Get("/id/{id}", ur.GetByUsrId)
	r.Get("/username/{username}", ur.GetByUsername)
	r.Post("/", ur.PostUser)
	r.Put("/{id}", ur.PutUser)
	r.Delete("/{id}", ur.DeleteUser)

	return r
}
