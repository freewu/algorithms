package main

// 2470. Number of Subarrays With LCM Equal to K
// Given an integer array nums and an integer k, 
// return the number of subarrays of nums where the least common multiple of the subarray's elements is k.

// A subarray is a contiguous non-empty sequence of elements within an array.

// The least common multiple of an array is the smallest positive integer that is divisible by all the array elements.

// Example 1:
// Input: nums = [3,6,2,7,1], k = 6
// Output: 4
// Explanation: The subarrays of nums where 6 is the least common multiple of all the subarray's elements are:
// - [3,6,2,7,1]
// - [3,6,2,7,1]
// - [3,6,2,7,1]
// - [3,6,2,7,1]

// Example 2:
// Input: nums = [3], k = 2
// Output: 0
// Explanation: There are no subarrays of nums where 2 is the least common multiple of all the subarray's elements.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i], k <= 1000

import "fmt"

func subarrayLCM(nums []int, k int) int {
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    lcm := func(x int, y int) int { return (x * y) / gcd(x, y) }
    res, n := 0, len(nums)
    for i := 0; i < n; i++ {
        if nums[i] == k { res++ }
        l := nums[i]
        for j := i + 1; j < n; j++ {
            l = lcm(l, nums[j])
            if l == k {
                res++
            } else if l > k {
                break
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,6,2,7,1], k = 6
    // Output: 4
    // Explanation: The subarrays of nums where 6 is the least common multiple of all the subarray's elements are:
    // - [(3),(6),2,7,1]
    // - [(3),(6),(2),7,1]
    // - [3,(6),2,7,1]
    // - [3,(6),(2),7,1]
    fmt.Println(subarrayLCM([]int{3,6,2,7,1}, 6)) // 4
    // Example 2:
    // Input: nums = [3], k = 2
    // Output: 0
    // Explanation: There are no subarrays of nums where 2 is the least common multiple of all the subarray's elements.
    fmt.Println(subarrayLCM([]int{3}, 2)) // 0

    fmt.Println(subarrayLCM([]int{1,2,3,4,5,6,7,8,9}, 2)) // 2
    fmt.Println(subarrayLCM([]int{9,8,7,6,5,4,3,2,1}, 2)) // 2
}