package main

// 503. Next Greater Element II
// Given a circular integer array nums (i.e., the next element of nums[nums.length - 1] is nums[0]), 
// return the next greater number for every element in nums.

// The next greater number of a number x is the first greater number to its traversing-order next in the array, 
// which means you could search circularly to find its next greater number. 
// If it doesn't exist, return -1 for this number.

// Example 1:
// Input: nums = [1,2,1]
// Output: [2,-1,2]
// Explanation: The first 1's next greater number is 2; 
// The number 2 can't find next greater number. 
// The second 1's next greater number needs to search circularly, which is also 2.

// Example 2:
// Input: nums = [1,2,3,4,3]
// Output: [2,3,4,-1,4]
 
// Constraints:
//     1 <= nums.length <= 10^4
//     -10^9 <= nums[i] <= 10^9

import "fmt"

// func nextGreaterElements(nums []int) []int {
//     res := []int{}
//     for i := 0; i < len(nums); i++ {
//         for j := 0; j < len(nums); j++ {
//             if nums[j] > nums[i] { // 找到了比自己大的数
//                 res = append(res, nums[j])
//                 break
//             }
//         }
//         if len(res) == i { // 没找到比自己大的数
//             res = append(res, -1)
//         }
//     }
//     return res
// }

// stack 
func nextGreaterElements(snums []int) []int {
    nums := append(snums, snums...)
    results, stack := make([]int, len(nums)), []int{}
    for i, num := range nums {
        if len(stack) == 0 {
            stack = append(stack, i)
        } else {
            for len(stack) > 0 && num > nums[stack[len(stack)-1]] {
                idx := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                results[idx] = num
            }
            stack = append(stack, i)
        }
    }
    for len(stack) > 0 {
        idx := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        results[idx] = -1
    }
    return results[:len(snums)]
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1]
    // Output: [2,-1,2]
    // Explanation: The first 1's next greater number is 2; 
    // The number 2 can't find next greater number. 
    // The second 1's next greater number needs to search circularly, which is also 2.
    fmt.Println(nextGreaterElements([]int{1,2,1})) // [2,-1,2]
    // Example 2:
    // Input: nums = [1,2,3,4,3]
    // Output: [2,3,4,-1,4]
    fmt.Println(nextGreaterElements([]int{1,2,3,4,3})) // [2,3,4,-1,4]

    fmt.Println(nextGreaterElements([]int{1,5,3,6,8})) // [5,6,6,8,-1]
}