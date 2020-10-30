package v1

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/goandfootball/test-api/pkg/response"
	"github.com/goandfootball/test-api/pkg/user"
	"net/http"
)

type URouter struct {
	Repository user.Repository
}

func (ur *URouter) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, err := ur.Repository.GetAll(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	response.EJSON(w, r, http.StatusOK, response.Map{"users": result})
}

func (ur *URouter) UPaths() http.Handler {
	r := chi.NewRouter()

	r.Get("/", ur.GetAll)

	return r
}
