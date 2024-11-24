package main

// 1994. The Number of Good Subsets
// You are given an integer array nums. 
// We call a subset of nums good if its product can be represented as a product of one or more distinct prime numbers.
//     For example, if nums = [1, 2, 3, 4]:
//         [2, 3], [1, 2, 3], and [1, 3] are good subsets with products 6 = 2*3, 6 = 2*3, and 3 = 3 respectively.
//         [1, 4] and [4] are not good subsets with products 4 = 2*2 and 4 = 2*2 respectively.

// Return the number of different good subsets in nums modulo 10^9 + 7.

// A subset of nums is any array that can be obtained by deleting some (possibly none or all) elements from nums. 
// Two subsets are different if and only if the chosen indices to delete are different.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: 6
// Explanation: The good subsets are:
// - [1,2]: product is 2, which is the product of distinct prime 2.
// - [1,2,3]: product is 6, which is the product of distinct primes 2 and 3.
// - [1,3]: product is 3, which is the product of distinct prime 3.
// - [2]: product is 2, which is the product of distinct prime 2.
// - [2,3]: product is 6, which is the product of distinct primes 2 and 3.
// - [3]: product is 3, which is the product of distinct prime 3.

// Example 2:
// Input: nums = [4,2,3,15]
// Output: 5
// Explanation: The good subsets are:
// - [2]: product is 2, which is the product of distinct prime 2.
// - [2,3]: product is 6, which is the product of distinct primes 2 and 3.
// - [2,15]: product is 30, which is the product of distinct primes 2, 3, and 5.
// - [3]: product is 3, which is the product of distinct prime 3.
// - [15]: product is 15, which is the product of distinct primes 3 and 5.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 30

import "fmt"

func numberOfGoodSubsets(nums []int) int {
    primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
    count := [31]int{}
    for _, v := range nums {
        count[v]++
    }
    res, n, mod := 0, 10, 1_000_000_007
    dp := make([]int, 1 << n)
    dp[0] = 1
    for i := 0; i < count[1]; i++ {
        dp[0] = dp[0] * 2 % mod
    }
    for i := 2; i < 31; i++ {
        if count[i] == 0 || i % 4 == 0 || i % 9 == 0 || i % 25 == 0 { continue }
        mask := 0
        for index, p := range primes {
            if i % p == 0 {
                mask |= 1 << index
            }
        }
        for state := 1 << n - 1; state > 0; state-- {
            if state & mask == mask {
                dp[state] = (dp[state] + dp[state ^ mask] * count[i]) % mod
            }
        }
    }
    for i := 1; i < 1 << n; i++ {
        res = (res + dp[i]) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4]
    // Output: 6
    // Explanation: The good subsets are:
    // - [1,2]: product is 2, which is the product of distinct prime 2.
    // - [1,2,3]: product is 6, which is the product of distinct primes 2 and 3.
    // - [1,3]: product is 3, which is the product of distinct prime 3.
    // - [2]: product is 2, which is the product of distinct prime 2.
    // - [2,3]: product is 6, which is the product of distinct primes 2 and 3.
    // - [3]: product is 3, which is the product of distinct prime 3.
    fmt.Println(numberOfGoodSubsets([]int{1,2,3,4})) // 6
    // Example 2:
    // Input: nums = [4,2,3,15]
    // Output: 5
    // Explanation: The good subsets are:
    // - [2]: product is 2, which is the product of distinct prime 2.
    // - [2,3]: product is 6, which is the product of distinct primes 2 and 3.
    // - [2,15]: product is 30, which is the product of distinct primes 2, 3, and 5.
    // - [3]: product is 3, which is the product of distinct prime 3.
    // - [15]: product is 15, which is the product of distinct primes 3 and 5.
    fmt.Println(numberOfGoodSubsets([]int{4,2,3,15})) // 5
}