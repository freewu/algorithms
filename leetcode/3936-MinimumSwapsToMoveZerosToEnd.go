package main

// 3936. Minimum Swaps to Move Zeros to End
// You are given an integer array nums.

// In one operation, you can choose any two distinct indices i and j and swap nums[i] and nums[j].

// Return an integer denoting the minimum number of operations required to move all 0s to the end of the array.

// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: 2
// Explanation:
// We perform the following swap operations:
// Swap nums[0] and nums[3], giving nums = [3, 1, 0, 0, 12].
// Swap nums[2] and nums[4], giving nums = [3, 1, 12, 0, 0].
// Thus, the answer is 2.

// Example 2:
// Input: nums = [0,1,0,2]
// Output: 1
// Explanation:
// We perform the following swap operations:
// Swap nums[0] and nums[3], giving nums = [2, 1, 0, 0].
// Thus, the answer is 1.

// Example 3:
// Input: nums = [1,2,0]
// Output: 0
// Explanation:
// The array already satisfies the condition. Therefore, no swap operations are needed.

// Constraints:
//     1 <= nums.length <= 100
//     0 <= nums[i] <= 100

import "fmt"

func minimumSwaps(nums []int) int {
    res, zeros, n := 0, 0,len(nums)
    for _, v := range nums {
        if v == 0 {
            zeros++
            if nums[n - zeros] != 0 {
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,1,0,3,12]
    // Output: 2
    // Explanation:
    // We perform the following swap operations:
    // Swap nums[0] and nums[3], giving nums = [3, 1, 0, 0, 12].
    // Swap nums[2] and nums[4], giving nums = [3, 1, 12, 0, 0].
    // Thus, the answer is 2.
    fmt.Println(minimumSwaps([]int{0,1,0,3,12})) // 2
    // Example 2:
    // Input: nums = [0,1,0,2]
    // Output: 1
    // Explanation:
    // We perform the following swap operations:
    // Swap nums[0] and nums[3], giving nums = [2, 1, 0, 0].
    // Thus, the answer is 1.
    fmt.Println(minimumSwaps([]int{0,1,0,2})) // 1
    // Example 3:
    // Input: nums = [1,2,0]
    // Output: 0
    // Explanation:
    // The array already satisfies the condition. Therefore, no swap operations are needed.
    fmt.Println(minimumSwaps([]int{1,2,0})) // 0

    fmt.Println(minimumSwaps([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minimumSwaps([]int{9,8,7,6,5,4,3,2,1})) // 0
}