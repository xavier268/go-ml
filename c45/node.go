// Package c45 contains the c4.5 algorithm.
package c45

type Node struct {
	parent *Node
	crit   int     // criteria
	cut    float64 // cut-off
	child  []*Node // up to 2 childs
	class  int     // class, if leaf
}

func (n *Node) IsLeaf() bool {
	return len(n.child) == 0
}

func (n *Node) GetClass() int {
	return n.class
}

func (n *Node) GetCriteria() int {
	return n.crit
}

func (n *Node) Select(v float64) *Node {
	if v <= n.cut {
		return n.child[0]
	} else {
		return n.child[1]
	}
}

func NewNode(parent *Node) *Node {
	n := new(Node)
	n.child = make([]*Node, 0, 2) // anticipate 2 childs max
	n.parent = parent
	return n
}
