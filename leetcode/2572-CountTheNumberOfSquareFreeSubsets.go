package main

// 2572. Count the Number of Square-Free Subsets
// You are given a positive integer 0-indexed array nums.

// A subset of the array nums is square-free if the product of its elements is a square-free integer.

// A square-free integer is an integer that is divisible by no square number other than 1.

// Return the number of square-free non-empty subsets of the array nums. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// A non-empty subset of nums is an array that can be obtained by deleting some (possibly none but not all) elements from nums. 
// Two subsets are different if and only if the chosen indices to delete are different.

// Example 1:
// Input: nums = [3,4,4,5]
// Output: 3
// Explanation: There are 3 square-free subsets in this example:
// - The subset consisting of the 0th element [3]. The product of its elements is 3, which is a square-free integer.
// - The subset consisting of the 3rd element [5]. The product of its elements is 5, which is a square-free integer.
// - The subset consisting of 0th and 3rd elements [3,5]. The product of its elements is 15, which is a square-free integer.
// It can be proven that there are no more than 3 square-free subsets in the given array.

// Example 2:
// Input: nums = [1]
// Output: 1
// Explanation: There is 1 square-free subset in this example:
// - The subset consisting of the 0th element [1]. The product of its elements is 1, which is a square-free integer.
// It can be proven that there is no more than 1 square-free subset in the given array.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 30

import "fmt"

func squareFreeSubsets(nums []int) int {
    sum, mod := 0, 1_000_000_007
    primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
    getMask := func(num int) int {
        mask := 0
        for i, prime := range primes {
            if num % (prime * prime) == 0 {
                mask = -1
                break
            } else if num % prime == 0 {
                mask = mask | (1 << i)
            }
        }
        return mask
    }
    dp := make([]int, 1 << 10)
    dp[0] = 1 // product 0 is always achievable
    for _, v := range nums {
        mask := getMask(v)
        if mask < 0 { continue }
        for i := 0; i < (1 << 10); i++ {
            if (mask & i) == 0 {
                dp[mask | i] += (dp[i] % mod)
            }
        }
    }
    for _, v := range dp {
        sum += v
    }
    return (sum - 1) % mod
}

func squareFreeSubsets1(nums []int) int {
    primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
    count := [31]int{}
    for _, v := range nums {
        count[v]++
    }
    res, n, mod := -1, 10, 1_000_000_007
    dp := make([]int, 1 << n)
    dp[0] = 1
    for i := 0; i < count[1]; i++ {
        dp[0] = dp[0] * 2 % mod
    }
    for i := 2; i < 31; i++ {
        if count[i] == 0 || i % 4 == 0 || i % 9 == 0 || i % 25 == 0 { continue }
        mask := 0
        for j, p := range primes {
            if i % p == 0 {
                mask |= 1 << j
            }
        }
        for j := 1 << n - 1; j > 0; j-- {
            if j & mask == mask {
                dp[j] = (dp[j] + dp[j ^ mask] * count[i]) % mod
            }
        }
    }
    for _, v := range dp {
        res = (res + v) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,4,4,5]
    // Output: 3
    // Explanation: There are 3 square-free subsets in this example:
    // - The subset consisting of the 0th element [3]. The product of its elements is 3, which is a square-free integer.
    // - The subset consisting of the 3rd element [5]. The product of its elements is 5, which is a square-free integer.
    // - The subset consisting of 0th and 3rd elements [3,5]. The product of its elements is 15, which is a square-free integer.
    // It can be proven that there are no more than 3 square-free subsets in the given array.
    fmt.Println(squareFreeSubsets([]int{3,4,4,5})) // 3
    // Example 2:
    // Input: nums = [1]
    // Output: 1
    // Explanation: There is 1 square-free subset in this example:
    // - The subset consisting of the 0th element [1]. The product of its elements is 1, which is a square-free integer.
    // It can be proven that there is no more than 1 square-free subset in the given array.
    fmt.Println(squareFreeSubsets([]int{1})) // 1

    fmt.Println(squareFreeSubsets([]int{1,2,3,4,5,6,7,8,9})) // 39
    fmt.Println(squareFreeSubsets([]int{9,8,7,6,5,4,3,2,1})) // 39

    fmt.Println(squareFreeSubsets1([]int{3,4,4,5})) // 3
    fmt.Println(squareFreeSubsets1([]int{1})) // 1
    fmt.Println(squareFreeSubsets1([]int{1,2,3,4,5,6,7,8,9})) // 39
    fmt.Println(squareFreeSubsets1([]int{9,8,7,6,5,4,3,2,1})) // 39
}