package user

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/goandfootball/test-api/pkg/responses"
	"github.com/goandfootball/test-api/pkg/user"
	"net/http"
	"strconv"
)

type URouter struct {
	Repository user.Repository
}

func (ur *URouter) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// 202010301943 TODO: search for http errors guide
	result, err := ur.Repository.SelectAllUsers(ctx)
	if err != nil {
		responses.ERROR(w, http.StatusNotImplemented, err)
	}

	responses.JSON(w, http.StatusOK, result)
}

func (ur *URouter) GetByUsrId(w http.ResponseWriter, r *http.Request) {
	var (
		ctx        context.Context
		paramValue string
		where      user.User
		errStr     error
	)

	ctx = r.Context()
	paramValue = chi.URLParam(r, "id")

	where.UsrId, errStr = strconv.Atoi(paramValue)
	if errStr != nil {
		responses.ERROR(w, http.StatusBadRequest, errStr)
	}

	result, err := ur.Repository.SelectUserByUsrId(ctx, where)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	responses.JSON(w, http.StatusOK, result)
}

func (ur *URouter) GetByUsername(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   context.Context
		where user.User
	)

	ctx = r.Context()
	where.Username = chi.URLParam(r, "username")

	result, err := ur.Repository.SelectUserByUsername(ctx, where)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	responses.JSON(w, http.StatusOK, result)
}

func (ur *URouter) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var (
		ctx        context.Context
		paramValue string
		where      user.User
		errStr     error
		errDel     error
	)

	ctx = r.Context()
	paramValue = chi.URLParam(r, "id")
	where.UsrId, errStr = strconv.Atoi(paramValue)
	if errStr != nil {
		responses.ERROR(w, http.StatusBadRequest, errStr)
	}

	errDel = ur.Repository.DeleteUserByUsrId(ctx, where)
	if errDel != nil {
		responses.ERROR(w, http.StatusNotModified, errDel)
	}

	responses.JSON(w, http.StatusOK, nil)
}
