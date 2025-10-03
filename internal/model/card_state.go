package model

import (
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
	Name() Status
	Validate(card *Card, evt Event) error
	Before(card *Card, evt Event) error
	Action(evt Event) (CardState, error)
	After(card *Card, evt Event) error
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

var stateRegistry = map[Status]func() CardState{
	StatusRequested: func() CardState { return &RequestedState{} },
	StatusActive:    func() CardState { return &ActiveState{} },
	StatusBlocked:   func() CardState { return &BlockedState{} },
	StatusClosed:    func() CardState { return &ClosedState{} },
}

func createState(status Status) (CardState, error) {
	if create, ok := stateRegistry[status]; ok {
		return create(), nil
	}

	return nil, ErrUnknownStatus
}

func (sm *CardSM) Transition(evt Event) error {
	state, err := createState(sm.input.card.Status)
	if err != nil {
		return err
	}

	if err := state.Validate(sm.input.card, evt); err != nil {
		return err
	}

	if err := state.Before(sm.input.card, evt); err != nil {
		return err
	}

	newState, err := state.Action(evt)
	if err != nil {
		return err
	}

	prevStatus := sm.input.card.Status
	sm.input.card.Status = newState.Name()

	if err := state.After(sm.input.card, evt); err != nil {
		sm.input.card.Status = prevStatus
		return err
	}

	return nil
}

type RequestedState struct{}

func (state *RequestedState) Name() Status {
	return StatusRequested
}

func (state *RequestedState) Validate(card *Card, evt Event) error {
	switch evt {
	case EventActivate:
		return nil
	default:
		return ErrInvalidTransition
	}
}

func (state *RequestedState) Before(card *Card, evt Event) error {
	return nil
}

func (state *RequestedState) Action(evt Event) (CardState, error) {
	switch evt {
	case EventActivate:
		return createState(StatusActive)
	default:
		panic(fmt.Sprintf("Unexpected event: %v", evt))
	}
}

func (state *RequestedState) After(card *Card, evt Event) error {
	return nil
}

type ActiveState struct{}

func (state *ActiveState) Name() Status {
	return StatusActive
}

func (state *ActiveState) Validate(card *Card, evt Event) error {
	switch evt {
	case EventBlock, EventClose:
		return nil
	default:
		return ErrInvalidTransition
	}
}

func (state *ActiveState) Before(card *Card, evt Event) error {
	return nil
}

func (state *ActiveState) Action(evt Event) (CardState, error) {
	switch evt {
	case EventBlock:
		return createState(StatusBlocked)
	case EventClose:
		return createState(StatusClosed)
	default:
		panic(fmt.Sprintf("Unexpected event: %v", evt))
	}
}

func (state *ActiveState) After(card *Card, evt Event) error {
	return nil
}

type BlockedState struct{}

func (state *BlockedState) Name() Status {
	return StatusBlocked
}

func (state *BlockedState) Validate(card *Card, evt Event) error {
	switch evt {
	case EventUnblock, EventClose:
		return nil
	default:
		return ErrInvalidTransition
	}
}

func (state *BlockedState) Before(card *Card, evt Event) error {
	return nil
}

func (state *BlockedState) Action(evt Event) (CardState, error) {
	switch evt {
	case EventUnblock:
		return createState(StatusActive)
	case EventClose:
		return createState(StatusClosed)
	default:
		panic(fmt.Sprintf("Unexpected event: %v", evt))
	}
}

func (state *BlockedState) After(card *Card, evt Event) error {
	return nil
}

type ClosedState struct{}

func (state *ClosedState) Name() Status {
	return StatusClosed
}

func (state *ClosedState) Validate(card *Card, evt Event) error {
	switch evt {
	default:
		return ErrInvalidTransition
	}
}

func (state *ClosedState) Before(card *Card, evt Event) error {
	return nil
}

func (state *ClosedState) Action(evt Event) (CardState, error) {
	switch evt {
	default:
		panic(fmt.Sprintf("Unexpected event: %v", evt))
	}
}

func (state *ClosedState) After(card *Card, evt Event) error {
	return nil
}
