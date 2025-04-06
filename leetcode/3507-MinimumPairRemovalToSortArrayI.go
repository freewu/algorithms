package main

// 3507. Minimum Pair Removal to Sort Array I
// Given an array nums, you can perform the following operation any number of times:
//     1. Select the adjacent pair with the minimum sum in nums. 
//        If multiple such pairs exist, choose the leftmost one.
//     2. Replace the pair with their sum.

// Return the minimum number of operations needed to make the array non-decreasing.

// An array is said to be non-decreasing if each element is greater than or equal to its previous element (if it exists).

// Example 1:
// Input: nums = [5,2,3,1]
// Output: 2
// Explanation:
// The pair (3,1) has the minimum sum of 4. After replacement, nums = [5,2,4].
// The pair (2,4) has the minimum sum of 6. After replacement, nums = [5,6].
// The array nums became non-decreasing in two operations.

// Example 2:
// Input: nums = [1,2,2]
// Output: 0
// Explanation:
// The array nums is already sorted.

// Constraints:
//     1 <= nums.length <= 50
//     -1000 <= nums[i] <= 1000

import "fmt"

func minimumPairRemoval(nums []int) int {
    helper := func(arr []int, n int) {
        index, sum := 0, nums[0] + nums[1]
        for i := 0; i < n-1; i++ {
            if nums[i] + nums[i+1] < sum {
                index, sum = i, nums[i] + nums[i+1]
            }
        }
        nums[index] = sum
        for i := index + 1; i < n - 1; i++ {
            nums[i] = nums[i+1]
        }
    }
    res, n := 0, len(nums)
    for i := 0; i < n-1; i++ {
        if nums[i] > nums[i+1] {
            res++
            helper(nums, n)
            i = -1
            n--
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,2,3,1]
    // Output: 2
    // Explanation:
    // The pair (3,1) has the minimum sum of 4. After replacement, nums = [5,2,4].
    // The pair (2,4) has the minimum sum of 6. After replacement, nums = [5,6].
    // The array nums became non-decreasing in two operations.
    fmt.Println(minimumPairRemoval([]int{5,2,3,1})) // 2
    // Example 2:
    // Input: nums = [1,2,2]
    // Output: 0
    // Explanation:
    // The array nums is already sorted.
    fmt.Println(minimumPairRemoval([]int{1,2,2})) // 0

    fmt.Println(minimumPairRemoval([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minimumPairRemoval([]int{9,8,7,6,5,4,3,2,1})) // 7
}