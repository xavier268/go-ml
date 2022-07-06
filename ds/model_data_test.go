package ds

import "math"

// Add examples, but do not change these !

var ti = []*Instance{
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
	{class: 200, data: []float64{1.5, math.NaN(), 10, 12, 13, 14}},
	{class: 300, data: []float64{1.5, math.NaN(), 10, 12, math.NaN(), 14}},
}
