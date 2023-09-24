from typing import List


class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        curr_sum = 0
        max_sum = float('-inf')
        for i, num in enumerate(nums):
            curr_sum += num
            if i > k - 1:
                curr_sum -= nums[i - k]
            if i >= k - 1:
                max_sum = max(max_sum, curr_sum)

        return max_sum / k


if __name__ == '__main__':
    sol = Solution()
    assert sol.findMaxAverage(nums=[1, 12, -5, -6, 50, 3], k=4) == 12.75
    assert sol.findMaxAverage(nums=[5], k=1) == 5.0
    assert sol.findMaxAverage(nums=[-1], k=1) == -1
