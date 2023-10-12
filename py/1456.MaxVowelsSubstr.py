class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        vowels = set('aeiou')  # 定義母音集合
        count = sum(1 for char in s[:k] if char in vowels)  # 使用列表生成式來計算母音數量

        # 初始化最大母音數量為目前的count
        max_count = count

        # 滑動窗口
        for i in range(k, len(s)):
            # 如果左邊的字母是母音，則減少count
            if s[i - k] in vowels:
                count -= 1
            # 如果右邊的字母是母音，則增加count
            if s[i] in vowels:
                count += 1
            # 更新最大母音數量
            max_count = max(max_count, count)

        return max_count


if __name__ == '__main__':
    sol = Solution()
    assert sol.maxVowels(s="abciiidef", k=3) == 3
    assert sol.maxVowels(s="aeiou", k=2) == 2
    assert sol.maxVowels(s="leetcode", k=3) == 2
