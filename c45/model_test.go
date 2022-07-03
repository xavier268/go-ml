package c45

import (
	"fmt"
	"math"
	"testing"
)

var ti = []*instance{
	{class: 0, data: []float64{0.01}},
	{class: 11, data: []float64{1.}},
	{class: 22, data: []float64{}},
	{class: 33, data: []float64{3.3, math.NaN(), 333.}},
	{class: 44, data: []float64{444}},
	{class: 55, data: []float64{math.NaN(), 5.55}},
	{class: 66, data: []float64{math.NaN(), 0.66}},
	{class: 77, data: []float64{math.NaN(), 7.77}},
	{class: 88, data: []float64{math.NaN(), 8.88}},
	{class: 99, data: []float64{math.NaN(), 9., .99}},
	{class: 100, data: []float64{1.5, math.NaN(), 10}},
}

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

func TestKMeans(t *testing.T) {
	d := NewDataset()
	fmt.Println(d)
	for _, it := range ti {
		fmt.Println("Added instance # ", d.AddInstance(it), ", entropy : ", d.Entropy())
	}
	fmt.Println(d)

	for k := 1; k < 4; k++ {
		showCentroids(d, k, 1e-99)
	}

	//showCentroids(d, 20, 0.0001)

}

func showCentroids(d Dataset, k int, epsilon float64) {
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - ")
	fmt.Printf("Computing %d centroids :\n", k)
	res := d.Centroids(k, epsilon)
	cc := d.Clusterize(res)
	for i, c := range res {
		fmt.Printf("Centroid # %d:\t%v\n", i, c)
		fmt.Println(cc[i])
	}
}
