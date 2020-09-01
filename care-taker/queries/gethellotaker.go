package queries

import (
	"fmt"

	domainQueries "github.com/markonesgava/take-care/care-taker/domain/queries"
	"github.com/markonesgava/take-care/domain/queries"
)

type getTakerHelloHandler struct {
}

func newGetTakerHelloHandler() *getTakerHelloHandler {
	return &getTakerHelloHandler{}
}

func (*getTakerHelloHandler) QueryName() string {
	return domainQueries.GetTakerHello{}.QueryName()
}

func (handler *getTakerHelloHandler) Handle(query queries.Querier) (interface{}, error) {
	getTakerHello := query.(domainQueries.GetTakerHello)
	if len(getTakerHello.Name()) == 0 {
		return "", fmt.Errorf("TakerName is required")
	}

	return fmt.Sprintf("Hello, %s 👋!", getTakerHello.Name()), nil
}
