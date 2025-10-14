package service

import (
	"card-service/gen/proto"
	"card-service/internal/adapter"
	"card-service/internal/apperr"
	"card-service/internal/model"
	"card-service/internal/repo"
	"context"

	"buf.build/go/protovalidate"
)

type CardService struct {
	//proto.UnimplementedCardServiceServer
	repo *repo.Repo
}

func NewCardService(r repo.IRepository) *CardService {
	return &CardService{repo: r}
}

func (cs *CardService) mustEmbedUnimplementedCardServiceServer() {}

func (cs *CardService) RequestCard(ctx context.Context, req *proto.RequestCardRequest) (*proto.RequestCardResponse, error) {
	if err := protovalidate.Validate(req); err != nil {
		return nil, err
	}

	count, err := cs.repo.CountCard(ctx, model.CardParams{UserID: req.GetUserId()})
	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, apperr.CardAlreadyExists
	}

	card := model.NewCard(req.GetUserId())
	if err := cs.repo.CreateCard(ctx, card); err != nil {
		return nil, err
	}

	return &proto.RequestCardResponse{Card: adapter.CardToProto(card)}, nil
}

func (cs *CardService) mutateCard(ctx context.Context, id int64, event model.Event) (*model.Card, error) {
	card, err := cs.repo.GetCard(ctx, model.CardParams{ID: id})
	if err != nil {
		return nil, err
	}

	state := model.NewCardSM(model.NewCardSMInput(card))

	if err := state.Transition(event); err != nil {
		return nil, err
	}

	if err := cs.repo.UpdateCardStatus(ctx, model.CardParams{ID: id}, card.Status); err != nil {
		return nil, err
	}

	return card, nil
}

func (cs *CardService) ActivateCard(ctx context.Context, req *proto.ActivateCardRequest) (*proto.ActivateCardResponse, error) {
	if err := protovalidate.Validate(req); err != nil {
		return nil, err
	}

	c, err := cs.mutateCard(ctx, req.GetId(), model.EventActivate)
	if err != nil {
		return nil, err
	}
	return &proto.ActivateCardResponse{Card: adapter.CardToProto(c)}, nil
}

func (cs *CardService) BlockCard(ctx context.Context, req *proto.BlockCardRequest) (*proto.BlockCardResponse, error) {
	if err := protovalidate.Validate(req); err != nil {
		return nil, err
	}

	c, err := cs.mutateCard(ctx, req.GetId(), model.EventBlock)
	if err != nil {
		return nil, err
	}
	return &proto.BlockCardResponse{Card: adapter.CardToProto(c)}, nil
}

func (cs *CardService) UnblockCard(ctx context.Context, req *proto.UnblockCardRequest) (*proto.UnblockCardResponse, error) {
	if err := protovalidate.Validate(req); err != nil {
		return nil, err
	}

	c, err := cs.mutateCard(ctx, req.GetId(), model.EventUnblock)
	if err != nil {
		return nil, err
	}
	return &proto.UnblockCardResponse{Card: adapter.CardToProto(c)}, nil
}

func (cs *CardService) CloseCard(ctx context.Context, req *proto.CloseCardRequest) (*proto.CloseCardResponse, error) {
	if err := protovalidate.Validate(req); err != nil {
		return nil, err
	}

	c, err := cs.mutateCard(ctx, req.GetId(), model.EventClose)
	if err != nil {
		return nil, err
	}
	return &proto.CloseCardResponse{Card: adapter.CardToProto(c)}, nil
}

func (cs *CardService) GetCard(ctx context.Context, req *proto.GetCardRequest) (*proto.GetCardResponse, error) {
	if err := protovalidate.Validate(req); err != nil {
		return nil, err
	}

	c, err := cs.repo.GetCard(ctx, model.CardParams{ID: req.GetId()})
	if err != nil {
		return nil, err
	}
	return &proto.GetCardResponse{Card: adapter.CardToProto(c)}, nil
}
