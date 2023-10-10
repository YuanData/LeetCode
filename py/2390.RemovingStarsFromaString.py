class Solution:
    def removeStars(self, s: str) -> str:
        stack = []
        for char in s:  # 遍歷字串s
            if char == '*':  # 如果字符是星號
                if stack:  # 且堆疊不為空
                    stack.pop()  # 就彈出堆疊的頂部元素
            else:  # 如果字符不是星號，就加入堆疊
                stack.append(char)
        return ''.join(stack)  # 將堆疊中的元素組成答案字串


if __name__ == '__main__':
    sol = Solution()
    assert sol.removeStars("leet**cod*e") == "lecoe"
    assert sol.removeStars("erase*****") == ""
