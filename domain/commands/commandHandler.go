package commands

// CommandHandler is the basic interface os an command handler
type CommandHandler interface {
	CommandName() string
	Handle(Commander) (interface{}, error)
}
