package mock

import (
	cardpb "card-service/gen/proto"
	mockrepo "card-service/gen/repo"
	"card-service/internal/adapter"
	"card-service/internal/errmsg"
	"card-service/internal/model"
	"card-service/internal/service"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func newServiceWithMock(t *testing.T) (*service.CardService, *mockrepo.IRepository) {
	repo := mockrepo.NewIRepository(t)
	svc := service.NewCardService(repo)
	return svc, repo
}

func TestRequestCard(t *testing.T) {
	testcases := []struct {
		name           string
		userIDs        []string
		expectedErrors []error
	}{
		{
			name:           "test case 1: success for different user ids",
			userIDs:        []string{"user-1", "user-2", "user-3"},
			expectedErrors: []error{nil, nil, nil},
		},
		{
			name:           "test case 2: failure for any two same user ids",
			userIDs:        []string{"user-1", "user-2", "user-1"},
			expectedErrors: []error{nil, nil, errmsg.CardAlreadyExists},
		},
	}

	for _, tc := range testcases {
		if len(tc.userIDs) != len(tc.expectedErrors) {
			panic("array of input and output must have same length")
		}

		t.Run(tc.name, func(t *testing.T) {
			svc, repo := newServiceWithMock(t)

			createdUsers := map[string]bool{}
			tmpCreatedUsers := map[string]bool{}
			for i := 0; i < len(tc.userIDs); i++ {
				repo.
					On("HasCreatedCard", mock.AnythingOfType("string")).
					Return(func(userID string) bool {
						_, created := createdUsers[userID]
						return created
					}, func(userID string) error {
						return nil
					}).
					Once()

				_, created := tmpCreatedUsers[tc.userIDs[i]]
				if created {
					continue
				}
				tmpCreatedUsers[tc.userIDs[i]] = true

				repo.
					On("CreateCard", mock.AnythingOfType("*model.Card")).
					Return(func(c *model.Card) error {
						createdUsers[c.UserID] = true
						return nil
					}).Once()
			}

			for i := 0; i < len(tc.userIDs); i++ {
				resp, err := svc.RequestCard(context.Background(), &cardpb.RequestCardRequest{UserId: tc.userIDs[i]})

				if tc.expectedErrors[i] != nil {
					require.Error(t, err)
					require.ErrorIs(t, err, tc.expectedErrors[i])
					require.Nil(t, resp)
				} else {
					require.NoError(t, err)
					require.NotNil(t, resp)

					card := adapter.ProtoToCard(resp.GetCard())
					require.NotEmpty(t, card.ID)
					require.Equal(t, tc.userIDs[i], card.UserID)
					require.Equal(t, int64(0), card.Debit)
					require.Equal(t, int64(0), card.Credit)
					require.Equal(t, model.StatusRequested, card.Status)
					require.WithinDuration(t, time.Now().UTC(), card.UpdatedAt, time.Second)
				}
			}

			repo.AssertExpectations(t)
			repo.AssertNotCalled(t, "UpdateCard", mock.Anything)
			repo.AssertNotCalled(t, "GetCard", mock.Anything)
		})
	}
}
