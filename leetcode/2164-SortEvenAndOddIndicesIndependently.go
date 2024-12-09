package main

// 2164. Sort Even and Odd Indices Independently
// You are given a 0-indexed integer array nums. Rearrange the values of nums according to the following rules:
//     1. Sort the values at odd indices of nums in non-increasing order.
//             For example, if nums = [4,1,2,3] before this step, it becomes [4,3,2,1] after. 
//             The values at odd indices 1 and 3 are sorted in non-increasing order.
//     2. Sort the values at even indices of nums in non-decreasing order.
//             For example, if nums = [4,1,2,3] before this step, it becomes [2,1,4,3] after. 
//             The values at even indices 0 and 2 are sorted in non-decreasing order.

// Return the array formed after rearranging the values of nums.

// Example 1:
// Input: nums = [4,1,2,3]
// Output: [2,3,4,1]
// Explanation: 
// First, we sort the values present at odd indices (1 and 3) in non-increasing order.
// So, nums changes from [4,1,2,3] to [4,3,2,1].
// Next, we sort the values present at even indices (0 and 2) in non-decreasing order.
// So, nums changes from [4,1,2,3] to [2,3,4,1].
// Thus, the array formed after rearranging the values is [2,3,4,1].

// Example 2:
// Input: nums = [2,1]
// Output: [2,1]
// Explanation: 
// Since there is exactly one odd index and one even index, no rearrangement of values takes place.
// The resultant array formed is [2,1], which is the same as the initial array. 

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"
import "sort"

func sortEvenOdd(nums []int) []int {
    even, odd := make([]int, 0), make([]int, 0)
    for i, v := range nums {
        if i % 2 == 0 {
            even = append(even, v)
        } else {
            odd = append(odd, v)
        }
    }
    sort.Ints(even)
    sort.Sort(sort.Reverse(sort.IntSlice(odd)))
    i, j := 0, 0
    for k := 0; k < len(nums); k++ {
        if k % 2 == 0 {
            nums[k] = even[i]
            i++
        } else {
            nums[k] = odd[j]
            j++
        }
    }
    return nums
}

func sortEvenOdd1(nums []int) []int {
    even, odd := make([]int, 0), make([]int, 0)
    for i, v := range nums {
        if i % 2 == 0 {
            even = append(even, v)
        } else {
            odd = append(odd, v)
        }
    }
    sort.Ints(odd)
    sort.Ints(even)
    i, j := 0, len(odd) - 1 // // 偶数，递增排序, 奇数，递减排序
    res := []int{}
    for k := 0; k < len(nums); k++ {
        if k % 2 == 1 {
            res = append(res, odd[j])
            j--
        } else {
            res = append(res, even[i])
            i++
        }
    }
    return res 
}

func main() {
    // Example 1:
    // Input: nums = [4,1,2,3]
    // Output: [2,3,4,1]
    // Explanation: 
    // First, we sort the values present at odd indices (1 and 3) in non-increasing order.
    // So, nums changes from [4,1,2,3] to [4,3,2,1].
    // Next, we sort the values present at even indices (0 and 2) in non-decreasing order.
    // So, nums changes from [4,1,2,3] to [2,3,4,1].
    // Thus, the array formed after rearranging the values is [2,3,4,1].
    fmt.Println(sortEvenOdd([]int{4,1,2,3})) // [2,3,4,1]
    // Example 2:
    // Input: nums = [2,1]
    // Output: [2,1]
    // Explanation: 
    // Since there is exactly one odd index and one even index, no rearrangement of values takes place.
    // The resultant array formed is [2,1], which is the same as the initial array. 
    fmt.Println(sortEvenOdd([]int{2,1})) // [2,1]

    fmt.Println(sortEvenOdd1([]int{4,1,2,3})) // [2,3,4,1]
    fmt.Println(sortEvenOdd1([]int{2,1})) // [2,1]
}