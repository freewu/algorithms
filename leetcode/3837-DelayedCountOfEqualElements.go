package main

// 3837. Delayed Count of Equal Elements
// You are given an integer array nums of length n and an integer k.

// For each index i, define the delayed count as the number of indices j such that:
//     1. i + k < j <= n - 1, and
//     2. nums[j] == nums[i]

// Return an array ans where ans[i] is the delayed count of index i.

// Example 1:
// Input: nums = [1,2,1,1], k = 1
// Output: [2,0,0,0]
// Explanation:
// i   | nums[i]   | possible j    | nums[j]   | satisfying(nums[j] == nums[i])    | ans[i]
// 0	| 1	        | [2, 3]	    | [1, 1]	| [2, 3]	                        | 2
// 1	| 2	        | [3]	        | [1]	    | []	                            | 0
// 2	| 1	        | []	        | []	    | []	                            | 0
// 3	| 1	        | []	        | []	    | []	                            | 0   
// Thus, ans = [2, 0, 0, 0]​​​​​​​.

// Example 2:
// Input: nums = [3,1,3,1], k = 0
// Output: [1,1,0,0]
// Explanation:
// i   | nums[i]   | possible j    | nums[j]   | satisfying(nums[j] == nums[i])    | ans[i]
// 0	| 3	        | [1, 2, 3]	    | [1, 3, 1]	| [2]	                            | 1
// 1	| 1	        | [2, 3]	    | [3, 1]	| [3]	                            | 1
// 2	| 3	        | [3]	        | [1]	    | []	                            | 0
// 3	| 1	        | []	        | []	    | []	                            | 0
// Thus, ans = [1, 1, 0, 0]​​​​​​​.

// Constraints:
//     1 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     0 <= k <= n - 1

import "fmt"

func delayedCount(nums []int, k int) []int {
    n := len(nums)
    res, mp := make([]int, n), make(map[int][]int)
    for i, v := range nums { // 预处理：统计每个数字的所有出现位置
        mp[v] = append(mp[v], i)
    }
    for i := 0; i < n; i++ { // 对每个i，二分查找 mp[nums[i]] 中 > i + k 的元素数量
        pos := mp[nums[i]]
        // 二分查找第一个 > i+k 的位置
        left, right := 0, len(pos)
        for left < right {
            mid := (left + right) / 2
            if pos[mid] > i + k {
                right = mid
            } else {
                left = mid + 1
            }
        }
        res[i] = len(pos) - left  // 数量 = 总长度 - left
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1,1], k = 1
    // Output: [2,0,0,0]
    // Explanation:
    // i   | nums[i]   | possible j    | nums[j]   | satisfying(nums[j] == nums[i])    | ans[i]
    // 0	| 1	        | [2, 3]	    | [1, 1]	| [2, 3]	                        | 2
    // 1	| 2	        | [3]	        | [1]	    | []	                            | 0
    // 2	| 1	        | []	        | []	    | []	                            | 0
    // 3	| 1	        | []	        | []	    | []	                            | 0   
    // Thus, ans = [2, 0, 0, 0]​​​​​​​.
    fmt.Println(delayedCount([]int{1,2,1,1}, 1)) // [2, 0, 0, 0]
    // Example 2:
    // Input: nums = [3,1,3,1], k = 0
    // Output: [1,1,0,0]
    // Explanation:
    // i   | nums[i]   | possible j    | nums[j]   | satisfying(nums[j] == nums[i])    | ans[i]
    // 0	| 3	        | [1, 2, 3]	    | [1, 3, 1]	| [2]	                            | 1
    // 1	| 1	        | [2, 3]	    | [3, 1]	| [3]	                            | 1
    // 2	| 3	        | [3]	        | [1]	    | []	                            | 0
    // 3	| 1	        | []	        | []	    | []	                            | 0
    // Thus, ans = [1, 1, 0, 0]​​​​​​​.
    fmt.Println(delayedCount([]int{3,1,3,1}, 0)) // [1, 1, 0, 0]

    fmt.Println(delayedCount([]int{1,2,3,4,5,6,7,8,9}, 1)) // [0 0 0 0 0 0 0 0 0]
    fmt.Println(delayedCount([]int{9,8,7,6,5,4,3,2,1}, 1)) // [0 0 0 0 0 0 0 0 0]
}