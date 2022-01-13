package producer

import "errors"

type Producer struct {
	strategy []interface{}
	events   []interface{}
}

func NewProducer() (*Producer, error) {
	return _, errors.New("nope")
}
