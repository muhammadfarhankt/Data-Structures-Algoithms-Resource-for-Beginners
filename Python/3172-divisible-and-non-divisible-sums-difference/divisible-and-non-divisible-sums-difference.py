class Solution:
    def differenceOfSums(self, n: int, m: int) -> int:
        div_sum = non_div_sum = 0
        for i in range(1, n+1):
            if i % m == 0:
                div_sum += i
            else:
                non_div_sum += i
        return non_div_sum - div_sum