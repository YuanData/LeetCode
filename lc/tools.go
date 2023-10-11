package lc

import "golang.org/x/exp/constraints"

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

func absIntG[T constraints.Integer](x T) T {
	if x < T(0) {
		return -x
	} else {
		return x
	}
}
