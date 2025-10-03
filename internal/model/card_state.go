package model

import (
	"card-service/internal/errmsg"
	"fmt"
)

type Event string

const (
	EventActivate Event = "activate"
	EventBlock    Event = "block"
	EventUnblock  Event = "unblock"
	EventRetire   Event = "retire"
	EventClose    Event = "close"
)

type CardState interface {
	Name() Status
	Before(card *Card, evt Event) error
	After(card *Card, evt Event) error
}

var stateRegistry = map[Status]func() CardState{
	StatusRequested: func() CardState { return &RequestedState{} },
	StatusActive:    func() CardState { return &ActiveState{} },
	StatusBlocked:   func() CardState { return &BlockedState{} },
	StatusClosed:    func() CardState { return &ClosedState{} },
}

func createState(status Status) CardState {
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
	next       Status
	conditions []ConditionFunc
}

func (csm *CardSM) Validate(evt Event) error {
	events, ok := transConditions[csm.input.card.Status]
	if !ok {
		panic(fmt.Sprintf("Card status %v not registered", csm.input.card.Status))
	}

	state, ok := events[evt]
	if !ok {
		return errmsg.CardInvalidStateTransition
	}

	for _, check := range state.conditions {
		if !check(csm.input.card) {
			return errmsg.CardInvalidStateTransition
		}
	}

	return nil
}

func (csm *CardSM) Action(evt Event) CardState {
	return createState(transConditions[csm.input.card.Status][evt].next)
}

func canBeRetired(card *Card) bool {
	return card.Debit > card.Credit
}

var transConditions = map[Status]map[Event]NextStateCondition{
	StatusRequested: {
		EventActivate: {next: StatusActive},
		EventClose:    {next: StatusClosed},
	}, StatusActive: {
		EventBlock:  {next: StatusBlocked},
		EventClose:  {next: StatusClosed},
		EventRetire: {next: StatusRetired, conditions: []ConditionFunc{canBeRetired}},
	}, StatusBlocked: {
		EventUnblock: {next: StatusActive},
		EventClose:   {next: StatusClosed},
		EventRetire:  {next: StatusRetired, conditions: []ConditionFunc{canBeRetired}},
	}, StatusRetired: {
		EventClose: {next: StatusClosed},
	}, StatusClosed: {},
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

func (state *RequestedState) Name() Status                       { return StatusRequested }
func (state *RequestedState) Before(card *Card, evt Event) error { return nil }
func (state *RequestedState) After(card *Card, evt Event) error  { return nil }

type ActiveState struct{}

func (state *ActiveState) Name() Status                       { return StatusActive }
func (state *ActiveState) Before(card *Card, evt Event) error { return nil }
func (state *ActiveState) After(card *Card, evt Event) error  { return nil }

type BlockedState struct{}

func (state *BlockedState) Name() Status                       { return StatusBlocked }
func (state *BlockedState) Before(card *Card, evt Event) error { return nil }
func (state *BlockedState) After(card *Card, evt Event) error  { return nil }

type RetiredState struct{}

func (state *RetiredState) Name() Status                       { return StatusRetired }
func (state *RetiredState) Before(card *Card, evt Event) error { return nil }
func (state *RetiredState) After(card *Card, evt Event) error  { return nil }

type ClosedState struct{}

func (state *ClosedState) Name() Status                       { return StatusClosed }
func (state *ClosedState) Before(card *Card, evt Event) error { return nil }
func (state *ClosedState) After(card *Card, evt Event) error  { return nil }
