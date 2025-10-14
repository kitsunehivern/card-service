package model

import (
	"time"
)

type CardStatus string

const (
	CardStatusRequested CardStatus = "requested"
	CardStatusActive    CardStatus = "active"
	CardStatusBlocked   CardStatus = "blocked"
	CardStatusExpired   CardStatus = "expired"
	CardStatusClosed    CardStatus = "closed"
)

type CardType string

const (
	CardTypeGold     CardType = "gold"
	CardTypeDiamond  CardType = "diamond"
	CardTypePlatinum CardType = "platinum"
)

//go:generate goqueryset -in card.go -out queryset_card.go

// gen:qs
type Card struct {
	ID             int64      `json:"id" db:"id" gorm:"primaryKey;<-:create"`
	UserID         int64      `json:"user_id" db:"user_id" gorm:"<-:create"`
	Type           CardType   `json:"type" db:"type" gorm:"<-:create"`
	Debit          int64      `json:"debit" db:"debit"`
	Credit         int64      `json:"credit" db:"credit"`
	ExpirationDate time.Time  `json:"expiration_date" db:"expiration_date" gorm:"<-create"`
	Status         CardStatus `json:"status" db:"status"`
	CreatedAt      time.Time  `json:"created_at" db:"updated_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`
}

type CardParams struct {
	ID     int64
	UserID int64
}

func NewCard(id int64, userID int64, type2 CardType) *Card {
	return &Card{
		ID:             id,
		UserID:         userID,
		Type:           type2,
		Debit:          0,
		Credit:         0,
		ExpirationDate: time.Now().UTC().Add(time.Minute * 5),
		Status:         CardStatusRequested,
	}
}
