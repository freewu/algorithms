package main

// 163. Missing Ranges
// You are given an inclusive range [lower, upper] and a sorted unique integer array nums, where all elements are within the inclusive range.
// A number x is considered missing if x is in the range [lower, upper] and x is not in nums.
// Return the shortest sorted list of ranges that exactly covers all the missing numbers. 
// That is, no element of nums is included in any of the ranges, and each missing number is covered by one of the ranges.

// Example 1:
// Input: nums = [0,1,3,50,75], lower = 0, upper = 99
// Output: [[2,2],[4,49],[51,74],[76,99]]
// Explanation: The ranges are:
// [2,2]
// [4,49]
// [51,74]
// [76,99]

// Example 2:
// Input: nums = [-1], lower = -1, upper = -1
// Output: []
// Explanation: There are no missing ranges since there are no missing numbers.
 
// Constraints:
//     -10^9 <= lower <= upper <= 10^9
//     0 <= nums.length <= 100
//     lower <= nums[i] <= upper
//     All the values of nums are unique.

import "fmt"

func findMissingRanges(nums []int, lower int, upper int) [][]int {
    res := [][]int{}
    // 把 lower & upper 也加到 nums 中
    nums = append(append([]int{lower - 1}, nums...), upper + 1)
    for i := 1; i < len(nums); i++ {
        // 如果不和前一个数据相连(差 1) 说明存在相差
        if nums[i] - nums[i-1] > 1 {
            res = append(res, []int{nums[i-1] + 1, nums[i] - 1})
        }
    }
    return res
}

func findMissingRanges1(nums []int, lower int, upper int) [][]int {
    if len(nums) == 0 {
        return [][]int{{ lower, upper }}
    }
    var res [][]int
    // 如果 nums[0] > lower 说明存在 [lower,nums[0] - 1] 的范围
    if nums[0] > lower {
        res = append(res, []int{lower, nums[0] - 1})
    }
    for i := 0; i < len(nums)-1; i++ {
        // 不是相连说明出现缺失的范围 [nums[i] + 1, nums[i+1] - 1]
        if nums[i+1] - nums[i] != 1 {
            res = append(res, []int{nums[i] + 1, nums[i+1] - 1})
        }
    }
    // 如果 nums[len(nums)-1] < upper 存在 [nums[len(nums)-1] + 1, upper]
    if nums[len(nums)-1] < upper {
        res = append(res, []int{nums[len(nums)-1]+1, upper})
    }
    return res
}

func main() {
    // Explanation: The ranges are:
    // [2,2]
    // [4,49]
    // [51,74]
    // [76,99]
    fmt.Println(findMissingRanges([]int{0,1,3,50,75}, 0, 99))

    // Explanation: There are no missing ranges since there are no missing numbers.
    fmt.Println(findMissingRanges([]int{-1}, -1, -1)) // []

    fmt.Println(findMissingRanges1([]int{0,1,3,50,75}, 0, 99))
    fmt.Println(findMissingRanges1([]int{-1}, -1, -1)) // []
}

