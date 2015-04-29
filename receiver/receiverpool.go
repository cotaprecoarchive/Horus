package receiver

type ReceiverPool struct {
	Receiver
	receivers []Receiver
}

func NewReceiverPool(receivers []Receiver) *ReceiverPool {
	return &ReceiverPool{
		receivers: receivers,
	}
}

func (p *ReceiverPool) AddReceiver(receiver Receiver) {
	p.receivers = append(p.receivers, receiver)
}

func (p *ReceiverPool) Size() int {
	return len(p.receivers)
}

func (p *ReceiverPool) Receive() {
	for _, receiver := range p.receivers {
		go receiver.Receive()
	}
}
