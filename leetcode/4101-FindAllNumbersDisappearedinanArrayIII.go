package main

// 4101. Find All Numbers Disappeared in an Array III
// You are given an integer array nums and two integers lower and upper.

// An integer is missing if it lies in the inclusive range [lower, upper] and does not appear in nums.

// Return a 2D integer array representing all missing integers grouped into contiguous ranges. 
// Represent each range as [start, end], including both endpoints, and return the ranges in increasing order of start.

// Consecutive missing integers must belong to the same range. 
// If there are no missing integers, return an empty array.

// Example 1:
// Input: nums = [3,9,7], lower = 1, upper = 12
// Output: [[1,2],[4,6],[8,8],[10,12]]
// Explanation:
// - The missing integers are [1, 2, 4, 5, 6, 8, 10, 11, 12].
// - Grouping the missing integers into the minimum number of contiguous ranges, we get [1, 2], [4, 6], [8, 8], and [10, 12].
// - Therefore, the answer is [[1, 2], [4, 6], [8, 8], [10, 12]].

// Example 2:
// Input: nums = [1,1,1000000000], lower = 5, upper = 1000000000
// Output: [[5,999999999]]
// Explanation:
// - The value 1 lies outside [lower, upper], so it does not cover any integer in the range.
// - The only integer of the range that appears in nums is 1000000000.
// - Every integer from 5 to 999999999 is missing, and they are consecutive, so they form the single range [5, 999999999].

// Example 3:
// Input: nums = [2,3,5], lower = 2, upper = 3
// Output: []
// Explanation:
// - Both 2 and 3 appear in nums, so there are no missing integers.
// - Therefore, the answer is [].

// Constraints:
//     - 1 <= nums.length <= 10^{5}
//     - 1 <= nums[i] <= 10^{9}
//     - 1 <= lower <= upper <= 10^{9}

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
    // - The missing integers are [1, 2, 4, 5, 6, 8, 10, 11, 12].
    // - Grouping the missing integers into the minimum number of contiguous ranges, we get [1, 2], [4, 6], [8, 8], and [10, 12].
    // - Therefore, the answer is [[1, 2], [4, 6], [8, 8], [10, 12]].
    fmt.Println(findDisappearedNumbers([]int{3,9,7}, 1, 12)) // [[1,2],[4,6],[8,8],[10,12]]
    // Example 2:
    // Input: nums = [1,1,1000000000], lower = 5, upper = 1000000000
    // Output: [[5,999999999]]
    // Explanation:
    // - The value 1 lies outside [lower, upper], so it does not cover any integer in the range.
    // - The only integer of the range that appears in nums is 1000000000.
    // - Every integer from 5 to 999999999 is missing, and they are consecutive, so they form the single range [5, 999999999].
    fmt.Println(findDisappearedNumbers([]int{1,1,1000000000}, 5, 1000000000)) // [[5,999999999]]
    // Example 3:
    // Input: nums = [2,3,5], lower = 2, upper = 3
    // Output: []
    // Explanation:
    // - Both 2 and 3 appear in nums, so there are no missing integers.
    // - Therefore, the answer is [].
    fmt.Println(findDisappearedNumbers([]int{2,3,5}, 2, 3)) // []

    fmt.Println(findDisappearedNumbers([]int{1,2,3,4,5,6,7,8,9}, 2, 3)) // []
    fmt.Println(findDisappearedNumbers([]int{9,8,7,6,5,4,3,2,1}, 2, 3)) // []
    fmt.Println(findDisappearedNumbers([]int{1,2,3,4,5,6,7,8,9}, 1, 1)) // []
    fmt.Println(findDisappearedNumbers([]int{9,8,7,6,5,4,3,2,1}, 1, 1)) // []
    fmt.Println(findDisappearedNumbers([]int{1,2,3,4,5,6,7,8,9}, 1, 1_000_000_000)) // [[10 1000000000]]
    fmt.Println(findDisappearedNumbers([]int{9,8,7,6,5,4,3,2,1}, 1, 1_000_000_000)) // [[10 1000000000]]
    fmt.Println(findDisappearedNumbers([]int{1,2,3,4,5,6,7,8,9}, 1_000_000_000, 1_000_000_000)) // [[1000000000 1000000000]]
    fmt.Println(findDisappearedNumbers([]int{9,8,7,6,5,4,3,2,1}, 1_000_000_000, 1_000_000_000)) // [[1000000000 1000000000]]

    fmt.Println(findDisappearedNumbers1([]int{3,9,7}, 1, 12)) // [[1,2],[4,6],[8,8],[10,12]]
    fmt.Println(findDisappearedNumbers1([]int{1,1,1000000000}, 5, 1000000000)) // [[5,999999999]]
    fmt.Println(findDisappearedNumbers1([]int{2,3,5}, 2, 3)) // []
    fmt.Println(findDisappearedNumbers1([]int{1,2,3,4,5,6,7,8,9}, 2, 3)) // []
    fmt.Println(findDisappearedNumbers1([]int{9,8,7,6,5,4,3,2,1}, 2, 3)) // []
    fmt.Println(findDisappearedNumbers1([]int{1,2,3,4,5,6,7,8,9}, 1, 1)) // []
    fmt.Println(findDisappearedNumbers1([]int{9,8,7,6,5,4,3,2,1}, 1, 1)) // []
    fmt.Println(findDisappearedNumbers1([]int{1,2,3,4,5,6,7,8,9}, 1, 1_000_000_000)) // [[10 1000000000]]
    fmt.Println(findDisappearedNumbers1([]int{9,8,7,6,5,4,3,2,1}, 1, 1_000_000_000)) // [[10 1000000000]]
    fmt.Println(findDisappearedNumbers1([]int{1,2,3,4,5,6,7,8,9}, 1_000_000_000, 1_000_000_000)) // [[1000000000 1000000000]]
    fmt.Println(findDisappearedNumbers1([]int{9,8,7,6,5,4,3,2,1}, 1_000_000_000, 1_000_000_000)) // [[1000000000 1000000000]]
}