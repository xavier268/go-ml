package c45

import "fmt"

var ErrUnknown = fmt.Errorf("unknown value")

type instance struct {
	data map[int]float64
}

func (is *instance) Set(att int, val float64) {
	is.data[att] = val
}

func (is *instance) Get(att int) (float64, error) {
	val, ok := is.data[att]
	if ok {
		return val, nil
	} else {
		return 0, ErrUnknown
	}
}

func NewInstance() Instance {
	is := new(instance)
	is.data = make(map[int]float64)
	return is
}
