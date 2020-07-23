package commands

import (
	"fmt"
	"testing"
)

func Test_Expect_String_with_length_when_call_the_handler(t *testing.T) {
	command := NewSendTakerHelloCommandHandler()
	name := "Garcia"
	_, greeting := command.Handle(name)

	if len(greeting) == 0 {
		t.Errorf("When sent %s as taker name didn't return a greeting", name)
	}
}

func Test_Expect_a_greeting_with_the_passed_name(t *testing.T) {
	command := NewSendTakerHelloCommandHandler()
	name := "Garcia"
	_, greeting := command.Handle(name)

	if fmt.Sprintf("Hello, %s 👋!", name) != greeting {
		t.Errorf("The greeting is not as expected")
	}
}
