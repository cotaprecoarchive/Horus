package receiver

import "github.com/CotaPreco/Horus/message"

type ReceiveStrategyChain struct {
	strategies []ReceiveStrategy
}

func NewReceiveStrategyChain(strategies []ReceiveStrategy) *ReceiveStrategyChain {
	return &ReceiveStrategyChain{
		strategies: strategies,
	}
}

// func (c *ReceiveStrategyChain) AddReceiveStrategy(strategy ReceiveStrategy) {
// 	c.strategies = append(c.strategies, strategy)
// }

func (c *ReceiveStrategyChain) CanReceive(message []byte) bool {
	for _, strategy := range c.strategies {
		if strategy.CanReceive(message) {
			return true
		}
	}

	return false
}

func (c *ReceiveStrategyChain) Receive(message []byte) message.MessageInterface {
	for _, strategy := range c.strategies {
		if strategy.CanReceive(message) {
			return strategy.Receive(message)
		}
	}

	return nil
}
