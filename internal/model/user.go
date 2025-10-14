package model

import (
	"time"
)

//go:generate goqueryset -in user.go -out queryset_user.go

// gen:qs
type User struct {
	ID           int64     `json:"id" db:"id" gorm:"primaryKey;<-create"`
	Name         string    `json:"name" db:"name"`
	PhoneNumber  string    `json:"phone_number" db:"phone_number" gorm:"<-create"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
	CreatedAt    time.Time `json:"created_at" db:"updated_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`
}

type UserParams struct {
	ID          int64
	PhoneNumber string
}

func NewUser(id int64, name string, phoneNumber string, passwordHash string) *User {
	return &User{
		ID:           id,
		Name:         name,
		PhoneNumber:  phoneNumber,
		PasswordHash: passwordHash,
	}
}
