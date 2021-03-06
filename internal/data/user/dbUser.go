package user

import (
	"context"
	"fmt"
	"github.com/goandfootball/moviereviews/internal/data"
	"github.com/goandfootball/moviereviews/pkg/user"
	"github.com/pkg/errors"
)

type DbUser struct {
	Data *data.Data
}

func (ud *DbUser) userExists(ctx context.Context, cond *user.User) bool {
	fmt.Println(cond.Username)
	count := ud.Data.Db.WithContext(ctx).Find(cond).RowsAffected
	fmt.Println("rows affected:", count)
	if count == 0 {
		return false
	}

	return true
}

func (ud *DbUser) SelectAllUsers(ctx context.Context) ([]user.User, error) {
	var dest []user.User

	err := ud.Data.Db.WithContext(ctx).Find(&dest).Error
	if err != nil {
		return []user.User{}, nil
	}

	return dest, nil
}

func (ud *DbUser) SelectUserByUsrId(ctx context.Context, cond *user.User) (user.User, error) {
	var result user.User

	exists := ud.userExists(ctx, cond)
	if exists == false {
		return user.User{}, errors.New("user doesn't exist")
	}

	err := ud.Data.Db.WithContext(ctx).Find(&result, cond).Error
	if err != nil {
		return user.User{}, err
	}

	return result, nil
}

func (ud *DbUser) SelectUserByUsername(ctx context.Context, cond *user.User) (user.User, error) {
	var result user.User

	exists := ud.userExists(ctx, cond)
	if exists == false {
		return user.User{}, errors.New("user doesn't exist")
	}

	err := ud.Data.Db.WithContext(ctx).Find(&result, cond).Error
	if err != nil {
		return user.User{}, err
	}

	return result, nil
}

func (ud *DbUser) InsertUser(ctx context.Context, value *user.User) error {
	var errBef, errCre error
	errBef = value.ValidateBeforeInsert()
	if errBef != nil {
		return errBef
	}

	errCre = ud.Data.Db.WithContext(ctx).Create(&value).Error
	if errCre != nil {
		return errCre
	}

	return nil
}

func (ud *DbUser) UpdateUser(ctx context.Context, model *user.User, updates *user.User) error {
	exists := ud.userExists(ctx, model)
	if exists == false {
		return errors.New("user doesn't exist")
	}

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

func (ud *DbUser) DeleteUserByUsrId(ctx context.Context, delete *user.User) error {
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
