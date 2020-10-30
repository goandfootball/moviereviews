package data

import (
	"context"
	"github.com/goandfootball/test-api/pkg/user"
)

type UserData struct {
	Data *Data
}

func (ud *UserData) GetAll(ctx context.Context) ([]user.User, error) {
	var modelUsers []user.User

	if err := ud.Data.Db.WithContext(ctx).Find(&modelUsers).Error; err != nil {
		return []user.User{}, nil
	}

	return modelUsers, nil
}
