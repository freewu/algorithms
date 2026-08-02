package main

// 4010. Maximize Pair Strength Using GCD
// You are given an integer array nums.

// Choose exactly one pair of distinct indices i and j. 
// The strength of the pair is defined as (nums[i] * nums[j]) / gcd(nums[i], nums[j])2.

// Return the maximum strength over all possible pairs.

// The term gcd(a, b) denotes the greatest common divisor of a and b.

// Example 1:
// Input: nums = [2,3,5]
// Output: 15
// Explanation:
// Choosing i = 1 and j = 2 gives strength (3 * 5) / gcd(3, 5)2 = 15 / 1 = 15, which is the maximum over all pairs.

// Example 2:
// Input: nums = [4,6,8]
// Output: 12
// Explanation:
// Choosing i = 1 and j = 2 gives strength (6 * 8) / gcd(6, 8)2 = 48 / 4 = 12, which is the maximum over all pairs.

// Example 3:
// Input: nums = [3,3]
// Output: 1
// Explanation:
// Choosing i = 0 and j = 1 gives strength (3 * 3) / gcd(3, 3)2 = 9 / 9 = 1, the maximum over all pairs.

// Constraints:
//     2 <= nums.length <= 2000
//     1 <= nums[i] <= 10^5

import "fmt"

func maxPairStrength(nums []int) int64 {
    res := int64(0)
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for i := 0; i < len(nums)-1; i++ {
        x := nums[i]
        for j := i + 1; j < len(nums); j++ {
            y := nums[j]
            z := (x * y) / (gcd(x, y) * gcd(x, y))
            if int64(z) > res {
                res = int64(z)
            }
        }
    }
    return res
}

func maxPairStrength1(nums []int) int64 {
    res, n := int64(0), len(nums)
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            v := int64(gcd(nums[i], nums[j]))
            res = max(res, (int64(nums[i]) * int64(nums[j])) / (v * v))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,3,5]
    // Output: 15
    // Explanation:
    // Choosing i = 1 and j = 2 gives strength (3 * 5) / gcd(3, 5)2 = 15 / 1 = 15, which is the maximum over all pairs.
    fmt.Println(maxPairStrength([]int{2,3,5})) // 15
    // Example 2:
    // Input: nums = [4,6,8]
    // Output: 12
    // Explanation:
    // Choosing i = 1 and j = 2 gives strength (6 * 8) / gcd(6, 8)2 = 48 / 4 = 12, which is the maximum over all pairs.
    fmt.Println(maxPairStrength([]int{4,6,8})) // 12
    // Example 3:
    // Input: nums = [3,3]
    // Output: 1
    // Explanation:
    // Choosing i = 0 and j = 1 gives strength (3 * 3) / gcd(3, 3)2 = 9 / 9 = 1, the maximum over all pairs.
    fmt.Println(maxPairStrength([]int{3,3})) // 1
  
    fmt.Println(maxPairStrength([]int{1,2,3,4,5,6,7,8,9})) // 72
    fmt.Println(maxPairStrength([]int{9,8,7,6,5,4,3,2,1})) // 72

    fmt.Println(maxPairStrength1([]int{2,3,5})) // 15
    fmt.Println(maxPairStrength1([]int{4,6,8})) // 12
    fmt.Println(maxPairStrength1([]int{3,3})) // 1
    fmt.Println(maxPairStrength1([]int{1,2,3,4,5,6,7,8,9})) // 72
    fmt.Println(maxPairStrength1([]int{9,8,7,6,5,4,3,2,1})) // 72
}