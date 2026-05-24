package main

// 3942. Minimum Operations to Sort a Permutation
// You are given an integer array nums of length n, where nums is a permutation of the numbers in the range [0..n - 1].

// You may perform only the following operations:
//     1. Reverse the entire array.
//     2. Rotate Left by One: Move the first element to the end of the array, and rest elements to left by one position.

// Return the minimum number of operations required to sort the array in increasing order.
// If it is not possible to sort the array using only the given operations, return -1.

// A permutation is a rearrangement of all the elements of an array.

// Example 1:
// Input: nums = [0,2,1]
// Output: 2
// Explanation:
// Rotate Left by one: [2, 1, 0]
// Reverse the array: [0, 1, 2]
// The array becomes sorted in 2 operations, which is minimal

// Example 2:
// Input: nums = [1,0,2]
// Output: 2
// Explanation:
// Reverse the array: [2, 0, 1]
// Rotate Left by one: [0, 1, 2]
// The array becomes sorted in 2 operations, which is minimal.

// Example 3:
// Input: nums = [2,0,1,3]
// Output: -1
// Explanation:
// It is impossible to reach [2, 0, 1, 3]. Thus, the answer is -1.

// Constraints:
//     1 <= n == nums.length <= 10^5
//     0 <= nums[i] <= n - 1
//     nums is a permutation of integers from 0 to n - 1.

import "fmt"

func minOperations(nums []int) int {
    res, count, p, n := 1 << 61, 0, 0, len(nums)
    for i := 1; i < n && count < 2; i++ {
        if nums[i-1] > nums[i] {
            count++
            p = i
        }
    }
    if count == 0 { // 已是递增
        return 0
    }
    if count == 1 && nums[0] > nums[n-1] { // 两个递增段
        res = min(p, n-p+2)
    }
    count, p = 0, 0
    for i := 1; i < n && count < 2; i++ {
        if nums[i-1] < nums[i] {
            count++
            p = i
        }
    }
    if count == 0 { // 已是递减
        return 1
    }
    if count == 1 && nums[0] < nums[n-1] { // 两个递减段
        res = min(res, p+1, n-p+1)
    }
    if res == 1 << 61 {
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,2,1]
    // Output: 2
    // Explanation:
    // Rotate Left by one: [2, 1, 0]
    // Reverse the array: [0, 1, 2]
    // The array becomes sorted in 2 operations, which is minimal
    fmt.Println(minOperations([]int{0,2,1})) // 2
    // Example 2:
    // Input: nums = [1,0,2]
    // Output: 2
    // Explanation:
    // Reverse the array: [2, 0, 1]
    // Rotate Left by one: [0, 1, 2]
    // The array becomes sorted in 2 operations, which is minimal
    fmt.Println(minOperations([]int{1,0,2})) // 2
    // Example 3:
    // Input: nums = [2,0,1,3]
    // Output: -1
    // Explanation:
    // It is impossible to reach [2, 0, 1, 3]. Thus, the answer is -1. 
    fmt.Println(minOperations([]int{2,0,1,3})) // -1

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // 1
}