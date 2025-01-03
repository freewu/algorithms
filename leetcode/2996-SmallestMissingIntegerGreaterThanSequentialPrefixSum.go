package main

// 2996. Smallest Missing Integer Greater Than Sequential Prefix Sum
// You are given a 0-indexed array of integers nums.

// A prefix nums[0..i] is sequential if, for all 1 <= j <= i, nums[j] = nums[j - 1] + 1. 
// In particular, the prefix consisting only of nums[0] is sequential.

// Return the smallest integer x missing from nums such that x is greater than 
// or equal to the sum of the longest sequential prefix.

// Example 1:
// Input: nums = [1,2,3,2,5]
// Output: 6
// Explanation: The longest sequential prefix of nums is [1,2,3] with a sum of 6. 6 is not in the array, therefore 6 is the smallest missing integer greater than or equal to the sum of the longest sequential prefix.

// Example 2:
// Input: nums = [3,4,5,1,12,14,13]
// Output: 15
// Explanation: The longest sequential prefix of nums is [3,4,5] with a sum of 12. 12, 13, and 14 belong to the array while 15 does not. Therefore 15 is the smallest missing integer greater than or equal to the sum of the longest sequential prefix.

// Constraints:
//     1 <= nums.length <= 50
//     1 <= nums[i] <= 50

import "fmt"
import "sort"

func missingInteger(nums []int) int {
    res := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] + 1 { break }
        res += nums[i]
    }
    sort.Ints(nums)
    for i := range nums {
        if nums[i] == res {
            res++
        }
    }
    return res
}

func missingInteger1(nums []int) int {
    set := make(map[int]bool)
    res, i := nums[0], 1
    set[nums[0]] = true
    for ; i < len(nums) && nums[i - 1] + 1 == nums[i]; i++ {
        set[nums[i]] = true
        res += nums[i]
    }
    for ; i < len(nums); i++ {
        set[nums[i]] = true
    }
    for set[res] {
        res++
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,2,5]
    // Output: 6
    // Explanation: The longest sequential prefix of nums is [1,2,3] with a sum of 6. 6 is not in the array, therefore 6 is the smallest missing integer greater than or equal to the sum of the longest sequential prefix.
    fmt.Println(missingInteger([]int{1,2,3,2,5})) // 6
    // Example 2:
    // Input: nums = [3,4,5,1,12,14,13]
    // Output: 15
    // Explanation: The longest sequential prefix of nums is [3,4,5] with a sum of 12. 12, 13, and 14 belong to the array while 15 does not. Therefore 15 is the smallest missing integer greater than or equal to the sum of the longest sequential prefix.
    fmt.Println(missingInteger([]int{3,4,5,1,12,14,13})) // 15

    fmt.Println(missingInteger([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(missingInteger([]int{9,8,7,6,5,4,3,2,1})) // 10

    fmt.Println(missingInteger1([]int{1,2,3,2,5})) // 6
    fmt.Println(missingInteger1([]int{3,4,5,1,12,14,13})) // 15
    fmt.Println(missingInteger1([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(missingInteger1([]int{9,8,7,6,5,4,3,2,1})) // 10
}