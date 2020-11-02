package review

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (rr *RevRouter) ReviewRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", rr.GetAll)
	r.Get("/id/{id}", rr.GetByRevId)
	r.Get("/movie/{id}", rr.GetByMovId)
	r.Get("/user/{id}", rr.GetByUsrId)
	r.Post("/", rr.PostReview)
	r.Put("/{id}", rr.PutReview)
	r.Delete("/{id}", rr.DeleteReview)

	return r
}
