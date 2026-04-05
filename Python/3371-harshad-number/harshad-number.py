class Solution:
    def sumOfTheDigitsOfHarshadNumber(self, x: int) -> int:
        digit_sum = 0
        temp = x
        while temp > 0:
            digit_sum += (temp % 10)
            temp = temp // 10
        if x % digit_sum ==  0:
            return digit_sum
        return -1