package c45

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/xavier268/go-c45/ds"
)

var t1 = []*ds.Instance{
	ds.NewInstance(1, []float64{1, 3, 8}),
	ds.NewInstance(1, []float64{5, 3, 17}),
	ds.NewInstance(1, []float64{9, 4, 17}),
	ds.NewInstance(2, []float64{3, -3, 10}),
	ds.NewInstance(2, []float64{2, 0, 1}),
	ds.NewInstance(2, []float64{2, 50, 1}),
	ds.NewInstance(3, []float64{2, 3, 5}),
}

func TestC45Visual(t *testing.T) {
	d := ds.NewDataset()
	for _, i := range t1 {
		d.AddInstance(i)
	}
	fmt.Println(d)
	t45 := NewC45(d)
	fmt.Println(t45)
}

func TestGoSyntaxVisual(_ *testing.T) {
	a := []float64{1, math.NaN(), -3, 55, math.NaN(), -18, math.Inf(+1), math.Inf(-1), 9}
	fmt.Println("Unsorted : ", a)
	sort.Float64Slice(a).Sort()
	fmt.Println("Sorted   : ", a)
}
