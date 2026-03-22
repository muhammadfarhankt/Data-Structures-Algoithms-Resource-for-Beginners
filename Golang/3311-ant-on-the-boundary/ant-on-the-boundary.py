class Solution:
    def returnToBoundaryCount(self, nums: List[int]) -> int:
        res = x = 0
        for v in nums:
            x += v
            res += x == 0
        return res