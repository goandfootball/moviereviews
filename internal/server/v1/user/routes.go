package user

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (ur *UsrRouter) UserRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", ur.GetAll)
	r.Get("/{id}", ur.GetByUsrId)
	r.Get("/username/{username}", ur.GetByUsername)
	r.Post("/", ur.PostUser)
	r.Put("/{id}", ur.PutUser)
	r.Delete("/{id}", ur.DeleteUser)

	return r
}
