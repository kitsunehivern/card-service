package adapter

import (
	cardpb "card-service/gen/proto"
	"card-service/internal/model"
	"time"
)

func CardToProto(card *model.Card) *cardpb.Card {
	return &cardpb.Card{
		Id:             card.ID,
		UserId:         card.UserID,
		Debit:          card.Debit,
		Credit:         card.Credit,
		ExpirationDate: card.ExpirationDate.Format(time.RFC3339),
		Status:         string(card.Status),
		CreatedAt:      card.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      card.UpdatedAt.Format(time.RFC3339),
	}
}

func ProtoToCard(card *cardpb.Card) *model.Card {
	expirationDate, err := time.Parse(time.RFC3339, card.GetExpirationDate())
	if err != nil {
		panic("Cannot parse timestamp")
	}

	createdAt, err := time.Parse(time.RFC3339, card.GetCreatedAt())
	if err != nil {
		panic("Cannot parse timestamp")
	}

	updatedAt, err := time.Parse(time.RFC3339, card.GetUpdatedAt())
	if err != nil {
		panic("Cannot parse timestamp")
	}

	return &model.Card{
		ID:             card.GetId(),
		UserID:         card.GetUserId(),
		Debit:          card.GetDebit(),
		Credit:         card.GetCredit(),
		ExpirationDate: expirationDate,
		Status:         model.Status(card.GetStatus()),
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
}
