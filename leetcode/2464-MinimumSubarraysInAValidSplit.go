package main

// 2464. Minimum Subarrays in a Valid Split
// You are given an integer array nums.
// Splitting of an integer array nums into subarrays is valid if:
//     the greatest common divisor of the first and last elements of each subarray is greater than 1, and
//     each element of nums belongs to exactly one subarray.

// Return the minimum number of subarrays in a valid subarray splitting of nums. 
// If a valid subarray splitting is not possible, return -1.

// Note that:
//     The greatest common divisor of two numbers is the largest positive integer that evenly divides both numbers.
//     A subarray is a contiguous non-empty part of an array.

// Example 1:
// Input: nums = [2,6,3,4,3]
// Output: 2
// Explanation: We can create a valid split in the following way: [2,6] | [3,4,3].
// - The starting element of the 1st subarray is 2 and the ending is 6. Their greatest common divisor is 2, which is greater than 1.
// - The starting element of the 2nd subarray is 3 and the ending is 3. Their greatest common divisor is 3, which is greater than 1.
// It can be proved that 2 is the minimum number of subarrays that we can obtain in a valid split.

// Example 2:
// Input: nums = [3,5]
// Output: 2
// Explanation: We can create a valid split in the following way: [3] | [5].
// - The starting element of the 1st subarray is 3 and the ending is 3. Their greatest common divisor is 3, which is greater than 1.
// - The starting element of the 2nd subarray is 5 and the ending is 5. Their greatest common divisor is 5, which is greater than 1.
// It can be proved that 2 is the minimum number of subarrays that we can obtain in a valid split.

// Example 3:
// Input: nums = [1,2,1]
// Output: -1
// Explanation: It is impossible to create valid split.
 
// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 10^5

import "fmt"

func validSubarraySplit(nums []int) int {
    n, inf := len(nums), 1 << 32 -1
    dp := make([]int, n + 1)
    for i := 0; i <= n; i++ {
        dp[i] = inf
    }
    dp[0] = 0
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            if gcd(nums[i - 1], nums[j - 1]) > 1 {
                dp[i] = min(dp[i], dp[j - 1] + 1)
            }
        }
    }
    if dp[n] == inf { return -1 }
    return dp[n]
}

// dfs
func validSubarraySplit1(nums []int) int {
    n, inf := len(nums), 1 << 32 -1
    dp := make([]int, n + 1)
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(i int) int
    dfs = func(i int) int {
        if i >= n { return 0 }
        if dp[i] > 0 { return dp[i] }
        res := inf
        for j := i; j < n; j++ {
            if gcd(nums[i], nums[j]) > 1 {
                res = min(res, 1 + dfs(j + 1))
            }
        }
        dp[i] = res
        return res
    }
    res := dfs(0)
    if res < inf { return res }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [2,6,3,4,3]
    // Output: 2
    // Explanation: We can create a valid split in the following way: [2,6] | [3,4,3].
    // - The starting element of the 1st subarray is 2 and the ending is 6. Their greatest common divisor is 2, which is greater than 1.
    // - The starting element of the 2nd subarray is 3 and the ending is 3. Their greatest common divisor is 3, which is greater than 1.
    // It can be proved that 2 is the minimum number of subarrays that we can obtain in a valid split.
    fmt.Println(validSubarraySplit([]int{2,6,3,4,3})) // 2
    // Example 2:
    // Input: nums = [3,5]
    // Output: 2
    // Explanation: We can create a valid split in the following way: [3] | [5].
    // - The starting element of the 1st subarray is 3 and the ending is 3. Their greatest common divisor is 3, which is greater than 1.
    // - The starting element of the 2nd subarray is 5 and the ending is 5. Their greatest common divisor is 5, which is greater than 1.
    // It can be proved that 2 is the minimum number of subarrays that we can obtain in a valid split.
    fmt.Println(validSubarraySplit([]int{3,5})) // 2
    // Example 3:
    // Input: nums = [1,2,1]
    // Output: -1
    // Explanation: It is impossible to create valid split.
    
    // Explanation: We can create a valid split in the following way: [3] | [5].
    // - The starting element of the 1st subarray is 3 and the ending is 3. Their greatest common divisor is 3, which is greater than 1.
    // - The starting element of the 2nd subarray is 5 and the ending is 5. Their greatest common divisor is 5, which is greater than 1.
    // It can be proved that 2 is the minimum number of subarrays that we can obtain in a valid split.
    fmt.Println(validSubarraySplit([]int{1,2,1})) // -1

    fmt.Println(validSubarraySplit([]int{1,2,3,4,5,6,7,8,9})) // -1
    fmt.Println(validSubarraySplit([]int{9,8,7,6,5,4,3,2,1})) // -1

    fmt.Println(validSubarraySplit1([]int{2,6,3,4,3})) // 2
    fmt.Println(validSubarraySplit1([]int{3,5})) // 2
    fmt.Println(validSubarraySplit1([]int{1,2,1})) // -1
    fmt.Println(validSubarraySplit1([]int{1,2,3,4,5,6,7,8,9})) // -1
    fmt.Println(validSubarraySplit1([]int{9,8,7,6,5,4,3,2,1})) // -1
}