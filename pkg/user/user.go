package user

import (
	"gorm.io/gorm"
	"time"

	validator "github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int       `json:"usr_id,omitempty" gorm:"column:usr_id;primary_key;<-:false"`
	FirstName    string    `json:"usr_first_name,omitempty" gorm:"column:usr_first_name;not null"`
	LastName     string    `json:"usr_last_name,omitempty" gorm:"column:usr_last_name;not null" validate:"necsfield=FirstName"`
	Username     string    `json:"usr_username,omitempty" gorm:"column:usr_username;unique" validate:"alphanumunicode"`
	Email        string    `json:"usr_email,omitempty" gorm:"column:usr_email;unique" validate:"email"`
	Password     string    `json:"usr_password,omitempty" gorm:"column:usr_password;not null"`
	Picture      string    `json:"usr_picture,omitempty" gorm:"column:usr_picture"`
	PasswordHash string    `json:"-" gorm:"-"`
	CreatedAt    time.Time `json:"usr_created_at,omitempty" gorm:"column:usr_created_at;not null;autoCreateTime:nano;<-:create"`
	UpdatedAt    time.Time `json:"usr_updated_at,omitempty" gorm:"column:usr_updated_at;autoUpdateTime:nano;<-:update"`
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

func (usr *User) ValidateBeforeInsert() error {
	val := validator.New()

	err := val.Struct(usr)
	if err != nil {
		return err
	}

	return nil
}

func (usr *User) Prepare() error {
	hashedPassword, err := hash(usr.Password)
	if err != nil {
		return err
	}

	usr.Password = hashedPassword

	return nil
}

func (usr User) BeforeUpdate(tx *gorm.DB) error {
	usr.UpdatedAt = time.Now()

	return nil
}
