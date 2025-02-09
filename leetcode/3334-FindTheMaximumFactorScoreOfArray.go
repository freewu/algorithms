package main

// 3334. Find the Maximum Factor Score of Array
// You are given an integer array nums.

// The factor score of an array is defined as the product of the LCM and GCD of all elements of that array.

// Return the maximum factor score of nums after removing at most one element from it.

// Note that both the LCM and GCD of a single number are the number itself, and the factor score of an empty array is 0.

// Example 1:
// Input: nums = [2,4,8,16]
// Output: 64
// Explanation:
// On removing 2, the GCD of the rest of the elements is 4 while the LCM is 16, which gives a maximum factor score of 4 * 16 = 64.

// Example 2:
// Input: nums = [1,2,3,4,5]
// Output: 60
// Explanation:
// The maximum factor score of 60 can be obtained without removing any elements.

// Example 3:
// Input: nums = [3]
// Output: 9

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 30

import "fmt"

func maxScore(nums []int) int64 {
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    lcm := func(x int, y int) int { return (x * y) / gcd(x, y) }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    n := len(nums) 
    sufGcd, sufLcm := make([]int, n + 1), make([]int, n + 1)
    sufLcm[n] = 1
    for i := n - 1; i >= 0; i-- {
        sufGcd[i], sufLcm[i] = gcd(sufGcd[i + 1], nums[i]), lcm(sufLcm[i + 1], nums[i])
    }
    res := sufGcd[0] * sufLcm[0]
    preGcd, preLcm := 0, 1
    for i := 0; i < n; i++ {
        res = max(res, gcd(preGcd, sufGcd[i + 1]) * lcm(preLcm, sufLcm[i + 1]))
        preGcd, preLcm = gcd(preGcd, nums[i]), lcm(preLcm, nums[i])
    }
    return int64(res)
}

func maxScore1(nums []int) int64 {
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    max := func(x, y int64) int64 { if x > y { return x; }; return y; }
    computeFactorScore := func(excl int) int64 {
        rGcd, rLcm := -1, int64(-1)
        for i, n := range nums {
            if i == excl { continue }
            if rGcd == -1 {
                rGcd, rLcm = n, int64(n)
                continue
            }
            rGcd, rLcm = gcd(rGcd, n), (rLcm * int64(n)) / int64(gcd(int(rLcm), n))
        }
        if rGcd == -1 {  return 0 }
        return int64(rGcd) * rLcm
    }
    res := computeFactorScore(-1)
    for i := range nums {
        res = max(res, computeFactorScore(i))
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,4,8,16]
    // Output: 64
    // Explanation:
    // On removing 2, the GCD of the rest of the elements is 4 while the LCM is 16, which gives a maximum factor score of 4 * 16 = 64.
    fmt.Println(maxScore([]int{2,4,8,16})) // 64
    // Example 2:
    // Input: nums = [1,2,3,4,5]
    // Output: 60
    // Explanation:
    // The maximum factor score of 60 can be obtained without removing any elements.
    fmt.Println(maxScore([]int{1,2,3,4,5})) // 60
    // Example 3:
    // Input: nums = [3]
    // Output: 9
    fmt.Println(maxScore([]int{3})) // 9

    fmt.Println(maxScore([]int{1,2,3,4,5,6,7,8,9})) // 2520
    fmt.Println(maxScore([]int{9,8,7,6,5,4,3,2,1})) // 2520

    fmt.Println(maxScore1([]int{2,4,8,16})) // 64
    fmt.Println(maxScore1([]int{1,2,3,4,5})) // 60
    fmt.Println(maxScore1([]int{3})) // 9
    fmt.Println(maxScore1([]int{1,2,3,4,5,6,7,8,9})) // 2520
    fmt.Println(maxScore1([]int{9,8,7,6,5,4,3,2,1})) // 2520
}