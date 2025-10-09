package test

import (
	"card-service/internal/errmsg"
	"card-service/internal/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCardTransition(t *testing.T) {
	testcases := []struct {
		name          string
		inputCard     *model.Card
		event         model.Event
		expectedCard  *model.Card
		expectedError error
	}{
		{
			name: "success for requested -> active",
			inputCard: &model.Card{
				ID:     1,
				UserID: "user-1",
				Debit:  0,
				Credit: 0,
				Status: model.StatusRequested,
			},
			event: model.EventActivate,
			expectedCard: &model.Card{
				ID:     1,
				UserID: "user-1",
				Debit:  0,
				Credit: 0,
				Status: model.StatusActive,
			},
			expectedError: nil,
		},
		{
			name: "failure for requested -> block",
			inputCard: &model.Card{
				ID:     2,
				UserID: "user-2",
				Debit:  10,
				Credit: 0,
				Status: model.StatusRequested,
			},
			event: model.EventBlock,
			expectedCard: &model.Card{
				ID:     2,
				UserID: "user-2",
				Debit:  10,
				Credit: 0,
				Status: model.StatusRequested,
			},
			expectedError: errmsg.CardInvalidStateTransition,
		},
		{
			name: "failure for active -> retired with debit < credit",
			inputCard: &model.Card{
				ID:     3,
				UserID: "user-3",
				Debit:  5,
				Credit: 10,
				Status: model.StatusActive,
			},
			event: model.EventRetire,
			expectedCard: &model.Card{
				ID:     3,
				UserID: "user-3",
				Debit:  5,
				Credit: 10,
				Status: model.StatusActive,
			},
			expectedError: errmsg.CardInvalidStateTransition,
		},
		{
			name: "failure for active -> retired with debit = credit",
			inputCard: &model.Card{
				ID:     4,
				UserID: "user-4",
				Debit:  10,
				Credit: 10,
				Status: model.StatusActive,
			},
			event: model.EventRetire,
			expectedCard: &model.Card{
				ID:     4,
				UserID: "user-4",
				Debit:  10,
				Credit: 10,
				Status: model.StatusActive,
			},
			expectedError: errmsg.CardInvalidStateTransition,
		},
		{
			name: "failure for active -> retired with debit > credit",
			inputCard: &model.Card{
				ID:     5,
				UserID: "user-5",
				Debit:  10,
				Credit: 5,
				Status: model.StatusActive,
			},
			event: model.EventRetire,
			expectedCard: &model.Card{
				ID:     5,
				UserID: "user-5",
				Debit:  10,
				Credit: 5,
				Status: model.StatusRetired,
			},
			expectedError: nil,
		},
	}

	for ti, tc := range testcases {
		t.Run(fmt.Sprintf("test case %v: %v", ti+1, tc.name), func(t *testing.T) {
			card := tc.inputCard
			sm := model.NewCardSM(model.NewCardSMInput(card))
			err := sm.Transition(tc.event)

			if tc.expectedError != nil {
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tc.expectedCard.ID, card.ID)
			require.Equal(t, tc.expectedCard.UserID, card.UserID)
			require.Equal(t, tc.expectedCard.Credit, card.Credit)
			require.Equal(t, tc.expectedCard.Debit, card.Debit)
			require.Equal(t, tc.expectedCard.Status, card.Status)
		})
	}
}
