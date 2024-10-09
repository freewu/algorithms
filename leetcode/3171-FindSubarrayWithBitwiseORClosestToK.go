package main

// 3171. Find Subarray With Bitwise OR Closest to K
// You are given an array nums and an integer k. 
// You need to find a subarray of nums such that the absolute difference between k and the bitwise OR of the subarray elements is as small as possible. 
// In other words, select a subarray nums[l..r] such that |k - (nums[l] OR nums[l + 1] ... OR nums[r])| is minimum.

// Return the minimum possible value of the absolute difference.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,2,4,5], k = 3
// Output: 0
// Explanation:
// The subarray nums[0..1] has OR value 3, which gives the minimum absolute difference |3 - 3| = 0.

// Example 2:
// Input: nums = [1,3,1,3], k = 2
// Output: 1
// Explanation:
// The subarray nums[1..1] has OR value 3, which gives the minimum absolute difference |3 - 2| = 1.

// Example 3:
// Input: nums = [1], k = 10
// Output: 9
// Explanation:
// There is a single subarray with OR value 1, which gives the minimum absolute difference |10 - 1| = 9.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= k <= 10^9

import "fmt"
import "sort"

// func minimumDifference(nums []int, k int) int {
//     res, n := 1 << 31, len(nums)
//     abs := func(x int) int { if x < 0 { return -x; }; return x; }
//     min := func (x, y int) int { if x < y { return x; }; return y; }
//     for i := 0; i < n; i++ {
//         for  j := i; j >= 0; j-- {
//             val := nums[i] & nums[j]
//             res = min(res, abs(k - val))
//             if i != j && val == nums[j] { break }
//             nums[j] = val
//         }
//     }
//     return res
// }

// func minimumDifference(nums []int, k int) int {
//     res, dp, n := 1 << 31, make([]int,100001), len(nums)
//     abs := func(x int) int { if x < 0 { return -x; }; return x; }
//     min := func (x, y int) int { if x < y { return x; }; return y; }
//     for i := 0; i < n; i++ {
//         for j := i; j < n; j++ {
//             val := nums[i]
//             val &= nums[j]
//             res = min(res, abs(k - val))
//             if val <= k || val == dp[j] { break }
//             dp[j] = val
//         }
//     }
//     return res
// }

// func minimumDifference(nums []int, k int) int {
//     res := 1 << 31
//     subArrayANDValues := make(map[int]bool)
//     abs := func(x int) int { if x < 0 { return -x; }; return x; }
//     min := func (x, y int) int { if x < y { return x; }; return y; }
//     for _, v := range nums {
//         newSubArrayANDValues := make(map[int]bool)
//         newSubArrayANDValues[v] = true
//         res = min(res, abs(v - k))
//         for v1, _ := range subArrayANDValues {
//             val := v1 & v
//             if _, ok := newSubArrayANDValues[val]; !ok {
//                 newSubArrayANDValues[val] = true
//                 res = min(res, abs(val - k))
//             }
//         }
//         subArrayANDValues = newSubArrayANDValues
//     }
//     return res
// }

func minimumDifference(nums []int, k int) int {
    res, n, bitsMaxPos := 1 << 31, len(nums), make([]int, 31)
    for i := range bitsMaxPos {
        bitsMaxPos[i] = -1
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < n; i++ {
        for j := 0; j <= 30; j++ {
            if nums[i]>>j & 1 == 1 {
                bitsMaxPos[j] = i
            }
        }
        posToBit := make([][2]int, 0)
        for j := 0; j <= 30; j++ {
            if bitsMaxPos[j] != -1 {
                posToBit = append(posToBit, [2]int{bitsMaxPos[j], j})
            }
        }
        sort.Slice(posToBit, func(a, b int) bool {
            return posToBit[a][0] > posToBit[b][0]
        })
        val := 0
        for j, p := 0, 0; j < len(posToBit); p = j {
            for j < len(posToBit) && posToBit[j][0] == posToBit[p][0] {
                val |= 1 << posToBit[j][1]
                j++
            }
            res = min(res, abs(val - k))
        }
    }
    return res
}

func minimumDifference1(nums []int, k int) int {
    res, left, bottom, rightOr := 1 << 31, 0, 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for right, v := range nums {
        rightOr |= v
        for left <= right && nums[left] | rightOr > k {
            res = min(res, (nums[left] | rightOr) - k)
            if bottom <= left {
                // 重新构建一个栈
                // 由于 left 即将移出窗口，只需计算到 left+1
                for i := right - 1; i > left; i-- {
                    nums[i] |= nums[i+1]
                }
                bottom = right
                rightOr = 0
            }
            left++
        }
        if left <= right {
            res = min(res, k - (nums[left] | rightOr))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,4,5], k = 3
    // Output: 0
    // Explanation:
    // The subarray nums[0..1] has OR value 3, which gives the minimum absolute difference |3 - 3| = 0.
    fmt.Println(minimumDifference([]int{1,2,4,5}, 3)) // 0
    // Example 2:
    // Input: nums = [1,3,1,3], k = 2
    // Output: 1
    // Explanation:
    // The subarray nums[1..1] has OR value 3, which gives the minimum absolute difference |3 - 2| = 1.
    fmt.Println(minimumDifference([]int{1,3,1,3}, 2)) // 1
    // Example 3:
    // Input: nums = [1], k = 10
    // Output: 9
    // Explanation:
    // There is a single subarray with OR value 1, which gives the minimum absolute difference |10 - 1| = 9.
    fmt.Println(minimumDifference([]int{1}, 10)) // 9

    fmt.Println(minimumDifference1([]int{1,2,4,5}, 3)) // 0
    fmt.Println(minimumDifference1([]int{1,3,1,3}, 2)) // 1
    fmt.Println(minimumDifference1([]int{1}, 10)) // 9
}