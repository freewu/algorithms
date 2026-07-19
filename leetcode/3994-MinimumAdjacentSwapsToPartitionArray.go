package main

// 3994. Minimum Adjacent Swaps to Partition Array
// You are given an integer array nums and two integers a and b such that a < b.

// An array is called good if it can be split into three contiguous parts, in this order, such that:
//     1. Every element in the first part is less than a.
//     2. Every element in the second part is in the range [a, b] inclusive.
//     3. Every element in the third part is greater than b.

// Any of the three parts may be empty.

// In one adjacent swap, you may swap two neighboring elements of nums.

// Return the minimum number of adjacent swaps required to make nums good. 
// Since the answer may be very large, return it modulo 10^9 + 7.
 
// Example 1:
// Input: nums = [1,3,2,4,5,6], a = 3, b = 4
// Output: 1
// Explanation:
// Swap nums[1] and nums[2]. The array becomes [1, 2, 3, 4, 5, 6].
// This array is good because it can be split into [1, 2], [3, 4], and [5, 6].

// Example 2:
// Input: nums = [9,7,5,3], a = 4, b = 8
// Output: 5
// Explanation:
// One sequence of optimal swaps is as follows:
// Swap nums[2] and nums[3]. The array becomes [9, 7, 3, 5].
// Swap nums[1] and nums[2]. The array becomes [9, 3, 7, 5].
// Swap nums[0] and nums[1]. The array becomes [3, 9, 7, 5].
// Swap nums[1] and nums[2]. The array becomes [3, 7, 9, 5].
// Swap nums[2] and nums[3]. The array becomes [3, 7, 5, 9].
// This array is good because it can be split into [3], [7, 5], and [9].

// Example 3:
// Input: nums = [3,7,5,9], a = 4, b = 8
// Output: 0
// Explanation:
// The array is already good. No swaps are needed.

// Constraints:
//     1 <= nums.length <= 10^5
//     ​​​​​​​1 <= nums[i] <= 10^9
//     1 <= a < b <= 10^9

import "fmt"

func minAdjacentSwaps(nums []int, a int, b int) int {
    res, mod := 0, 1_000_000_007
    p1, p2, p3 := 0, 0, 0
    for _, v := range nums {
        if v < a {
            res = (res + p2 + p3) % mod
            p1++
        } else if b < v {
            p3++
        } else {
            res = (res + p3) % mod
            p2++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,3,2,4,5,6], a = 3, b = 4
    // Output: 1
    // Explanation:
    // Swap nums[1] and nums[2]. The array becomes [1, 2, 3, 4, 5, 6].
    // This array is good because it can be split into [1, 2], [3, 4], and [5, 6].
    fmt.Println(minAdjacentSwaps([]int{1,3,2,4,5,6}, 3, 4)) // 1
    // Example 2:
    // Input: nums = [9,7,5,3], a = 4, b = 8
    // Output: 5
    // Explanation:
    // One sequence of optimal swaps is as follows:
    // Swap nums[2] and nums[3]. The array becomes [9, 7, 3, 5].
    // Swap nums[1] and nums[2]. The array becomes [9, 3, 7, 5].
    // Swap nums[0] and nums[1]. The array becomes [3, 9, 7, 5].
    // Swap nums[1] and nums[2]. The array becomes [3, 7, 9, 5].
    // Swap nums[2] and nums[3]. The array becomes [3, 7, 5, 9].
    // This array is good because it can be split into [3], [7, 5], and [9].
    fmt.Println(minAdjacentSwaps([]int{9,7,5,3}, 4, 8)) // 5
    // Example 3:
    // Input: nums = [3,7,5,9], a = 4, b = 8
    // Output: 0
    // Explanation:
    // The array is already good. No swaps are needed.
    fmt.Println(minAdjacentSwaps([]int{3,7,5,9}, 4, 8)) // 0

    fmt.Println(minAdjacentSwaps([]int{1,2,3,4,5,6,7,8,9}, 4, 8)) // 0
    fmt.Println(minAdjacentSwaps([]int{9,8,7,6,5,4,3,2,1}, 4, 8)) // 23
}