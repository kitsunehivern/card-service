package adapter

import (
	"card-service/gen/proto"
	"card-service/internal/model"
	"time"
)

func CardToProto(card *model.Card) *proto.Card {
	return &proto.Card{
		Id:             card.ID,
		UserId:         card.UserID,
		Type:           string(card.Type),
		Debit:          card.Debit,
		Credit:         card.Credit,
		ExpirationDate: card.ExpirationDate.Format(time.RFC3339),
		Status:         string(card.Status),
		CreatedAt:      card.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      card.UpdatedAt.Format(time.RFC3339),
	}
}

func ProtoToCard(card *proto.Card) *model.Card {
	expirationDate, _ := time.Parse(time.RFC3339, card.GetExpirationDate())
	createdAt, _ := time.Parse(time.RFC3339, card.GetCreatedAt())
	updatedAt, _ := time.Parse(time.RFC3339, card.GetUpdatedAt())

	return &model.Card{
		ID:             card.GetId(),
		UserID:         card.GetUserId(),
		Type:           model.CardType(card.GetType()),
		Debit:          card.GetDebit(),
		Credit:         card.GetCredit(),
		ExpirationDate: expirationDate,
		Status:         model.CardStatus(card.GetStatus()),
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
}
