class Solution:
    def reverseVowels(self, s: str) -> str:
        lst = list(s)
        n = len(s)
        vowels = list("aeiouAEIOU")
        l, r = 0, n - 1
        while l < r:
            while l < r and lst[l] not in vowels:
                l += 1
            while r > l and lst[r] not in vowels:
                r -= 1
            lst[l], lst[r] = lst[r], lst[l]  # swapping the vowels

            l += 1
            r -= 1
        return "".join(lst)


if __name__ == '__main__':
    sol = Solution()
    assert sol.reverseVowels("hello") == "holle"
    assert sol.reverseVowels("leetcode") == "leotcede"
