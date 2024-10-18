package main

// 2229. Check if an Array Is Consecutive
// Given an integer array nums, return true if nums is consecutive, otherwise return false.

// An array is consecutive if it contains every number in the range [x, x + n - 1] (inclusive), 
// where x is the minimum number in the array and n is the length of the array.

// Example 1:
// Input: nums = [1,3,4,2]
// Output: true
// Explanation:
// The minimum value is 1 and the length of nums is 4.
// All of the values in the range [x, x + n - 1] = [1, 1 + 4 - 1] = [1, 4] = (1, 2, 3, 4) occur in nums.
// Therefore, nums is consecutive.

// Example 2:
// Input: nums = [1,3]
// Output: false
// Explanation:
// The minimum value is 1 and the length of nums is 2.
// The value 2 in the range [x, x + n - 1] = [1, 1 + 2 - 1], = [1, 2] = (1, 2) does not occur in nums.
// Therefore, nums is not consecutive.

// Example 3:
// Input: nums = [3,5,4]
// Output: true
// Explanation:
// The minimum value is 3 and the length of nums is 3.
// All of the values in the range [x, x + n - 1] = [3, 3 + 3 - 1] = [3, 5] = (3, 4, 5) occur in nums.
// Therefore, nums is consecutive.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^5

import "fmt"
import "sort"

func isConsecutive(nums []int) bool {
    sort.Ints(nums)
    for i := 1; i < len(nums); i++ {
        if nums[i] - nums[i - 1] != 1 { return false }
    }
    return true
}

func isConsecutive1(nums []int) bool {
    mp := make(map[int]bool)
    x, y, n := 100005, -1, len(nums)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        if mp[v] { return false }
        mp[v] = true
        x, y = min(x, v), max(y, v)
        if y - x + 1 > n { return false } 
    }
    return true
}

func main() {
    // Example 1:
    // Input: nums = [1,3,4,2]
    // Output: true
    // Explanation:
    // The minimum value is 1 and the length of nums is 4.
    // All of the values in the range [x, x + n - 1] = [1, 1 + 4 - 1] = [1, 4] = (1, 2, 3, 4) occur in nums.
    // Therefore, nums is consecutive.
    fmt.Println(isConsecutive([]int{1,3,4,2})) // true
    // Example 2:
    // Input: nums = [1,3]
    // Output: false
    // Explanation:
    // The minimum value is 1 and the length of nums is 2.
    // The value 2 in the range [x, x + n - 1] = [1, 1 + 2 - 1], = [1, 2] = (1, 2) does not occur in nums.
    // Therefore, nums is not consecutive.
    fmt.Println(isConsecutive([]int{1,3})) // false
    // Example 3:
    // Input: nums = [3,5,4]
    // Output: true
    // Explanation:
    // The minimum value is 3 and the length of nums is 3.
    // All of the values in the range [x, x + n - 1] = [3, 3 + 3 - 1] = [3, 5] = (3, 4, 5) occur in nums.
    // Therefore, nums is consecutive.
    fmt.Println(isConsecutive([]int{3,5,4})) // true

    fmt.Println(isConsecutive1([]int{1,3,4,2})) // true
    fmt.Println(isConsecutive1([]int{1,3})) // false
    fmt.Println(isConsecutive1([]int{3,5,4})) // true
}