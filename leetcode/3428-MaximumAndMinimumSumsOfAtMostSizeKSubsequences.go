package main

// 3428. Maximum and Minimum Sums of at Most Size K Subsequences
// You are given an integer array nums and a positive integer k. 
// Return the sum of the maximum and minimum elements of all subsequences of nums with at most k elements.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [1,2,3], k = 2
// Output: 24
// Explanation:
// The subsequences of nums with at most 2 elements are:
// Subsequence	Minimum	Maximum	Sum
// [1]	1	1	2
// [2]	2	2	4
// [3]	3	3	6
// [1, 2]	1	2	3
// [1, 3]	1	3	4
// [2, 3]	2	3	5
// Final Total	 	 	24
// The output would be 24.

// Example 2:
// Input: nums = [5,0,6], k = 1
// Output: 22
// Explanation:
// For subsequences with exactly 1 element, the minimum and maximum values are the element itself. 
// Therefore, the total is 5 + 5 + 0 + 0 + 6 + 6 = 22.

// Example 3:
// Input: nums = [1,1,1], k = 2
// Output: 12
// Explanation:
// The subsequences [1, 1] and [1] each appear 3 times. 
// For all of them, the minimum and maximum are both 1. 
// Thus, the total is 12.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^9
//     1 <= k <= min(70, nums.length)

import "fmt"
import "sort"

// const mod = 1_000_000_007

// var frac []int

// func init() {
//     frac = make([]int, 1e5+1)
//     frac[0] = 1
//     for i := 1; i <= 1e5; i++ {
//         frac[i] = frac[i-1] * i % mod
//     }
// }

// func minMaxSums(nums []int, k int) int {
//     sort.Ints(nums)
//     res, s, n := 0, 1, len(nums)
//     pow := func(x, y, mod int) int {
//         res := 1
//         for y > 0 {
//             if y & 1 == 1 {
//                 res = res * x % mod
//             }
//             x = x * x % mod
//             y >>= 1
//         }
//         return res
//     }
//     comb := func(n, k int) int {
//         if n < k { return 0 }
//         return frac[n] * pow(frac[k], mod - 2, mod) % mod * pow(frac[n-k], mod - 2, mod) % mod
//     }
//     for i := 0; i < n; i++ {
//         res = (res + (nums[i] + nums[n - i - 1]) * s) % mod
//         s = (2*s - comb(i, k-1)) % mod
//         if s < 0 {
//             s += mod
//         }
//     }
//     return res
// }

const mod = 1_000_000_007
const mx = 100_000

var f [mx]int    // f[i] = i!
var invF [mx]int // invF[i] = i!^-1

func init() {
    f[0] = 1
    pow := func(x, n int) int {
        res := 1
        for ; n > 0; n /= 2 {
            if n%2 > 0 {
                res = res * x % mod
            }
            x = x * x % mod
        }
        return res
    }
    for i := 1; i < mx; i++ {
        f[i] = f[i-1] * i % mod
    }
    invF[mx-1] = pow(f[mx-1], mod-2)
    for i := mx - 1; i > 0; i-- {
        invF[i-1] = invF[i] * i % mod
    }
}

func minMaxSums(nums []int, k int) int {
    sort.Ints(nums)
    comb := func(n, m int) int {
        if m > n { return 0 }
        return f[n] * invF[m] % mod * invF[n-m] % mod
    }
    res, s := 0, 1
    for i, v := range nums {
        res = (res + s * (v + nums[len(nums) - i - 1])) % mod
        s = (s * 2 - comb(i, k - 1) + mod) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3], k = 2
    // Output: 24
    // Explanation:
    // The subsequences of nums with at most 2 elements are:
    // Subsequence	Minimum	Maximum	Sum
    // [1]	1	1	2
    // [2]	2	2	4
    // [3]	3	3	6
    // [1, 2]	1	2	3
    // [1, 3]	1	3	4
    // [2, 3]	2	3	5
    // Final Total	 	 	24
    // The output would be 24.
    fmt.Println(minMaxSums([]int{1,2,3}, 2)) // 24
    // Example 2:
    // Input: nums = [5,0,6], k = 1
    // Output: 22
    // Explanation:
    // For subsequences with exactly 1 element, the minimum and maximum values are the element itself. 
    // Therefore, the total is 5 + 5 + 0 + 0 + 6 + 6 = 22.
    fmt.Println(minMaxSums([]int{5,0,6}, 1)) // 22
    // Example 3:
    // Input: nums = [1,1,1], k = 2
    // Output: 12
    // Explanation:
    // The subsequences [1, 1] and [1] each appear 3 times. 
    // For all of them, the minimum and maximum are both 1. 
    // Thus, the total is 12.
    fmt.Println(minMaxSums([]int{1,1,1}, 2)) // 12

    fmt.Println(minMaxSums([]int{1,2,3,4,5,6,7,8,9}, 2)) // 450
    fmt.Println(minMaxSums([]int{9,8,7,6,5,4,3,2,1}, 2)) // 450
}