package main

// 1920. Build Array from Permutation
// Given a zero-based permutation nums (0-indexed), 
// build an array ans of the same length where ans[i] = nums[nums[i]] for each 0 <= i < nums.length and return it.

// A zero-based permutation nums is an array of distinct integers from 0 to nums.length - 1 (inclusive).

// Example 1:
// Input: nums = [0,2,1,5,3,4]
// Output: [0,1,2,4,5,3]
// Explanation: The array ans is built as follows: 
// ans = [nums[nums[0]], nums[nums[1]], nums[nums[2]], nums[nums[3]], nums[nums[4]], nums[nums[5]]]
//     = [nums[0], nums[2], nums[1], nums[5], nums[3], nums[4]]
//     = [0,1,2,4,5,3]

// Example 2:
// Input: nums = [5,0,1,2,3,4]
// Output: [4,5,0,1,2,3]
// Explanation: The array ans is built as follows:
// ans = [nums[nums[0]], nums[nums[1]], nums[nums[2]], nums[nums[3]], nums[nums[4]], nums[nums[5]]]
//     = [nums[5], nums[0], nums[1], nums[2], nums[3], nums[4]]
//     = [4,5,0,1,2,3]

// Constraints:
//     1 <= nums.length <= 1000
//     0 <= nums[i] < nums.length
//     The elements in nums are distinct.

// Follow-up: Can you solve it without using an extra space (i.e., O(1) memory)?

import "fmt"

func buildArray(nums []int) []int {
    res := make([]int, len(nums))
    for i, v := range nums {
        res[i] = nums[v]
    }
    return res
}

func buildArray1(nums []int) []int {
    n := len(nums)
    for i := range nums {
        nums[i] = nums[i] + n * (nums[nums[i]] % n)
    }
    for i := range nums {
        nums[i] = nums[i] / n
    }
    return nums
}

func main() {
    // Example 1:
    // Input: nums = [0,2,1,5,3,4]
    // Output: [0,1,2,4,5,3]
    // Explanation: The array ans is built as follows: 
    // ans = [nums[nums[0]], nums[nums[1]], nums[nums[2]], nums[nums[3]], nums[nums[4]], nums[nums[5]]]
    //     = [nums[0], nums[2], nums[1], nums[5], nums[3], nums[4]]
    //     = [0,1,2,4,5,3]
    fmt.Println(buildArray([]int{0,2,1,5,3,4})) // [0,1,2,4,5,3]
    // Example 2:
    // Input: nums = [5,0,1,2,3,4]
    // Output: [4,5,0,1,2,3]
    // Explanation: The array ans is built as follows:
    // ans = [nums[nums[0]], nums[nums[1]], nums[nums[2]], nums[nums[3]], nums[nums[4]], nums[nums[5]]]
    //     = [nums[5], nums[0], nums[1], nums[2], nums[3], nums[4]]
    //     = [4,5,0,1,2,3]
    fmt.Println(buildArray([]int{5,0,1,2,3,4})) // [4,5,0,1,2,3]

    fmt.Println(buildArray([]int{0,1,2,3,4,5,6,7,8,9})) // [0 1 2 3 4 5 6 7 8 9]
    fmt.Println(buildArray([]int{9,8,7,6,5,4,3,2,1,0})) // [0 1 2 3 4 5 6 7 8 9]

    fmt.Println(buildArray1([]int{0,2,1,5,3,4})) // [0,1,2,4,5,3]
    fmt.Println(buildArray1([]int{5,0,1,2,3,4})) // [4,5,0,1,2,3]
    fmt.Println(buildArray1([]int{0,1,2,3,4,5,6,7,8,9})) // [0 1 2 3 4 5 6 7 8 9]
    fmt.Println(buildArray1([]int{9,8,7,6,5,4,3,2,1,0})) // [0 1 2 3 4 5 6 7 8 9]
}