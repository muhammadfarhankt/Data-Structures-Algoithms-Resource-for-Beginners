class Solution:
    def checkPerfectNumber(self, num: int) -> bool:
        if num == 1:
            return False
        sum = 1
        for i in range(2, int(num ** 0.5) + 1):
            if num % i == 0:
                sum += i + num // i
        if sum == num:
            return True
        return False