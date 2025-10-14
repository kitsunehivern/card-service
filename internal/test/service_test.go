package test

import (
	imock "card-service/gen/mock"
	cardpb "card-service/gen/proto"
	"card-service/internal/adapter"
	"card-service/internal/apperr"
	"card-service/internal/model"
	"card-service/internal/service"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func newServiceWithMock(t *testing.T) (*service.CardService, *imock.MockICardRepo) {
	repo := imock.NewMockICardRepo(t)
	svc := service.NewCardService(repo)
	return svc, repo
}

func TestRequestCardService(t *testing.T) {
	testcases := []struct {
		name           string
		userIDs        []int64
		expectedErrors []error
	}{
		{
			name:           "success for different user ids",
			userIDs:        []int64{1, 2, 3},
			expectedErrors: []error{nil, nil, nil},
		},
		{
			name:           "failure for any two same user ids",
			userIDs:        []int64{1, 2, 1},
			expectedErrors: []error{nil, nil, apperr.CardAlreadyExists},
		},
	}

	for ti, tc := range testcases {
		if len(tc.userIDs) != len(tc.expectedErrors) {
			panic("array of input and output must have same length")
		}

		t.Run(fmt.Sprintf("test case %v: %v", ti+1, tc.name), func(t *testing.T) {
			svc, repo := newServiceWithMock(t)

			createdUsers := map[int64]bool{}
			tmpCreatedUsers := map[int64]bool{}
			for i := 0; i < len(tc.userIDs); i++ {
				repo.
					On("CountCardByUserID", mock.Anything, mock.AnythingOfType("int64")).
					Return(func(ctx context.Context, userID int64) (int, error) {
						_, ok := createdUsers[userID]
						if ok {
							return 1, nil
						}
						return 0, nil
					}).
					Once()

				_, created := tmpCreatedUsers[tc.userIDs[i]]
				if created {
					continue
				}
				tmpCreatedUsers[tc.userIDs[i]] = true

				repo.
					On("CreateCard", mock.Anything, mock.AnythingOfType("*model.Card")).
					Return(func(ctx context.Context, c *model.Card) error {
						createdUsers[c.UserID] = true
						return nil
					}).Once()
			}

			for i := 0; i < len(tc.userIDs); i++ {
				resp, err := svc.RequestCard(context.Background(), &cardpb.RequestCardRequest{UserId: tc.userIDs[i]})

				if tc.expectedErrors[i] != nil {
					require.ErrorIs(t, err, tc.expectedErrors[i])
					require.Nil(t, resp)
				} else {
					require.NoError(t, err)
					require.NotNil(t, resp)

					card := adapter.ProtoToCard(resp.GetCard())
					require.NotEqual(t, int64(0), card.ID)
					require.Equal(t, tc.userIDs[i], card.UserID)
					require.Equal(t, int64(0), card.Debit)
					require.Equal(t, int64(0), card.Credit)
					require.Equal(t, model.CardStatusRequested, card.Status)
				}
			}

			repo.AssertExpectations(t)
		})
	}
}
