package user

import "context"

type Repository interface {
	// 202010311123 TODO: add resource for reset user password
	SelectAllUsers(ctx context.Context) ([]User, error)
	SelectUserByUsrId(ctx context.Context, cond *User) (User, error)
	SelectUserByUsername(ctx context.Context, cond *User) (User, error)
	InsertUser(ctx context.Context, value *User) error
	UpdateUser(ctx context.Context, model *User, updates *User) error
	DeleteUserByUsrId(ctx context.Context, delete *User) error
}
