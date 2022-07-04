package km

import (
	"fmt"
	"testing"

	"github.com/xavier268/go-ml/ds"
)

func TestCentroid1(t *testing.T) {

	d := ds.NewDataset()
	for i := 0; i < 2; i++ {
		d.AddInstance(ti2[i])
	}

	c := compute1Centroid(d, d.GetNatt())

	if c.GetVal(0) != 1.25 || c.GetVal(1) != 1.95 {
		fmt.Printf("The centroid of the full dataset %v is invalid :\n%v\n", d, c)
		t.Fatal("Invalid centroid computation")
	}
}

func TestCentroid2(t *testing.T) {

	d := ds.NewDataset()
	for i := 0; i < 2; i++ {
		d.AddInstance(ti[i])
	}
	c := compute1Centroid(d, d.GetNatt())
	should := ds.NewInstance(0, []float64{0.8})
	if !c.Almost(should) {
		fmt.Printf("The centroid of the full dataset %v is invalid :\n%v\n", d, c)
		t.Fatal("Invalid centroid computation")
	}
}

func TestCentroid3(t *testing.T) {

	d := ds.NewDataset()
	for i := 0; i < 3; i++ {
		d.AddInstance(ti[i])
	}
	c := compute1Centroid(d, d.GetNatt())
	should := ds.NewInstance(0, []float64{0.8})
	if !c.Almost(should) {
		fmt.Printf("The centroid of the full dataset %v is invalid :\n%v\n", d, c)
		t.Fatal("Invalid centroid computation")
	}
}

func TestCentroid4(t *testing.T) {

	d := ds.NewDataset()
	for i := 0; i < 4; i++ {
		d.AddInstance(ti[i])
	}
	c := compute1Centroid(d, d.GetNatt())
	should := ds.NewInstance(0, []float64{1.6, 0., 333.})
	if !c.Almost(should) {
		fmt.Printf("The centroid of the full dataset %v is invalid :\n%v\n", d, c)
		t.Fatal("Invalid centroid computation")
	}
}

func TestCentroid5(t *testing.T) {

	d := ds.NewDataset()
	for i := 0; i < 5; i++ {
		d.AddInstance(ti[i])
	}
	c := compute1Centroid(d, d.GetNatt())
	should := ds.NewInstance(0, []float64{112.2, 0., 333.})
	if !c.Almost(should) {
		fmt.Printf("The centroid of the full dataset %v is invalid :\n%v\n", d, c)
		t.Fatal("Invalid centroid computation")
	}
}

func TestKMeansVisual(t *testing.T) {
	var d *ds.Dataset

	d = ds.NewDataset()
	for _, it := range ti2 {
		d.AddInstance(it)
	}
	fmt.Println(d)
	for k := 1; k < 4; k++ {
		showCentroids(d, k)
	}

	d = ds.NewDataset()
	for _, it := range ti {
		d.AddInstance(it)
	}
	fmt.Println(d)
	for k := 1; k < 4; k++ {
		showCentroids(d, k)
	}
	showCentroids(d, 20)

}

func showCentroids(d *ds.Dataset, k int) {
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - ")
	fmt.Printf("Computing %d centroids for :\n%v\n", k, d)
	km := NewKMean(d, k)
	km.Dump(d)
}
