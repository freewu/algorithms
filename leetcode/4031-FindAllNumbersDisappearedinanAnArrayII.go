package main 

// 4031. Find All Numbers Disappeared in an Array II
// You are given an integer array nums and two integers lower and upper.

// A missing integer is an integer in the inclusive range [lower, upper] that does not appear in nums.

// Return a 2D integer array where each element is of the form [start, end], representing a contiguous range of missing integers. 
// Return the ranges in increasing order. 
// If there are no missing integers, return an empty array.

// Note: Consecutive missing integers should be grouped into a single range.

// Example 1:
// Input: nums = [3,9,7], lower = 1, upper = 12
// Output: [[1,2],[4,6],[8,8],[10,12]]
// Explanation:
// The missing integers are [1, 2, 4, 5, 6, 8, 10, 11, 12].
// Grouping the missing integers into the minimum number of contiguous ranges, we get [1, 2], [4, 6], [8, 8], and [10, 12].
// Therefore, the answer is [[1, 2], [4, 6], [8, 8], [10, 12]].

// Example 2:
// Input: nums = [1,1], lower = 5, upper = 7
// Output: [[5,7]]
// Explanation:
// The missing integers are [5, 6, 7].
// Grouping the missing integers into the minimum number of contiguous ranges, we get [5, 7].
// Therefore, the answer is [[5, 7]].

// Example 3:
// Input: nums = [2,3,5], lower = 2, upper = 3
// Output: []
// Explanation:
// There are no missing integers.
// Therefore, the answer is [].

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     1 <= lower <= upper <= 10^5

import "fmt"
import "sort"

func findDisappearedNumbers(nums []int, lower, upper int) [][]int {
    res := [][]int{}
    nums = append(nums, lower - 1, upper + 1)
    sort.Ints(nums)
    l, r := sort.SearchInts(nums, lower), sort.SearchInts(nums, upper+1)
    for i := l; i <= r; i++ {
        if nums[i]-nums[i-1] > 1 {
            res = append(res, []int{nums[i-1] + 1, nums[i] - 1})
        }
    }
    return res
}

func findDisappearedNumbers1(nums []int, lower int, upper int) [][]int {
    res, i := [][]int{}, 0
    seen := make([]bool, upper-lower+1)
    for _, v := range nums {
        if v >= lower && v <= upper {
            seen[v-lower] = true
        }
    }
    for i < len(seen) {
        if seen[i] {
            i++
            continue
        }
        l := i
        for i < len(seen) && !seen[i] {
            i++
        }
        res = append(res, []int{ lower + l, lower + i - 1 })
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,9,7], lower = 1, upper = 12
    // Output: [[1,2],[4,6],[8,8],[10,12]]
    // Explanation:
    // The missing integers are [1, 2, 4, 5, 6, 8, 10, 11, 12].
    // Grouping the missing integers into the minimum number of contiguous ranges, we get [1, 2], [4, 6], [8, 8], and [10, 12].
    // Therefore, the answer is [[1, 2], [4, 6], [8, 8], [10, 12]].
    fmt.Println(findDisappearedNumbers([]int{3,9,7}, 1, 12)) // [[1,2],[4,6],[8,8],[10,12]]
    // Example 2:
    // Input: nums = [1,1], lower = 5, upper = 7
    // Output: [[5,7]]
    // Explanation:
    // The missing integers are [5, 6, 7].
    // Grouping the missing integers into the minimum number of contiguous ranges, we get [5, 7].
    // Therefore, the answer is [[5, 7]].
    fmt.Println(findDisappearedNumbers([]int{1,1}, 5, 7)) // [[5,7]]
    // Example 3:
    // Input: nums = [2,3,5], lower = 2, upper = 3
    // Output: []
    // Explanation:
    // There are no missing integers.
    // Therefore, the answer is [].
    fmt.Println(findDisappearedNumbers([]int{2,3,5}, 2, 3)) // []

    fmt.Println(findDisappearedNumbers([]int{1,2,3,4,5,6,7,8,9}, 2, 3)) // []
    fmt.Println(findDisappearedNumbers([]int{9,8,7,6,5,4,3,2,1}, 2, 3)) // []

    fmt.Println(findDisappearedNumbers1([]int{3,9,7}, 1, 12)) // [[1,2],[4,6],[8,8],[10,12]]
    fmt.Println(findDisappearedNumbers1([]int{1,1}, 5, 7)) // [[5,7]]
    fmt.Println(findDisappearedNumbers1([]int{2,3,5}, 2, 3)) // []
    fmt.Println(findDisappearedNumbers1([]int{1,2,3,4,5,6,7,8,9}, 2, 3)) // []
    fmt.Println(findDisappearedNumbers1([]int{9,8,7,6,5,4,3,2,1}, 2, 3)) // []
}