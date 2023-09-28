from typing import List


class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        count_dic = {}
        for num in arr:
            count_dic[num] = count_dic.get(num, 0) + 1

        count_set = set(count_dic.values())
        return len(count_set) == len(count_dic.values())


if __name__ == '__main__':
    sol = Solution()
    assert sol.uniqueOccurrences([1, 2, 2, 1, 1, 3]) == True
    assert sol.uniqueOccurrences([1, 2]) == False
    assert sol.uniqueOccurrences([-3, 0, 1, -3, 1, 1, 1, -3, 10, 0]) == True
