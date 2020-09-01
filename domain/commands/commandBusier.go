package commands

// CommandBusier is the basic interface for command busier
type CommandBusier interface {
	Send(Commander) error
	Attach(CommandHandler) error
}
