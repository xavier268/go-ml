// Package c45 contains the c4.5 algorithm.
package c45

import (
	"fmt"
	"strings"
)

type Node struct {
	// parent *Node
	crit  int     // criteria attribute
	cut   float64 // cut-off value of the attribute
	child []*Node // up to 2 childs
	class int     // class, if leaf
	count int     // count of training examples at this node
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

func (n *Node) String() string {
	var sb strings.Builder
	n.string(&sb, "")
	return sb.String()

}

func (n *Node) string(sb *strings.Builder, pad string) {
	if len(n.child) == 0 {
		fmt.Fprintf(sb, "%s (class = %d, count=%d) \n", pad, n.class, n.count)
		return
	}

	fmt.Fprintf(sb, "%satt#%d < %f ? (count=%d)\n", pad, n.crit, n.cut, n.count)
	fmt.Fprintf(sb, "   %syes>", pad)
	n.child[0].string(sb, pad+"   ")
	fmt.Fprintf(sb, "   %sno >", pad)
	n.child[1].string(sb, pad+"   ")
}
