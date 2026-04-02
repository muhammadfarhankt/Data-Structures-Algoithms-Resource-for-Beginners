class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        rev = 0
        temp = x
        while temp > 0:
            rev = (temp % 10) + (rev * 10)
            temp = temp // 10
        if rev == x:
            return True
        else:
            return False