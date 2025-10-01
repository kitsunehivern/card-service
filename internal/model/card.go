package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	StatusRequested = "REQUESTED"
	StatusActive    = "ACTIVE"
	StatusBlocked   = "BLOCK"
	StatusClosed    = "CLOSED"
)

var (
	ErrNotFound     = errors.New("model not found")
	ErrInvalidState = errors.New("invalid state transition")
)

type Card struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Status    string    `json:"status" db:"status"`
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

func (c *Card) touch() {
	c.UpdatedAt = time.Now().UTC()
}

func (c *Card) Activate() error {
	switch c.Status {
	case StatusRequested:
		c.Status = StatusActive
		c.touch()
		return nil
	default:
		return ErrInvalidState
	}
}

func (c *Card) Block() error {
	switch c.Status {
	case StatusActive:
		c.Status = StatusBlocked
		c.touch()
		return nil
	default:
		return ErrInvalidState
	}
}

func (c *Card) Unblock() error {
	switch c.Status {
	case StatusBlocked:
		c.Status = StatusActive
		c.touch()
		return nil
	default:
		return ErrInvalidState
	}
}

func (c *Card) Close() error {
	switch c.Status {
	case StatusActive, StatusBlocked:
		c.Status = StatusClosed
		c.touch()
		return nil
	default:
		return ErrInvalidState
	}
}
