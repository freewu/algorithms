package main

// 3379. Transformed Array
// You are given an integer array nums that represents a circular array. 
// Your task is to create a new array result of the same size, following these rules:

// For each index i (where 0 <= i < nums.length), perform the following independent actions:
//     1. If nums[i] > 0: Start at index i and move nums[i] steps to the right in the circular array. 
//        Set result[i] to the value of the index where you land.
//     2. If nums[i] < 0: Start at index i and move abs(nums[i]) steps to the left in the circular array. 
//        Set result[i] to the value of the index where you land.
//     3. If nums[i] == 0: Set result[i] to nums[i].

// Return the new array result.

// Note: Since nums is circular, moving past the last element wraps around to the beginning, and moving before the first element wraps back to the end.

// Example 1:
// Input: nums = [3,-2,1,1]
// Output: [1,1,1,3]
// Explanation:
// For nums[0] that is equal to 3, If we move 3 steps to right, we reach nums[3]. So result[0] should be 1.
// For nums[1] that is equal to -2, If we move 2 steps to left, we reach nums[3]. So result[1] should be 1.
// For nums[2] that is equal to 1, If we move 1 step to right, we reach nums[3]. So result[2] should be 1.
// For nums[3] that is equal to 1, If we move 1 step to right, we reach nums[0]. So result[3] should be 3.

// Example 2:
// Input: nums = [-1,4,-1]
// Output: [-1,-1,4]
// Explanation:
// For nums[0] that is equal to -1, If we move 1 step to left, we reach nums[2]. So result[0] should be -1.
// For nums[1] that is equal to 4, If we move 4 steps to right, we reach nums[2]. So result[1] should be -1.
// For nums[2] that is equal to -1, If we move 1 step to left, we reach nums[1]. So result[2] should be 4.

// Constraints:
//     1 <= nums.length <= 100
//     -100 <= nums[i] <= 100

import "fmt"

func constructTransformedArray(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    for i, v := range nums {
        if v > 0 { // 从下标 i 开始，向 右 移动 nums[i] 步，在循环数组中落脚的下标对应的值赋给 res[i]
            res[i] = nums[(i + v) % n]
        } else if v < 0 { // 从下标 i 开始，向 左 移动 abs(nums[i]) 步，在循环数组中落脚的下标对应的值赋给 res[i]
            s := (-v) % n
            m := i - s
            if m < 0 {
                m += n
            }
            res[i] = nums[m]
        } else { // 0
            res[i] = v
        }
    }
    return res
}

func constructTransformedArray1(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    for i := 0; i < n; i++ {
        step := nums[i]
        step = ((step + i) % n + n) % n
        res[i] = nums[step]
    }
    return res
}

func constructTransformedArray2(nums []int) []int {
    n := len(nums)
    res := []int{}
    for i := 0; i < n; i++ {
        res = append (res, nums[(i + nums[i] % n + n) % n])
    }
    return res
}

func constructTransformedArray3(nums []int) []int {
    n := len(nums)
    bign := n * 100
    res := make([]int, n)
    for i, v := range nums {
        res[i] = nums[(i + v + bign) % n]
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,-2,1,1]
    // Output: [1,1,1,3]
    // Explanation:
    // For nums[0] that is equal to 3, If we move 3 steps to right, we reach nums[3]. So result[0] should be 1.
    // For nums[1] that is equal to -2, If we move 2 steps to left, we reach nums[3]. So result[1] should be 1.
    // For nums[2] that is equal to 1, If we move 1 step to right, we reach nums[3]. So result[2] should be 1.
    // For nums[3] that is equal to 1, If we move 1 step to right, we reach nums[0]. So result[3] should be 3.
    fmt.Println(constructTransformedArray([]int{3,-2,1,1})) // [1,1,1,3]
    // Example 2:
    // Input: nums = [-1,4,-1]
    // Output: [-1,-1,4]
    // Explanation:
    // For nums[0] that is equal to -1, If we move 1 step to left, we reach nums[2]. So result[0] should be -1.
    // For nums[1] that is equal to 4, If we move 4 steps to right, we reach nums[2]. So result[1] should be -1.
    // For nums[2] that is equal to -1, If we move 1 step to left, we reach nums[1]. So result[2] should be 4.
    fmt.Println(constructTransformedArray([]int{-1,4,-1})) // [-1,-1,4]

    fmt.Println(constructTransformedArray([]int{1,2,3,4,5,6,7,8,9})) // [2 4 6 8 1 3 5 7 9]
    fmt.Println(constructTransformedArray([]int{9,8,7,6,5,4,3,2,1})) // [9 9 9 9 9 9 9 9 9]

    fmt.Println(constructTransformedArray1([]int{3,-2,1,1})) // [1,1,1,3]
    fmt.Println(constructTransformedArray1([]int{-1,4,-1})) // [-1,-1,4]
    fmt.Println(constructTransformedArray1([]int{1,2,3,4,5,6,7,8,9})) // [2 4 6 8 1 3 5 7 9]
    fmt.Println(constructTransformedArray1([]int{9,8,7,6,5,4,3,2,1})) // [9 9 9 9 9 9 9 9 9]

    fmt.Println(constructTransformedArray2([]int{3,-2,1,1})) // [1,1,1,3]
    fmt.Println(constructTransformedArray2([]int{-1,4,-1})) // [-1,-1,4]
    fmt.Println(constructTransformedArray2([]int{1,2,3,4,5,6,7,8,9})) // [2 4 6 8 1 3 5 7 9]
    fmt.Println(constructTransformedArray2([]int{9,8,7,6,5,4,3,2,1})) // [9 9 9 9 9 9 9 9 9]

    fmt.Println(constructTransformedArray3([]int{3,-2,1,1})) // [1,1,1,3]
    fmt.Println(constructTransformedArray3([]int{-1,4,-1})) // [-1,-1,4]
    fmt.Println(constructTransformedArray3([]int{1,2,3,4,5,6,7,8,9})) // [2 4 6 8 1 3 5 7 9]
    fmt.Println(constructTransformedArray3([]int{9,8,7,6,5,4,3,2,1})) // [9 9 9 9 9 9 9 9 9]
}