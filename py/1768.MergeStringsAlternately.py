class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        len1 = len(word1)
        len2 = len(word2)

        output = []
        i = 0
        while i < len1 or i < len2:
            if i < len1:
                output.append(word1[i])
            if i < len2:
                output.append(word2[i])
            i += 1
        return ''.join(output)


if __name__ == '__main__':
    sol = Solution()
    assert sol.mergeAlternately("abc", "pqr") == "apbqcr"
    assert sol.mergeAlternately("ab", "pqrs") == "apbqrs"
    assert sol.mergeAlternately("abcd", "pq") == "apbqcd"
