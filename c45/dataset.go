package c45

import (
	"fmt"
	"math"
	"strings"
)

type dataset struct {
	data      []Instance // Instances are shared as much as possible between dataset. Instances are immutable.
	selection []int      // which instances are part of the subset, possibly duplicated
	natt      int
}

func (ds *dataset) GetNatt() int {
	return ds.natt
}

func (ds *dataset) AddInstance(inst Instance) int {
	id := len(ds.data)
	ds.selection = append(ds.selection, id)
	ds.data = append(ds.data, inst)
	if ds.natt < inst.Natt() {
		ds.natt = inst.Natt()
	}
	return id
}

func (ds *dataset) Natt() int {
	return ds.natt
}

func (ds *dataset) DuplicateInstance(id ...int) {
	ds.selection = append(ds.selection, id...)
}

func (ds *dataset) GetInstance(i int) Instance {
	return (ds.data)[i]
}

func (ds *dataset) Split(f SplitFunc) (Dataset, Dataset) {
	s1, s2 := make([]int, 0, 2+len(ds.selection)/2), make([]int, 0, 2+len(ds.selection)/2)
	for _, i := range ds.selection {
		if f(ds.data[i]) {
			s1 = append(s1, i)
		} else {
			s2 = append(s2, i)
		}
	}
	d1 := ds.Subset(s1)
	d2 := ds.Subset(s2)
	return d1, d2
}

func (ds *dataset) Subset(idx []int) Dataset {
	dd := new(dataset)
	dd.data = ds.data // shared instance
	dd.selection = idx
	return dd
}

func NewDataset() Dataset {
	ds := new(dataset)
	return ds
}

func (ds *dataset) String() string {
	var sb strings.Builder
	//fmt.Fprintf(&sb, "Dataset contains a selection of %d instances out of %d : %v\n", len(ds.selection), len(ds.data), ds.selection)
	fmt.Fprintf(&sb, "Dataset contains a selection of %d instances among %d unique instances (entropy : %f)\n", len(ds.selection), len(ds.data), ds.Entropy())
	for _, i := range ds.selection {
		fmt.Fprintf(&sb, "#%d:\t%s\n", i, ds.data[i])
	}
	return sb.String()
}

func (ds *dataset) Entropy() (ent float64) {
	n, m := ds.countClasses()
	fn := float64(n)

	for _, c := range m {
		x := float64(c) / fn
		ent += x * math.Log2(x)
	}
	return -ent
}

func (ds *dataset) Dump(msg ...interface{}) {
	fmt.Println("------------------------------------")
	fmt.Print("Dump: ")
	fmt.Println(msg...)
	fmt.Println("------------------------------------")
	fmt.Println(ds)
	n, det := ds.countClasses()
	fmt.Printf("data : \n%v\nselection :\n%v\nclass repartition :\n%v (total : %d) \n", ds.data, ds.selection, det, n)
	fmt.Println("------------------------------------")
}

// Retun the total selected instances and a map from class -> nb of instances in class (including duplicates)
func (ds *dataset) countClasses() (ttl int, detail map[int]int) {

	m := make(map[int]int)
	for _, s := range ds.selection {
		c := ds.data[s].GetClass()
		m[c] = m[c] + 1
	}
	return len(ds.selection), m
}
