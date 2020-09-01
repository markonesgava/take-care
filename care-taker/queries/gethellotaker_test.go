package queries

import (
	"fmt"
	"testing"

	domainQueries "github.com/markonesgava/take-care/care-taker/domain/queries"
)

func Test_Expect_String_with_length_when_call_the_handler(t *testing.T) {
	handler := newGetTakerHelloHandler()

	name := "Garcia"

	greeting, _ := handler.Handle(domainQueries.NewGetTakerHello(name))

	if len(greeting.(string)) == 0 {
		t.Errorf("When sent %s as taker name didn't return a greeting", name)
	}
}

func Test_Expect_a_greeting_with_the_passed_name(t *testing.T) {
	handler := newGetTakerHelloHandler()
	name := "Garcia"

	greeting, _ := handler.Handle(domainQueries.NewGetTakerHello(name))

	if fmt.Sprintf("Hello, %s 👋!", name) != greeting {
		t.Errorf("The greeting is not as expected")
	}
}
