package review

import "context"

type Repository interface {
	SelectAllReviews(ctx context.Context) ([]Review, error)
	SelectReviewByRevId(ctx context.Context, cond Review) (Review, error)
	SelectReviewsByMovId(ctx context.Context, cond Review) ([]Review, error)
	SelectReviewsByUsrId(ctx context.Context, cond Review) ([]Review, error)
	InsertReview(ctx context.Context, create *Review) error
	UpdateReview(ctx context.Context, model *Review, updates *Review) error
	DeleteReviewByRevId(ctx context.Context, delete Review) error
}
