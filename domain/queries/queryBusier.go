package queries

// QueryBusier is the basic interface for a QueryBus
type QueryBusier interface {
	Send(Querier) (interface{}, error)
	Attach(QueryHandler) error
}
