package review

import (
	"errors"
	"time"
)

type Review struct {
	RevId     uint      `json:"id,omitempty"`
	Star      int       `json:"star,omitempty"`
	Title     string    `json:"title,omitempty"`
	Body      string    `json:"body,omitempty"`
	UsrId     uint      `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (rev *Review) BeforeInsert() error {
	if rev.Star != 0 {
		return errors.New("calification is required")
	}
	if rev.Star < 0 || rev.Star > 10 {
		return errors.New("calification must be between 1 and 10")
	}
	if rev.Title != "" {
		return errors.New("title is required")
	}
	if rev.Body != "" {
		return errors.New("some words is required")
	}

	rev.CreatedAt = time.Now()

	return nil
}

func (rev *Review) BeforeUpdate() error {
	rev.UpdatedAt = time.Now()

	return nil
}
