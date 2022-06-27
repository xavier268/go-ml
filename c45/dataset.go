package c45

type dataset struct {
	data []Instance
}

func (ds *dataset) AddInstance(inst Instance) {
	ds.data = append(ds.data, inst)
}

func (ds *dataset) GetInstance(i int) Instance {
	return ds.data[i]
}

func NewDataset() Dataset {
	ds := new(dataset)
	return ds
}
