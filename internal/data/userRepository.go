package data

import (
	"context"
	"errors"
	"github.com/goandfootball/test-api/pkg/user"
)

type UserRepository struct {
	Data *Data
}

func (ud *UserRepository) userExists(ctx context.Context, model user.User) bool {
	count := ud.Data.Db.WithContext(ctx).Select(&model).RowsAffected
	if count == 0 {
		return false
	}

	return true
}

func (ud *UserRepository) SelectAllUsers(ctx context.Context) ([]user.User, error) {
	var dest []user.User

	err := ud.Data.Db.WithContext(ctx).Find(&dest).Error
	if err != nil {
		return []user.User{}, nil
	}

	return dest, nil
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

func (ud *UserRepository) InsertUser(ctx context.Context, new *user.User) error {
	var errBef, errCre error

	errBef = new.BeforeInsert(ud.Data.Db)
	if errBef != nil {
		return errBef
	}

	errCre = ud.Data.Db.WithContext(ctx).Create(&new).Error
	if errCre != nil {
		return errCre
	}

	return nil
}

func (ud *UserRepository) UpdateUser(ctx context.Context, model *user.User, updates *user.User) error {
	errBef := updates.BeforeUpdate(ud.Data.Db)
	if errBef != nil {
		return errBef
	}

	errUpd := ud.Data.Db.WithContext(ctx).Model(&model).Updates(&updates).Error
	if errUpd != nil {
		return errUpd
	}
	return nil
}

func (ud *UserRepository) DeleteUserByUsrId(ctx context.Context, delete user.User) error {
	exists := ud.userExists(ctx, delete)
	if exists == false {
		return errors.New("user doesn't exist")
	}

	err := ud.Data.Db.WithContext(ctx).Delete(&delete).Error
	if err != nil {
		return err
	}

	return nil
}
