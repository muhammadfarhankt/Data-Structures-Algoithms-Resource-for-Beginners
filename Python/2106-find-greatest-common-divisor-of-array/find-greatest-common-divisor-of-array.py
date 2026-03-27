class Solution:
    def findGCD(self, nums: List[int]) -> int:
        large = small = nums[0]
        gcd = 1
        for num in nums:
            if num > large:
                large = num
            if num < small:
                small = num
        for i in range(small, 1, -1):
            if small % i == 0 and large % i == 0:
                gcd = i
                break
        return gcd