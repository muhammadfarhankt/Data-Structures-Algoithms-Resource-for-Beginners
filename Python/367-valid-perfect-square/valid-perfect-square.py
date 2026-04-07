class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        sqrt = int(num ** 0.5)
        if sqrt * sqrt == num:
            return True
        return False