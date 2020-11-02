package review

import (
	"context"
	"github.com/goandfootball/test-api/internal/data"
	"github.com/goandfootball/test-api/pkg/review"
	"github.com/pkg/errors"
)

type DbReview struct {
	Data *data.Data
}

// 202011011439 TODO: validate pointers usage on gorm
func (dr *DbReview) reviewExists(ctx context.Context, review *review.Review) bool {
	count := dr.Data.Db.WithContext(ctx).Select(&review).RowsAffected
	if count == 0 {
		return false
	}

	return true
}

func (dr *DbReview) SelectAllReviews(ctx context.Context) ([]review.Review, error) {
	var reviews []review.Review

	err := dr.Data.Db.WithContext(ctx).Find(&reviews).Error
	if err != nil {
		return []review.Review{}, err
	}

	return reviews, nil
}

func (dr *DbReview) SelectReviewByRevId(ctx context.Context, cond review.Review) (review.Review, error) {
	var dest review.Review

	err := dr.Data.Db.WithContext(ctx).First(&dest, cond).Error
	if err != nil {
		return review.Review{}, err
	}

	return dest, nil
}

func (dr *DbReview) SelectReviewsByMovId(ctx context.Context, cond review.Review) ([]review.Review, error) {
	var dest []review.Review

	err := dr.Data.Db.WithContext(ctx).Find(&dest, cond).Error
	if err != nil {
		return []review.Review{}, err
	}

	return dest, nil
}

func (dr *DbReview) SelectReviewsByUsrId(ctx context.Context, cond review.Review) ([]review.Review, error) {
	var dest []review.Review

	err := dr.Data.Db.WithContext(ctx).Find(&dest, cond).Error
	if err != nil {
		return []review.Review{}, err
	}

	return dest, nil
}

func (dr *DbReview) InsertReview(ctx context.Context, create *review.Review) error {
	errBef := create.BeforeInsert(dr.Data.Db)
	if errBef != nil {
		return errBef
	}

	err := dr.Data.Db.WithContext(ctx).Create(&create).Error
	if err != nil {
		return err
	}

	return nil
}

func (dr *DbReview) UpdateReview(ctx context.Context, model *review.Review, updates *review.Review) error {
	errBef := updates.BeforeUpdate(dr.Data.Db)
	if errBef != nil {
		return errBef
	}

	errUpd := dr.Data.Db.WithContext(ctx).Model(&model).Updates(&updates).Error
	if errUpd != nil {
		return errUpd
	}

	return nil
}

func (dr *DbReview) DeleteReviewByRevId(ctx context.Context, delete review.Review) error {
	exists := dr.reviewExists(ctx, &delete)
	if exists == false {
		return errors.New("review doesn't exists")
	}

	err := dr.Data.Db.WithContext(ctx).Delete(&delete).Error
	if err != nil {
		return err
	}

	return nil
}