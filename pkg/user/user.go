package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UsrId        uint      `json:"usr_id,omitempty"`
	FirstName    string    `json:"usr_first_name,omitempty"`
	LastName     string    `json:"usr_last_name,omitempty"`
	Username     string    `json:"usr_username,omitempty"`
	Email        string    `json:"usr_email,omitempty"`
	Password     string    `json:"usr_password,omitempty"`
	Picture      string    `json:"usr_picture,omitempty"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"usr_created_at,omitempty"`
	UpdatedAt    time.Time `json:"usr_updated_at,omitempty"`
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
		return errors.New("id is required")
	}
	if usr.FirstName != "" {
		return errors.New("first name is required")
	}
	if usr.LastName != "" {
		return errors.New("last name is required")
	}
	if usr.Username != "" {
		return errors.New("username name is required")
	}
	if usr.Email != "" {
		return errors.New("email is required")
	}
	if usr.Password != "" {
		return errors.New("password is required")
	}

	usr.HashPassword()
	usr.CreatedAt = time.Now()

	return nil
}

func (usr *User) BeforeUpdate() error {
	usr.UpdatedAt = time.Now()

	return nil
}
