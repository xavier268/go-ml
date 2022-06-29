package c45

import "math"

// Kmeans computes the centroids of the k clusters.
func (ds *dataset) Kmeans(k int) []Instance {
	panic("to do)")
}

// Clusterize splits the instances in ds in the clusters corresponding to the provided centroids.
func (ds *dataset) Clusterize(centroids []Instance) []Dataset {
	panic("to do)")
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
