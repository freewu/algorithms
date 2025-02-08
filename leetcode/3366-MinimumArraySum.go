package main

// 3366. Minimum Array Sum
// You are given an integer array nums and three integers k, op1, and op2.

// You can perform the following operations on nums:
//     1. Operation 1: Choose an index i and divide nums[i] by 2, rounding up to the nearest whole number. 
//        You can perform this operation at most op1 times, and not more than once per index.
//     2. Operation 2: Choose an index i and subtract k from nums[i], but only if nums[i] is greater than or equal to k. 
//        You can perform this operation at most op2 times, and not more than once per index.

// Note: Both operations can be applied to the same index, but at most once each.

// Return the minimum possible sum of all elements in nums after performing any number of operations.

// Example 1:
// Input: nums = [2,8,3,19,3], k = 3, op1 = 1, op2 = 1
// Output: 23
// Explanation:
// Apply Operation 2 to nums[1] = 8, making nums[1] = 5.
// Apply Operation 1 to nums[3] = 19, making nums[3] = 10.
// The resulting array becomes [2, 5, 3, 10, 3], which has the minimum possible sum of 23 after applying the operations.

// Example 2:
// Input: nums = [2,4,3], k = 3, op1 = 2, op2 = 1
// Output: 3
// Explanation:
// Apply Operation 1 to nums[0] = 2, making nums[0] = 1.
// Apply Operation 1 to nums[1] = 4, making nums[1] = 2.
// Apply Operation 2 to nums[2] = 3, making nums[2] = 0.
// The resulting array becomes [1, 2, 0], which has the minimum possible sum of 3 after applying the operations.

// Constraints:
//     1 <= nums.length <= 100
//     0 <= nums[i] <= 10^5
//     0 <= k <= 10^5
//     0 <= op1, op2 <= nums.length

import "fmt"
import "sort"

func minArraySum(nums []int, d int, op1 int, op2 int) int {
    n, inf := len(nums), 1 << 31
    dp := make([][][]int, n + 1)
    for i := range dp {
        dp[i] = make([][]int, op1 + 1)
        for j := range dp[i] {
            dp[i][j] = make([]int, op2 + 1)
            for k := range dp[i][j] {
                dp[i][j][k] = inf
            }
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    dp[0][0][0] = 0
    for i := 1; i <= n; i++ {
        v := nums[i - 1]
        for j := 0; j <= op1; j++ {
            for k := 0; k <= op2; k++ {
                dp[i][j][k] = dp[i-1][j][k] + v
                if j > 0 {
                    dp[i][j][k] = min(dp[i][j][k], dp[i-1][j-1][k] + (v + 1)/2)
                }
                if k > 0 && v >= d {
                    dp[i][j][k] = min(dp[i][j][k], dp[i-1][j][k-1] + (v - d))
                }
                if j > 0 && k > 0 {
                    y := (v + 1) / 2
                    if y >= d {
                        dp[i][j][k] = min(dp[i][j][k], dp[i-1][j-1][k-1] + (y - d))
                    }
                    if v >= d {
                        dp[i][j][k] = min(dp[i][j][k], dp[i-1][j-1][k-1] + (v - d + 1) / 2)
                    }
                }
            }
        }
    }
    res := inf
    for j := 0; j <= op1; j++ {
        for k := 0; k <= op2; k++ {
            res = min(res, dp[n][j][k])
        }
    }
    return res
}

func minArraySum1(nums []int, k, op1, op2 int) int {
    sort.Ints(nums)
    high, low := sort.SearchInts(nums, k * 2 - 1), sort.SearchInts(nums, k)
    // 在 [2k-1,∞) 中的数，直接先除再减（从大到小操作）
    for i := len(nums) - 1; i >= high; i-- {
        if op1 > 0 {
            nums[i] = (nums[i] + 1) / 2
            op1--
        }
        if op2 > 0 {
            nums[i] -= k
            op2--
        }
    }
    // 在 [k,2k-2] 中的数，先把小的数 -k
    count, odd := make(map[int]int), 0
    for i := low; i < high; i++ {
        if op2 > 0 {
            nums[i] -= k
            if k % 2 > 0 && nums[i] % 2 > 0 {
                // nums[i] 原来是偶数，后面有机会把这次 -k 操作留给奇数，得到更小的答案
                count[nums[i]]++
            }
            op2--
        } else {
            odd += nums[i] % 2 // 没有执行 -k 的奇数
        }
    }
    // 重新排序（注：这里可以改用合并两个有序数组的做法）
    sort.Ints(nums[:high])
    res := 0
    if k % 2 > 0 {
        // 调整，对于 [k,2k-2] 中 -k 后还要再 /2 的数，如果原来是偶数，改成给奇数 -k 再 /2，这样答案可以减一
        for i := high - op1; i < high && odd > 0; i++ {
            x := nums[i]
            if count[x] > 0 {
                count[x]--
                if count[x] == 0 {
                    delete(count, x)
                }
                odd--
                res--
            }
        }
    }
    // 最后，从大到小执行操作 1
    for i := high - 1; i >= 0 && op1 > 0; i-- {
        nums[i] = (nums[i] + 1) / 2
        op1--
    }
    for _, v := range nums {
        res += v
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,8,3,19,3], k = 3, op1 = 1, op2 = 1
    // Output: 23
    // Explanation:
    // Apply Operation 2 to nums[1] = 8, making nums[1] = 5.
    // Apply Operation 1 to nums[3] = 19, making nums[3] = 10.
    // The resulting array becomes [2, 5, 3, 10, 3], which has the minimum possible sum of 23 after applying the operations.
    fmt.Println(minArraySum([]int{2,8,3,19,3}, 3, 1, 1)) // 23
    // Example 2:
    // Input: nums = [2,4,3], k = 3, op1 = 2, op2 = 1
    // Output: 3
    // Explanation:
    // Apply Operation 1 to nums[0] = 2, making nums[0] = 1.
    // Apply Operation 1 to nums[1] = 4, making nums[1] = 2.
    // Apply Operation 2 to nums[2] = 3, making nums[2] = 0.
    // The resulting array becomes [1, 2, 0], which has the minimum possible sum of 3 after applying the operations.
    fmt.Println(minArraySum([]int{2,4,3}, 3, 2, 1)) // 3

    fmt.Println(minArraySum([]int{1,2,3,4,5,6,7,8,9}, 3, 2, 1)) // 34
    fmt.Println(minArraySum([]int{9,8,7,6,5,4,3,2,1}, 3, 2, 1)) // 34

    fmt.Println(minArraySum1([]int{2,8,3,19,3}, 3, 1, 1)) // 23
    fmt.Println(minArraySum1([]int{2,4,3}, 3, 2, 1)) // 3
    fmt.Println(minArraySum1([]int{1,2,3,4,5,6,7,8,9}, 3, 2, 1)) // 34
    fmt.Println(minArraySum1([]int{9,8,7,6,5,4,3,2,1}, 3, 2, 1)) // 34
}