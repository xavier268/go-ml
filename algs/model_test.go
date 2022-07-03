package algs

import (
	"fmt"
	"math"
	"testing"
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

	fmt.Println("Duplicating instances 2 & 1")
	d.DuplicateInstance(2, 1)
	fmt.Println(d)
	d.(*dataset).Dump("with duplicated instances")

	s := d.Subset([]int{2, 1, 2})
	fmt.Println("Subset #2, #1, #2:\n", s)

	d1, d2 := d.Split(func(ist Instance) bool { return ist.GetClass() <= 22 })
	fmt.Println("Splitting on class <= 22 :\n", d1, d2)

}