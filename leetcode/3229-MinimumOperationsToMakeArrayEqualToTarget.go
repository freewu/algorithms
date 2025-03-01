package main

// 3229. Minimum Operations to Make Array Equal to Target
// You are given two positive integer arrays nums and target, of the same length.

// In a single operation, you can select any subarray of nums and increment each element within that subarray by 1 or decrement each element within that subarray by 1.

// Return the minimum number of operations required to make nums equal to the array target.

// Example 1:
// Input: nums = [3,5,1,2], target = [4,6,2,4]
// Output: 2
// Explanation:
// We will perform the following operations to make nums equal to target:
// - Increment nums[0..3] by 1, nums = [4,6,2,3].
// - Increment nums[3..3] by 1, nums = [4,6,2,4].

// Example 2:
// Input: nums = [1,3,2], target = [2,1,4]
// Output: 5
// Explanation:
// We will perform the following operations to make nums equal to target:
// - Increment nums[0..0] by 1, nums = [2,3,2].
// - Decrement nums[1..1] by 1, nums = [2,2,2].
// - Decrement nums[1..1] by 1, nums = [2,1,2].
// - Increment nums[2..2] by 1, nums = [2,1,3].
// - Increment nums[2..2] by 1, nums = [2,1,4].

// Constraints:
//     1 <= nums.length == target.length <= 10^5
//     1 <= nums[i], target[i] <= 10^8

import "fmt"

func minimumOperations(nums []int, target []int) int64 {
    res, incr, decr := 0, 0, 0
    for i := range nums {
        diff := (target[i] - nums[i])
        if diff > 0 {
            if incr < diff {
                res += diff - incr
            }
            incr, decr = diff, 0
        } else if diff < 0 {
            if diff < decr {
                res += decr - diff
            }
            decr, incr = diff, 0
        } else {
            incr, decr = 0, 0
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [3,5,1,2], target = [4,6,2,4]
    // Output: 2
    // Explanation:
    // We will perform the following operations to make nums equal to target:
    // - Increment nums[0..3] by 1, nums = [4,6,2,3].
    // - Increment nums[3..3] by 1, nums = [4,6,2,4].
    fmt.Println(minimumOperations([]int{3,5,1,2}, []int{4,6,2,4})) // 2
    // Example 2:
    // Input: nums = [1,3,2], target = [2,1,4]
    // Output: 5
    // Explanation:
    // We will perform the following operations to make nums equal to target:
    // - Increment nums[0..0] by 1, nums = [2,3,2].
    // - Decrement nums[1..1] by 1, nums = [2,2,2].
    // - Decrement nums[1..1] by 1, nums = [2,1,2].
    // - Increment nums[2..2] by 1, nums = [2,1,3].
    // - Increment nums[2..2] by 1, nums = [2,1,4].
    fmt.Println(minimumOperations([]int{1,3,2}, []int{2,1,4})) // 5

    fmt.Println(minimumOperations([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minimumOperations([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 16
    fmt.Println(minimumOperations([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 16
    fmt.Println(minimumOperations([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 0
}