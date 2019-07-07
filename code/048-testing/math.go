package davemath

func Average(x ...float64) float64 {
	var sum float64
	if len(x) == 0 {
		return sum
	}

	for _, v := range x {
		sum += v
	}
	return sum / float64(len(x))
}
