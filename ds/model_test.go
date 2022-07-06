package ds

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"golang.org/x/exp/slices"
)

func TestModelVisual(t *testing.T) {

	i := NewInstance(11, []float64{5.2, 4.2, 3.2, 2.2, 1.2, math.NaN(), -22.2, -33.3})
	fmt.Println(i)

	d := NewDataset()
	fmt.Println(d)
	for _, it := range ti {
		fmt.Println("Added instance # ", d.AddInstance(it), ", entropy : ", d.Entropy())
	}
	fmt.Println(d)

	d1, d2 := d.Split(func(ist *Instance) bool { return ist.GetClass() <= 22 })
	fmt.Println("Splitting on class <= 22 :\n", d1, d2)

}

func TestMajority(t *testing.T) {
	data := [][]int{
		{5, 4, 6, 9, 7, 3, 2, 5, 2, 1, 6, 6, 4, 6}, {6},
		{5}, {5},
		{2, 3}, {3, 2},
		{}, {},
	}

	for i := 0; i+1 < len(data); i += 2 {
		got := Majority(data[i])
		sort.IntSlice(got).Sort()
		fmt.Println(got)
		want := data[i+1]
		sort.IntSlice(want).Sort()
		if !slices.Equal(want, got) {
			t.Fatalf("error : wanted %v, but got %v", want, got)
		}
	}
}

func TestChooseOneInt(t *testing.T) {

	data := []int{2, 4, 6, 12, 30, 40, 8, 2, 4, 4, 2}
	for i := 0; i < 30; i++ {
		v := ChooseOneInt(data...)
		if v%2 != 0 || v <= 0 || v > 40 {
			t.Fatalf("expecting an even value")
		}
	}
	v := ChooseOneInt(4)
	if v != 4 {
		t.Fatal("erreur choosing with only one option")
	}
}

func TestNormalizeVisual(t *testing.T) {

	d := NewDataset()
	for _, it := range ti {
		d.AddInstance(it)
	}
	fmt.Println("Raw :\n", d)
	m, v, c := d.MeanVarCount()
	fmt.Println("Mean  : ", m)
	fmt.Println("Var   : ", v)
	fmt.Println("Count : ", c)

	dn := d.Normalize()
	fmt.Println("Normalized : \n", dn)
	nm, nv, nc := dn.MeanVarCount()
	fmt.Println("Mean  : ", nm)
	fmt.Println("Var   : ", nv)
	fmt.Println("Count : ", nc)

	for a := 0; a < dn.GetNatt(); a++ {
		v1, v2 := nc.GetVal(a), c.GetVal(a)
		if !math.IsNaN(v1) && math.Abs(v1-v2) > Precision {
			t.Fatal("Failed normalization count", a, v1, v2)
		}
		v1, v2 = nm.GetVal(a), 0.
		if !math.IsNaN(v1) && math.Abs(v1-v2) > Precision {
			t.Fatal("Failed normalization mean")
		}
		v1, v2 = nv.GetVal(a), 1.0
		if !math.IsNaN(v1) && math.Abs(v1) > Precision && math.Abs(v1-v2) > Precision { // var, if it exists, should be 0 or 1
			t.Fatal("Failed normalization var")
		}
	}
}
