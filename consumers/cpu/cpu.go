package cpu

import (
	"fmt"
)

type Consumer interface {
	Do() error
}

type CPU struct {
	data chan interface{}
	ctrl chan interface{}
}

func NewCPU(data chan interface{}, ctrl chan interface{}) (*CPU, error) {
	return &CPU{data: data, ctrl: ctrl}, nil
}

func (c CPU) Do() error {
	for {
		select {
		case random := <-c.data:
			fmt.Println(random)
		case _ = <-c.ctrl:
			break
		}
	}
}
