package ds

// Transform Dataset by applying a 'kernel' function to each instance.
func (ds *Dataset) Transform(kernel func(inst *Instance) *Instance) *Dataset {

	d := NewDataset()
	for _, i := range ds.GetInstances() {
		d.AddInstance(kernel(i))
	}
	return d
}

// Merge d2 into ds
func (ds *Dataset) Merge(d2 *Dataset) {
	if d2 == nil {
		return
	}
	for _, i := range d2.GetInstances() {
		ds.AddInstance(i)
	}
}

// Clone a Dataset (instances are NOT cloned, they are immutable)
func (ds *Dataset) Clone() *Dataset {
	dd := NewDataset()
	for _, i := range ds.GetInstances() {
		dd.AddInstance(i)
	}
	return dd
}
