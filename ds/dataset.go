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

// GetNatt provides an upper bound for the number of attributes in Instances.
// it is ok to query a non allocated attribute, but you will get a NaN.
func (ds *Dataset) GetNatt() int {
	return ds.natt
}

// AddInstance adds an instance to the dataset. Instances being immutable, only a pointer to it is used, making it space efficient.
func (ds *Dataset) AddInstance(inst *Instance) int {

	id := len(ds.data)
	ds.data = append(ds.data, inst)
	if ds.natt < inst.GetNatt() {
		ds.natt = inst.GetNatt()
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

// MeanVar returns 3 Instances that respectiveley contain, per attributes,  the mean of the defined attributes, the variance, and the count.
// Only defined attributes are taken into account, independently.
// If no attribute defined, count is 0, mean and var are NaN.
// If only one attribute value is defined, count is 1, mean is the value, var is NaN.
func (ds *Dataset) MeanVarCount() (*Instance, *Instance, *Instance) {
	ma, va, ca := make([]float64, ds.GetNatt()), make([]float64, ds.GetNatt()), make([]float64, ds.GetNatt())
	for _, in := range ds.GetInstances() {
		for a := range in.data {
			v := in.GetVal(a)
			if !math.IsNaN(v) {
				ma[a] = ma[a] + v
				va[a] = va[a] + v*v
				ca[a] = ca[a] + 1.0
			} // ignore NaNs
		}
	}

	for a := 0; a < ds.GetNatt(); a++ {
		switch ca[a] {
		case 0:
			ma[a], va[a] = math.NaN(), math.NaN()
		case 1:
			va[a] = math.NaN()
			ma[a] = ma[a] / ca[a]
		default:
			ma[a] = ma[a] / ca[a]
			va[a] = (va[a]/ca[a] - ma[a]*ma[a]) * ca[a] / (ca[a] - 1)
		}
	}
	return NewInstance(0, ma), NewInstance(0, va), NewInstance(0, ca)
}

// Normalize creates a new Dataset where instances attributes have a zero-mean and unit-variance.
func (ds *Dataset) Normalize() *Dataset {
	d := NewDataset()
	m, v, c := ds.MeanVarCount()
	for _, ins := range ds.GetInstances() {
		newi := make([]float64, ins.GetNatt())
		for i := range newi {
			n := c.GetVal(i)
			if math.IsNaN(ins.GetVal(i)) {
				newi[i] = math.NaN()
				continue
			}
			switch n {
			case 0:
				newi[i] = math.NaN()
			case 1:
				newi[i] = 0 // a single value, normalized is always 0.
			default:
				if v.GetVal(i) > Precision {
					newi[i] = (ins.GetVal(i) - m.GetVal(i)) / math.Sqrt(v.GetVal(i))
				} else { // variance almost 0, all numers almost equals ...
					newi[i] = (ins.GetVal(i) - m.GetVal(i))
				}
			}
		}
		d.AddInstance(NewInstance(0, newi))
	}
	return d
}
