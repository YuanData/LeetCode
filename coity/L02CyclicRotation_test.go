package coity

import (
	"reflect"
	"testing"
)

func cyclicRotation(A []int, K int) []int {
	N := len(A)
	if N == 0 || K%N == 0 {
		return A // Return original array if A is empty or K is a multiple of N
	}

	K = K % N           // If K > N, we only need to rotate K % N times
	B := make([]int, N) // Create a new array B with the same length as A

	for i := 0; i < N; i++ {
		B[(i+K)%N] = A[i] // Calculate the new index for each element after rotation
	}

	return B
}

func TestSolution(t *testing.T) {
	tests := []struct {
		A        []int
		K        int
		expected []int
	}{
		{[]int{3, 8, 9, 7, 6}, 3, []int{9, 7, 6, 3, 8}},
		{[]int{0, 0, 0}, 1, []int{0, 0, 0}},
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		{[]int{}, 1, []int{}},
	}

	for _, test := range tests {
		result := cyclicRotation(test.A, test.K)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("cyclicRotation(%v, %v) => %v, expected %v", test.A, test.K, result, test.expected)
		}
	}
}
