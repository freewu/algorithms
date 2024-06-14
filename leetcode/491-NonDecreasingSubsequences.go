package main

// 491. Non-decreasing Subsequences
// Given an integer array nums, return all the different possible non-decreasing subsequences of the given array with at least two elements. 
// You may return the answer in any order.

// Example 1:
// Input: nums = [4,6,7,7]
// Output: [[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]

// Example 2:
// Input: nums = [4,4,3,2,1]
// Output: [[4,4]]
 
// Constraints:
//     1 <= nums.length <= 15
//     -100 <= nums[i] <= 100

import "fmt"

func findSubsequences(nums []int) [][]int {
    res := [][]int{}
    var dfs func(int, []int)
    dfs = func(l int, sl []int) {
        if len(sl) > 1 {
            cp := make([]int, len(sl))
            copy(cp, sl)
            res = append(res, cp)
        }
        seen := make(map[int]bool, len(nums)-l)
        for r := l; r < len(nums); r++ {
            if (l > 0 && nums[r] < nums[l-1]) || seen[nums[r]] {
                continue
            }
            seen[nums[r]] = true
            dfs(r + 1, append(sl, nums[r]))
        }
    }
    dfs(0, []int{})
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,6,7,7]
    // Output: [[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
    fmt.Println(findSubsequences([]int{4,6,7,7})) // [[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
    // Example 2:
    // Input: nums = [4,4,3,2,1]
    // Output: [[4,4]]
    fmt.Println(findSubsequences([]int{4,4,3,2,1})) // [[4,4]]
}