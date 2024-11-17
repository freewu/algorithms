package main

// 3082. Find the Sum of the Power of All Subsequences
// You are given an integer array nums of length n and a positive integer k.

// The power of an array of integers is defined as the number of subsequences with their sum equal to k.

// Return the sum of power of all subsequences of nums.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input:  nums = [1,2,3], k = 3 
// Output:  6 
// Explanation:
// There are 5 subsequences of nums with non-zero power:
// The subsequence [1,2,3] has 2 subsequences with sum == 3: [1,2,3] and [1,2,3].
// The subsequence [1,2,3] has 1 subsequence with sum == 3: [1,2,3].
// The subsequence [1,2,3] has 1 subsequence with sum == 3: [1,2,3].
// The subsequence [1,2,3] has 1 subsequence with sum == 3: [1,2,3].
// The subsequence [1,2,3] has 1 subsequence with sum == 3: [1,2,3].
// Hence the answer is 2 + 1 + 1 + 1 + 1 = 6.

// Example 2:
// Input:  nums = [2,3,3], k = 5 
// Output:  4 
// Explanation:
// There are 3 subsequences of nums with non-zero power:
// The subsequence [2,3,3] has 2 subsequences with sum == 5: [2,3,3] and [2,3,3].
// The subsequence [2,3,3] has 1 subsequence with sum == 5: [2,3,3].
// The subsequence [2,3,3] has 1 subsequence with sum == 5: [2,3,3].
// Hence the answer is 2 + 1 + 1 = 4.

// Example 3:
// Input:  nums = [1,2,3], k = 7 
// Output:  0 
// Explanation: There exists no subsequence with sum 7. Hence all subsequences of nums have power = 0.

// Constraints:
//     1 <= n <= 100
//     1 <= nums[i] <= 10^4
//     1 <= k <= 100

import "fmt"

func sumOfPower(nums []int, k int) int {
    dp, prev, mod := make([]int, k + 1), make([]int, k + 1), 1_000_000_007
    prev[0] = 1
    for _, v := range nums {
        for i := 0; i <= k; i++ {
            dp[i] = (prev[i] * 2) % mod
            if i - v >= 0 {
                dp[i] += prev[i - v]
                dp[i] %= mod
            }
        }
        copy(prev, dp)
    }
    return prev[k]
}

func main() {
    // Example 1:
    // Input:  nums = [1,2,3], k = 3 
    // Output:  6 
    // Explanation:
    // There are 5 subsequences of nums with non-zero power:
    // The subsequence [1,2,3] has 2 subsequences with sum == 3: [1,2,3] and [1,2,3].
    // The subsequence [1,2,3] has 1 subsequence with sum == 3: [1,2,3].
    // The subsequence [1,2,3] has 1 subsequence with sum == 3: [1,2,3].
    // The subsequence [1,2,3] has 1 subsequence with sum == 3: [1,2,3].
    // The subsequence [1,2,3] has 1 subsequence with sum == 3: [1,2,3].
    // Hence the answer is 2 + 1 + 1 + 1 + 1 = 6.
    fmt.Println(sumOfPower([]int{1,2,3}, 3)) // 6
    // Example 2:
    // Input:  nums = [2,3,3], k = 5 
    // Output:  4 
    // Explanation:
    // There are 3 subsequences of nums with non-zero power:
    // The subsequence [2,3,3] has 2 subsequences with sum == 5: [2,3,3] and [2,3,3].
    // The subsequence [2,3,3] has 1 subsequence with sum == 5: [2,3,3].
    // The subsequence [2,3,3] has 1 subsequence with sum == 5: [2,3,3].
    // Hence the answer is 2 + 1 + 1 = 4.
    fmt.Println(sumOfPower([]int{2,3,3}, 5)) // 4
    // Example 3:
    // Input:  nums = [1,2,3], k = 7 
    // Output:  0 
    // Explanation: There exists no subsequence with sum 7. Hence all subsequences of nums have power = 0.
    fmt.Println(sumOfPower([]int{1,2,3}, 7)) // 0
}