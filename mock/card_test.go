package mock

import (
	cardpb "card-service/gen/proto"
	mockrepo "card-service/gen/repo"
	"card-service/internal/model"
	"card-service/internal/service"
	"context"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestRequestCard_Succeeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockrepo.NewMockRepository(ctrl)

	var created *model.Card
	mockRepo.EXPECT().
		Create(gomock.AssignableToTypeOf(&model.Card{})).
		DoAndReturn(func(c *model.Card) error {
			if c.ID == "" {
				t.Errorf("expected non-empty id, not empty one")
			}

			if c.UserID != "ABC" {
				t.Errorf("exptected user_id=%v, got %v", "ABC", c.UserID)
			}

			if c.Status != model.StatusRequested {
				t.Errorf("expected status=%v, got %v", model.StatusRequested, c.Status)
			}

			created = c
			return nil
		})

	svc := service.NewCardService(mockRepo)
	resp, err := svc.RequestCard(context.Background(), &cardpb.RequestCardRequest{UserId: "ABC"})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.GetCard().GetId() == "" {
		t.Errorf("expected non-empty id, not empty one")
	}

	if resp.GetCard().GetUserId() != "ABC" {
		t.Errorf("expected user_id=%v, got %v", "ABC", resp.GetCard().GetUserId())
	}

	if model.Status(resp.GetCard().GetStatus()) != model.StatusRequested {
		t.Errorf("expected status=%v, got %v", model.StatusRequested, resp.GetCard().GetStatus())
	}

	if created == nil {
		t.Fatalf("expected card to be created")
	}

	if created.ID == "" {
		t.Errorf("expected non-empty id, not empty one")
	}

	if created.UserID != "ABC" {
		t.Errorf("expected user_id=%v, got %v", "ABC", created.UserID)
	}

	if created.Status != model.StatusRequested {
		t.Errorf("expected status=%v, got %v", model.StatusRequested, created.Status)
	}
}

func TestActivateCard_FromRequest_ToActive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockrepo.NewMockRepository(ctrl)

	card := &model.Card{
		ID:     "ID",
		UserID: "ABC",
		Status: model.StatusRequested,
	}

	gomock.InOrder(
		mockRepo.EXPECT().
			Get("ID").
			Return(card, nil),
		mockRepo.EXPECT().
			Update(gomock.AssignableToTypeOf(&model.Card{})).
			DoAndReturn(func(c *model.Card) error {
				if c.Status != model.StatusActive {
					t.Fatalf("expected status=%v, got %v", model.StatusActive, c.Status)
				}

				return nil
			}),
	)

	svc := service.NewCardService(mockRepo)
	_, err := svc.ActivateCard(context.Background(), &cardpb.ActivateCardRequest{Id: "ID"})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
