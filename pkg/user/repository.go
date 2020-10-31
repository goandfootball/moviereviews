package user

import "context"

// Repository handle the CRUD operations with Users.
type Repository interface {
	SelectAllUsers(ctx context.Context) ([]User, error)
	SelectUserByUsrId(ctx context.Context, where User) (User, error)
	SelectUserByUsername(ctx context.Context, where User) (User, error)
	DeleteUserByUsrId(ctx context.Context, where User) error
	/*
		Create(ctx context.Context, user *User) error
		Update(ctx context.Context, id uint, usr User) error
	*/
}
