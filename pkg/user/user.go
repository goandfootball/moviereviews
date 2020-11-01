package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UsrId        int       `json:"usr_id,omitempty" gorm:"column:usr_id;primary_key"`
	FirstName    string    `json:"usr_first_name,omitempty" gorm:"column:usr_first_name"`
	LastName     string    `json:"usr_last_name,omitempty" gorm:"column:usr_last_name"`
	Username     string    `json:"usr_username,omitempty" gorm:"column:usr_username;unique"`
	Email        string    `json:"usr_email,omitempty" gorm:"column:usr_email;unique"`
	Password     string    `json:"usr_password,omitempty" gorm:"column:usr_password"`
	Picture      string    `json:"usr_picture,omitempty" gorm:"column:usr_picture"`
	PasswordHash string    `json:"-" gorm:"-"`
	CreatedAt    time.Time `json:"usr_created_at,omitempty" gorm:"column:usr_created_at"`
	UpdatedAt    time.Time `json:"usr_updated_at,omitempty" gorm:"column:usr_updated_at"`
}

func (usr *User) HashPassword() error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	usr.PasswordHash = string(passwordHash)

	return nil
}

func (usr User) PasswordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(usr.PasswordHash), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func (usr *User) BeforeInsert() error {
	if usr.UsrId != 0 {
		return errors.New("id is generated automatically in database")
	}
	if usr.FirstName == "" {
		return errors.New("first name is required")
	}
	if usr.LastName == "" {
		return errors.New("last name is required")
	}
	if usr.Username == "" {
		return errors.New("username name is required")
	}
	if usr.Email == "" {
		return errors.New("email is required")
	}
	if usr.Password == "" {
		return errors.New("password is required")
	}

	err := usr.HashPassword()
	if err != nil {
		return err
	}
	usr.Password = usr.PasswordHash
	usr.CreatedAt = time.Now()

	return nil
}

func (usr *User) BeforeUpdate() error {
	// 202010311214 TODO: add validations
	usr.UpdatedAt = time.Now()

	return nil
}
