package user

import (
	"github.com/goandfootball/test-api/pkg/responses"
	"github.com/goandfootball/test-api/pkg/user"
	"net/http"
)

type URouter struct {
	Repository user.Repository
}

func (ur *URouter) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// 202010301943 TODO: search for http errors guide
	result, err := ur.Repository.GetAll(ctx)
	if err != nil {
		responses.ERROR(w, http.StatusNotImplemented, err)
	}

	responses.JSON(w, http.StatusOK, result)
}
