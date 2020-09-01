package events

type EventBusier interface {
	Publish(events ...Eventer) error
	Attach(EventHadler) error
}
