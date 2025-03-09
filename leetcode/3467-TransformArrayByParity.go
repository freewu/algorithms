package main

// 3467. Transform Array by Parity
// You are given an integer array nums. 
// Transform nums by performing the following operations in the exact order specified:
//     1. Replace each even number with 0.
//     2. Replace each odd numbers with 1.
//     3. Sort the modified array in non-decreasing order.

// Return the resulting array after performing these operations.

// Example 1:
// Input: nums = [4,3,2,1]
// Output: [0,0,1,1]
// Explanation:
// Replace the even numbers (4 and 2) with 0 and the odd numbers (3 and 1) with 1. Now, nums = [0, 1, 0, 1].
// After sorting nums in non-descending order, nums = [0, 0, 1, 1].

// Example 2:
// Input: nums = [1,5,1,4,2]
// Output: [0,0,1,1,1]
// Explanation:
// Replace the even numbers (4 and 2) with 0 and the odd numbers (1, 5 and 1) with 1. Now, nums = [1, 1, 1, 0, 0].
// After sorting nums in non-descending order, nums = [0, 0, 1, 1, 1].

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 1000

import "fmt"
import "sort"

func transformArray(nums []int) []int {
    res := make([]int, len(nums))
    for i, v := range nums {
        // if v % 2 == 0 { // Replace each even number with 0.
        //     res[i] = 0
        // } else { // Replace each odd numbers with 1.
        //     res[i] = 1
        // }
        res[i] = (v % 2)
    }
    sort.Ints(res)
    return res
}

// 双指针
func transformArray1(nums []int) []int {
    i, j := 0, len(nums) - 1
    for i <= j {
        if nums[i] % 2 == 0 {
            nums[i] = 0
            i++
        } else {
            nums[i] = nums[j] // 交换重点
            nums[j] = 1
            j--
        }
    }
    return nums
}

// 统计出 0 / 1 数量
func transformArray2(nums []int) []int {
    count, n :=[2]int{}, len(nums)
    for _, v := range nums {
        count[v % 2]++
    }
    res, index := make([]int, n), 0
    for i := 0; i < count[0]; i++ {
        res[index] = 0
        index++
    }
    for i := 0; i < count[1]; i++ {
        res[index] = 1
        index++
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,3,2,1]
    // Output: [0,0,1,1]
    // Explanation:
    // Replace the even numbers (4 and 2) with 0 and the odd numbers (3 and 1) with 1. Now, nums = [0, 1, 0, 1].
    // After sorting nums in non-descending order, nums = [0, 0, 1, 1].
    fmt.Println(transformArray([]int{4,3,2,1})) // [0,0,1,1]
    // Example 2:
    // Input: nums = [1,5,1,4,2]
    // Output: [0,0,1,1,1]
    // Explanation:
    // Replace the even numbers (4 and 2) with 0 and the odd numbers (1, 5 and 1) with 1. Now, nums = [1, 1, 1, 0, 0].
    // After sorting nums in non-descending order, nums = [0, 0, 1, 1, 1].
    fmt.Println(transformArray([]int{1,5,1,4,2})) // [0,0,1,1,1]

    fmt.Println(transformArray([]int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 1 1 1 1 1]
    fmt.Println(transformArray([]int{9,8,7,6,5,4,3,2,1})) // [0 0 0 0 1 1 1 1 1]

    fmt.Println(transformArray1([]int{4,3,2,1})) // [0,0,1,1]
    fmt.Println(transformArray1([]int{1,5,1,4,2})) // [0,0,1,1,1]
    fmt.Println(transformArray1([]int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 1 1 1 1 1]
    fmt.Println(transformArray1([]int{9,8,7,6,5,4,3,2,1})) // [0 0 0 0 1 1 1 1 1]

    fmt.Println(transformArray2([]int{4,3,2,1})) // [0,0,1,1]
    fmt.Println(transformArray2([]int{1,5,1,4,2})) // [0,0,1,1,1]
    fmt.Println(transformArray2([]int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 1 1 1 1 1]
    fmt.Println(transformArray2([]int{9,8,7,6,5,4,3,2,1})) // [0 0 0 0 1 1 1 1 1]
}