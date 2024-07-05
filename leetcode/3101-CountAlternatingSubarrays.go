package main

// 3101. Count Alternating Subarrays
// You are given a binary array nums.
// We call a subarray alternating if no two adjacent elements in the subarray have the same value.
// Return the number of alternating subarrays in nums.

// Example 1:
// Input: nums = [0,1,1,1]
// Output: 5
// Explanation:
// The following subarrays are alternating: [0], [1], [1], [1], and [0,1].

// Example 2:
// Input: nums = [1,0,1,0]
// Output: 10
// Explanation:
// Every subarray of the array is alternating. There are 10 possible subarrays that we can choose.

// Constraints:
//     1 <= nums.length <= 10^5
//     nums[i] is either 0 or 1.

import "fmt"

func countAlternatingSubarrays(nums []int) int64 {
    res, i, j, n := 0, 0, 0, len(nums) - 1
    for j <= n {
        for j < n && nums[j] != nums[j+1] {
            j++
        }
        l := j - i + 1
        res += l * (l + 1) / 2
        i = j
        i++
        j++
    }
    return int64(res)
}

func countAlternatingSubarrays1(nums []int) int64 {
    res, cur, pre := 0, 0, -1
    for _, v := range nums {
        if pre != v {
            cur++
        } else {
            cur = 1
        }
        pre = v
        res += cur
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [0,1,1,1]
    // Output: 5
    // Explanation:
    // The following subarrays are alternating: [0], [1], [1], [1], and [0,1].
    fmt.Println(countAlternatingSubarrays([]int{0,1,1,1})) // 5
    // Example 2:
    // Input: nums = [1,0,1,0]
    // Output: 10
    // Explanation:
    // Every subarray of the array is alternating. There are 10 possible subarrays that we can choose.
    fmt.Println(countAlternatingSubarrays([]int{1,0,1,0})) // 10

    fmt.Println(countAlternatingSubarrays1([]int{0,1,1,1})) // 5
    fmt.Println(countAlternatingSubarrays1([]int{1,0,1,0})) // 10
}