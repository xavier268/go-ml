package c45

import "math"

// Add examples, but do not change these !

var ti = []*instance{
	{class: 00, data: []float64{0.6}},
	{class: 11, data: []float64{1.}},
	{class: 22, data: []float64{}},
	{class: 33, data: []float64{3.2, math.NaN(), 333.}},
	{class: 44, data: []float64{444}},
	{class: 55, data: []float64{math.NaN(), 5.55}},
	{class: 66, data: []float64{math.NaN(), 0.66}},
	{class: 77, data: []float64{math.NaN(), 7.77}},
	{class: 88, data: []float64{math.NaN(), 8.88}},
	{class: 99, data: []float64{math.NaN(), 9., .99}},
	{class: 100, data: []float64{1.5, math.NaN(), 10}},
}

var ti2 = []*instance{
	{class: 0, data: []float64{1.2, 2.0}},
	{class: 1, data: []float64{1.3, 1.9}},
	{class: 2, data: []float64{0.9, 2.1}},

	{class: 3, data: []float64{12, 2.0}},
	{class: 4, data: []float64{13, 5.0}},
	{class: 5, data: []float64{11, 2.0}},
	{class: 6, data: []float64{1.0, 8.0}},
}
