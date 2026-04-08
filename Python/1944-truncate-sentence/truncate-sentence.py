class Solution:
    def truncateSentence(self, s: str, k: int) -> str:
        count = 0
        for i, char in enumerate(s):
            if char == ' ':
                count += 1
            if count == k:
                return s[:i]
        return s