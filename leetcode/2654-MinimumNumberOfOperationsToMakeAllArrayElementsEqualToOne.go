package main

// 2654. Minimum Number of Operations to Make All Array Elements Equal to 1
// You are given a 0-indexed array nums consisiting of positive integers. 
// You can do the following operation on the array any number of times:
//     1. Select an index i such that 0 <= i < n - 1 and replace either of nums[i] or nums[i+1] with their gcd value.

// Return the minimum number of operations to make all elements of nums equal to 1. 
// If it is impossible, return -1.

// The gcd of two integers is the greatest common divisor of the two integers.

// Example 1:
// Input: nums = [2,6,3,4]
// Output: 4
// Explanation: We can do the following operations:
// - Choose index i = 2 and replace nums[2] with gcd(3,4) = 1. Now we have nums = [2,6,1,4].
// - Choose index i = 1 and replace nums[1] with gcd(6,1) = 1. Now we have nums = [2,1,1,4].
// - Choose index i = 0 and replace nums[0] with gcd(2,1) = 1. Now we have nums = [1,1,1,4].
// - Choose index i = 2 and replace nums[3] with gcd(1,4) = 1. Now we have nums = [1,1,1,1].

// Example 2:
// Input: nums = [2,10,6,14]
// Output: -1
// Explanation: It can be shown that it is impossible to make all the elements equal to 1.

// Constraints:
//     2 <= nums.length <= 50
//     1 <= nums[i] <= 10^6

import "fmt"

func minOperations(nums []int) int {
    res, n, one := 1 << 31, len(nums), 0
    for _, v := range nums {
        if v == 1 {
            one++
        }
    }
    if one > 0 { return n - one }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for i := 0; i < n; i++ {
        v := nums[i]
        for j := i + 1; j < n; j++ {
            v = gcd(v, nums[j])
            if v == 1 {
                res = min(res, j - i)
                break
            }
        }
        if v != 1 {
            break
        }
    }
    if res == 1 << 31 { return -1 }
    return n - 1 + res
}

func minOperations1(nums []int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    all, one, n := 0, 0, len(nums)
    for _, v := range nums {
        all = gcd(all, v)
        if v == 1 {
            one++
        }
    }
    if all > 1 { return -1 }
    if one > 0 { return n - one }
    mn := n
    for i := range nums {
        g := 0
        for j, v := range nums[i:] {
            g = gcd(g, v)
            if g == 1 {
                mn = min(mn, j)
                break
            }
        }
    }
    return mn + n - 1
}

func main() {
    // Example 1:
    // Input: nums = [2,6,3,4]
    // Output: 4
    // Explanation: We can do the following operations:
    // - Choose index i = 2 and replace nums[2] with gcd(3,4) = 1. Now we have nums = [2,6,1,4].
    // - Choose index i = 1 and replace nums[1] with gcd(6,1) = 1. Now we have nums = [2,1,1,4].
    // - Choose index i = 0 and replace nums[0] with gcd(2,1) = 1. Now we have nums = [1,1,1,4].
    // - Choose index i = 2 and replace nums[3] with gcd(1,4) = 1. Now we have nums = [1,1,1,1].
    fmt.Println(minOperations([]int{2,6,3,4})) // 4
    // Example 2:
    // Input: nums = [2,10,6,14]
    // Output: -1
    // Explanation: It can be shown that it is impossible to make all the elements equal to 1.
    fmt.Println(minOperations([]int{2,10,6,14})) // -1

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // 8
    fmt.Println(minOperations([]int{1,1,1,1})) // 0

    fmt.Println(minOperations1([]int{2,6,3,4})) // 4
    fmt.Println(minOperations1([]int{2,10,6,14})) // -1
    fmt.Println(minOperations1([]int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(minOperations1([]int{9,8,7,6,5,4,3,2,1})) // 8
    fmt.Println(minOperations1([]int{1,1,1,1})) // 0
}