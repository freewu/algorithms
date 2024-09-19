package main

// 1250. Check If It Is a Good Array
// Given an array nums of positive integers. 
// Your task is to select some subset of nums, multiply each element by an integer and add all these numbers. 
// The array is said to be good if you can obtain a sum of 1 from the array by any possible subset and multiplicand.

// Return True if the array is good otherwise return False.

// Example 1:
// Input: nums = [12,5,7,23]
// Output: true
// Explanation: Pick numbers 5 and 7.
// 5*3 + 7*(-2) = 1

// Example 2:
// Input: nums = [29,6,10]
// Output: true
// Explanation: Pick numbers 29, 6 and 10.
// 29*1 + 6*(-3) + 10*(-1) = 1

// Example 3:
// Input: nums = [3,6]
// Output: false

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

func isGoodArray(nums []int) bool {
    res := nums[0]
    gcd := func(j,k int) int {
        a, b := k, j
        if j > k {
            a, b = j, k 
        }
        for true {
            if a % b == 0 { return b }
            a, b = b, a % b
        }
        return -100
    }
    for i := 1; i < len(nums); i++ {
        res = gcd(res,nums[i])
        if res == 1 { break}
    }
    return res == 1
}

func isGoodArray1(nums []int) bool {
    // 裴蜀定理
    // g=gcd(a,b), 则对于任何的x,y都满足 ax+by是g的倍数,特别的,存在整数x,y,使得 ax+by=g,可以推广到多个整数
    // 所以题目转换为,是否存在子数组gcd==1, 因为gcd的性质(越多数参与,越小),所以从前往后遍历一遍即可
    gcd := func(a, b int) int {
        for b != 0 {
            a, b = b, a % b
        }
        return a
    }
    res := 0
    for _, v := range nums {
        res = gcd(res, v)
        if res == 1 {
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: nums = [12,5,7,23]
    // Output: true
    // Explanation: Pick numbers 5 and 7.
    // 5*3 + 7*(-2) = 1
    fmt.Println(isGoodArray([]int{12,5,7,23})) // true
    // Example 2:
    // Input: nums = [29,6,10]
    // Output: true
    // Explanation: Pick numbers 29, 6 and 10.
    // 29*1 + 6*(-3) + 10*(-1) = 1
    fmt.Println(isGoodArray([]int{29,6,10})) // true
    // Example 3:
    // Input: nums = [3,6]
    // Output: false
    fmt.Println(isGoodArray([]int{3,6})) // false

    fmt.Println(isGoodArray1([]int{12,5,7,23})) // true
    fmt.Println(isGoodArray1([]int{29,6,10})) // true
    fmt.Println(isGoodArray1([]int{3,6})) // false
}