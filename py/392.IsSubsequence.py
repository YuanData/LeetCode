class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        for c in s:
            i = t.find(c)
            if i == -1:
                return False
            else:
                t = t[i + 1:]
        return True


if __name__ == '__main__':
    sol = Solution()
    assert sol.isSubsequence("abc", "ahbgdc")
    assert not sol.isSubsequence("axc", "ahbgdc")
