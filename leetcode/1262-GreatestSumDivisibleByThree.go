package main

// 1262. Greatest Sum Divisible by Three
// Given an integer array nums, 
// return the maximum possible sum of elements of the array such that it is divisible by three.

// Example 1:
// Input: nums = [3,6,5,1,8]
// Output: 18
// Explanation: Pick numbers 3, 6, 1 and 8 their sum is 18 (maximum sum divisible by 3).

// Example 2:
// Input: nums = [4]
// Output: 0
// Explanation: Since 4 is not divisible by 3, do not pick any number.

// Example 3:
// Input: nums = [1,2,3,4,4]
// Output: 12
// Explanation: Pick numbers 1, 3, 4 and 4 their sum is 12 (maximum sum divisible by 3).

// Constraints:
//     1 <= nums.length <= 4 * 10^4
//     1 <= nums[i] <= 10^4

import "fmt"

func maxSumDivThree(nums []int) int {
    res, one, two := 0, 10000, 10000
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < len(nums); i++ {
        res += nums[i]
        if nums[i] % 3 == 1 {
            two, one = min(two, one + nums[i]), min(one, nums[i])
        }
        if nums[i] % 3 == 2 {
            one, two = min(one, two + nums[i]), min(two, nums[i])
        }
    }
    if res % 3 == 0 { return res }
    if res % 3 == 1 { return res - one }
    return res - two
}

func maxSumDivThree1(nums []int) int {
    dp := make([]int, 3)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        t := make([]int, 3)
        copy(t, dp)
        for i := 0; i < 3; i++ {
            dp[(v + t[i]) % 3] = max(dp[(v + t[i]) % 3], v + t[i])
        }
    }
    return dp[0]
}

func main() {
    // Example 1:
    // Input: nums = [3,6,5,1,8]
    // Output: 18
    // Explanation: Pick numbers 3, 6, 1 and 8 their sum is 18 (maximum sum divisible by 3).
    fmt.Println(maxSumDivThree([]int{3,6,5,1,8})) // 18
    // Example 2:
    // Input: nums = [4]
    // Output: 0
    // Explanation: Since 4 is not divisible by 3, do not pick any number.
    fmt.Println(maxSumDivThree([]int{4})) // 0
    // Example 3:
    // Input: nums = [1,2,3,4,4]
    // Output: 12
    // Explanation: Pick numbers 1, 3, 4 and 4 their sum is 12 (maximum sum divisible by 3).
    fmt.Println(maxSumDivThree([]int{1,2,3,4,4})) // 12

    fmt.Println(maxSumDivThree([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxSumDivThree([]int{9,8,7,6,5,4,3,2,1})) // 45

    fmt.Println(maxSumDivThree1([]int{3,6,5,1,8})) // 18
    fmt.Println(maxSumDivThree1([]int{4})) // 0
    fmt.Println(maxSumDivThree1([]int{1,2,3,4,4})) // 12
    fmt.Println(maxSumDivThree1([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxSumDivThree1([]int{9,8,7,6,5,4,3,2,1})) // 45
}