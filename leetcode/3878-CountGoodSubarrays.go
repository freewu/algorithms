package main

// 3878. Count Good Subarrays
// You are given an integer array nums.

// A subarray is called good if the bitwise OR of all its elements is equal to at least one element present in that subarray.

// Return the number of good subarrays in nums.

// Here, the bitwise OR of two integers a and b is denoted by a | b.

// Example 1:
// Input: nums = [4,2,3]
// Output: 4
// Explanation:
// The subarrays of nums are:
// Subarray    | Bitwise OR    | Present in Subarray
// [4]         | 4 = 4	        | Yes
// [2]         | 2 = 2	        | Yes
// [3]         | 3 = 3	        | Yes
// [4, 2]      | 4 | 2 = 6	    | No
// [2, 3]      | 2 | 3 = 3	    | Yes 
// [4, 2, 3]   | 4 | 2 | 3 = 7	| No  
// Thus, the good subarrays of nums are [4], [2], [3] and [2, 3]. Thus, the answer is 4.

// Example 2:
// Input: nums = [1,3,1]
// Output: 6
// Explanation:
// Any subarray of nums containing 3 has bitwise OR equal to 3, and subarrays containing only 1 have bitwise OR equal to 1.
// In both cases, the result is present in the subarray, so all subarrays are good, and the answer is 6.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^9

import "fmt"

func countGoodSubarrays(nums []int) int64 {
    type Pair struct{ or, left int } // 子数组或值，最小左端点
    res, orLeft, last := 0, []Pair{}, map[int]int{}
    for i, v := range nums {
        last[v] = i
        // 计算以 i 为右端点的子数组或值
        for j := range orLeft {
            orLeft[j].or |= v
        }
        // v 单独一个数作为子数组
        orLeft = append(orLeft, Pair{v, i})
        // 原地去重（相同或值只保留最左边的）
        // 原理见力扣 26. 删除有序数组中的重复项
        index := 1
        for j := 1; j < len(orLeft); j++ {
            if orLeft[j].or != orLeft[j-1].or {
                orLeft[index] = orLeft[j]
                index++
            }
        }
        orLeft = orLeft[:index]
        for k, p := range orLeft {
            orVal := p.or
            left := p.left
            right := i
            if k < len(orLeft)-1 {
                right = orLeft[k+1].left - 1
            }
            // 对于左端点在 [left, right]，右端点为 i 的子数组，OR 值都是 orVal
            j, ok := last[orVal]
            if ok && j >= left {
                res += (min(right, j) - left + 1)
            }
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [4,2,3]
    // Output: 4
    // Explanation:
    // The subarrays of nums are:
    // Subarray    | Bitwise OR    | Present in Subarray
    // [4]         | 4 = 4	        | Yes
    // [2]         | 2 = 2	        | Yes
    // [3]         | 3 = 3	        | Yes
    // [4, 2]      | 4 | 2 = 6	    | No
    // [2, 3]      | 2 | 3 = 3	    | Yes 
    // [4, 2, 3]   | 4 | 2 | 3 = 7	| No  
    // Thus, the good subarrays of nums are [4], [2], [3] and [2, 3]. Thus, the answer is 4.
    fmt.Println(countGoodSubarrays([]int{4,2,3})) // 4
    // Example 2:
    // Input: nums = [1,3,1]
    // Output: 6
    // Explanation:
    // Any subarray of nums containing 3 has bitwise OR equal to 3, and subarrays containing only 1 have bitwise OR equal to 1.
    // In both cases, the result is present in the subarray, so all subarrays are good, and the answer is 6.
    fmt.Println(countGoodSubarrays([]int{1,3,1}))  // 6

    fmt.Println(countGoodSubarrays([]int{1,2,3,4,5,6,7,8,9})) // 19
    fmt.Println(countGoodSubarrays([]int{9,8,7,6,5,4,3,2,1})) // 19
}