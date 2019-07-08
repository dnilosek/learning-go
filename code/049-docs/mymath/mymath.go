package mymath

// Sum over any number of ints
func Sum(x ...int) int {
	var s int
	for _, v := range x {
		s += v
	}
	return s
}
