package lc

func twoSum(nums []int, target int) []int {
	rst := []int{}
	numMap := map[int]int{}

	for i, ele := range nums {
		if numMap[target-ele] != 0 {
			rst = append(rst, i)
			rst = append(rst, numMap[target-ele]-1)
			return rst
		}
		numMap[ele] = i + 1
	}
	return rst
}
