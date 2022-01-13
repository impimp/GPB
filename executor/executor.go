package executor

import (
	"GBP/consumers/cpu"
	"GBP/producer"
	"sync"
)

type Executor struct {
	producer producer.Producer
	consumer cpu.Consumer
}

func NewExecutor(p producer.Producer, c cpu.Consumer) (*Executor, error) {

	executor := Executor{producer: p, consumer: c}

	return &executor, nil
}

func (e Executor) Do() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		e.producer.Do()

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		e.consumer.Do()

	}()

	wg.Wait()

}
