package commands

import (
	"fmt"
)

type SendTakerHello struct {
}

func NewSendTakerHelloCommandHandler() *SendTakerHello {
	return &SendTakerHello{}
}

func (c *SendTakerHello) Handle(takerName string) (error, string) {
	if len(takerName) == 0 {
		return fmt.Errorf("TakerName is required"), ""
	}

	return nil, fmt.Sprintf("Hello, %s 👋!", takerName)
}
