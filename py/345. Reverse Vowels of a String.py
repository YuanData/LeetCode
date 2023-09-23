class Solution:
    def reverseVowels(self, s: str) -> str:
        lst = list(s)
        n = len(s)
        vowels = "aeiouAEIOU"
        l, r = 0, n - 1
        while l < r:
            if l < r and lst[l] not in vowels:
                l += 1
            elif r > l and lst[r] not in vowels:
                r -= 1
            else:
                lst[l], lst[r] = lst[r], lst[l]  # swapping the vowels
                l += 1
                r -= 1
        return "".join(lst)


if __name__ == '__main__':
    sol = Solution()
    assert sol.reverseVowels("hello") == "holle"
    assert sol.reverseVowels("leetcode") == "leotcede"
