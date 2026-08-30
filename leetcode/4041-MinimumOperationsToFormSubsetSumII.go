package main 

// 4041. Minimum Operations to Form Subset Sum II
// You are given an integer array nums and an integer sum.

// In one operation, choose an element with current value x and replace it with either 2 * x or floor(x / 2).

// For each element, multiplication and division operations may be performed in any order.

// Return the minimum number of operations needed so that some subset of the resulting array has a sum exactly equal to sum. 
// If it is impossible, return -1.

// The floor() function returns the integer part of the division.

// Example 1:
// Input: nums = [10,2], sum = 13
// Output: 3
// Explanation:
// Divide nums[0] = 10 once: 10 → 5, costing 1 operation.
// Multiply nums[1] = 2 twice: 2 → 4 → 8, costing 2 operations.
// After these operations, nums = [5, 8]. The subset {5, 8} sums to 13 using 3 operations in total.

// Example 2:
// Input: nums = [6,3], sum = 8
// Output: 2
// Explanation:​​​​​​​
// Turn nums[1] = 3 into 2 using 2 operations:
// Divide nums[1] to get 1.
// Multiply nums[1] = 1 to get 2.
// After these operations, nums = [6, 2]. The subset {6, 2} sums to 8 using 2 operations in total.

// Example 3:
// Input: nums = [2,2], sum = 7
// Output: -1
// Explanation:
// No sequence of operations lets a subset of nums sum to 7, so the answer is -1.
 
// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 500
//     1 <= sum <= 5000

import "fmt"

// // Wrong Answer 965 / 999 testcases passed
// func minOperations(nums[]int,sum int) int {
//     inf := 1 << 61
//     dp := make([]int,sum + 1)
//     for i := 1; i <= sum; i++ {
//         dp[i] = inf
//     }
//     for _, x := range nums {
//         mp := make(map[int]int)
//         for v,c := x, 0; v <= sum; v,c = v * 2, c + 1 {
//             if p,ok:=mp[v];!ok||c<p{
//                 mp[v]=c
//             }
//         }
//         for v,c := x,0; v >= 1; v,c = v / 2,c + 1 {
//             if p,ok := mp[v]; !ok || c < p {
//                 mp[v]=c
//             }
//         }
//         nd := append([]int{},dp...)
//         for v,c := range mp {
//             for j := sum; j >= v; j-- {
//                 if dp[j-v] + c < nd[j] {
//                     nd[j] = dp[j-v] + c
//                 }
//             }
//         }
//         dp = nd
//     }
//     if dp[sum] >= inf {
//         return -1
//     }
//     return dp[sum]
// }

func minOperations(nums []int, sum int) int {
    f := make([]int, sum+1)
    for i := 1; i <= sum; i++ {
        f[i] = 1 << 61 // 避免加法溢出
    }
    for _, x := range nums {
        for i := sum; i > 0; i-- {
            // 回想一下，0-1 背包是选或不选，状态转移方程为 f[i] = min(f[i], f[i-物品体积] + 物品价值)
            // 本题是分组背包，要枚举选哪个物品（枚举除法操作次数为 a，乘法操作次数为 b）
            for a := 0; x>>a > 0; a++ {
                for b := 0; x>>a<<b <= i; b++ {
                    // 物品体积为 x>>a<<b，价值为 a+b
                    f[i] = min(f[i], f[i-x>>a<<b]+a+b)
                }
            }
        }
        if f[sum] == 0 { // 优化：提前返回
            return 0
        }
    }
    if f[sum] == 1 << 61 {
        return -1
    }
    return f[sum]
}

func main() {
    // Example 1:
    // Input: nums = [10,2], sum = 13
    // Output: 3
    // Explanation:
    // Divide nums[0] = 10 once: 10 → 5, costing 1 operation.
    // Multiply nums[1] = 2 twice: 2 → 4 → 8, costing 2 operations.
    // After these operations, nums = [5, 8]. The subset {5, 8} sums to 13 using 3 operations in total.
    fmt.Println(minOperations([]int{10,2}, 13)) // 3
    // Example 2:
    // Input: nums = [6,3], sum = 8
    // Output: 2
    // Explanation:​​​​​​​
    // Turn nums[1] = 3 into 2 using 2 operations:
    // Divide nums[1] to get 1.
    // Multiply nums[1] = 1 to get 2.
    // After these operations, nums = [6, 2]. The subset {6, 2} sums to 8 using 2 operations in total.
    fmt.Println(minOperations([]int{6,3}, 8)) // 2
    // Example 3:
    // Input: nums = [2,2], sum = 7
    // Output: -1
    // Explanation:
    // No sequence of operations lets a subset of nums sum to 7, so the answer is -1.
    fmt.Println(minOperations([]int{2,2}, 7)) // -1

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9}, 4)) // 0
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1}, 4)) // 0
}