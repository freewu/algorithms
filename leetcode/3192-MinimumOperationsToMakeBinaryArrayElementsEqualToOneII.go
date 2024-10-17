package main

// 3192. Minimum Operations to Make Binary Array Elements Equal to One II
// You are given a binary array nums.

// You can do the following operation on the array any number of times (possibly zero):
//     Choose any index i from the array and flip all the elements from index i to the end of the array.

// Flipping an element means changing its value from 0 to 1, and from 1 to 0.

// Return the minimum number of operations required to make all elements in nums equal to 1.

// Example 1:
// Input: nums = [0,1,1,0,1]
// Output: 4
// Explanation:
// We can do the following operations:
//     Choose the index i = 1. The resulting array will be nums = [0,0,0,1,0].
//     Choose the index i = 0. The resulting array will be nums = [1,1,1,0,1].
//     Choose the index i = 4. The resulting array will be nums = [1,1,1,0,0].
//     Choose the index i = 3. The resulting array will be nums = [1,1,1,1,1].

// Example 2:
// Input: nums = [1,0,0,0]
// Output: 1
// Explanation:
// We can do the following operation:
// Choose the index i = 1. The resulting array will be nums = [1,1,1,1].

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 1

import "fmt"

func minOperations(nums []int) int {
    res, flipped := 0, 0
    for i := 0; i < len(nums); i++ {
        // Determine the current state considering the number of flips so far
        state := nums[i]
        if flipped % 2 != 0 { state = 1 - nums[i] }
        if state == 0 { // Perform a flip operation
            res++
            flipped++
        }
    }
    return res
}

func minOperations1(nums []int) int {
    res := 0
    for i := len(nums) - 2; i >= 0; i-- {
        if nums[i] != nums[i+1] {
            res++
        }
    }
    if nums[0] == 1 { return res }
    return res + 1
}

func minOperations2(nums []int) int {
    res, val := 0, 0
    for _, v := range nums {
        v ^= val
        if v == 0 {
            val ^= 1
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,1,1,0,1]
    // Output: 4
    // Explanation:
    // We can do the following operations:
    //     Choose the index i = 1. The resulting array will be nums = [0,0,0,1,0].
    //     Choose the index i = 0. The resulting array will be nums = [1,1,1,0,1].
    //     Choose the index i = 4. The resulting array will be nums = [1,1,1,0,0].
    //     Choose the index i = 3. The resulting array will be nums = [1,1,1,1,1].
    fmt.Println(minOperations([]int{0,1,1,0,1})) // 4
    // Example 2:
    // Input: nums = [1,0,0,0]
    // Output: 1
    // Explanation:
    // We can do the following operation:
    // Choose the index i = 1. The resulting array will be nums = [1,1,1,1].
    fmt.Println(minOperations([]int{1,0,0,0})) // 1

    fmt.Println(minOperations1([]int{0,1,1,0,1})) // 4
    fmt.Println(minOperations1([]int{1,0,0,0})) // 1

    fmt.Println(minOperations2([]int{0,1,1,0,1})) // 4
    fmt.Println(minOperations2([]int{1,0,0,0})) // 1
}