package main

// 3347. Maximum Frequency of an Element After Performing Operations II
// You are given an integer array nums and two integers k and numOperations.

// You must perform an operation numOperations times on nums, where in each operation you:
//     Select an index i that was not selected in any previous operations.
//     Add an integer in the range [-k, k] to nums[i].

// Return the maximum possible frequency of any element in nums after performing the operations.

// Example 1:
// Input: nums = [1,4,5], k = 1, numOperations = 2
// Output: 2
// Explanation:
// We can achieve a maximum frequency of two by:
//     Adding 0 to nums[1], after which nums becomes [1, 4, 5].
//     Adding -1 to nums[2], after which nums becomes [1, 4, 4].

// Example 2:
// Input: nums = [5,11,20,20], k = 5, numOperations = 1
// Output: 2
// Explanation:
// We can achieve a maximum frequency of two by:
//     Adding 0 to nums[1].

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     0 <= k <= 10^9
//     0 <= numOperations <= nums.length

import "fmt"
import "sort"

func maxFrequency(nums []int, k int, numOperations int) int {
    sort.Ints(nums)
    res, n := 0, len(nums)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, l, r := 0, 0, 0; i < n; {
        j := i
        for l < j && nums[j]-nums[l] > k  { l++ }
        r = max(r, i)
        for r < n && nums[r]-nums[j] <= k { r++ }
        for i < n && nums[i] == nums[j] { i++ }
        res = max(res, i - j + min(j - l + r - i, numOperations))
    }
    for l, r := 0, 0; r < n; r++ { // 所有的数，都变化的情况
        for l < r && nums[r]-nums[l] > 2 * k { l++ }
        res = max(res, min(numOperations, r - l + 1))
    }
    return res
}

func maxFrequency1(nums []int, k, numOperations int) int {
    sort.Ints(nums)
    res, count, left, right, left2, n  := 0, 0, 0, 0, 0, len(nums)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range nums {
        for nums[left2] < v - k * 2 { left2++ }
        res = max(res, min(i - left2 + 1, numOperations))
        count++
        if i < n - 1 && v == nums[i + 1] { continue }// 循环直到连续相同段的末尾，这样可以统计出 v的出现次数
        for nums[left] < v - k { left++ }
        for right < n && nums[right] <= v + k { right++ }
        res, count = max(res, min(right - left, count + numOperations)), 0
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,4,5], k = 1, numOperations = 2
    // Output: 2
    // Explanation:
    // We can achieve a maximum frequency of two by:
    //     Adding 0 to nums[1]. nums becomes [1, 4, 5].
    //     Adding -1 to nums[2]. nums becomes [1, 4, 4].
    fmt.Println(maxFrequency([]int{1,4,5}, 1, 2)) // 2
    // Example 2: 
    // Input: nums = [5,11,20,20], k = 5, numOperations = 1
    // Output: 2
    // Explanation:
    // We can achieve a maximum frequency of two by:
    //     Adding 0 to nums[1].
    fmt.Println(maxFrequency([]int{5,11,20,20}, 5, 1)) // 2

    fmt.Println(maxFrequency([]int{1,2,3,4,5,6,7,8,9}, 1, 2)) // 3
    fmt.Println(maxFrequency([]int{9,8,7,6,5,4,3,2,1}, 1, 2)) // 3

    fmt.Println(maxFrequency1([]int{1,4,5}, 1, 2)) // 2
    fmt.Println(maxFrequency1([]int{5,11,20,20}, 5, 1)) // 2
    fmt.Println(maxFrequency1([]int{1,2,3,4,5,6,7,8,9}, 1, 2)) // 3
    fmt.Println(maxFrequency1([]int{9,8,7,6,5,4,3,2,1}, 1, 2)) // 3
}