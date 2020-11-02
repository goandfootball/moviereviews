package user

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/goandfootball/test-api/pkg/responses"
	"github.com/goandfootball/test-api/pkg/user"
	"net/http"
	"strconv"
)

type UsrRouter struct {
	Repository user.Repository
}

func (ur *UsrRouter) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// 202010301943 TODO: search for http errors guide
	result, err := ur.Repository.SelectAllUsers(ctx)
	if err != nil {
		responses.ERROR(w, http.StatusNotImplemented, err)
	}

	responses.JSON(w, http.StatusOK, result)
}

func (ur *UsrRouter) GetByUsrId(w http.ResponseWriter, r *http.Request) {
	var (
		ctx        context.Context
		paramValue string
		where      user.User
		errStr     error
	)

	ctx = r.Context()
	paramValue = chi.URLParam(r, "id")

	where.Id, errStr = strconv.Atoi(paramValue)
	if errStr != nil {
		responses.ERROR(w, http.StatusBadRequest, errStr)
	}

	result, err := ur.Repository.SelectUserByUsrId(ctx, where)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	responses.JSON(w, http.StatusOK, result)
}

func (ur *UsrRouter) GetByUsername(w http.ResponseWriter, r *http.Request) {
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

func (ur *UsrRouter) PostUser(w http.ResponseWriter, r *http.Request) {
	var newUser user.User

	ctx := r.Context()

	errDec := json.NewDecoder(r.Body).Decode(&newUser)
	if errDec != nil {
		responses.ERROR(w, http.StatusBadRequest, errDec)
		return
	}

	errIns := ur.Repository.InsertUser(ctx, &newUser)
	if errIns != nil {
		responses.ERROR(w, http.StatusBadRequest, errIns)
		return
	}

	responses.JSON(w, http.StatusCreated, nil)
}

func (ur *UsrRouter) PutUser(w http.ResponseWriter, r *http.Request) {
	var (
		paramValue     string
		model, updates user.User

		errStr error
	)

	ctx := r.Context()

	paramValue = chi.URLParam(r, "id")
	model.Id, errStr = strconv.Atoi(paramValue)
	if errStr != nil {
		responses.ERROR(w, http.StatusBadRequest, errStr)
	}

	errDec := json.NewDecoder(r.Body).Decode(&updates)
	if errDec != nil {
		responses.ERROR(w, http.StatusBadRequest, errDec)
	}
	/*
		errBef := updates.BeforeUpdate()
		if errBef != nil {
			responses.ERROR(w, http.StatusBadRequest, errDec)
		}
	*/

	errUpd := ur.Repository.UpdateUser(ctx, &model, &updates)
	if errUpd != nil {
		responses.ERROR(w, http.StatusNotModified, errDec)
	}

	responses.JSON(w, http.StatusOK, nil)
}

func (ur *UsrRouter) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var (
		paramValue string
		where      user.User
		errStr     error
	)

	ctx := r.Context()

	paramValue = chi.URLParam(r, "id")
	where.Id, errStr = strconv.Atoi(paramValue)
	if errStr != nil {
		responses.ERROR(w, http.StatusBadRequest, errStr)
		return
	}

	errDel := ur.Repository.DeleteUserByUsrId(ctx, where)
	if errDel != nil {
		responses.ERROR(w, http.StatusNotModified, errDel)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
