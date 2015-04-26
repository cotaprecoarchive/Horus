package command

type CommandBus interface {
	Dispatch(cmd Command)
	PushHandler(handler CommandHandler)
}
