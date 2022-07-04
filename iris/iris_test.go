package iris

import (
	"fmt"
	"testing"

	"github.com/xavier268/go-ml/c45"
	"github.com/xavier268/go-ml/km"
)

func TestVisualIrisC45(t *testing.T) {

	iris, cc := NewIrisDataset()
	//iris.Dump("Iris dataset content")
	fmt.Println(cc)

	train, test := iris.SampleSplit(0.25) // 25% training, 75% testing.
	t45 := c45.NewC45(train)

	fmt.Println(t45)

	var good, bad float64

	for _, i := range test.GetInstances() {
		cl := t45.Classify(i)
		if cl == i.GetClass() {
			good++
		} else {
			bad++
		}
	}
	fmt.Printf("Training set        : %d / %d\t%2.1f%%\n", train.NbInstances(), iris.NbInstances(), 100.*float64(train.NbInstances())/float64(iris.NbInstances()))
	fmt.Printf("Testing set         : %d / %d\t%2.1f%%\n", test.NbInstances(), iris.NbInstances(), 100.*float64(test.NbInstances())/float64(iris.NbInstances()))
	fmt.Printf("Good classification : %.0f \t%2.1f%%\n", good, 100.*good/(good+bad))
	fmt.Printf("Bad classification  : %.0f \t%2.1f%%\n", bad, 100.*bad/(good+bad))

}

func TestVisualIrisKMeans(t *testing.T) {

	iris, cc := NewIrisDataset()
	//iris.Dump("Iris dataset content")
	fmt.Println(cc)

	train, _ := iris.SampleSplit(0.5) // 25% training, 75% testing.

	kk := km.NewKMean(train, 10)
	kk.Dump(train)

}
