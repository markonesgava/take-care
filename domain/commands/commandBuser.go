package commands

// CommandBuser is the basic interface for command busier
type CommandBuser interface {
	Send(Commander) (interface{}, error)
	Attach(CommandHandler) error
}
