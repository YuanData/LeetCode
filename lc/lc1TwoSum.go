package lc

func twoSum(nums []int, target int) []int {
	rst := []int{}
	numMap := map[int]int{}

	for i, ele := range nums {
		if numMap[target-ele] != 0 {
			// if htere esists 2 ele which sum is equal target
			// then add idx to result
			rst = append(rst, i)
			rst = append(rst, numMap[target-ele]-1)
			return rst
		}
		numMap[ele] = i + 1
	}
	return rst
}
