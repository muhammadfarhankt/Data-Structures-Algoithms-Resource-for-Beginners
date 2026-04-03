class Solution:
    def subtractProductAndSum(self, n: int) -> int:
        sum = 0
        product = 1
        while n > 0:
            quot = n % 10
            sum += quot
            product *= quot
            n = n // 10
        return product - sum