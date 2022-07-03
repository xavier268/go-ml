package algs

import (
	"fmt"
	"testing"
)

func TestCentroid1(t *testing.T) {

	d := NewDataset()
	for i := 0; i < 2; i++ {
		d.AddInstance(ti2[i])
	}
	c := d.Centroid(d.(*dataset).natt)

	if c.GetVal(0) != 1.25 || c.GetVal(1) != 1.95 {
		fmt.Printf("The centroid of the full dataset %v is invalid :\n%v\n", d, c)
		t.Fatal("Invalid centroid computation")
	}
}

func TestCentroid2(t *testing.T) {

	d := NewDataset()
	for i := 0; i < 2; i++ {
		d.AddInstance(ti[i])
	}
	c := d.Centroid(d.(*dataset).natt)
	should := NewInstance(0, []float64{0.8})
	if !c.Almost(should, 0.00000001) {
		fmt.Printf("The centroid of the full dataset %v is invalid :\n%v\n", d, c)
		t.Fatal("Invalid centroid computation")
	}
}

func TestCentroid3(t *testing.T) {

	d := NewDataset()
	for i := 0; i < 3; i++ {
		d.AddInstance(ti[i])
	}
	c := d.Centroid(d.(*dataset).natt)
	should := NewInstance(0, []float64{0.8})
	if !c.Almost(should, 0.00000001) {
		fmt.Printf("The centroid of the full dataset %v is invalid :\n%v\n", d, c)
		t.Fatal("Invalid centroid computation")
	}
}

func TestCentroid4(t *testing.T) {

	d := NewDataset()
	for i := 0; i < 4; i++ {
		d.AddInstance(ti[i])
	}
	c := d.Centroid(d.(*dataset).natt)
	should := NewInstance(0, []float64{1.6, 0., 333.})
	if !c.Almost(should, 0.00000001) {
		fmt.Printf("The centroid of the full dataset %v is invalid :\n%v\n", d, c)
		t.Fatal("Invalid centroid computation")
	}
}

func TestCentroid5(t *testing.T) {

	d := NewDataset()
	for i := 0; i < 5; i++ {
		d.AddInstance(ti[i])
	}
	c := d.Centroid(d.(*dataset).natt)
	should := NewInstance(0, []float64{112.2, 0., 333.})
	if !c.Almost(should, 0.00000001) {
		fmt.Printf("The centroid of the full dataset %v is invalid :\n%v\n", d, c)
		t.Fatal("Invalid centroid computation")
	}
}

func TestKMeansVisual(t *testing.T) {
	var d Dataset

	d = NewDataset()
	for _, it := range ti2 {
		d.AddInstance(it)
	}
	fmt.Println(d)
	for k := 1; k < 4; k++ {
		showCentroids(d, k, 1e-99)
	}

	d = NewDataset()
	for _, it := range ti {
		d.AddInstance(it)
	}
	fmt.Println(d)
	for k := 1; k < 4; k++ {
		showCentroids(d, k, 1e-99)
	}
	showCentroids(d, 20, 1e-99)

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
