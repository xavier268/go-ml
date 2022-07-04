//Package km contains the kmean algorithm.
package kmeans

import (
	"fmt"
	"math"

	"github.com/xavier268/go-ml/ds"
)

type KMean struct {
	k         int // numebr of centroids
	centroids []*ds.Instance
	natt      int // nuber of attributes
}

// NewKMean define k clusters on the ds Dataset with defined precision.
func NewKMean(dst *ds.Dataset, k int) *KMean {
	km := new(KMean)
	km.k, km.natt = k, dst.GetNatt()
	km.centroids = make([]*ds.Instance, k)

	// initialize the centroids with real instances
	ists := dst.GetInstances()
	for ci := 0; ci < k; ci++ { // iterate centroids
		l := len(ists)
		c := make([]float64, km.natt) // prepare a new centroid value
		for i := 0; i < km.natt; i++ {
			v := ists[ci%l].GetVal(i)
			if !math.IsNaN(v) {
				c[i] = v
				//fmt.Println("d = ", d)
				//fmt.Printf("c.data[%d] <- d.GetVal(%d) = %f\n", i, i, d.GetVal(i))
			}
		}

		km.centroids[ci] = ds.NewInstance(0, c)
		fmt.Println("Initial centroid setup :\t", km.centroids[ci]) // debug
	}

	// Iterate until centroid computation stabilizes.
	changed := true
	for changed {
		changed = false
		dd := km.partition(dst)
		for i := range km.centroids {
			newc := compute1Centroid(dd[i], km.natt) // compute a new centroid
			if !km.centroids[i].Almost(newc) {
				changed = true
				fmt.Printf("centroid #%d changed from %v to %v\n", i, km.centroids[i], newc)
				km.centroids[i] = newc
			}
		}
	}
	return km
}

// partition splits the instances in ds into the clusters corresponding to the provided centroids.
func (km *KMean) partition(dst *ds.Dataset) []*ds.Dataset {

	// handle special cases
	if dst == nil || len(km.centroids) == 0 {
		return nil // no  Dataset, beacuse no centroids
	}
	if len(km.centroids) == 1 {
		return []*ds.Dataset{dst} // single  Dataset, identical to input.
	}

	// initialize partition
	dd := make([]*ds.Dataset, len(km.centroids))
	for i := range dd {
		dd[i] = ds.NewDataset()
	}

	for _, ii := range dst.GetInstances() {
		bestci := km.GetClusterId(ii)
		dd[bestci].AddInstance(ii)
	}
	return dd
}

// Compute the centroid of the given dataset, using the provided number of attributes, natt
func compute1Centroid(dst *ds.Dataset, natt int) *ds.Instance {

	m := make([]float64, natt)
	if dst.NbInstances() == 0 {
		return ds.NewInstance(0, m)
	}
	for a := range m {
		n := 0.0
		for i := 0; i < dst.NbInstances(); i++ {
			v := dst.GetInstance(i).GetVal(a)
			if !math.IsNaN(v) {
				m[a] = m[a] + v
				n += 1.0
			}
		}
		if n != 0 {
			m[a] = m[a] / n
		}
	}
	return ds.NewInstance(0, m)
}

// GetClusterId provides the id of the cluster the inst Instance belongs to.
func (km *KMean) GetClusterId(inst *ds.Instance) int {
	var bestd float64
	var bestci int = -1

	for ci, cc := range km.centroids {
		d2 := cc.D2(inst)
		//fmt.Println("   distance from ", cc, "to  ", ii, "is ", d2)
		if bestci < 0 || d2 < bestd {
			bestci = ci
			bestd = d2
		}
	}
	return bestci
}

// GetClusterCenter provides the centroid the inst Instance is closest to.
func (km *KMean) GetClusterCenter(inst *ds.Instance) *ds.Instance {
	return km.centroids[km.GetClusterId(inst)]
}

// GetCentroids retunrs the (up to k) computed centroids.
func (km *KMean) GetCentroids() []*ds.Instance {
	return km.centroids
}
