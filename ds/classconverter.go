package ds

import "fmt"

// ClassConverter maps between class names (string) and class ids (int).
type ClassConverter struct {
	m map[string]int
	a []string
}

// NewClassConverter creates a ClassConverter ready to use.
func NewClassConverter() *ClassConverter {
	cc := new(ClassConverter)
	cc.m = make(map[string]int, 4)
	cc.a = make([]string, 0, 4)
	return cc
}

// ToId gets the id of a class name, and register the name with a new ID if class is unknown yet.
func (cc *ClassConverter) ToId(clname string) int {
	if id, ok := cc.m[clname]; ok {
		return id
	}
	id := len(cc.m)
	cc.m[clname] = id
	cc.a = append(cc.a, clname)
	return id
}

// ToString returns the string name, or "unknownClass" if unknown.
func (cc *ClassConverter) ToString(id int) string {
	if id < len(cc.a) {
		return cc.a[id]
	}
	return "unknownClass"
}
// String gets the class names in a human readable format.
func (cc *ClassConverter) String() string {
	return fmt.Sprintf("%d classes : %q", len(cc.a), cc.a)
}
