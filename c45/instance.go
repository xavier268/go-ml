package c45

import (
	"fmt"
	"strings"
)

var ErrUnknown = fmt.Errorf("unknown value")

type instance struct {
	data  []float64 // using NaN for unknown numbers
	class int
}

func (is *instance) GetClass() int {
	return is.class
}

func (is *instance) GetVal(att int) float64 {
	return is.data[att]
}

func (is *instance) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "<%3d>\t", is.class)
	for _, v := range is.data {
		fmt.Fprintf(&sb, "%3.2f ", v)
	}
	return sb.String()
}

// NewInstance containg the provided values, for each attribute.
// Unknown values should be NaN.
// Once created, an instance should not be modified.
func NewInstance(class int, values []float64) Instance {
	is := new(instance)
	is.data = values
	is.class = class
	return is
}
