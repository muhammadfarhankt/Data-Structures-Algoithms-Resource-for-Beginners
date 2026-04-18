class Solution:
    def climbStairs(self, n: int) -> int:
        if n == 1:
            return 1
        elif n == 2:
            return 2
        a = b = 1
        fib = 0
        for i in range(3, n+1):
            fib = a + b
            a = b
            b = fib
        return fib+a