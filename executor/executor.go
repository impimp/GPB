package executor

import "errors"

type Executor struct {
	producers []interface{}
	consumers []interface{}
}

func NewExecutor() (*Executor, error) {
	return _, errors.New("nope")
}
