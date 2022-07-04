package ds

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// Precision when comparing float64 values.
var Precision float64 = 1e-15

type Instance struct {
	data  []float64 // using NaN for unknown numbers
	class int
}

func (is *Instance) GetClass() int {
	return is.class
}

func (is *Instance) Natt() int {
	return len(is.data)
}

// Will return NaN if not set, not 0.
func (is *Instance) GetVal(att int) float64 {
	if att >= len(is.data) {
		return math.NaN()
	}
	return is.data[att]
}

func (is *Instance) String() string {
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
func NewInstance(class int, values []float64) *Instance {
	is := new(Instance)
	is.data = values
	is.class = class
	return is
}

func NewRandomInstance(rd *rand.Rand, natt int) *Instance {

	ist := new(Instance)
	ist.data = make([]float64, natt)
	for i := 0; i < natt; i++ {
		ist.data[i] = rd.Float64() // between 0. and 1.
	}
	return ist
}

// Squarred L2 distance between is and b. Distance is zero for NaN attributes.
func (is *Instance) D2(b *Instance) float64 {
	if b == nil {
		return 0.
	}
	d2 := 0.0
	cnt := 0 // useful coordinates
	for a, v := range is.data {
		vv := b.GetVal(a)
		if math.IsNaN(v) || math.IsNaN(vv) {
			continue // ignore NnN comparisons
		}
		vv = v - vv
		d2 = d2 + vv*vv
		cnt++
	}
	if cnt == 0 {
		return 0.
	}
	return d2 / float64(cnt)
}

func (is Instance) Equal(b *Instance) bool {
	natt := b.Natt()
	if is.Natt() > natt {
		natt = is.Natt()
	}
	// now n2 is  max (n1, n2)
	for i := 0; i < natt; i++ {
		v1, v2 := is.GetVal(i), b.GetVal(i)
		if math.IsNaN(v1) && math.IsNaN(v2) {
			continue
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}

// Less is is < b. By convention, NaN < ...
func (is Instance) Less(b Instance) bool {
	natt := b.Natt()
	if is.Natt() > natt {
		natt = is.Natt()
	}
	// now n2 is  max (n1, n2)
	for i := 0; i < natt; i++ {
		v1, v2 := is.GetVal(i), b.GetVal(i)
		if math.IsNaN(v1) && math.IsNaN(v2) {
			continue
		}
		if v1 != v2 {
			return (v1 < v2)
		}
	}
	return false

}

func (is *Instance) Almost(b *Instance) bool {
	if b == nil {
		return false
	}
	natt := b.Natt()
	if is.Natt() > natt {
		natt = is.Natt()
	}
	// now n2 is  max (n1, n2)
	for i := 0; i < natt; i++ {
		v1, v2 := is.GetVal(i), b.GetVal(i)
		if math.IsNaN(v1) && math.IsNaN(v2) {
			continue
		}
		if math.Abs(v1-v2) > Precision {
			return false
		}
	}
	return true
}
