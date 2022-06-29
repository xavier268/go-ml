package c45

import (
	"fmt"
	"math"
	"math/rand"
)

// Computes the centroids of the k clusters.
func (ds *dataset) Centroids(k int) []Instance {
	epsilon := 0.00001
	centroids := make([]Instance, k)
	natt := ds.Natt()
	rd := rand.New(rand.NewSource(42)) // reproductible random generator
	// Initialize with the k first centroids in dictionnary.
	for ci := range centroids {
		centroids[ci] = NewRandomInstance(rd, natt)
	}
	changed := true
	for changed {
		changed = false
		dd := ds.Clusterize(centroids)
		for i := range centroids {
			newc := dd[i].Centroid()
			if !centroids[i].Almost(newc, epsilon) {
				changed = true
				fmt.Println("centroid changed ", newc, centroids[i])
				centroids[i] = newc
			}
		}
	}
	return centroids
}

// Clusterize splits the instances in ds in the clusters corresponding to the provided centroids.
func (ds *dataset) Clusterize(centroids []Instance) []Dataset {

	dd := make([]Dataset, len(centroids))
	for i := range dd {
		dd[i] = NewDataset()
	}

	for _, s := range ds.selection {
		var bestd float64
		var bestci int = -1
		ii := ds.data[s]
		fmt.Println("Where to put ", ii)
		for ci, cc := range centroids {
			d2 := cc.D2(ii)
			//fmt.Println("   distance from ", cc, "to  ", ii, "is ", d2)
			if bestci < 0 || d2 < bestd {
				bestci = ci
				bestd = d2
			}
		}
		fmt.Println("putting ii in cluster # ", bestci)
		dd[bestci].AddInstance(ii)
	}
	return dd
}

// Compute the centroid of the given dataset
func (ds *dataset) Centroid() Instance {

	n := 0.0
	m := make([]float64, ds.natt)
	for _, s := range ds.selection {
		for a := range m {
			if !math.IsNaN(ds.data[s].GetVal(a)) {
				m[a] = m[a] + ds.data[s].GetVal(a)
			}
			n++
		}
	}
	if n != 0 {
		for i := range m {
			m[i] = m[i] / n
		}
	}
	return NewInstance(0, m)
}
