package main

// 面试题 16.24. Pairs With Sum LCCI
// Design an algorithm to find all pairs of integers within an array which sum to a specified value.

// Example 1:
// Input: nums = [5,6,5], target = 11
// Output: [[5,6]]

// Example 2:
// Input: nums = [5,6,5,6], target = 11
// Output: [[5,6],[5,6]]

// Note:
//     nums.length <= 100000
//     -10^5 <= nums[i], target <= 10^5

import "fmt"
import "sort"

func pairSums(nums []int, target int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    for i, j := 0, len(nums) - 1; i < j; {
        if nums[i] + nums[j] == target {
            res = append(res, []int{nums[i], nums[j]})
            i++
            j--
        } else if nums[i] + nums[j] > target {
            j--
        } else if nums[i] + nums[j] < target {
            i++
        }
    }
    return res
}

func pairSums1(nums []int, target int) [][]int {
    limit, n := 100000, len(nums)
    res, arr := make([][]int, 0, n / 2 + 1), make([]int, 2 * limit + 1)
    for _, v := range nums {
        find := target - v
        index := find + limit
        if arr[index] > 0 {
            res = append(res, []int{ v, find })
            arr[index]--
        } else {
            arr[v + limit]++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,6,5], target = 11
    // Output: [[5,6]]
    fmt.Println(pairSums([]int{5,6,5}, 11)) // [[5,6]]
    // Example 2:
    // Input: nums = [5,6,5,6], target = 11
    // Output: [[5,6],[5,6]]
    fmt.Println(pairSums([]int{5,6,5,6}, 11)) // [[5,6],[5,6]]

    fmt.Println(pairSums([]int{1,2,3,4,5,6,7,8,9}, 11)) // [[2 9] [3 8] [4 7] [5 6]]
    fmt.Println(pairSums([]int{9,8,7,6,5,4,3,2,1}, 11)) // [[2 9] [3 8] [4 7] [5 6]]

    fmt.Println(pairSums1([]int{5,6,5}, 11)) // [[5,6]]
    fmt.Println(pairSums1([]int{5,6,5,6}, 11)) // [[5,6],[5,6]]
    fmt.Println(pairSums1([]int{1,2,3,4,5,6,7,8,9}, 11)) // [[2 9] [3 8] [4 7] [5 6]]
    fmt.Println(pairSums1([]int{9,8,7,6,5,4,3,2,1}, 11)) // [[2 9] [3 8] [4 7] [5 6]]
}