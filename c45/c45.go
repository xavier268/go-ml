package c45

import (
	"math"
	"sort"

	"github.com/xavier268/go-c45/ds"
)

// NewC45 creates and train a new, unpruned, c4.5 tree trained with the provided dataset.
func NewC45(dst *ds.Dataset) *Node {
	tree := new(Node)
	grow(dst, tree, 1e-15)
	return tree
}

// grow the tree downward, using provided dataset.
func grow(dst *ds.Dataset, tree *Node, epsilon float64) {
	tree.count = dst.NbInstances()
	if tree.count == 0 {
		panic("internal error - unexpected empty subset")
	}
	if tree.count == 1 || math.Abs(dst.Entropy()) <= epsilon {
		tree.class = dst.GetClass(0)
		return // already a leaf
	}

	att, ctf := getBestCriteria(dst)
	tree.crit = att
	tree.cut = ctf
	if att < 0 || ctf == math.Inf(-1) {
		tree.class = dst.GetClass(0)
		return // can't do better !
	}

	d1, d2 := dst.Split(func(inst *ds.Instance) bool {
		return inst.GetVal(att) < ctf
	})

	n1, n2 := new(Node), new(Node)
	tree.child = []*Node{n1, n2}
	grow(d1, n1, epsilon)
	grow(d2, n2, epsilon)
}

// find best criteria and cutoff for given dataset
func getBestCriteria(dst *ds.Dataset) (attribute int, cutoff float64) {
	nat := dst.GetNatt()
	ins := dst.GetInstances()

	// select attribute
	besta := -1                   // best attribute so far
	bestctf := math.Inf(-1)       // best ever cutoff
	bestgainratio := math.Inf(-1) // used to select best attribute

	// try each attribute ...
	for a := 0; a < nat; a++ {

		// select cutoff for the given attribute
		bestctfa := math.Inf(-1)  // relative best cutoff value for a given attribute
		bestgaina := math.Inf(-1) // best gain reached with this combination of attribute/cutoff.
		bestsplitratioa := 1.     // used intermediate value

		// get  list of cutoff  values for this attribute
		m := make(map[float64]bool, len(ins))
		for _, it := range ins {
			v := it.GetVal(a)
			if !math.IsNaN(v) { // only remember actual values
				m[it.GetVal(a)] = true
			}
		}
		av := make([]float64, 0, len(ins))
		for k := range m {
			av = append(av, k)
		}
		//fmt.Println("Attr : ", a, av)
		av = cutoffs(av)
		//fmt.Println("   ---> Possible cutofs : ", av) // debug

		// retain best cutoff value, based on gain only
		for _, c := range av {
			g, s := gainSplit(dst, a, c)
			if g > bestgaina {
				bestgaina = g
				bestctfa = c
				bestsplitratioa = s
			}
		}

		// use selected cutoff to evaluate the gain ratio score
		gr := bestgaina / bestsplitratioa
		if gr > bestgainratio {
			bestctf = bestctfa
			bestgainratio = gr
			besta = a
		}

	}
	return besta, bestctf
}

// gainSplit and splitratio, used selecting a cutoff and an attribute.
// higher gainSplit and loqer splits are better.
func gainSplit(dst *ds.Dataset, att int, ctf float64) (g, s float64) {

	d1, d2 := dst.Split(func(inst *ds.Instance) bool {
		return inst.GetVal(att) < ctf
	})

	n1, n2, n := float64(d1.NbInstances()), float64(d2.NbInstances()), float64(dst.NbInstances())

	g = dst.Entropy() - (d1.Entropy()*n1+d2.Entropy()*n2)/n        // relative gain
	s = 1e-30 + (n1*math.Log2(n1)+n2*math.Log2(n2))/n*math.Log2(n) // split ratio
	return g, s
}

// compute a list of cutoff values from a non sorted, but depuplicated slice of values, with  NaN removed.
func cutoffs(ctf []float64) []float64 {
	if len(ctf) <= 1 { // If one or less value, no critera will work !
		return []float64{}
	}
	sort.Float64Slice(ctf).Sort()
	for i := 0; i < len(ctf)-1; i++ {
		ctf[i] = (ctf[i] + ctf[i+1]) / 2
	}
	return ctf[:len(ctf)-1]
}

// Classify an unknown instance using the trained tree.
func (tree *Node) Classify(inst *ds.Instance) (class int) {

	if tree.IsLeaf() {
		return tree.GetClass()
	}
	a, c := tree.crit, tree.cut
	if inst.GetVal(a) < c {
		return tree.child[0].Classify(inst)
	}
	return tree.child[1].Classify(inst)
}
