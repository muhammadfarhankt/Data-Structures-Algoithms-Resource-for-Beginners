class Solution:
    def commonFactors(self, a: int, b: int) -> int:
        minimum = min(a, b)
        factors = 0
        for i in range(1, minimum+1):
            if a % i == 0 and b % i == 0:
                factors += 1
        return factors