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
		Status:    card.Status,
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
		Status:    card.GetStatus(),
		UpdatedAt: updatedAt,
	}
}
