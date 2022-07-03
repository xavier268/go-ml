package c45

import (
	"math"

	"github.com/xavier268/go-c45/ds"
)

// NewC45 creates a new, unpruned, c4.5 tree trained with the provided dataset.
func NewC45(dst *ds.Dataset) *Node {
	tree := new(Node)
	grow(dst, tree)
	return tree
}

// grow the tree downward, using provided dataset.
func grow(dst *ds.Dataset, tree *Node) {
	if dst.NbInstances() <= 1 || dst.Entropy() == 0 {
		tree.class = dst.GetClass(0)
		return // already a leaf
	}

	att, ctf := getBestCriteria(dst)
	if att < 0 || ctf == math.Inf(-1) {
		return // can't do better !
	}

	d1, d2 := dst.Split(func(inst *ds.Instance) bool {
		return inst.GetVal(att) < ctf
	})
	tree.crit = att
	tree.cut = ctf

	n1, n2 := new(Node), new(Node)
	tree.child = []*Node{n1, n2}
	grow(d1, n1)
	grow(d2, n2)
}

// find best criteria and cutoff for given dataset
func getBestCriteria(dst *ds.Dataset) (attribute int, cutoff float64) {
	nat := dst.GetNatt()
	ins := dst.GetInstances()

	besta := -1               // best attribute so far
	bestctf := math.Inf(-1)   // best cutoff value so far
	bestscore := math.Inf(-1) // best score reached with this combination of attribute/cutoff.

	// try each attribute ...
	for a := 0; a < nat; a++ {

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
		av = cutoffs(av)

		// retain best cutoff value
		for _, c := range av {
			sc := score(dst, a, c)
			if sc > bestscore {
				bestscore = sc
				bestctf = c
				besta = a
			}
		}
	}
	return besta, bestctf
}

// score if selecting a cutoff and an attribute.
// higher scores are better.
func score(dst *ds.Dataset, att int, ctf float64) float64 {

	d1, d2 := dst.Split(func(inst *ds.Instance) bool {
		return inst.GetVal(att) < ctf
	})

	sc := dst.Entropy() - (d1.Entropy()*float64(d1.NbInstances())+d2.Entropy()*float64(d2.NbInstances()))/float64(dst.NbInstances()) // relative gain
	return sc
}

// compute a list of cutoff values from a non sorted, depuplicated slice of values, with non NaN value left.
func cutoffs(ctf []float64) []float64 {
	if len(ctf) <= 1 { // If one or less vale, no critera will work !
		return []float64{}
	}
	for i := 0; i < len(ctf)-1; i++ {
		ctf[i] = (ctf[i] + ctf[i+1]) / 2
	}
	return ctf[1:]
}
