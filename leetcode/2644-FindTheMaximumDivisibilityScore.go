package main

// 2644. Find the Maximum Divisibility Score
// You are given two 0-indexed integer arrays nums and divisors.
// The divisibility score of divisors[i] is the number of indices j such that nums[j] is divisible by divisors[i].
// Return the integer divisors[i] with the maximum divisibility score. If there is more than one integer with the maximum score, return the minimum of them.

// Example 1:
// Input: nums = [4,7,9,3,9], divisors = [5,2,3]
// Output: 3
// Explanation: The divisibility score for every element in divisors is:
// The divisibility score of divisors[0] is 0 since no number in nums is divisible by 5.
// The divisibility score of divisors[1] is 1 since nums[0] is divisible by 2.
// The divisibility score of divisors[2] is 3 since nums[2], nums[3], and nums[4] are divisible by 3.
// Since divisors[2] has the maximum divisibility score, we return it.

// Example 2:
// Input: nums = [20,14,21,10], divisors = [5,7,5]
// Output: 5
// Explanation: The divisibility score for every element in divisors is:
// The divisibility score of divisors[0] is 2 since nums[0] and nums[3] are divisible by 5.
// The divisibility score of divisors[1] is 2 since nums[1] and nums[2] are divisible by 7.
// The divisibility score of divisors[2] is 2 since nums[0] and nums[3] are divisible by 5.
// Since divisors[0], divisors[1], and divisors[2] all have the maximum divisibility score, we return the minimum of them (i.e., divisors[2]).

// Example 3:
// Input: nums = [12], divisors = [10,16]
// Output: 10
// Explanation: The divisibility score for every element in divisors is:
// The divisibility score of divisors[0] is 0 since no number in nums is divisible by 10.
// The divisibility score of divisors[1] is 0 since no number in nums is divisible by 16.
// Since divisors[0] and divisors[1] both have the maximum divisibility score, we return the minimum of them (i.e., divisors[0]).
 
// Constraints:
//     1 <= nums.length, divisors.length <= 1000
//     1 <= nums[i], divisors[i] <= 10^9

import "fmt"
import "slices"

func maxDivScore(nums []int, divisors []int) int {
    res, score, scores := 0, -1 << 32 - 1, make([]int,len(divisors))
    for i, d := range divisors {
        for _, v := range nums {
            if v % d == 0 { // 能被整除
                scores[i]++
            }
        }
    }
    for i, v := range scores {
        if v > score { // 可整除性得分 最大的整数 divisors[i]
            score = v
            res = divisors[i] 
        } else if v == score { // 如果有多个整数具有最大得分，则返回数值最小的一个。
            if divisors[i] < res {
                res = divisors[i]
            }
        }
    }
    return res
}

func maxDivScore1(nums []int, divisors []int) int {
    res, dup, maxCnt := 0, 0, -1
    slices.SortFunc(nums, func(a, b int) int { return b - a })
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            dup++
        }
    }
    slices.Sort(divisors)
    for _, d := range divisors {
        if (maxCnt - dup + 1) * d > nums[0] {
            break
        }
        cnt := 0
        for _, x := range nums {
            if x < d {
                break
            }
            if x % d == 0 {
                cnt++
            }
        }
        if cnt > maxCnt {
            maxCnt = cnt
            res = d
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,7,9,3,9], divisors = [5,2,3]
    // Output: 3
    // Explanation: The divisibility score for every element in divisors is:
    // The divisibility score of divisors[0] is 0 since no number in nums is divisible by 5.
    // The divisibility score of divisors[1] is 1 since nums[0] is divisible by 2.
    // The divisibility score of divisors[2] is 3 since nums[2], nums[3], and nums[4] are divisible by 3.
    // Since divisors[2] has the maximum divisibility score, we return it.
    fmt.Println(maxDivScore([]int{4,7,9,3,9}, []int{5,2,3})) // 3
    // Example 2:
    // Input: nums = [20,14,21,10], divisors = [5,7,5]
    // Output: 5
    // Explanation: The divisibility score for every element in divisors is:
    // The divisibility score of divisors[0] is 2 since nums[0] and nums[3] are divisible by 5.
    // The divisibility score of divisors[1] is 2 since nums[1] and nums[2] are divisible by 7.
    // The divisibility score of divisors[2] is 2 since nums[0] and nums[3] are divisible by 5.
    // Since divisors[0], divisors[1], and divisors[2] all have the maximum divisibility score, we return the minimum of them (i.e., divisors[2]).
    fmt.Println(maxDivScore([]int{20,14,21,10}, []int{5,7,5})) // 5
    // Example 3:
    // Input: nums = [12], divisors = [10,16]
    // Output: 10
    // Explanation: The divisibility score for every element in divisors is:
    // The divisibility score of divisors[0] is 0 since no number in nums is divisible by 10.
    // The divisibility score of divisors[1] is 0 since no number in nums is divisible by 16.
    // Since divisors[0] and divisors[1] both have the maximum divisibility score, we return the minimum of them (i.e., divisors[0]).
    fmt.Println(maxDivScore([]int{12}, []int{10,16})) // 10

    fmt.Println(maxDivScore1([]int{4,7,9,3,9}, []int{5,2,3})) // 3
    fmt.Println(maxDivScore1([]int{20,14,21,10}, []int{5,7,5})) // 5
    fmt.Println(maxDivScore1([]int{12}, []int{10,16})) // 10
}