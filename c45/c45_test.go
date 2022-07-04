package c45

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/xavier268/go-ml/ds"
)

var t1 = []*ds.Instance{
	ds.NewInstance(11, []float64{1, 3, 8}),
	ds.NewInstance(11, []float64{5, 3, 17}),
	ds.NewInstance(11, []float64{9, 4, 17}),
	ds.NewInstance(22, []float64{3, -3, 10}),
	ds.NewInstance(22, []float64{2, 0, 1}),
	ds.NewInstance(22, []float64{2, 50, 1}),
	ds.NewInstance(33, []float64{2, 3, 5}),
}

func TestC45Visual(t *testing.T) {
	d := ds.NewDataset()
	for _, i := range t1 {
		d.AddInstance(i)
	}
	fmt.Println(d)
	t45 := NewC45(d)
	fmt.Println(t45)

	// Verify classification ?
	for n, i := range t1 {
		cl := t45.Classify(i)
		if cl != i.GetClass() {
			t.Fatalf("failed classifiction #%d for %v -> %d", n, i, cl)
		}
	}
}

func TestCutoff(t *testing.T) {

	data := [][]float64{ // [test#][testdata] - before/after
		{5, 7, 3}, {4, 6},
		{5, -7, 3}, {-2, 4},
		{5, math.Inf(-1), 3}, {math.Inf(-1), 4},
		{5, 7}, {6},
		{5}, {},
		{}, {},
	}

	for i := 0; i+1 < len(data); i += 2 {
		from, should := data[i], data[i+1]
		got := cutoffs(from)
		if !sort.SliceIsSorted(got, func(i, j int) bool {
			return got[i] < got[j]
		}) {
			t.Fatalf("cutofs are not sorted : %v", got)
		}
		if !equal(got, should) {
			t.Fatalf("Unexpected cutof #%d got : %f, should %v", i/2, got, should)
		}
	}
}

func equal(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > 0.00001 {
			return false
		}
	}
	return true
}

func TestGoSyntaxVisual(_ *testing.T) {
	a := []float64{1, math.NaN(), -3, 55, math.NaN(), -18, math.Inf(+1), math.Inf(-1), 9}
	fmt.Println("Unsorted : ", a)
	sort.Float64Slice(a).Sort()
	fmt.Println("Sorted   : ", a)
}
