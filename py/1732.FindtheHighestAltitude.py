from typing import List


class Solution:
    def largestAltitude(self, gain: List[int]) -> int:
        current_altitude, max_altitude = 0, 0  # 初始化最高海拔和当前海拔
        for g in gain:
            current_altitude += g  # 使用 Prefix Sum 更新当前海拔
            max_altitude = max(max_altitude, current_altitude)  # 更新最高海拔
        return max_altitude


if __name__ == '__main__':
    sol = Solution()
    assert sol.largestAltitude([-5, 1, 5, 0, -7]) == 1
    assert sol.largestAltitude([-4, -3, -2, -1, 4, 3, 2]) == 0
