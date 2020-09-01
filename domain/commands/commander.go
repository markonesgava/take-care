package commands

// Commander is the basic interface of an command
type Commander interface {
	CommandName() string
}
