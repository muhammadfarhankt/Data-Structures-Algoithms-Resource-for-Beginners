func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    mid := left + (right - left) / 2
    for left <= right {
        mid = (left + right) / 2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
