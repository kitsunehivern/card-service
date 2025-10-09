package model

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	StatusRequested Status = "requested"
	StatusActive    Status = "active"
	StatusBlocked   Status = "blocked"
	StatusRetired   Status = "retired"
	StatusClosed    Status = "closed"
)

//go:generate goqueryset -in card.go

// gen:qs
type Card struct {
	gorm.Model
	ID        int64     `json:"id" db:"id" gorm:"primaryKey;autoIncrement"`
	UserID    string    `json:"user_id" db:"user_id" gorm:"uniqueIndex"`
	Debit     int64     `json:"debit" db:"debit"`
	Credit    int64     `json:"credit" db:"credit"`
	Status    Status    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"updated_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
}

func NewCard(userId string) *Card {
	return &Card{
		UserID:    userId,
		Debit:     0,
		Credit:    0,
		Status:    StatusRequested,
		DeletedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
