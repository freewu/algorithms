package  main

// 2962. Count Subarrays Where Max Element Appears at Least K Times
// You are given an integer array nums and a positive integer k.
// Return the number of subarrays where the maximum element of nums appears at least k times in that subarray.
// A subarray is a contiguous sequence of elements within an array.

// Example 1:
// Input: nums = [1,3,2,3,3], k = 2
// Output: 6
// Explanation: The subarrays that contain the element 3 at least 2 times are: [1,3,2,3], [1,3,2,3,3], [3,2,3], [3,2,3,3], [2,3,3] and [3,3].

// Example 2:
// Input: nums = [1,4,2,1], k = 3
// Output: 0
// Explanation: No subarray contains the element 4 at least 3 times.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^6
//     1 <= k <= 10^5

import "fmt"
import "slices"

func countSubarrays(nums []int, k int) int64 {
    res := 0
    count, left, mx := 0, 0, slices.Max(nums)
    for _, x := range nums {
        // 统计最大值出现次数
        if x == mx {
            count++
        }
        // 如果此时 count = k，则不断右移左指针 left，直到窗口内的 mx 的出现次数小于 k为止。
        // 此时，对于右端点为 right 且左端点小于 left 的子数组, mx 的出现次数都至少为 k，把答案增加 left。
        for count == k {
            if nums[left] == mx {
                count--
            }
            left++
        }
        res += left
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,3,2,3,3], k = 2
    // Output: 6
    // Explanation: The subarrays that contain the element 3 at least 2 times are: [1,3,2,3], [1,3,2,3,3], [3,2,3], [3,2,3,3], [2,3,3] and [3,3].
    fmt.Println(countSubarrays([]int{1,3,2,3,3}, 2)) // 6
    // Example 2:
    // Input: nums = [1,4,2,1], k = 3
    // Output: 0
    // Explanation: No subarray contains the element 4 at least 3 times.
    fmt.Println(countSubarrays([]int{1,4,2,1}, 3)) // 0

    fmt.Println(countSubarrays([]int{1,2,3,4,5,6,7,8,9}, 1)) // 9
    fmt.Println(countSubarrays([]int{9,8,7,6,5,4,3,2,1}, 1)) // 9
}