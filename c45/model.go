package c45

type Dataset interface {
	GetInstance(i int) Instance
	AddInstance(inst Instance)
}

type Instance interface {
	Get(att int) (float64, error) // Get attribute, error if unknown
	Set(att int, val float64)
}

type Node interface {

	// Navigating finished tree
	IsLeaf() bool
	GetClass() int             // only meaningful for leafs
	GetCriteria() (att int)    // what criteria do we manage ?
	Select(value float64) Node // Select child node corresponding to the value for the given criteria

}
