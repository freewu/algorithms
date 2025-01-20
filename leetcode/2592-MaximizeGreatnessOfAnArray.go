package main

// 2592. Maximize Greatness of an Array
// You are given a 0-indexed integer array nums. 
// You are allowed to permute nums into a new array perm of your choosing.

// We define the greatness of nums be the number of indices 0 <= i < nums.length for which perm[i] > nums[i].

// Return the maximum possible greatness you can achieve after permuting nums.

// Example 1:
// Input: nums = [1,3,5,2,1,3,1]
// Output: 4
// Explanation: One of the optimal rearrangements is perm = [2,5,1,3,3,1,1].
// At indices = 0, 1, 3, and 4, perm[i] > nums[i]. Hence, we return 4.

// Example 2:
// Input: nums = [1,2,3,4]
// Output: 3
// Explanation: We can prove the optimal perm is [2,3,4,1].
// At indices = 0, 1, and 2, perm[i] > nums[i]. Hence, we return 3.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^9

import "fmt"
import "sort"

func maximizeGreatness(nums []int) int {
    sort.Slice(nums, func(i, j int) bool { // 从大到小
        return nums[j] > nums[i]
    })
    res := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[res] {
            res++
        }
    }
    return res
}

func maximizeGreatness1(nums []int) int {
    sort.Ints(nums)
    res := 0
    for _, v := range nums {
        if v > nums[res] {
            res++
        }
    }
    return res
}

func maximizeGreatness2(nums []int) int {
    sort.Ints(nums)
    res, right := 0, len(nums) - 1
    for i := right; i >= 0;i-- {
        right--
        for right >= 0 && nums[right] >= nums[i] {
            right--
        }
        if right < 0 { break }
        res++
    }
    return res
}


func main() {
    // Example 1:
    // Input: nums = [1,3,5,2,1,3,1]
    // Output: 4
    // Explanation: One of the optimal rearrangements is perm = [2,5,1,3,3,1,1].
    // At indices = 0, 1, 3, and 4, perm[i] > nums[i]. Hence, we return 4.
    fmt.Println(maximizeGreatness([]int{1,3,5,2,1,3,1})) // 4
    // Example 2:
    // Input: nums = [1,2,3,4]
    // Output: 3
    // Explanation: We can prove the optimal perm is [2,3,4,1].
    // At indices = 0, 1, and 2, perm[i] > nums[i]. Hence, we return 3.
    fmt.Println(maximizeGreatness([]int{1,2,3,4})) // 3

    fmt.Println(maximizeGreatness([]int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(maximizeGreatness([]int{9,8,7,6,5,4,3,2,1})) // 8

    fmt.Println(maximizeGreatness1([]int{1,3,5,2,1,3,1})) // 4
    fmt.Println(maximizeGreatness1([]int{1,2,3,4})) // 3
    fmt.Println(maximizeGreatness1([]int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(maximizeGreatness1([]int{9,8,7,6,5,4,3,2,1})) // 8

    fmt.Println(maximizeGreatness2([]int{1,3,5,2,1,3,1})) // 4
    fmt.Println(maximizeGreatness2([]int{1,2,3,4})) // 3
    fmt.Println(maximizeGreatness2([]int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(maximizeGreatness2([]int{9,8,7,6,5,4,3,2,1})) // 8
}