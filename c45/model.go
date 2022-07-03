package c45

import "fmt"

type SplitFunc func(inst Instance) bool

type Dataset interface {
	GetInstance(i int) Instance           // using absolute instance id
	AddInstance(inst Instance) int        // return the absolute instance id
	DuplicateInstance(id ...int)          // duplicates existing instances using their absolute ids. Only impacts selection.
	Split(f SplitFunc) (Dataset, Dataset) // split dataset based upon the function result. The fisrt for true, the second for false.
	Entropy() float64                     // (binary) entropy provides the minimum quantity of information in bits to select a given class.
	Subset([]int) Dataset                 // select using the absolute instance Ids. Duplicates are allowed.
	GetNatt() int                         // number of attributes (upper bound, could be lower)
	fmt.Stringer
	Clusterer
}

// Instances are immutable, once created.
type Instance interface {
	GetVal(att int) float64 // Get attribute, NaN if unknown
	GetClass() int          // Retrieve class, if available. Default to 0.
	Natt() int              // max number of attribute (could be lower)
	D2(b Instance) float64  // L2 distance, per useful coordinate (to take into account NaN)
	Equal(b Instance) bool
	Almost(b Instance, epsilon float64) bool
	Less(b Instance) bool
	fmt.Stringer
}

type Node interface {

	// Navigating finished tree
	IsLeaf() bool
	GetClass() int             // only meaningful for leafs
	GetCriteria() (att int)    // what criteria do we manage ?
	Select(value float64) Node // Select child node corresponding to the value for the given criteria
}

type Clusterer interface {
	Centroid(natt int) Instance                  // find centroid of a Dataset, with the specified number of attributes
	Centroids(k int, epsilon float64) []Instance // find k centroids, given an epsilon precision
	Clusterize(centroids []Instance) []Dataset   // split by cluster, using provided centroids
}
