package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusRequested Status = "requested"
	StatusActive    Status = "active"
	StatusBlocked   Status = "blocked"
	StatusClosed    Status = "closed"
)

var (
	ErrNotFound          = errors.New("card not found")
	ErrUnknownStatus     = errors.New("unknown status")
	ErrInvalidTransition = errors.New("invalid status transition")
)

type Card struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Status    Status    `json:"status" db:"status"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func New(userId string) *Card {
	return &Card{
		ID:        uuid.NewString(),
		UserID:    userId,
		Status:    StatusRequested,
		UpdatedAt: time.Now().UTC(),
	}
}

func (card *Card) touch() {
	card.UpdatedAt = time.Now().UTC()
}
