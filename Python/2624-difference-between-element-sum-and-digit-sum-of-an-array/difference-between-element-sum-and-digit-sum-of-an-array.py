class Solution:
    def differenceOfSum(self, nums: List[int]) -> int:
        element_sum = digit_sum = 0
        for num in nums:
            element_sum += num
            while num > 0:
                digit_sum += (num % 10)
                num = num // 10
        if element_sum > digit_sum:
            return element_sum - digit_sum
        else:
            return digit_sum - element_sum