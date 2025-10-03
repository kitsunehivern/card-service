package service

import (
	cardpb "card-service/gen/proto"
	"card-service/internal/adapter"
	"card-service/internal/model"
	"card-service/internal/repo"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CardService struct {
	repo repo.Repository
}

func NewCardService(r repo.Repository) *CardService {
	return &CardService{repo: r}
}

func (cs *CardService) RequestCard(ctx context.Context, req *cardpb.RequestCardRequest) (*cardpb.RequestCardResponse, error) {
	card := model.New(req.GetUserId())
	if err := cs.repo.Create(card); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &cardpb.RequestCardResponse{Card: adapter.CardToProto(card)}, nil
}

func (cs *CardService) mutateCard(ctx context.Context, id string, event model.Event) (*model.Card, error) {
	card, err := cs.repo.Get(id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	state := model.NewCardSM(model.NewCardSMInput(card))

	if err := state.Transition(event); err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if err := cs.repo.Update(card); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return card, nil
}

func (cs *CardService) ActivateCard(ctx context.Context, req *cardpb.ActivateCardRequest) (*cardpb.ActivateCardResponse, error) {
	c, err := cs.mutateCard(ctx, req.GetId(), model.EventActivate)
	if err != nil {
		return nil, err
	}
	return &cardpb.ActivateCardResponse{Card: adapter.CardToProto(c)}, nil
}

func (cs *CardService) BlockCard(ctx context.Context, req *cardpb.BlockCardRequest) (*cardpb.BlockCardResponse, error) {
	c, err := cs.mutateCard(ctx, req.GetId(), model.EventBlock)
	if err != nil {
		return nil, err
	}
	return &cardpb.BlockCardResponse{Card: adapter.CardToProto(c)}, nil
}

func (cs *CardService) UnblockCard(ctx context.Context, req *cardpb.UnblockCardRequest) (*cardpb.UnblockCardResponse, error) {
	c, err := cs.mutateCard(ctx, req.GetId(), model.EventUnblock)
	if err != nil {
		return nil, err
	}
	return &cardpb.UnblockCardResponse{Card: adapter.CardToProto(c)}, nil
}

func (cs *CardService) CloseCard(ctx context.Context, req *cardpb.CloseCardRequest) (*cardpb.CloseCardResponse, error) {
	c, err := cs.mutateCard(ctx, req.GetId(), model.EventClose)
	if err != nil {
		return nil, err
	}
	return &cardpb.CloseCardResponse{Card: adapter.CardToProto(c)}, nil
}

func (cs *CardService) GetCard(ctx context.Context, req *cardpb.GetCardRequest) (*cardpb.GetCardResponse, error) {
	c, err := cs.repo.Get(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &cardpb.GetCardResponse{Card: adapter.CardToProto(c)}, nil
}
