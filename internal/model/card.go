package model

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusRequested Status = "requested"
	StatusActive    Status = "active"
	StatusBlocked   Status = "blocked"
	StatusRetired   Status = "retired"
	StatusClosed    Status = "closed"
)

type Card struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Debit     int64     `json:"debit" db:"debit"`
	Credit    int64     `json:"credit" db:"credit"`
	Status    Status    `json:"status" db:"status"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func New(userId string) *Card {
	return &Card{
		ID:        uuid.NewString(),
		UserID:    userId,
		Debit:     0,
		Credit:    0,
		Status:    StatusRequested,
		UpdatedAt: time.Now().UTC(),
	}
}
