package model

import (
	"card-service/internal/apperr"
	"fmt"
)

type Event string

const (
	EventActivate Event = "activate"
	EventBlock    Event = "block"
	EventUnblock  Event = "unblock"
	EventClose    Event = "close"
)

type CardState interface {
	Name() CardStatus
	Before(card *Card, evt Event) error
	After(card *Card, evt Event) error
}

var stateRegistry = map[CardStatus]func() CardState{
	CardStatusRequested: func() CardState { return &RequestedState{} },
	CardStatusActive:    func() CardState { return &ActiveState{} },
	CardStatusBlocked:   func() CardState { return &BlockedState{} },
	CardStatusExpired:   func() CardState { return &ExpiredState{} },
	CardStatusClosed:    func() CardState { return &ClosedState{} },
}

func createState(status CardStatus) CardState {
	if create, ok := stateRegistry[status]; ok {
		return create()
	}

	panic(fmt.Sprintf("Card status %v not registered", status))
}

type CardSMInput struct {
	card *Card
}

func NewCardSMInput(card *Card) CardSMInput {
	return CardSMInput{card: card}
}

type CardSM struct {
	input CardSMInput
}

func NewCardSM(input CardSMInput) *CardSM {
	return &CardSM{input: input}
}

type ConditionFunc func(*Card) bool

type NextStateCondition struct {
	next       CardStatus
	conditions []ConditionFunc
}

func (csm *CardSM) Validate(evt Event) error {
	events, ok := transConditions[csm.input.card.Status]
	if !ok {
		panic(fmt.Sprintf("Card status %v not registered", csm.input.card.Status))
	}

	state, ok := events[evt]
	if !ok {
		return apperr.CardInvalidStateTransition
	}

	for _, check := range state.conditions {
		if !check(csm.input.card) {
			return apperr.CardInvalidStateTransition
		}
	}

	return nil
}

func (csm *CardSM) Action(evt Event) CardState {
	return createState(transConditions[csm.input.card.Status][evt].next)
}

func canBeClosed(card *Card) bool {
	return card.Debit >= card.Credit
}

var transConditions = map[CardStatus]map[Event]NextStateCondition{
	CardStatusRequested: {
		EventActivate: {next: CardStatusActive},
		EventClose:    {next: CardStatusClosed, conditions: []ConditionFunc{canBeClosed}},
	}, CardStatusActive: {
		EventBlock: {next: CardStatusBlocked},
		EventClose: {next: CardStatusClosed, conditions: []ConditionFunc{canBeClosed}},
	}, CardStatusBlocked: {
		EventUnblock: {next: CardStatusActive},
		EventClose:   {next: CardStatusClosed, conditions: []ConditionFunc{canBeClosed}},
	}, CardStatusExpired: {
		EventClose: {next: CardStatusClosed, conditions: []ConditionFunc{canBeClosed}},
	}, CardStatusClosed: {},
}

func (csm *CardSM) Transition(evt Event) error {
	state := createState(csm.input.card.Status)

	if err := csm.Validate(evt); err != nil {
		return err
	}

	if err := state.Before(csm.input.card, evt); err != nil {
		return err
	}

	newState := csm.Action(evt)
	prevStatus := csm.input.card.Status
	csm.input.card.Status = newState.Name()

	if err := state.After(csm.input.card, evt); err != nil {
		csm.input.card.Status = prevStatus
		return err
	}

	return nil
}

type RequestedState struct{}

func (state *RequestedState) Name() CardStatus                   { return CardStatusRequested }
func (state *RequestedState) Before(card *Card, evt Event) error { return nil }
func (state *RequestedState) After(card *Card, evt Event) error  { return nil }

type ActiveState struct{}

func (state *ActiveState) Name() CardStatus                   { return CardStatusActive }
func (state *ActiveState) Before(card *Card, evt Event) error { return nil }
func (state *ActiveState) After(card *Card, evt Event) error  { return nil }

type BlockedState struct{}

func (state *BlockedState) Name() CardStatus                   { return CardStatusBlocked }
func (state *BlockedState) Before(card *Card, evt Event) error { return nil }
func (state *BlockedState) After(card *Card, evt Event) error  { return nil }

type ExpiredState struct{}

func (state *ExpiredState) Name() CardStatus                   { return CardStatusExpired }
func (state *ExpiredState) Before(card *Card, evt Event) error { return nil }
func (state *ExpiredState) After(card *Card, evt Event) error  { return nil }

type ClosedState struct{}

func (state *ClosedState) Name() CardStatus                   { return CardStatusClosed }
func (state *ClosedState) Before(card *Card, evt Event) error { return nil }
func (state *ClosedState) After(card *Card, evt Event) error  { return nil }
