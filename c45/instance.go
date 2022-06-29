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

func (is *instance) NAtt() int {
	return len(is.data)
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

// L2 distance between is and b. Distance is zero for NaN attributes.
func (is *instance) D2(b Instance) float64 {
	if b == nil {
		return 0.
	}
	d2 := 0.0
	for a, v := range is.data {
		if math.IsNaN(v) {
			continue
		}
		vv := b.GetVal(a)
		if !math.IsNaN(vv) {
			continue
		}
		vv = v - vv
		d2 = d2 + vv*vv
	}
	return d2
}
