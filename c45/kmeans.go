package c45

import (
	"fmt"
	"math"
)

// Computes the centroids of the k clusters, with the provided (epsilon) precision.
func (ds *dataset) Centroids(k int, epsilon float64) []Instance {
	centroids := make([]Instance, k)
	natt := ds.Natt()
	for ci := 0; ci < k; ci++ {
		l := len(ds.data)
		c := make([]float64, natt)
		for i := 0; i < natt; i++ {
			v := ds.data[ci%l].GetVal(i)
			if !math.IsNaN(v) {
				c[i] = v
				//fmt.Println("d = ", d)
				//fmt.Printf("c.data[%d] <- d.GetVal(%d) = %f\n", i, i, d.GetVal(i))
			}
		}

		centroids[ci] = NewInstance(0, c)
		fmt.Println("Initial centroid setup :\t", centroids[ci])
	}
	/*rd := rand.New(rand.NewSource(42))
	for ci := range centroids {
		centroids[ci] = NewRandomInstance(rd, natt)
		fmt.Println("Initial centroid setup :\n", centroids[ci])
	}
	*/
	changed := true
	for changed {
		changed = false
		dd := ds.Clusterize(centroids)
		for i := range centroids {
			newc := dd[i].Centroid(natt)
			if !centroids[i].Almost(newc, epsilon) {
				changed = true
				fmt.Printf("centroid #%d changed from %v to %v\n", i, newc, centroids[i])
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

// Compute the centroid of the given dataset, using the provided number of attributes, natt
func (ds *dataset) Centroid(natt int) Instance {

	m := make([]float64, natt)
	if len(ds.selection) == 0 {
		return NewInstance(0, m)
	}
	for a := range m {
		n := 0.0
		for _, s := range ds.selection {
			if !math.IsNaN(ds.data[s].GetVal(a)) {
				m[a] = m[a] + ds.data[s].GetVal(a)
				n += 1.0
			}
		}
		if n != 0 {
			m[a] = m[a] / n
		}
	}

	return NewInstance(0, m)
}
