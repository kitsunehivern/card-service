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
		CreatedAt: card.CreatedAt.Format(time.RFC3339),
		UpdatedAt: card.UpdatedAt.Format(time.RFC3339),
		DeletedAt: card.DeletedAt.Format(time.RFC3339),
	}
}

func ProtoToCard(card *cardpb.Card) *model.Card {
	createdAt, err := time.Parse(time.RFC3339, card.GetCreatedAt())
	if err != nil {
		createdAt = time.Now().UTC()
	}
	updatedAt, err := time.Parse(time.RFC3339, card.GetUpdatedAt())
	if err != nil {
		updatedAt = time.Now().UTC()
	}
	deletedAt, err := time.Parse(time.RFC3339, card.GetDeletedAt())
	if err != nil {
		deletedAt = time.Now().UTC()
	}

	return &model.Card{
		ID:        card.GetId(),
		UserID:    card.GetUserId(),
		Debit:     card.GetDebit(),
		Credit:    card.GetCredit(),
		Status:    model.Status(card.GetStatus()),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
