package commands

import (
	"fmt"
	"sync"
)

//CommandBus is a implementation of CommandBusier
type CommandBus struct {
	CommandBusier
	mutex           sync.Mutex
	commandHandlers map[string]CommandHandler
}

//NewCommandBus create a new Command bus
func NewCommandBus() CommandBusier {
	return &CommandBus{
		commandHandlers: make(map[string]CommandHandler),
	}
}

//Attach includes an handle
func (qb *CommandBus) Attach(handler CommandHandler) error {
	qb.mutex.Lock()
	defer qb.mutex.Unlock()

	if qb.commandHandlers[handler.CommandName()] != nil {
		return fmt.Errorf("Command %s already attached", handler.CommandName())
	}

	qb.commandHandlers[handler.CommandName()] = handler
	return nil
}

//Send execute a Command on an attached handle
func (qb *CommandBus) Send(command Commander) error {
	if qb.commandHandlers[command.CommandName()] == nil {
		return fmt.Errorf("no handler attached to %s command", command.CommandName())
	}

	handler := qb.commandHandlers[command.CommandName()]
	return handler.Handle(command)
}
