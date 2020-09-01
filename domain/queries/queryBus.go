package queries

import (
	"fmt"
	"sync"
)

// QueryBus is a implementation of QueryBusier
type QueryBus struct {
	QueryBusier
	mutex         sync.Mutex
	queryHandlers map[string]QueryHandler
}

//NewQueryBus create a new query bus
func NewQueryBus() QueryBusier {
	return &QueryBus{
		queryHandlers: make(map[string]QueryHandler),
	}
}

// Attach includes an handle
func (qb *QueryBus) Attach(handler QueryHandler) error {
	qb.mutex.Lock()
	defer qb.mutex.Unlock()

	if qb.queryHandlers[handler.QueryName()] != nil {
		return fmt.Errorf("query %s already attached", handler.QueryName())
	}

	qb.queryHandlers[handler.QueryName()] = handler
	return nil
}

// Send execute a query on an attached handle
func (qb *QueryBus) Send(query Querier) (interface{}, error) {
	if qb.queryHandlers[query.QueryName()] == nil {
		return nil, fmt.Errorf("no handler attached to %s query", query.QueryName())
	}

	handlerPointer := qb.queryHandlers[query.QueryName()]
	return handlerPointer.Handle(query)
}
