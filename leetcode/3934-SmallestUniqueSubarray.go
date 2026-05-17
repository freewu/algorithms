package main

// 3934. Smallest Unique Subarray
// You are given an integer array nums.

// Find the minimum length of a subarray that is not identical to any other subarray in nums.

// Return an integer denoting the minimum possible length of such a subarray.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Two subarrays are considered identical if they have the same length and the same elements in corresponding positions.

// Example 1:
// Input: nums = [3,3,3]
// Output: 3
// Explanation:
// Subarrays of length 1: [3] → appears 3 times
// Subarrays of length 2: [3, 3] → appears 2 times
// Subarrays of length 3: [3, 3, 3] → appears once
// The subarray [3, 3, 3] is unique, so the smallest unique subarray length is 3.

// Example 2:
// Input: nums = [2,1,2,3,3]
// Output: 1
// Explanation:
// Subarrays of length 1:
// [2] → appears 2 times
// [1] → appears once
// [3] → appears 2 times
// The subarray [1] is unique, so the smallest unique subarray length is 1.

// Example 3:
// Input: nums = [1,1,2,2,1]
// Output: 2
// Explanation:
// Subarrays of length 1:
// [1] → appears 3 times
// [2] → appears 2 times
// Subarrays of length 2:
// [1, 1] → appears once
// [1, 2] → appears once
// [2, 2] → appears once
// [2, 1] → appears once
// There is at least one subarray of length 2 that is unique, so the smallest unique subarray length is 2.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func smallestUniqueSubarray(nums []int) int {
    n := len(nums)
    check := func(l int) bool {
        counts := make(map[int]int, n - l + 1)
        unique := 0
        add := func(h int) {
            counts[h]++
            if counts[h] == 1 {
                unique++
            } else if counts[h] == 2 {
                unique--
            }
        }
        const base, mod = 131, 1 << 31 - 1
        h, pow := 0, 1
        for i := range l {
            pow = pow * base % mod
            h = (h*base + nums[i]) % mod
        }
        add(h)
        for i := l; i < n; i++ {
            h = (h * base - (nums[i-l]) * pow % mod + mod + nums[i]) % mod
            add(h)
        }
        return unique > 0
    }
    low, high := 1, n
    for low < high {
        mid := (low + high) / 2
        if check(mid) {
            high = mid
        } else {
            low = mid + 1
        }
    }
    return low
}

func main() {
    // Example 1:
    // Input: nums = [3,3,3]
    // Output: 3
    // Explanation:
    // Subarrays of length 1: [3] → appears 3 times
    // Subarrays of length 2: [3, 3] → appears 2 times
    // Subarrays of length 3: [3, 3, 3] → appears once
    // The subarray [3, 3, 3] is unique, so the smallest unique subarray length is 3.
    fmt.Println(smallestUniqueSubarray([]int{3,3,3})) // 3
    // Example 2:
    // Input: nums = [2,1,2,3,3]
    // Output: 1
    // Explanation:
    // Subarrays of length 1:
    // [2] → appears 2 times
    // [1] → appears once
    // [3] → appears 2 times
    // The subarray [1] is unique, so the smallest unique subarray length is 1.
    fmt.Println(smallestUniqueSubarray([]int{2,1,2,3,3})) // 1  
    // Example 3:
    // Input: nums = [1,1,2,2,1]
    // Output: 2
    // Explanation:
    // Subarrays of length 1:
    // [1] → appears 3 times
    // [2] → appears 2 times
    // Subarrays of length 2:
    // [1, 1] → appears once
    // [1, 2] → appears once
    // [2, 2] → appears once
    // [2, 1] → appears once
    // There is at least one subarray of length 2 that is unique, so the smallest unique subarray length is 2.
    fmt.Println(smallestUniqueSubarray([]int{1,1,2,2,1})) // 2 

    fmt.Println(smallestUniqueSubarray([]int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(smallestUniqueSubarray([]int{9,8,7,6,5,4,3,2,1})) // 1 
}

