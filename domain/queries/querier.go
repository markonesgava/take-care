package queries

// Querier is the base interface of an Query
type Querier interface {
	QueryName() string
}
