// Package ds contains the data structures, common to all algorithms.
package ds

import (
	"fmt"
	"math"
	"strings"
)

type Dataset struct {
	data []*Instance // Instances are shared as much as possible between dataset. Instances are immutable.
	natt int
}

func (ds *Dataset) GetNatt() int {
	return ds.natt
}

func (ds *Dataset) AddInstance(inst *Instance) int {

	id := len(ds.data)
	ds.data = append(ds.data, inst)
	if ds.natt < inst.Natt() {
		ds.natt = inst.Natt()
	}
	return id
}

func (ds *Dataset) Natt() int {
	return ds.natt
}

func (ds *Dataset) GetInstance(i int) *Instance {
	if i < 0 || i >= len(ds.data) {
		return nil
	}
	return (ds.data)[i]
}

type SplitFunc func(*Instance) bool

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

func NewDataset() *Dataset {
	ds := new(Dataset)
	return ds
}

func (ds *Dataset) String() string {
	var sb strings.Builder
	//fmt.Fprintf(&sb, "Dataset contains a selection of %d instances out of %d : %v\n", len(ds.selection), len(ds.data), ds.selection)
	fmt.Fprintf(&sb, "Dataset contains %d  instances (entropy : %f)\n", len(ds.data), ds.Entropy())
	for i, inst := range ds.data {
		fmt.Fprintf(&sb, "#%d:\t%s\n", i, inst)
	}
	return sb.String()
}

func (ds *Dataset) Entropy() (ent float64) {
	n, m := ds.countClasses()
	fn := float64(n)

	for _, c := range m {
		x := float64(c) / fn
		ent += x * math.Log2(x)
	}
	return -ent
}

func (ds *Dataset) Dump(msg ...any) {
	fmt.Println("------------------------------------")
	fmt.Print("Dump: ")
	fmt.Println(msg...)
	fmt.Println("------------------------------------")
	fmt.Println(ds)
	n, det := ds.countClasses()
	fmt.Printf("data : \n%v\nclass repartition :\n%v (total : %d) \n", ds.data, det, n)
	fmt.Println("------------------------------------")
}

// Retun the total selected instances and a map from class -> nb of instances in class (including duplicates)
func (ds *Dataset) countClasses() (ttl int, detail map[int]int) {

	m := make(map[int]int)
	for _, s := range ds.data {
		c := s.GetClass()
		m[c] = m[c] + 1
	}
	return len(ds.data), m
}

func (ds *Dataset) GetInstances() []*Instance {
	return ds.data
}

func (ds *Dataset) NbInstances() int {
	return len(ds.data)
}
