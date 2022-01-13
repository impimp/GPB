package producer

import (
	"math/rand"
	"time"
)

type Producer struct {
	data chan interface{}
	ctrl chan interface{}

	strategy []interface{}
	events   []interface{}
}

func NewProducer(data chan interface{}, ctrl chan interface{}) (*Producer, error) {
	return &Producer{data: data, ctrl: ctrl}, nil
}

func (p Producer) Do() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		select {
		case p.data <- random.Intn(1024):
		case _ = <-p.ctrl:
			break
		}
	}

}
