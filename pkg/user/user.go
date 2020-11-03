package user

import (
	"errors"
	"gorm.io/gorm"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int       `json:"usr_id,omitempty" gorm:"column:usr_id;primary_key;<-:false"`
	FirstName    string    `json:"usr_first_name,omitempty" gorm:"column:usr_first_name;not null"`
	LastName     string    `json:"usr_last_name,omitempty" gorm:"column:usr_last_name;not null"`
	Username     string    `json:"usr_username,omitempty" gorm:"column:usr_username;unique"`
	Email        string    `json:"usr_email,omitempty" gorm:"column:usr_email;unique"`
	Password     string    `json:"usr_password,omitempty" gorm:"column:usr_password;not null"`
	Picture      string    `json:"usr_picture,omitempty" gorm:"column:usr_picture"`
	PasswordHash string    `json:"-" gorm:"-"`
	CreatedAt    time.Time `json:"usr_created_at,omitempty" gorm:"column:usr_created_at;<-:create;not null"`
	UpdatedAt    time.Time `json:"usr_updated_at,omitempty" gorm:"column:usr_updated_at"`
}

func hash(password string) (string, error) {
	bytePassword := []byte(password)

	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	result := string(passwordHash)

	return result, nil
}

func (usr User) PasswordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(usr.PasswordHash), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func (usr *User) Prepare() error {
	if usr.Id != 0 {
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

	hashedPassword, err := hash(usr.Password)
	if err != nil {
		return err
	}

	usr.Password = hashedPassword
	usr.CreatedAt = time.Now()

	return nil
}

func (usr User) BeforeUpdate(tx *gorm.DB) error {
	usr.UpdatedAt = time.Now()

	return nil
}
