package mock

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mockrepo "card-service/gen/repo"
	"card-service/internal/model"
	"card-service/internal/server"
	"card-service/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func TestRequestCard_HTTP(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockrepo.NewMockRepository(ctrl)
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

			return nil
		})

	svc := service.NewCardService(mockRepo)
	r := server.NewRouter(svc)

	body, _ := json.Marshal(map[string]string{"user_id": "ABC"})
	req := httptest.NewRequest(http.MethodPost, "/card/request", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var created model.Card
	if err := json.Unmarshal([]byte(rec.Body.String()), &created); err != nil {
		t.Fatalf("expected card json, got %v", err)
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
