package review

import (
	"github.com/go-chi/chi"
	"strconv"

	"encoding/json"
	"github.com/goandfootball/moviereviews/pkg/review"
	"net/http"
)

type RevRouter struct {
	Repository review.Repository
}

const (
	paramRevId      string = "id"
	paramMovId      string = "id"
	paramUsrId      string = "id"
	headerKey              = "content-type"
	contentTypeJSON        = "application/json"
)

func (rr *RevRouter) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, errSel := rr.Repository.SelectAllReviews(ctx)
	if errSel != nil {
		http.Error(w, errSel.Error(), http.StatusBadRequest)
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

func (rr *RevRouter) GetByRevId(w http.ResponseWriter, r *http.Request) {
	var (
		cond   review.Review
		errStr error
	)

	ctx := r.Context()
	paramsValue := chi.URLParam(r, paramRevId)

	cond.Id, errStr = strconv.Atoi(paramsValue)
	if errStr != nil {
		http.Error(w, errStr.Error(), http.StatusBadRequest)
		return
	}

	result, errSel := rr.Repository.SelectReviewByRevId(ctx, &cond)
	if errSel != nil {
		http.Error(w, errSel.Error(), http.StatusBadRequest)
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

func (rr *RevRouter) GetByMovId(w http.ResponseWriter, r *http.Request) {
	var (
		cond   review.Review
		errStr error
	)

	ctx := r.Context()
	paramMovId := chi.URLParam(r, paramMovId)

	cond.MovId, errStr = strconv.Atoi(paramMovId)
	if errStr != nil {
		http.Error(w, errStr.Error(), http.StatusBadRequest)
		return
	}

	result, errSel := rr.Repository.SelectReviewsByMovId(ctx, &cond)
	if errSel != nil {
		http.Error(w, errSel.Error(), http.StatusBadRequest)
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

func (rr *RevRouter) GetByUsrId(w http.ResponseWriter, r *http.Request) {
	var (
		cond   review.Review
		errStr error
	)

	ctx := r.Context()
	paramMovId := chi.URLParam(r, paramUsrId)

	cond.MovId, errStr = strconv.Atoi(paramMovId)
	if errStr != nil {
		http.Error(w, errStr.Error(), http.StatusBadRequest)
		return
	}

	result, errSel := rr.Repository.SelectReviewsByUsrId(ctx, &cond)
	if errSel != nil {
		http.Error(w, errSel.Error(), http.StatusBadRequest)
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

func (rr *RevRouter) PostReview(w http.ResponseWriter, r *http.Request) {
	var create review.Review

	ctx := r.Context()

	errDec := json.NewDecoder(r.Body).Decode(&create)
	if errDec != nil {
		http.Error(w, errDec.Error(), http.StatusBadRequest)
		return
	}

	errIns := rr.Repository.InsertReview(ctx, &create)
	if errIns != nil {
		http.Error(w, errIns.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(headerKey, contentTypeJSON)
	w.WriteHeader(http.StatusCreated)
}

func (rr *RevRouter) PutReview(w http.ResponseWriter, r *http.Request) {
	var (
		cond    review.Review
		updates review.Review

		errStr error
	)

	ctx := r.Context()
	paramValue := chi.URLParam(r, paramRevId)

	cond.Id, errStr = strconv.Atoi(paramValue)
	if errStr != nil {
		http.Error(w, errStr.Error(), http.StatusBadRequest)
		return
	}

	errDec := json.NewDecoder(r.Body).Decode(&updates)
	if errDec != nil {
		http.Error(w, errDec.Error(), http.StatusBadRequest)
		return
	}

	errUpd := rr.Repository.UpdateReview(ctx, &cond, &updates)
	if errUpd != nil {
		http.Error(w, errUpd.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(headerKey, contentTypeJSON)
	w.WriteHeader(http.StatusCreated)
	// 202011032130 TODO: Add output with updates result on database function
}

func (rr *RevRouter) DeleteReview(w http.ResponseWriter, r *http.Request) {
	var (
		cond   review.Review
		errStr error
	)

	ctx := r.Context()
	paramValue := chi.URLParam(r, paramRevId)

	cond.Id, errStr = strconv.Atoi(paramValue)
	if errStr != nil {
		http.Error(w, errStr.Error(), http.StatusBadRequest)
		return
	}

	errDel := rr.Repository.DeleteReviewByRevId(ctx, &cond)
	if errDel != nil {
		http.Error(w, errDel.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set(headerKey, contentTypeJSON)
	w.WriteHeader(http.StatusCreated)
}
