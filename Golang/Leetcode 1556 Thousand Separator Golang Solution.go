Intuition

We want to format an integer by inserting a thousands separator (a period) every three digits from the end. If the number is less than 1000, there's no need to add a separator.

Runtime : 0ms - Memory : 1.85MB - Beats 100.00% of users with Go

Approach

Convert the integer n into a string using strconv.Itoa(n).
Check if n is less than 1000. If it is, return the original number as a string.
If n is 1000 or greater, iterate through the string representation of n in a loop.
Starting from the third-to-last character of the string, insert a period after every three characters.
Return the modified string as the formatted number.

Complexity

Time complexity: O(n), where n is the number of digits in the input integer n.
Space complexity: O(n) since we create a string to store the number's representation.

Code

func thousandSeparator(n int) string {
    numString := strconv.Itoa(n)
    if n < 1000 {
        return numString
    }
    for i:=len(numString)-3; i>0; i-=3 {
        numString = numString[:i]+"."+numString[i:]
    }
    return numString
} 
