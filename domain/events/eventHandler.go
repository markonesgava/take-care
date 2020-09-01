package events

type EventHadler interface {
	Handle(event Eventer) error
	EventNames() []string
}
