package main

// 3644. Maximum K to Sort a Permutation
// You are given an integer array nums of length n, 
// where nums is a permutation of the numbers in the range [0..n - 1].

// You may swap elements at indices i and j only if nums[i] AND nums[j] == k, 
// where AND denotes the bitwise AND operation and k is a non-negative integer.

// Return the maximum value of k such that the array can be sorted in non-decreasing order using any number of such swaps. 
// If nums is already sorted, return 0.

// Example 1:
// Input: nums = [0,3,2,1]
// Output: 1
// Explanation:
// Choose k = 1. Swapping nums[1] = 3 and nums[3] = 1 is allowed since nums[1] AND nums[3] == 1, resulting in a sorted permutation: [0, 1, 2, 3].

// Example 2:
// Input: nums = [0,1,3,2]
// Output: 2
// Explanation:
// Choose k = 2. Swapping nums[2] = 3 and nums[3] = 2 is allowed since nums[2] AND nums[3] == 2, resulting in a sorted permutation: [0, 1, 2, 3].

// Example 3:
// Input: nums = [3,2,1,0]
// Output: 0
// Explanation:
// Only k = 0 allows sorting since no greater k allows the required swaps where nums[i] AND nums[j] == k.

// Constraints:
//     1 <= n == nums.length <= 10^5
//     0 <= nums[i] <= n - 1
//     nums is a permutation of integers from 0 to n - 1.

import "fmt"

func sortPermutation(nums []int) int {
    var v *int
    for i := 0; i < len(nums); i++ {
        if nums[i] != i && v == nil {
            v = &nums[i]
        } else if nums[i] != i {
            *v &= nums[i]
        }
    }
    if v == nil {return 0}
    return *v
}

func sortPermutation1(nums []int) int {
    if nums[0] != 0 {
        return 0
    }
    res, n:= -1, len(nums)
    for i := 0; i < n; i++ {
        if nums[i] != i {
            res &= nums[i]
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    return max(res, 0)
}

func main() {
    // Example 1:
    // Input: nums = [0,3,2,1]
    // Output: 1
    // Explanation:
    // Choose k = 1. Swapping nums[1] = 3 and nums[3] = 1 is allowed since nums[1] AND nums[3] == 1, resulting in a sorted permutation: [0, 1, 2, 3].
    fmt.Println(sortPermutation([]int{0,3,2,1})) // 1
    // Example 2:
    // Input: nums = [0,1,3,2]
    // Output: 2
    // Explanation:
    // Choose k = 2. Swapping nums[2] = 3 and nums[3] = 2 is allowed since nums[2] AND nums[3] == 2, resulting in a sorted permutation: [0, 1, 2, 3].
    fmt.Println(sortPermutation([]int{0,1,3,2})) // 2
    // Example 3:
    // Input: nums = [3,2,1,0]
    // Output: 0
    // Explanation:
    // Only k = 0 allows sorting since no greater k allows the required swaps where nums[i] AND nums[j] == k.
    fmt.Println(sortPermutation([]int{3,2,1,0})) // 0

    fmt.Println(sortPermutation([]int{0,1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(sortPermutation([]int{9,8,7,6,5,4,3,2,1,0})) // 0

    fmt.Println(sortPermutation1([]int{0,3,2,1})) // 1
    fmt.Println(sortPermutation1([]int{0,1,3,2})) // 2
    fmt.Println(sortPermutation1([]int{3,2,1,0})) // 0
    fmt.Println(sortPermutation1([]int{0,1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(sortPermutation1([]int{9,8,7,6,5,4,3,2,1,0})) // 0
}