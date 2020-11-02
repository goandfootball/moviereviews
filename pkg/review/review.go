package review

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Review struct {
	Id        int       `json:"rev_id,omitempty" gorm:"column:rev_id;primary_key"`
	MovId     int       `json:"mov_id,omitempty" gorm:"column:mov_id;not null"`
	UsrId     uint      `json:"usr_id,omitempty" gorm:"column:usr_id;not null"`
	Star      int       `json:"rev_star,omitempty" gorm:"column:rev_star;not null"`
	Title     string    `json:"rev_title,omitempty" gorm:"column:rev_title;not null"`
	Body      string    `json:"rev_body,omitempty" gorm:"column:rev_body;not null"`
	CreatedAt time.Time `json:"rev_created_at,omitempty" gorm:"column:rev_created_at;<-:create;not null"`
	UpdatedAt time.Time `json:"rev_updated_at,omitempty" gorm:"column:rev_updated_at"`
}

func (rev *Review) BeforeInsert(tx *gorm.DB) error {
	if rev.Star == 0 {
		return errors.New("qualification is required")
	}
	if rev.Star < 0 || rev.Star > 10 {
		return errors.New("qualification must be between 1 and 10")
	}
	if rev.Title == "" {
		return errors.New("title is required")
	}
	if rev.Body == "" {
		return errors.New("some words is required")
	}

	rev.CreatedAt = time.Now()

	return nil
}

func (rev *Review) BeforeUpdate(tx *gorm.DB) error {
	rev.UpdatedAt = time.Now()

	return nil
}
