// Package ds contains the data structures, common to all algorithms.
package ds

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// Dataset is the main objet to store the data instances, with the class attached if relevant.
type Dataset struct {
	data []*Instance // Instances are shared as much as possible between dataset. Instances are immutable.
	natt int
}

// GetNatt provides an uppe bound for the number of attributes in Instances.
// it is ok to query a non allocated attribute, but you will get a NaN.
func (ds *Dataset) GetNatt() int {
	return ds.natt
}

// AddInstance adds an instance to the dataset. Instances being immutable, only a pointer to it is used, making it space efficient.
func (ds *Dataset) AddInstance(inst *Instance) int {

	id := len(ds.data)
	ds.data = append(ds.data, inst)
	if ds.natt < inst.Natt() {
		ds.natt = inst.Natt()
	}
	return id
}

// GetInstance gets a pointer to the i-th instance.
func (ds *Dataset) GetInstance(i int) *Instance {
	if i < 0 || i >= len(ds.data) {
		return nil
	}
	return (ds.data)[i]
}

// GetClass returns the class of the i-th instance, or 0 if not set.
func (ds *Dataset) GetClass(i int) int {
	if i < 0 || i >= len(ds.data) {
		return 0
	}
	return ds.data[i].class
}

// SplitFunc is used to separate a Dataset in 2 Datasets.
// The first one returned will contain the 'true' values of the function.
type SplitFunc func(*Instance) bool

// Split the Dataset in 2 Datasets, using the provided SplitFunc.
func (ds *Dataset) Split(f SplitFunc) (*Dataset, *Dataset) {
	d1, d2 := NewDataset(), NewDataset()
	for _, inst := range ds.data {
		if f(inst) {
			d1.AddInstance(inst)
		} else {
			d2.AddInstance(inst)
		}
	}
	return d1, d2
}

// NewDataset creates a new Dataset ready to use, but with no instances yet.
// Use AddInstance to add them.
func NewDataset() *Dataset {
	ds := new(Dataset)
	return ds
}

// String provides a human readable format for the Dataset.
func (ds *Dataset) String() string {
	var sb strings.Builder
	//fmt.Fprintf(&sb, "Dataset contains a selection of %d instances out of %d : %v\n", len(ds.selection), len(ds.data), ds.selection)
	fmt.Fprintf(&sb, "Dataset contains %d  instances (entropy : %f)\n", len(ds.data), ds.Entropy())
	for i, inst := range ds.data {
		fmt.Fprintf(&sb, "#%d:\t%s\n", i, inst)
	}
	return sb.String()
}

// Entropy defines the quantity of information, in bits, attached with the classes of the instances in the Dataset.
func (ds *Dataset) Entropy() (ent float64) {
	n, m := ds.CountClasses()
	fn := float64(n)

	for _, c := range m {
		x := float64(c) / fn
		ent += x * math.Log2(x)
	}
	return -ent
}

// Dump will diplay the provided messages as title, and then print detailed information about the Dataset.
// Used for debugging.
func (ds *Dataset) Dump(msg ...any) {
	fmt.Println("------------------------------------")
	fmt.Print("Dump: ")
	fmt.Println(msg...)
	fmt.Println("------------------------------------")
	fmt.Println(ds)
	n, det := ds.CountClasses()
	fmt.Printf("data : \n%v\nclass repartition :\n%v (total : %d) \n", ds.data, det, n)
	fmt.Println("------------------------------------")
}

// CountClasses return the total selected instances and a map from class -> nb of instances in class (including duplicates)
func (ds *Dataset) CountClasses() (ttl int, detail map[int]int) {

	m := make(map[int]int)
	for _, s := range ds.data {
		c := s.GetClass()
		m[c] = m[c] + 1
	}
	return len(ds.data), m
}

// GetInstances returns an array with all the instances from the Dataset.
func (ds *Dataset) GetInstances() []*Instance {
	return ds.data
}

// NbInstances is the number of Instance(s) loaded.
func (ds *Dataset) NbInstances() int {
	return len(ds.data)
}

// SampleSplit splits the dataset in 2 parts, with the provided percentage in the first dataset (approx).
// If percent is 1.0, all samples will be in the first dataset, if 0.0, non.
func (ds *Dataset) SampleSplit(percent float64) (*Dataset, *Dataset) {
	rd := rand.New(rand.NewSource(time.Now().Unix()))

	d1, d2 := ds.Split(func(*Instance) bool {
		return rd.Float64() < percent
	})
	return d1, d2
}
