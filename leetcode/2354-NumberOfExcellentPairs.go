package main

// 2354. Number of Excellent Pairs
// You are given a 0-indexed positive integer array nums and a positive integer k.

// A pair of numbers (num1, num2) is called excellent if the following conditions are satisfied:
//     1. Both the numbers num1 and num2 exist in the array nums.
//     2. The sum of the number of set bits in num1 OR num2 and num1 AND num2 is greater than or equal to k, 
//        where OR is the bitwise OR operation and AND is the bitwise AND operation.

// Return the number of distinct excellent pairs.

// Two pairs (a, b) and (c, d) are considered distinct if either a != c or b != d. For example, (1, 2) and (2, 1) are distinct.

// Note that a pair (num1, num2) such that num1 == num2 can also be excellent if you have at least one occurrence of num1 in the array.

// Example 1:
// Input: nums = [1,2,3,1], k = 3
// Output: 5
// Explanation: The excellent pairs are the following:
// - (3, 3). (3 AND 3) and (3 OR 3) are both equal to (11) in binary. The total number of set bits is 2 + 2 = 4, which is greater than or equal to k = 3.
// - (2, 3) and (3, 2). (2 AND 3) is equal to (10) in binary, and (2 OR 3) is equal to (11) in binary. The total number of set bits is 1 + 2 = 3.
// - (1, 3) and (3, 1). (1 AND 3) is equal to (01) in binary, and (1 OR 3) is equal to (11) in binary. The total number of set bits is 1 + 2 = 3.
// So the number of excellent pairs is 5.

// Example 2:
// Input: nums = [5,1,1], k = 10
// Output: 0
// Explanation: There are no excellent pairs for this array.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= k <= 60

import "fmt"
import "math/bits"

func countExcellentPairs(nums []int, k int) int64 {
    count, mp := [32]int{}, make(map[int]bool)
    for i := 0; i < len(nums); i++ {
        if mp[nums[i]] { continue }
        mp[nums[i]] = true
        n := bits.OnesCount(uint(nums[i]))
        count[n]++
    }
    res := 0
    for i := 0; i <= 31; i++ {
        if i + i >= k {
            res += count[i] * count[i]
        }
        for j := i + 1; j <= 31; j++ {
            if i + j >= k {
                res += 2 * count[i] * count[j]
            }
        }
    }
    return int64(res)
}

func countExcellentPairs1(nums []int, k int) int64 {
    used, bits := make(map[int]bool), make([]int, 32)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    popcount := func(n int) int {
        n = (n & 0x55555555) + ((n >> 1) & 0x55555555)
        n = (n & 0x33333333) + ((n >> 2) & 0x33333333)
        n = (n & 0x0f0f0f0f) + ((n >> 4) & 0x0f0f0f0f)
        n = (n & 0x00ff00ff) + ((n >> 8) & 0x00ff00ff)
        return (n & 0x0000ffff) + ((n >> 16) & 0x0000ffff)
    }
    for _, v := range nums {
        if used[v] { continue }
        used[v] = true
        bits[popcount(v)]++
    }
    res := int64(0)
    for i := 0; i < 32; i++ {
        for j := max(0, k - i); j < 32; j++ {
            res += int64(bits[i] * bits[j])
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,1], k = 3
    // Output: 5
    // Explanation: The excellent pairs are the following:
    // - (3, 3). (3 AND 3) and (3 OR 3) are both equal to (11) in binary. The total number of set bits is 2 + 2 = 4, which is greater than or equal to k = 3.
    // - (2, 3) and (3, 2). (2 AND 3) is equal to (10) in binary, and (2 OR 3) is equal to (11) in binary. The total number of set bits is 1 + 2 = 3.
    // - (1, 3) and (3, 1). (1 AND 3) is equal to (01) in binary, and (1 OR 3) is equal to (11) in binary. The total number of set bits is 1 + 2 = 3.
    // So the number of excellent pairs is 5.
    fmt.Println(countExcellentPairs([]int{1,2,3,1}, 3)) // 5
    // Example 2:
    // Input: nums = [5,1,1], k = 10
    // Output: 0
    // Explanation: There are no excellent pairs for this array.
    fmt.Println(countExcellentPairs([]int{5,1,1}, 10)) // 0

    fmt.Println(countExcellentPairs1([]int{1,2,3,1}, 3)) // 5
    fmt.Println(countExcellentPairs1([]int{5,1,1}, 10)) // 0
}