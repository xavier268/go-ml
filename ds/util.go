package ds

import "math/rand"

// Majority selects the majority vote, from the vote values.
func Majority(vote []int) []int {
	if len(vote) == 0 {
		return []int{}
	}
	if len(vote) == 1 {
		return []int{vote[0]}
	}
	// vote counter
	cnt := make(map[int]int)
	max := 0
	for _, v := range vote {
		vv := cnt[v] + 1
		if vv > max {
			max = vv
		}
		cnt[v] = vv
	}
	var res []int
	for v, c := range cnt {
		if c == max {
			res = append(res, v)
		}
	}
	return res
}

// ChooseOneInt selects a value among the provided values, with a uniform distribution.
func ChooseOneInt(values ...int) int {
	if len(values) == 0 {
		panic("choice is impossible, no value provided")
	}
	return values[rand.Intn(len(values))]
}
