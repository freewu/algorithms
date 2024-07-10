package main

// 2972. Count the Number of Incremovable Subarrays II
// You are given a 0-indexed array of positive integers nums.

// A subarray of nums is called incremovable if nums becomes strictly increasing on removing the subarray. 
// For example, the subarray [3, 4] is an incremovable subarray of [5, 3, 4, 6, 7] because removing this subarray changes the array [5, 3, 4, 6, 7] to [5, 6, 7] which is strictly increasing.

// Return the total number of incremovable subarrays of nums.
// Note that an empty array is considered strictly increasing.
// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: 10
// Explanation: The 10 incremovable subarrays are: [1], [2], [3], [4], [1,2], [2,3], [3,4], [1,2,3], [2,3,4], and [1,2,3,4], because on removing any one of these subarrays nums becomes strictly increasing. Note that you cannot select an empty subarray.

// Example 2:
// Input: nums = [6,5,7,8]
// Output: 7
// Explanation: The 7 incremovable subarrays are: [5], [6], [5,7], [6,5], [5,7,8], [6,5,7] and [6,5,7,8].
// It can be shown that there are only 7 incremovable subarrays in nums.

// Example 3:
// Input: nums = [8,7,6,6]
// Output: 3
// Explanation: The 3 incremovable subarrays are: [8,7,6], [7,6,6], and [8,7,6,6]. Note that [8,7] is not an incremovable subarray because after removing [8,7] nums becomes [6,6], which is sorted in ascending order but not strictly increasing.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

func incremovableSubarrayCount(nums []int) int64 {
    n, leftIncreasing, rightIncreasing := len(nums), 1, 1
    for i := 1; i < n; i++ {
        if nums[i] > nums[i - 1] {
            leftIncreasing++
        } else {
            break
        }
    }
    if leftIncreasing == n { // simplier to just return the answer when the whole nums is a strictly increasing sequence
        return int64((1 + n) * n / 2)
    }
    for i := n - 2; i >= 0; i-- {
        if nums[i + 1] > nums[i] {
            rightIncreasing++
        } else {
            break
        }
    }
    // find the patten of removing a subarray starts from i != 0 to i != n - 1
    twinIncreasing, r := 0, n - rightIncreasing
    for l := 0; l < leftIncreasing; l++ {
        for r < n && nums[l] >= nums[r] {
            r++
        }
        if r == n {
            break
        }
        twinIncreasing += (n - r)
    }
    // remember to add the case removing all numbers
    return int64(1 + leftIncreasing + rightIncreasing + twinIncreasing);
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4]
    // Output: 10
    // Explanation: The 10 incremovable subarrays are: [1], [2], [3], [4], [1,2], [2,3], [3,4], [1,2,3], [2,3,4], and [1,2,3,4], because on removing any one of these subarrays nums becomes strictly increasing. Note that you cannot select an empty subarray.
    fmt.Println(incremovableSubarrayCount([]int{1,2,3,4})) // 10
    // Example 2:
    // Input: nums = [6,5,7,8]
    // Output: 7
    // Explanation: The 7 incremovable subarrays are: [5], [6], [5,7], [6,5], [5,7,8], [6,5,7] and [6,5,7,8].
    // It can be shown that there are only 7 incremovable subarrays in nums.
    fmt.Println(incremovableSubarrayCount([]int{6,5,7,8})) // 7
    // Example 3:
    // Input: nums = [8,7,6,6]
    // Output: 3
    // Explanation: The 3 incremovable subarrays are: [8,7,6], [7,6,6], and [8,7,6,6]. Note that [8,7] is not an incremovable subarray because after removing [8,7] nums becomes [6,6], which is sorted in ascending order but not strictly increasing.
    fmt.Println(incremovableSubarrayCount([]int{8,7,6,6})) // 3
}