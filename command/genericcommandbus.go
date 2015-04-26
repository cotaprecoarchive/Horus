package command

type GenericCommandBus struct {
	handlers []CommandHandler
}

func NewGenericCommandBus() CommandBus {
	return &GenericCommandBus{
		handlers: []CommandHandler{},
	}
}

func (bus *GenericCommandBus) Dispatch(cmd Command) {
	for _, handler := range bus.handlers {
		if handler.CanHandle(cmd) {
			handler.Handle(cmd)
		}
	}
}

func (bus *GenericCommandBus) PushHandler(handler CommandHandler) {
	bus.handlers = append(bus.handlers, handler)
}
