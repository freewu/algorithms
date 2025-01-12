package main

// 2447. Number of Subarrays With GCD Equal to K
// Given an integer array nums and an integer k, 
// return the number of subarrays of nums where the greatest common divisor of the subarray's elements is k.

// A subarray is a contiguous non-empty sequence of elements within an array.

// The greatest common divisor of an array is the largest integer that evenly divides all the array elements.

// Example 1:
// Input: nums = [9,3,1,2,6,3], k = 3
// Output: 4
// Explanation: The subarrays of nums where 3 is the greatest common divisor of all the subarray's elements are:
// - [9,3,1,2,6,3]
// - [9,3,1,2,6,3]
// - [9,3,1,2,6,3]
// - [9,3,1,2,6,3]

// Example 2:
// Input: nums = [4], k = 7
// Output: 0
// Explanation: There are no subarrays of nums where 7 is the greatest common divisor of all the subarray's elements.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i], k <= 10^9

import "fmt"

func subarrayGCD(nums []int, k int) int {
    res, n := 0, len(nums)
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for i := 0; i < n; i++ {
        v := nums[i]
        for j := i; j < n; j++ {
            v = gcd(v, nums[j])
            if v == k {
                res++
            } else if v < k { // the gcd of the subarray starts from i will not get any bigger
                break // if it is smaller than k, then simply abort the search
            }
        }
    }
    return res
}

func subarrayGCD1(nums []int, k int) int {
    res, n := 0, len(nums)
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for i := 0; i < n; i++ {
        v := nums[i]
        for j := i; j < n; j++ {
            v = gcd(v, nums[j])
            if v < k  { break }
            if v == k { res++ }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [9,3,1,2,6,3], k = 3
    // Output: 4
    // Explanation: The subarrays of nums where 3 is the greatest common divisor of all the subarray's elements are:
    // - [9,3,1,2,6,3]
    // - [9,3,1,2,6,3]
    // - [9,3,1,2,6,3]
    // - [9,3,1,2,6,3]
    fmt.Println(subarrayGCD([]int{9,3,1,2,6,3}, 3)) // 4
    // Example 2:
    // Input: nums = [4], k = 7
    // Output: 0
    // Explanation: There are no subarrays of nums where 7 is the greatest common divisor of all the subarray's elements.
    fmt.Println(subarrayGCD([]int{4}, 7)) // 0

    fmt.Println(subarrayGCD([]int{1,2,3,4,5,6,7,8,9}, 4)) // 1
    fmt.Println(subarrayGCD([]int{9,8,7,6,5,4,3,2,1}, 4)) // 1

    fmt.Println(subarrayGCD1([]int{9,3,1,2,6,3}, 3)) // 4
    fmt.Println(subarrayGCD1([]int{4}, 7)) // 0
    fmt.Println(subarrayGCD1([]int{1,2,3,4,5,6,7,8,9}, 4)) // 1
    fmt.Println(subarrayGCD1([]int{9,8,7,6,5,4,3,2,1}, 4)) // 1
}