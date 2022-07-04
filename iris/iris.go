package iris

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/xavier268/go-ml/ds"
)

// NewIrisDataset create a Dataset containing the full iris dataset.
func NewIrisDataset() (*ds.Dataset, *ds.ClassConverter) {

	f, err := os.Open("iris.txt")
	if err != nil {
		panic("could not open the iris.txt file")
	}
	defer f.Close()

	/*
		type iris struct {
			SepalLength float64
			SepalWidth  float64
			Petallength float64
			PetalWidth  float64
			ClassString string
		}
	*/

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	d := ds.NewDataset()
	cc := ds.NewClassConverter()

	for _, l := range lines {
		data := make([]float64, 4)
		for i := 0; i < 4; i++ {
			data[i], err = strconv.ParseFloat(l[i], 64)
			if err != nil {
				panic(err)
			}
			cl := cc.ToId(l[4])
			d.AddInstance(ds.NewInstance(cl, data))
		}
	}

	return d, cc

}
