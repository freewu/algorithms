package main

// 905. Sort Array By Parity
// Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
// Return any array that satisfies this condition.

// Example 1:
// Input: nums = [3,1,2,4]
// Output: [2,4,3,1]
// Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

// Example 2:
// Input: nums = [0]
// Output: [0]

// Constraints:
//     1 <= nums.length <= 5000
//     0 <= nums[i] <= 5000

import "fmt"

func sortArrayByParity(nums []int) []int {
    for i, j := 0, 0; j < len(nums); j++ {
        if nums[j] % 2 == 0 { // 如果是偶数交换到前面去
            // t := nums[i]
            // nums[i] = nums[j]
            // nums[j] = t
            nums[i],nums[j] =  nums[j],  nums[i]
            i++
        }
    }
    return nums
}

func sortArrayByParity1(nums []int) []int {
    odd, even:= map[int]int{}, map[int]int{}
    for _, v := range nums {
        if v % 2 == 1 {
            odd[v]++
        } else {
            even[v]++
        }
    }
    i := 0
    for k, v := range even {
        for v > 0 {
            nums[i] = k
            i++
            v--
        }
    }
    for k, v := range odd {
        for v > 0 {
            nums[i] = k
            i++
            v--
        }
    }
    return nums
}

func main() {
    // Example 1:
    // Input: nums = [3,1,2,4]
    // Output: [2,4,3,1]
    // Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
    fmt.Println(sortArrayByParity([]int{3,1,2,4})) // [2,4,3,1]
    // Example 2:
    // Input: nums = [0]
    // Output: [0]
    fmt.Println(sortArrayByParity([]int{0})) // [0]

    fmt.Println(sortArrayByParity1([]int{3,1,2,4})) // [2,4,3,1]
    fmt.Println(sortArrayByParity1([]int{0})) // [0]
}