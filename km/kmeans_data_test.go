package kmeans

import (
	"math"

	"github.com/xavier268/go-ml/ds"
)

// Add examples, but do not change these !

var ti = []*ds.Instance{
	ds.NewInstance(00, []float64{0.6}),
	ds.NewInstance(11, []float64{1.}),
	ds.NewInstance(22, []float64{}),
	ds.NewInstance(33, []float64{3.2, math.NaN(), 333.}),
	ds.NewInstance(44, []float64{444}),
	ds.NewInstance(55, []float64{math.NaN(), 5.55}),
	ds.NewInstance(66, []float64{math.NaN(), 0.66}),
	ds.NewInstance(77, []float64{math.NaN(), 7.77}),
	ds.NewInstance(88, []float64{math.NaN(), 8.88}),
	ds.NewInstance(99, []float64{math.NaN(), 9., .99}),
	ds.NewInstance(100, []float64{1.5, math.NaN(), 10}),
}

var ti2 = []*ds.Instance{
	ds.NewInstance(0, []float64{1.2, 2.0}),
	ds.NewInstance(1, []float64{1.3, 1.9}),
	ds.NewInstance(2, []float64{0.9, 2.1}),

	ds.NewInstance(3, []float64{12, 2.0}),
	ds.NewInstance(4, []float64{13, 5.0}),
	ds.NewInstance(5, []float64{11, 2.0}),
	ds.NewInstance(6, []float64{1.0, 8.0}),
}
