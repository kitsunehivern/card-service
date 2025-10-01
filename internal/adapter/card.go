package adapter

import (
	cardpb "card-service/gen/proto"
	"card-service/internal/model"
	"time"
)

func CardToProto(c *model.Card) *cardpb.Card {
	return &cardpb.Card{
		Id:        c.ID,
		UserId:    c.UserID,
		Status:    c.Status,
		UpdatedAt: c.UpdatedAt.Format(time.RFC3339),
	}
}

func ProtoToCard(c *cardpb.Card) *model.Card {
	updatedAt, err := time.Parse(time.RFC3339, c.GetUpdatedAt())
	if err != nil {
		updatedAt = time.Now().UTC()
	}

	return &model.Card{
		ID:        c.GetId(),
		UserID:    c.GetUserId(),
		Status:    c.GetStatus(),
		UpdatedAt: updatedAt,
	}
}
