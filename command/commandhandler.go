package command

type CommandHandler interface {
	CanHandle(cmd Command) bool
	Handle(cmd Command)
}
