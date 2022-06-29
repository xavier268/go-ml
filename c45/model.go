package c45

import "fmt"

type SplitFunc func(inst Instance) bool

type Dataset interface {
	GetInstance(i int) Instance    // using absolute instance id
	AddInstance(inst Instance) int // return the absolute instance id
	DuplicateInstance(id ...int)   // duplicates existing instances using their absolute ids. Only impacts selection.
	Split(f SplitFunc) (Dataset, Dataset)
	Entropy() float64
	Subset([]int) Dataset // select using the absolute instance Ids. Duplicates are allowed.
	fmt.Stringer
}

// Instances are immutable, once created.
type Instance interface {
	GetVal(att int) float64 // Get attribute, NaN if unknown
	GetClass() int
	fmt.Stringer
}

type Node interface {

	// Navigating finished tree
	IsLeaf() bool
	GetClass() int             // only meaningful for leafs
	GetCriteria() (att int)    // what criteria do we manage ?
	Select(value float64) Node // Select child node corresponding to the value for the given criteria
}
