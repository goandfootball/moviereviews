package review

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Review, error)
	GetOne(ctx context.Context, id uint) (Review, error)
	GetByUser(ctx context.Context, userID uint) ([]Review, error)
	Create(ctx context.Context, rev *Review) error
	Update(ctx context.Context, id uint, rev Review) error
	Delete(ctx context.Context, id uint) error
}
