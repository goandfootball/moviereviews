package data

import (
	"context"
	"github.com/goandfootball/test-api/pkg/user"
)

type UserRepository struct {
	Data *Data
}

func (ud *UserRepository) SelectAllUsers(ctx context.Context) ([]user.User, error) {
	var modelUsers []user.User

	if err := ud.Data.Db.WithContext(ctx).Find(&modelUsers).Error; err != nil {
		return []user.User{}, nil
	}

	return modelUsers, nil
}

func (ud *UserRepository) SelectUserByUsrId(ctx context.Context, where user.User) (user.User, error) {
	var result user.User

	err := ud.Data.Db.WithContext(ctx).First(&result, where).Error
	if err != nil {
		return user.User{}, err
	}

	return result, nil
}

func (ud *UserRepository) SelectUserByUsername(ctx context.Context, where user.User) (user.User, error) {
	var result user.User

	err := ud.Data.Db.WithContext(ctx).First(&result, where).Error
	if err != nil {
		return user.User{}, err
	}

	return result, nil
}

func (ud *UserRepository) DeleteUserByUsrId(ctx context.Context, where user.User) error {
	var model user.User
	// 202030102315 TODO: when user id not exists return ok 200, fix...
	err := ud.Data.Db.WithContext(ctx).Delete(model, &where)
	if err != nil {
		return err.Error
	}

	return nil
}
