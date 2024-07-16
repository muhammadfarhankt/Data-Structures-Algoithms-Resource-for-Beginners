func lengthOfLongestSubstring(s string) int {
    maxLen := 0
    if len(s) == 0 {
        return 0
    } else if len(s) == 1 {
        return 1
    }
    count := make([]int, 256)
    for right, left := 0, 0; right < len(s); right++ {
        count[s[right]]++
        for count[s[right]] > 1 {
            count[s[left]]--
            left++
        }
        if right-left+1 > maxLen {
            maxLen = right - left + 1
        }
    }
    return maxLen
}