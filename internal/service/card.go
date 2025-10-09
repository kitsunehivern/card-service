package service

import (
	cardpb "card-service/gen/proto"
	"card-service/internal/adapter"
	"card-service/internal/errmsg"
	"card-service/internal/model"
	"card-service/internal/repo"
	"context"

	"buf.build/go/protovalidate"
)

type CardService struct {
	cardpb.UnimplementedCardServiceServer
	repo repo.IRepository
}

func NewCardService(r repo.IRepository) *CardService {
	return &CardService{repo: r}
}

func (cs *CardService) mustEmbedUnimplementedCardServiceServer() {}

func (cs *CardService) RequestCard(ctx context.Context, req *cardpb.RequestCardRequest) (*cardpb.RequestCardResponse, error) {
	if err := protovalidate.Validate(req); err != nil {
		return nil, err
	}

	count, err := cs.repo.CountCardByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, errmsg.CardAlreadyExists
	}

	card := model.NewCard(req.GetUserId())
	if err := cs.repo.CreateCard(ctx, card); err != nil {
		return nil, err
	}

	return &cardpb.RequestCardResponse{Card: adapter.CardToProto(card)}, nil
}

func (cs *CardService) mutateCard(ctx context.Context, id int64, event model.Event) (*model.Card, error) {
	card, err := cs.repo.GetCardByID(ctx, id)
	if err != nil {
		return nil, err
	}

	state := model.NewCardSM(model.NewCardSMInput(card))

	if err := state.Transition(event); err != nil {
		return nil, err
	}

	if err := cs.repo.UpdateCardStatus(ctx, id, card.Status); err != nil {
		return nil, err
	}

	return card, nil
}

func (cs *CardService) ActivateCard(ctx context.Context, req *cardpb.ActivateCardRequest) (*cardpb.ActivateCardResponse, error) {
	if err := protovalidate.Validate(req); err != nil {
		return nil, err
	}

	c, err := cs.mutateCard(ctx, req.GetId(), model.EventActivate)
	if err != nil {
		return nil, err
	}
	return &cardpb.ActivateCardResponse{Card: adapter.CardToProto(c)}, nil
}

func (cs *CardService) BlockCard(ctx context.Context, req *cardpb.BlockCardRequest) (*cardpb.BlockCardResponse, error) {
	if err := protovalidate.Validate(req); err != nil {
		return nil, err
	}

	c, err := cs.mutateCard(ctx, req.GetId(), model.EventBlock)
	if err != nil {
		return nil, err
	}
	return &cardpb.BlockCardResponse{Card: adapter.CardToProto(c)}, nil
}

func (cs *CardService) UnblockCard(ctx context.Context, req *cardpb.UnblockCardRequest) (*cardpb.UnblockCardResponse, error) {
	if err := protovalidate.Validate(req); err != nil {
		return nil, err
	}

	c, err := cs.mutateCard(ctx, req.GetId(), model.EventUnblock)
	if err != nil {
		return nil, err
	}
	return &cardpb.UnblockCardResponse{Card: adapter.CardToProto(c)}, nil
}

func (cs *CardService) RetireCard(ctx context.Context, req *cardpb.RetireCardRequest) (*cardpb.RetireCardResponse, error) {
	if err := protovalidate.Validate(req); err != nil {
		return nil, err
	}

	c, err := cs.mutateCard(ctx, req.GetId(), model.EventRetire)
	if err != nil {
		return nil, err
	}
	return &cardpb.RetireCardResponse{Card: adapter.CardToProto(c)}, nil
}

func (cs *CardService) CloseCard(ctx context.Context, req *cardpb.CloseCardRequest) (*cardpb.CloseCardResponse, error) {
	if err := protovalidate.Validate(req); err != nil {
		return nil, err
	}

	c, err := cs.mutateCard(ctx, req.GetId(), model.EventClose)
	if err != nil {
		return nil, err
	}
	return &cardpb.CloseCardResponse{Card: adapter.CardToProto(c)}, nil
}

func (cs *CardService) GetCard(ctx context.Context, req *cardpb.GetCardRequest) (*cardpb.GetCardResponse, error) {
	if err := protovalidate.Validate(req); err != nil {
		return nil, err
	}

	c, err := cs.repo.GetCardByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &cardpb.GetCardResponse{Card: adapter.CardToProto(c)}, nil
}
