package c45

import (
	"math"
)

// Computes the centroids of the k clusters, with the provided (epsilon) precision.
func (ds *dataset) Centroids(k int, epsilon float64) []Instance {
	centroids := make([]Instance, k)
	natt := ds.Natt()
	for ci := range centroids {
		c := new(instance)
		c.data = make([]float64, natt)
		for i, d := range ds.data {
			if !math.IsNaN(d.GetVal(i)) {
				c.data[i] = d.GetVal(i)
			}
		}
		centroids[ci] = c
	}
	changed := true
	for changed {
		changed = false
		dd := ds.Clusterize(centroids)
		for i := range centroids {
			newc := dd[i].Centroid()
			if !centroids[i].Almost(newc, epsilon) {
				changed = true
				//fmt.Println("centroid changed ", newc, centroids[i])
				centroids[i] = newc
			}
		}
	}
	return centroids
}

// Clusterize splits the instances in ds into the clusters corresponding to the provided centroids.
func (ds *dataset) Clusterize(centroids []Instance) []Dataset {

	dd := make([]Dataset, len(centroids))
	for i := range dd {
		dd[i] = NewDataset()
	}

	for _, s := range ds.selection {
		var bestd float64
		var bestci int = -1
		ii := ds.data[s]
		//fmt.Println("Where to put ", ii)
		for ci, cc := range centroids {
			d2 := cc.D2(ii)
			//fmt.Println("   distance from ", cc, "to  ", ii, "is ", d2)
			if bestci < 0 || d2 < bestd {
				bestci = ci
				bestd = d2
			}
		}
		//fmt.Println("putting ii in cluster # ", bestci)
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
