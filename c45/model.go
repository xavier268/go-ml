package c45

type SplitFunc func(inst Instance) bool
type Dataset interface {
	GetInstance(i int) Instance
	AddInstance(inst Instance)
	Split(f SplitFunc) (Dataset, Dataset)
	Entropy() float64
}

type Instance interface {
	GetVal(att int) (float64, error) // Get attribute, error if unknown
	GetClass() int
	SetVal(att int, val float64)
	SetClass(int)
}

type Node interface {

	// Navigating finished tree
	IsLeaf() bool
	GetClass() int             // only meaningful for leafs
	GetCriteria() (att int)    // what criteria do we manage ?
	Select(value float64) Node // Select child node corresponding to the value for the given criteria
}
