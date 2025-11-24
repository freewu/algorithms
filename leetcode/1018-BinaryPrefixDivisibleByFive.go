package main

// 1018. Binary Prefix Divisible By 5
// You are given a binary array nums (0-indexed).

// We define xi as the number whose binary representation is the subarray nums[0..i] (from most-significant-bit to least-significant-bit).
//     For example, if nums = [1,0,1], then x0 = 1, x1 = 2, and x2 = 5.

// Return an array of booleans answer where answer[i] is true if xi is divisible by 5.

// Example 1:
// Input: nums = [0,1,1]
// Output: [true,false,false]
// Explanation: The input numbers in binary are 0, 01, 011; which are 0, 1, and 3 in base-10.
// Only the first number is divisible by 5, so answer[0] is true.

// Example 2:
// Input: nums = [1,1,1]
// Output: [false,false,false]

// Constraints:
//     1 <= nums.length <= 10^5
//     nums[i] is either 0 or 1.

import "fmt"

func prefixesDivBy5(nums []int) []bool {
    n, res := 0, make([]bool, len(nums))
    for i, v := range nums {
        n = (n << 1 + v) % 5
        res[i] = (n == 0)
    }
    return res
}

func prefixesDivBy51(nums []int) []bool {
    n, res := 0, []bool{}
    for _, v := range nums {
        n = n * 2 + v
        n = n % 5
        res = append(res, n == 0)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,1,1]
    // Output: [true,false,false]
    // Explanation: The input numbers in binary are 0, 01, 011; which are 0, 1, and 3 in base-10.
    // Only the first number is divisible by 5, so answer[0] is true.
    fmt.Println(prefixesDivBy5([]int{0,1,1})) // [true,false,false]
    // Example 2:
    // Input: nums = [1,1,1]
    // Output: [false,false,false]
    fmt.Println(prefixesDivBy5([]int{1,1,1})) // [false,false,false]

    fmt.Println(prefixesDivBy5([]int{1,2,3,4,5,6,7,8,9})) // [false false false false false true false false false]
    fmt.Println(prefixesDivBy5([]int{9,8,7,6,5,4,3,2,1})) // [false false false false false true false false false]

    fmt.Println(prefixesDivBy51([]int{0,1,1})) // [true,false,false]
    fmt.Println(prefixesDivBy51([]int{1,1,1})) // [false,false,false]
    fmt.Println(prefixesDivBy51([]int{1,2,3,4,5,6,7,8,9})) // [false false false false false true false false false]
    fmt.Println(prefixesDivBy51([]int{9,8,7,6,5,4,3,2,1})) // [false false false false false true false false false]
}