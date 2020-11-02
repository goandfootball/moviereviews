package review

import (
	"github.com/go-chi/chi"
	"strconv"

	"encoding/json"
	"github.com/goandfootball/test-api/pkg/responses"
	"github.com/goandfootball/test-api/pkg/review"
	"net/http"
)

type RevRouter struct {
	Repository review.Repository
}

const (
	constRevId string = "id"
	constMovId string = "id"
	constUsrId string = "id"
)

func (rr *RevRouter) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, errSel := rr.Repository.SelectAllReviews(ctx)
	if errSel != nil {
		responses.ERROR(w, http.StatusBadRequest, errSel)
		return
	}

	responses.JSON(w, http.StatusOK, result)
}

func (rr *RevRouter) GetByRevId(w http.ResponseWriter, r *http.Request) {
	var (
		cond   review.Review
		errStr error
	)

	ctx := r.Context()
	paramsValue := chi.URLParam(r, constRevId)

	cond.Id, errStr = strconv.Atoi(paramsValue)
	if errStr != nil {
		responses.ERROR(w, http.StatusBadRequest, errStr)
		return
	}

	result, errSel := rr.Repository.SelectReviewByRevId(ctx, cond)
	if errSel != nil {
		responses.ERROR(w, http.StatusBadRequest, errSel)
		return
	}

	responses.JSON(w, http.StatusOK, result)
}

func (rr *RevRouter) GetByMovId(w http.ResponseWriter, r *http.Request) {
	var (
		cond   review.Review
		errStr error
	)

	ctx := r.Context()
	paramMovId := chi.URLParam(r, constMovId)

	cond.MovId, errStr = strconv.Atoi(paramMovId)
	if errStr != nil {
		responses.ERROR(w, http.StatusBadRequest, errStr)
		return
	}

	result, errSel := rr.Repository.SelectReviewsByMovId(ctx, cond)
	if errSel != nil {
		responses.ERROR(w, http.StatusBadRequest, errSel)
		return
	}

	responses.JSON(w, http.StatusOK, result)
}

func (rr *RevRouter) GetByUsrId(w http.ResponseWriter, r *http.Request) {
	var (
		cond   review.Review
		errStr error
	)

	ctx := r.Context()
	paramMovId := chi.URLParam(r, constUsrId)

	cond.MovId, errStr = strconv.Atoi(paramMovId)
	if errStr != nil {
		responses.ERROR(w, http.StatusBadRequest, errStr)
		return
	}

	result, errSel := rr.Repository.SelectReviewsByUsrId(ctx, cond)
	if errSel != nil {
		responses.ERROR(w, http.StatusBadRequest, errSel)
		return
	}

	responses.JSON(w, http.StatusOK, result)
}

func (rr *RevRouter) PostReview(w http.ResponseWriter, r *http.Request) {
	var create review.Review

	ctx := r.Context()

	errDec := json.NewDecoder(r.Body).Decode(&create)
	if errDec != nil {
		responses.ERROR(w, http.StatusBadRequest, errDec)
		return
	}

	errIns := rr.Repository.InsertReview(ctx, &create)
	if errIns != nil {
		responses.ERROR(w, http.StatusBadRequest, errIns)
		return
	}

	responses.JSON(w, http.StatusCreated, nil)
}

func (rr *RevRouter) PutReview(w http.ResponseWriter, r *http.Request) {
	var (
		cond    review.Review
		updates review.Review

		errStr error
	)

	ctx := r.Context()
	paramValue := chi.URLParam(r, constRevId)

	cond.Id, errStr = strconv.Atoi(paramValue)
	if errStr != nil {
		responses.ERROR(w, http.StatusBadRequest, errStr)
		return
	}

	errDec := json.NewDecoder(r.Body).Decode(&updates)
	if errDec != nil {
		responses.ERROR(w, http.StatusBadRequest, errDec)
		return
	}

	errUpd := rr.Repository.UpdateReview(ctx, &cond, &updates)
	if errUpd != nil {
		responses.ERROR(w, http.StatusBadRequest, errUpd)
		return
	}

	responses.JSON(w, http.StatusAccepted, nil)
}

func (rr *RevRouter) DeleteReview(w http.ResponseWriter, r *http.Request) {
	var (
		deleteRev review.Review
		errStr    error
	)

	ctx := r.Context()
	paramValue := chi.URLParam(r, constRevId)

	deleteRev.Id, errStr = strconv.Atoi(paramValue)
	if errStr != nil {
		responses.ERROR(w, http.StatusBadRequest, errStr)
		return
	}

	err := rr.Repository.DeleteReviewByRevId(ctx, deleteRev)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
