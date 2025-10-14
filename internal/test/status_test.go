package test

import (
	"card-service/internal/apperr"
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
				UserID: 1,
				Type:   model.CardTypeGold,
				Debit:  0,
				Credit: 0,
				Status: model.CardStatusRequested,
			},
			event: model.EventActivate,
			expectedCard: &model.Card{
				ID:     1,
				UserID: 1,
				Type:   model.CardTypeGold,
				Debit:  0,
				Credit: 0,
				Status: model.CardStatusActive,
			},
			expectedError: nil,
		},
		{
			name: "failure for requested -> block",
			inputCard: &model.Card{
				ID:     2,
				UserID: 2,
				Type:   model.CardTypeGold,
				Debit:  10,
				Credit: 0,
				Status: model.CardStatusRequested,
			},
			event: model.EventBlock,
			expectedCard: &model.Card{
				ID:     2,
				UserID: 2,
				Type:   model.CardTypeGold,
				Debit:  10,
				Credit: 0,
				Status: model.CardStatusRequested,
			},
			expectedError: apperr.CardInvalidStateTransition,
		},
		{
			name: "failure for active -> closed with debit < credit",
			inputCard: &model.Card{
				ID:     3,
				UserID: 3,
				Type:   model.CardTypeDiamond,
				Debit:  5,
				Credit: 10,
				Status: model.CardStatusActive,
			},
			event: model.EventClose,
			expectedCard: &model.Card{
				ID:     3,
				UserID: 3,
				Type:   model.CardTypeDiamond,
				Debit:  5,
				Credit: 10,
				Status: model.CardStatusActive,
			},
			expectedError: apperr.CardInvalidStateTransition,
		},
		{
			name: "success for active -> closed with debit = credit",
			inputCard: &model.Card{
				ID:     4,
				UserID: 4,
				Type:   model.CardTypeDiamond,
				Debit:  10,
				Credit: 10,
				Status: model.CardStatusActive,
			},
			event: model.EventClose,
			expectedCard: &model.Card{
				ID:     4,
				UserID: 4,
				Type:   model.CardTypeDiamond,
				Debit:  10,
				Credit: 10,
				Status: model.CardStatusClosed,
			},
			expectedError: nil,
		},
		{
			name: "success for active -> closed with debit > credit",
			inputCard: &model.Card{
				ID:     5,
				UserID: 5,
				Type:   model.CardTypePlatinum,
				Debit:  10,
				Credit: 5,
				Status: model.CardStatusActive,
			},
			event: model.EventClose,
			expectedCard: &model.Card{
				ID:     5,
				UserID: 5,
				Type:   model.CardTypePlatinum,
				Debit:  10,
				Credit: 5,
				Status: model.CardStatusClosed,
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
			require.Equal(t, tc.expectedCard.Type, card.Type)
			require.Equal(t, tc.expectedCard.Credit, card.Credit)
			require.Equal(t, tc.expectedCard.Debit, card.Debit)
			require.Equal(t, tc.expectedCard.Status, card.Status)
		})
	}
}
