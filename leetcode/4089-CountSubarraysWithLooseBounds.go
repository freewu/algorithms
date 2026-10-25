package main

// 4089. Count Subarrays With Loose Bounds
// You are given an integer array nums and two integers minK and maxK.

// A loose-bound subarray is a subarray in which every element is greater than or equal to minK and less than or equal to maxK.

// Return the number of loose-bound subarrays in nums.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,4,7,3], minK = 1, maxK = 5
// Output: 4
// Explanation:
// The loose-bound subarrays are [1], [4], [3], and [1, 4]. 
// The minimum value in these subarrays is at least minK = 1 and the maximum value is at most maxK = 5.

// Example 2:
// Input: nums = [1,1], minK = 1, maxK = 1
// Output: 3
// Explanation:
// The loose-bound subarrays are [1], [1], and [1, 1]. 
// The minimum value in these subarrays is at least minK = 1 and the maximum value is at most maxK = 1.

// Constraints:
//     - 2 <= nums.length <= 10^{5}
//     - 1 <= nums[i], minK, maxK <= 10^{6}

import "fmt"

func countSubarrays(nums []int, minK, maxK int) int64 {
    res, end := int64(0), -1
    for i, x := range nums {
        if x < minK || x > maxK {
            end = i
        } else {
            res += int64(i - end)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,4,7,3], minK = 1, maxK = 5
    // Output: 4
    // Explanation:
    // The loose-bound subarrays are [1], [4], [3], and [1, 4]. 
    // The minimum value in these subarrays is at least minK = 1 and the maximum value is at most maxK = 5.
    fmt.Println(countSubarrays([]int{1,4,7,3}, 1, 5)) // 4
    // Example 2:
    // Input: nums = [1,1], minK = 1, maxK = 1
    // Output: 3
    // Explanation:
    // The loose-bound subarrays are [1], [1], and [1, 1]. 
    // The minimum value in these subarrays is at least minK = 1 and the maximum value is at most maxK = 1.
    fmt.Println(countSubarrays([]int{1,1}, 1, 1)) // 3

    fmt.Println(countSubarrays([]int{1,2,3,4,5,6,7,8,9}, 1, 1)) // 3
    fmt.Println(countSubarrays([]int{9,8,7,6,5,4,3,2,1}, 1, 1)) // 3
}