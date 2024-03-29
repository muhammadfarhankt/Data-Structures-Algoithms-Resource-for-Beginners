// Solution 1: Using Array
func maximumLengthSubstring(s string) int {
    // Initialize variables
    maxLen := 0
    charArr := make([]int, 26) // Array to store character counts

    // Define sliding window using two pointers: left and right
    for left, right := 0, 0; right < len(s); right++ {
        // Increment character count for the current character
        charArr[int(s[right]-'a')]++

        // Shrink the window until all characters in the window have at most two occurrences
        for charArr[int(s[right]-'a')] > 2 {
            charArr[int(s[left]-'a')]-- // Decrement character count for the character at the left pointer
            left++                       // Move the left pointer to the right
        }

        // Update maximum length substring
        maxLen = max(maxLen, right-left+1)
    }

    return maxLen
}

// Solution 2: Using Map
func maximumLengthSubstring(s string) int {
    // Initialize variables
    maxLen, left := 0, 0
    charMap := make(map[byte]int) // Map to store character counts

    // Define sliding window using two pointers: left and right
    for right := 0; right < len(s); right++ {
        // Increment character count for the current character
        charMap[s[right]]++

        // Shrink the window until all characters in the window have at most two occurrences
        for charMap[s[right]] > 2 {
            charMap[s[left]]-- // Decrement character count for the character at the left pointer
            left++             // Move the left pointer to the right
        }

        // Update maximum length substring
        maxLen = max(maxLen, right-left+1)
    }

    return maxLen
}
