package user

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/goandfootball/moviereviews/pkg/user"
	"net/http"
	"strconv"
)

type UsrRouter struct {
	Repository user.Repository
}

const (
	paramUsrId       string = "id"
	paramUsrUsername string = "username"
	headerKey               = "content-type"
	contentTypeJSON         = "application/json"
)

func (ur *UsrRouter) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, err := ur.Repository.SelectAllUsers(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(headerKey, contentTypeJSON)
	w.WriteHeader(http.StatusOK)
	errEnc := json.NewEncoder(w).Encode(&result)
	if errEnc != nil {
		http.Error(w, errEnc.Error(), http.StatusInternalServerError)
		return
	}
}

func (ur *UsrRouter) GetByUsrId(w http.ResponseWriter, r *http.Request) {
	var (
		ctx        context.Context
		paramValue string
		cond       user.User
		errStr     error
	)

	ctx = r.Context()
	paramValue = chi.URLParam(r, paramUsrId)

	cond.Id, errStr = strconv.Atoi(paramValue)
	if errStr != nil {
		http.Error(w, errStr.Error(), http.StatusBadRequest)
		return
	}

	result, err := ur.Repository.SelectUserByUsrId(ctx, &cond)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(headerKey, contentTypeJSON)
	w.WriteHeader(http.StatusOK)
	errEnc := json.NewEncoder(w).Encode(&result)
	if errEnc != nil {
		http.Error(w, errEnc.Error(), http.StatusInternalServerError)
		return
	}
}

func (ur *UsrRouter) GetByUsername(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  context.Context
		cond user.User
	)

	ctx = r.Context()
	cond.Username = chi.URLParam(r, paramUsrUsername)

	result, err := ur.Repository.SelectUserByUsername(ctx, &cond)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(headerKey, contentTypeJSON)
	w.WriteHeader(http.StatusOK)
	errEnc := json.NewEncoder(w).Encode(&result)
	if errEnc != nil {
		http.Error(w, errEnc.Error(), http.StatusInternalServerError)
		return
	}
}

func (ur *UsrRouter) PostUser(w http.ResponseWriter, r *http.Request) {
	var newUser user.User

	ctx := r.Context()

	errDec := json.NewDecoder(r.Body).Decode(&newUser)
	if errDec != nil {
		http.Error(w, errDec.Error(), http.StatusBadRequest)
		return
	}

	errIns := ur.Repository.InsertUser(ctx, &newUser)
	if errIns != nil {
		http.Error(w, errIns.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(headerKey, contentTypeJSON)
	w.WriteHeader(http.StatusCreated)
}

func (ur *UsrRouter) PutUser(w http.ResponseWriter, r *http.Request) {
	var (
		paramValue     string
		model, updates user.User

		errStr error
	)

	ctx := r.Context()

	paramValue = chi.URLParam(r, paramUsrId)
	model.Id, errStr = strconv.Atoi(paramValue)
	if errStr != nil {
		http.Error(w, errStr.Error(), http.StatusBadRequest)
		return
	}

	errDec := json.NewDecoder(r.Body).Decode(&updates)
	if errDec != nil {
		http.Error(w, errDec.Error(), http.StatusBadRequest)
		return
	}

	errUpd := ur.Repository.UpdateUser(ctx, &model, &updates)
	if errUpd != nil {
		http.Error(w, errUpd.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(headerKey, contentTypeJSON)
	w.WriteHeader(http.StatusCreated)
	// 202011032130 TODO: Add output with updates result on database function
}

func (ur *UsrRouter) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var (
		paramValue string
		cond       user.User
		errStr     error
	)

	ctx := r.Context()

	paramValue = chi.URLParam(r, paramUsrId)
	cond.Id, errStr = strconv.Atoi(paramValue)
	if errStr != nil {
		http.Error(w, errStr.Error(), http.StatusBadRequest)
		return
	}

	errDel := ur.Repository.DeleteUserByUsrId(ctx, &cond)
	if errDel != nil {
		http.Error(w, errDel.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(headerKey, contentTypeJSON)
	w.WriteHeader(http.StatusCreated)
}
