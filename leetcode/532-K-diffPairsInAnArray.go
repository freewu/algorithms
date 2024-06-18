package main

// 532. K-diff Pairs in an Array
// Given an array of integers nums and an integer k, return the number of unique k-diff pairs in the array.
// A k-diff pair is an integer pair (nums[i], nums[j]), where the following are true:
//     0 <= i, j < nums.length
//     i != j
//     |nums[i] - nums[j]| == k

// Notice that |val| denotes the absolute value of val.

// Example 1:
// Input: nums = [3,1,4,1,5], k = 2
// Output: 2
// Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
// Although we have two 1s in the input, we should only return the number of unique pairs.

// Example 2:
// Input: nums = [1,2,3,4,5], k = 1
// Output: 4
// Explanation: There are four 1-diff pairs in the array, (1, 2), (2, 3), (3, 4) and (4, 5).

// Example 3:
// Input: nums = [1,3,1,5,4], k = 0
// Output: 1
// Explanation: There is one 0-diff pair in the array, (1, 1).

// Constraints:
//     1 <= nums.length <= 10^4
//     -10^7 <= nums[i] <= 10^7
//     0 <= k <= 10^7

import "fmt"

func findPairs(nums []int, k int) int {
    if k < 0 || len(nums) == 0 {
        return 0
    }
    res, mp := 0, make(map[int]int, len(nums))
    for _, v := range nums { // 统计出现的次数
        mp[v]++
    }
    for key := range mp {
        if k == 0 && mp[key] > 1 { // 如果 k 要求为 0 需要找到出现了两次的 组成如: (2,2) (5,5)
            res++
            continue
        }
        if k > 0 && mp[key + k] > 0 { // 找到了 key + k 的数 如 key为 1 k 为2  (1, 3)  key 为 4 k 为 2 (4, 6)
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,1,4,1,5], k = 2
    // Output: 2
    // Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
    // Although we have two 1s in the input, we should only return the number of unique pairs.
    fmt.Println(findPairs([]int{3,1,4,1,5}, 2)) // 2
    // Example 2:
    // Input: nums = [1,2,3,4,5], k = 1
    // Output: 4
    // Explanation: There are four 1-diff pairs in the array, (1, 2), (2, 3), (3, 4) and (4, 5).
    fmt.Println(findPairs([]int{1,2,3,4,5}, 1)) // 4
    // Example 3:
    // Input: nums = [1,3,1,5,4], k = 0
    // Output: 1
    // Explanation: There is one 0-diff pair in the array, (1, 1).
    fmt.Println(findPairs([]int{1,3,1,5,4}, 0)) // 1
}