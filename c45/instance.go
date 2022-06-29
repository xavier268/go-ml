package c45

import (
	"fmt"
	"math"
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

// Will return NaN if not set, not 0.
func (is *instance) GetVal(att int) float64 {
	if att >= len(is.data) {
		return math.NaN()
	}
	return is.data[att]
}

func (is *instance) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "<%d> (", is.class)
	for _, v := range is.data {
		fmt.Fprintf(&sb, " %3.2f;", v)
	}
	fmt.Fprint(&sb, " ),")
	return sb.String()
}

// NewInstance containg the provided values, for each attribute.
// Unknown values should be NaN, or outside the slice range.
// Once created, an instance should not be modified.
func NewInstance(class int, values []float64) Instance {
	is := new(instance)
	is.data = values
	is.class = class
	return is
}
