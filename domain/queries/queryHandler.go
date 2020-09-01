package queries

// QueryHandler is the basic interface for an query handler
type QueryHandler interface {
	QueryName() string
	Handle(Querier) (interface{}, error)
}
