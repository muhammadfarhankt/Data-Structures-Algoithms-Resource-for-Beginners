class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        size = len(nums)
        total_sum = size * (size + 1) // 2
        nums_sum = 0
        for num in nums:
            nums_sum += num
        return total_sum - nums_sum