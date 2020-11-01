package user

import "context"

type Repository interface {
	// 202010311123 TODO: add resource for reset user password
	SelectAllUsers(ctx context.Context) ([]User, error)
	SelectUserByUsrId(ctx context.Context, where User) (User, error)
	SelectUserByUsername(ctx context.Context, where User) (User, error)
	InsertUser(ctx context.Context, new *User) error
	UpdateUser(ctx context.Context, model *User, updates *User) (User, error)
	DeleteUserByUsrId(ctx context.Context, where User) error
	/*
		Update(ctx context.Context, id uint, usr User) error
	*/
}
