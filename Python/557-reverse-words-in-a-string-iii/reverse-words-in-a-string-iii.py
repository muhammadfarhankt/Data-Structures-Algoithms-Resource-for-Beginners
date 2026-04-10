class Solution:
    def reverseWords(self, s: str) -> str:
        start = end = 0
        s = list(s)
        length = len(s)
        for index, char in enumerate(s):
            if char == " " or index == length - 1:
                end = index - 1 if char == " " else index
                for i in range(start, (start + end + 1) // 2):
                    temp = s[i]
                    s[i] = s[start + end - i]
                    s[start + end - i] = temp
                start = index + 1
        return "".join(s)