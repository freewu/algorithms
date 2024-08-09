package main

// 805. Split Array With Same Average
// You are given an integer array nums.
// You should move each element of nums into one of the two arrays A and B such that A and B are non-empty, and average(A) == average(B).
// Return true if it is possible to achieve that and false otherwise.
// Note that for an array arr, average(arr) is the sum of all the elements of arr over the length of arr.

// Example 1:
// Input: nums = [1,2,3,4,5,6,7,8]
// Output: true
// Explanation: We can split the array into [1,4,5,8] and [2,3,6,7], and both of them have an average of 4.5.

// Example 2:
// Input: nums = [3,1]
// Output: false

// Constraints:
//     1 <= nums.length <= 30
//     0 <= nums[i] <= 10^4

import "fmt"

func splitArraySameAverage(nums []int) bool {
    total, n := 0, len(nums)
    for _, num := range nums {
        total += num
    }
    dp := make([][]bool, n/2+1)
    for i := range dp {
        dp[i] = make([]bool, total * n / 2 / n + 1)
    }
    dp[0][0] = true
    for _, num := range nums {
        for cnt := n/2; cnt >= 1; cnt-- {
            for sum := total*n/2/n; sum >= num; sum-- {
                dp[cnt][sum] = dp[cnt][sum] || dp[cnt-1][sum-num]
            }
        }
    }
    for i := 1; i <= n/2; i++ {
        if total*i%n == 0 && dp[i][total*i/n] {
            return true
        }
    }
    return false
}

func splitArraySameAverage1(nums []int) bool {
    n := len(nums)
    if n == 1 {
        return false
    }
    sum := 0
    for _, x := range nums {
        sum += x
    }
    for i := range nums {
        nums[i] = nums[i] * n - sum
    }
    m := n / 2
    left := map[int]bool{}
    for i := 1; i < 1 << m; i++ {
        total := 0
        for j, x := range nums[:m] {
            if i >> j & 1 > 0 {
                total += x
            }
        }
        if total == 0 {
            return true
        }
        left[total] = true
    }
    rsum := 0
    for _, x := range nums[m:] {
        rsum += x
    }
    for i := 1; i < 1 << (n-m); i++ {
        total := 0
        for j, x := range nums[m:] {
            if i >> j&1 > 0 {
                total += x
            }
        }
        if total == 0 || rsum != total && left[-total] {
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5,6,7,8]
    // Output: true
    // Explanation: We can split the array into [1,4,5,8] and [2,3,6,7], and both of them have an average of 4.5.
    fmt.Println(splitArraySameAverage([]int{1,2,3,4,5,6,7,8})) // true
    // Example 2:
    // Input: nums = [3,1]
    // Output: false
    fmt.Println(splitArraySameAverage([]int{3,1})) // false

    fmt.Println(splitArraySameAverage1([]int{1,2,3,4,5,6,7,8})) // true
    fmt.Println(splitArraySameAverage1([]int{3,1})) // false
}