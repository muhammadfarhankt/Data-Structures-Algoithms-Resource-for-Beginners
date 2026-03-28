class Solution:
    def numberGame(self, nums: List[int]) -> List[int]:
        nums.sort()
        size = len(nums)
        for i in range(0, size, 2):
            temp = nums[i]
            nums[i] = nums[i+1]
            nums[i+1] = temp
        return nums