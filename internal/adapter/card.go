package adapter

import (
	cardpb "card-service/gen/proto"
	"card-service/internal/model"
	"time"
)

func CardToProto(card *model.Card) *cardpb.Card {
	return &cardpb.Card{
		Id:        card.ID,
		UserId:    card.UserID,
		Debit:     card.Debit,
		Credit:    card.Credit,
		Status:    string(card.Status),
		UpdatedAt: card.UpdatedAt.Format(time.RFC3339),
	}
}

func ProtoToCard(card *cardpb.Card) *model.Card {
	updatedAt, err := time.Parse(time.RFC3339, card.GetUpdatedAt())
	if err != nil {
		updatedAt = time.Now().UTC()
	}

	return &model.Card{
		ID:        card.GetId(),
		UserID:    card.GetUserId(),
		Debit:     card.GetDebit(),
		Credit:    card.GetCredit(),
		Status:    model.Status(card.GetStatus()),
		UpdatedAt: updatedAt,
	}
}
