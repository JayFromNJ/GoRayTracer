package mathf

import "math"

func Float64Equals(a, b float64) bool {
	threshold := 0.00001

	if math.Abs(a-b) < threshold {
		return true
	}
	return false
}

func Odd(a, b int) bool {
	if ((a + b) % 2) == 0 {
		return false
	}
	return true
}

func ArrayEquals(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if Float64Equals(a[i], b[i]) == false {
			return false
		}
	}

	return true
}
