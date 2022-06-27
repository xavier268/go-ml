package c45

import "fmt"

var ErrUnknown = fmt.Errorf("unknown value")

type instance struct {
	data  map[int]float64
	class int
}

func (is *instance) SetVal(att int, val float64) {
	is.data[att] = val
}
func (is *instance) SetClass(cl int) {
	is.class = cl
}

func (is *instance) GetClass() int {
	return is.class
}

func (is *instance) GetVal(att int) (float64, error) {
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
