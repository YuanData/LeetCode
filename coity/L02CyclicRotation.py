from typing import List


def cyclic_rotation(A: List, K: int):
    n = len(A)
    if n == 0 or K % n == 0:
        return A  # Return original array if A is empty or K is a multiple of n

    K = K % n  # If K > n, we only need to rotate K % n times

    # Rotate the list by shifting all elements to the right by K positions.
    return A[-K:] + A[:-K]


def test_cyclic_rotation():
    assert cyclic_rotation([3, 8, 9, 7, 6], 3) == [9, 7, 6, 3, 8]
    assert cyclic_rotation([0, 0, 0], 1) == [0, 0, 0]
    assert cyclic_rotation([1, 2, 3, 4], 4) == [1, 2, 3, 4]
    assert cyclic_rotation([], 1) == []
    assert cyclic_rotation([1, 2, 3, 4], 2) == [3, 4, 1, 2]
    print("All test cases pass")


if __name__ == "__main__":
    test_cyclic_rotation()
