from typing import List


class Solution:
    def maxArea(self, height: List[int]) -> int:
        maxArea = 0
        left, right = 0, len(height) - 1
        while left < right:
            currentArea = (right - left) * min(height[left], height[right])
            maxArea = max(currentArea, maxArea)
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1
        return maxArea


if __name__ == '__main__':
    sol = Solution()
    assert sol.maxArea(height=[1, 8, 6, 2, 5, 4, 8, 3, 7]) == 49
    assert sol.maxArea(height=[1, 1]) == 1
