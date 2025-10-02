package model

type Event string

const (
	EventActivate Event = "activate"
	EventBlock    Event = "block"
	EventUnblock  Event = "unblock"
	EventClose    Event = "close"
)

type CardState interface {
	Name() Status
	Validate(card *Card) error
	Before(card *Card, evt Event) error
	After(card *Card, evt Event) error
	Action(evt Event) (Status, error)
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

func (card *Card) Transition(evt Event) error {
	state, err := createState(card.Status)
	if err != nil {
		return err
	}

	if err := state.Validate(card); err != nil {
		return err
	}

	if err := state.Before(card, evt); err != nil {
		return err
	}

	newStatus, err := state.Action(evt)
	if err != nil {
		return err
	}

	prevStatus := card.Status
	card.Status = newStatus

	if err := state.After(card, evt); err != nil {
		card.Status = prevStatus
		return err
	}

	return nil
}

type RequestedState struct{}

func (state *RequestedState) Name() Status                       { return StatusRequested }
func (state *RequestedState) Validate(card *Card) error          { return nil }
func (state *RequestedState) Before(card *Card, evt Event) error { return nil }
func (state *RequestedState) After(card *Card, evt Event) error  { return nil }
func (state *RequestedState) Action(evt Event) (Status, error) {
	switch evt {
	case EventActivate:
		return StatusActive, nil
	default:
		return StatusNull, ErrInvalidTransition
	}
}

type ActiveState struct{}

func (state *ActiveState) Name() Status                       { return StatusActive }
func (state *ActiveState) Validate(card *Card) error          { return nil }
func (state *ActiveState) Before(card *Card, evt Event) error { return nil }
func (state *ActiveState) After(card *Card, evt Event) error  { return nil }
func (state *ActiveState) Action(evt Event) (Status, error) {
	switch evt {
	case EventBlock:
		return StatusBlocked, nil
	case EventClose:
		return StatusClosed, nil
	default:
		return StatusNull, ErrInvalidTransition
	}
}

type BlockedState struct{}

func (state *BlockedState) Name() Status                       { return StatusBlocked }
func (state *BlockedState) Validate(card *Card) error          { return nil }
func (state *BlockedState) Before(card *Card, evt Event) error { return nil }
func (state *BlockedState) After(card *Card, evt Event) error  { return nil }
func (state *BlockedState) Action(evt Event) (Status, error) {
	switch evt {
	case EventUnblock:
		return StatusActive, nil
	case EventClose:
		return StatusClosed, nil
	default:
		return StatusNull, ErrInvalidTransition
	}
}

type ClosedState struct{}

func (state *ClosedState) Name() Status                       { return StatusClosed }
func (state *ClosedState) Validate(card *Card) error          { return nil }
func (state *ClosedState) Before(card *Card, evt Event) error { return nil }
func (state *ClosedState) After(card *Card, evt Event) error  { return nil }
func (state *ClosedState) Action(evt Event) (Status, error) {
	switch evt {
	default:
		return StatusNull, ErrInvalidTransition
	}
}
