package main

// 3685. Subsequence Sum After Capping Elements
// You are given an integer array nums of size n and a positive integer k.

// An array capped by value x is obtained by replacing every element nums[i] with min(nums[i], x).

// For each integer x from 1 to n, determine whether it is possible to choose a subsequence from the array capped by x such that the sum of the chosen elements is exactly k.

// Return a 0-indexed boolean array answer of size n, where answer[i] is true if it is possible when using x = i + 1, and false otherwise.

// Example 1:
// Input: nums = [4,3,2,4], k = 5
// Output: [false,false,true,true]
// Explanation:
// For x = 1, the capped array is [1, 1, 1, 1]. Possible sums are 1, 2, 3, 4, so it is impossible to form a sum of 5.
// For x = 2, the capped array is [2, 2, 2, 2]. Possible sums are 2, 4, 6, 8, so it is impossible to form a sum of 5.
// For x = 3, the capped array is [3, 3, 2, 3]. A subsequence [2, 3] sums to 5, so it is possible.
// For x = 4, the capped array is [4, 3, 2, 4]. A subsequence [3, 2] sums to 5, so it is possible.

// Example 2:
// Input: nums = [1,2,3,4,5], k = 3
// Output: [true,true,true,true,true]
// Explanation:
// For every value of x, it is always possible to select a subsequence from the capped array that sums exactly to 3.

// Constraints:
//     1 <= n == nums.length <= 4000
//     1 <= nums[i] <= n
//     1 <= k <= 4000

import "fmt"
import "sort"
import "math/big"

func subsequenceSumAfterCapping(nums []int, k int) []bool {
    sort.Ints(nums)
    n, i, f := len(nums), 0, big.NewInt(1)
    res := make([]bool, n)
    u := new(big.Int).Lsh(big.NewInt(1), uint(k + 1))
    u.Sub(u, big.NewInt(1))
    for x := 1; x <= n; x++ {
        // 增量地考虑所有等于 x 的数
        for i < n && nums[i] == x {
            shifted := new(big.Int).Lsh(f, uint(nums[i]))
            f.Or(f, shifted).And(f, u) // And(f, u) 保证 f 的二进制长度 <= k+1
            i++
        }
        // 枚举（从大于 x 的数中）选了 j 个 x
        for j := range min(n - i, k / x) + 1 {
            if f.Bit(k - j * x) > 0 {
                res[x-1] = true
                break
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,3,2,4], k = 5
    // Output: [false,false,true,true]
    // Explanation:
    // For x = 1, the capped array is [1, 1, 1, 1]. Possible sums are 1, 2, 3, 4, so it is impossible to form a sum of 5.
    // For x = 2, the capped array is [2, 2, 2, 2]. Possible sums are 2, 4, 6, 8, so it is impossible to form a sum of 5.
    // For x = 3, the capped array is [3, 3, 2, 3]. A subsequence [2, 3] sums to 5, so it is possible.
    // For x = 4, the capped array is [4, 3, 2, 4]. A subsequence [3, 2] sums to 5, so it is possible.
    fmt.Println(subsequenceSumAfterCapping([]int{4,3,2,4}, 5)) // [false,false,true,true]
    // Example 2:
    // Input: nums = [1,2,3,4,5], k = 3
    // Output: [true,true,true,true,true]
    // Explanation:
    // For every value of x, it is always possible to select a subsequence from the capped array that sums exactly to 3.
    fmt.Println(subsequenceSumAfterCapping([]int{1,2,3,4,5}, 3)) // [true,true,true,true,true]

    fmt.Println(subsequenceSumAfterCapping([]int{1,2,3,4,5,6,7,8,9}, 2)) // [true true true true true true true true true]
    fmt.Println(subsequenceSumAfterCapping([]int{9,8,7,6,5,4,3,2,1}, 2)) // [true true true true true true true true true]
}