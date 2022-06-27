package c45

import "math"

type dataset struct {
	data []Instance
}

func (ds *dataset) AddInstance(inst Instance) {
	ds.data = append(ds.data, inst)
}

func (ds *dataset) GetInstance(i int) Instance {
	return ds.data[i]
}

func (ds *dataset) Split(f SplitFunc) (Dataset, Dataset) {
	d1, d2 := NewDataset(), NewDataset()
	for _, d := range ds.data {
		if f(d) {
			d1.AddInstance(d)
		} else {
			d2.AddInstance(d)
		}
	}
	return d1, d2
}

func NewDataset() Dataset {
	ds := new(dataset)
	return ds
}

func (ds *dataset) Entropy() (ent float64) {
	// TODO sum of p(x)Log(x) for all classes ?
	n := float64(len(ds.data)) // ttl occurences
	m := make(map[int]int)     // count occurences per class
	for _, d := range ds.data {
		c := d.GetClass()
		m[c] = m[c] + 1
	}

	for _, c := range m {
		x := float64(c) / n
		ent += x * math.Log2(x)
	}
	return -ent

}
