package lc

import "golang.org/x/exp/constraints"

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}
