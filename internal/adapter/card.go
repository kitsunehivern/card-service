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
