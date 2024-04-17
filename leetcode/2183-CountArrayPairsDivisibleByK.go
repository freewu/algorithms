package main

// 2183. Count Array Pairs Divisible by K
// Given a 0-indexed integer array nums of length n and an integer k, 
// return the number of pairs (i, j) such that:
//     0 <= i < j <= n - 1 and
//     nums[i] * nums[j] is divisible by k.
 
// Example 1:
// Input: nums = [1,2,3,4,5], k = 2
// Output: 7
// Explanation: 
// The 7 pairs of indices whose corresponding products are divisible by 2 are
// (0, 1), (0, 3), (1, 2), (1, 3), (1, 4), (2, 3), and (3, 4).
// Their products are 2, 4, 6, 8, 10, 12, and 20 respectively.
// Other pairs such as (0, 2) and (2, 4) have products 3 and 15 respectively, which are not divisible by 2.    

// Example 2:
// Input: nums = [1,2,3,4], k = 5
// Output: 0
// Explanation: There does not exist any pair of indices whose corresponding product is divisible by 5.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i], k <= 10^5

import "fmt"

// func countPairs(nums []int, k int) int64 {
//     if len(nums) < k {
//         return 0
//     }
//     if len(nums) == k {
//         return 1
//     }
//     res := 0

//     return int64(res)
// }

func countPairs(nums []int, k int) int64 {
    counts := make(map[int]int)
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for _, num := range nums {
        counts[gcd(num, k)]++
    }
    res := 0
    for x, xCount := range counts {
        for y, yCount := range counts {
            if x <= y && (x*y) % k == 0 {
                if x == y {
                    res += xCount * (xCount - 1) / 2
                } else {
                    res += xCount * yCount
                }
            }
        }
    }
    return int64(res)
}

func main() {
    // The 7 pairs of indices whose corresponding products are divisible by 2 are
    // (0, 1), (0, 3), (1, 2), (1, 3), (1, 4), (2, 3), and (3, 4).
    // Their products are 2, 4, 6, 8, 10, 12, and 20 respectively.
    // Other pairs such as (0, 2) and (2, 4) have products 3 and 15 respectively, which are not divisible by 2.    
    fmt.Println(countPairs([]int{1,2,3,4,5}, 2)) // 7
    // Explanation: There does not exist any pair of indices whose corresponding product is divisible by 5.
    fmt.Println(countPairs([]int{1,2,3,4}, 5)) // 0
}