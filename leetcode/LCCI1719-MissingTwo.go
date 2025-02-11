package main

// 面试题 17.19. Missing Two LCCI
// You are given an array with all the numbers from 1 to N appearing exactly once, except for two number that is missing. 
// How can you find the missing number in O(N) time and 0(1) space?

// You can return the missing numbers in any order.

// Example 1:
// Input: [1]
// Output: [2,3]

// Example 2:
// Input: [2,3]
// Output: [1,4]

// Note:
//     nums.length <= 30000

import "fmt"

func missingTwo(nums []int) []int {
    xorSum, n := 0, len(nums) + 2
    for _, v := range nums {
        xorSum ^= v
    }
    for i := 1; i <= n; i++ {
        xorSum ^= i
    }
    type1, type2, lsb := 0, 0, xorSum & -xorSum
    for _, v := range nums {
        if v & lsb > 0 {
            type1 ^= v
        } else {
            type2 ^= v
        }
    }
    for i := 1; i <= n; i++ {
        if i & lsb > 0 {
            type1 ^= i
        } else {
            type2 ^= i
        }
    }
    return []int{ type1, type2 }
}

func missingTwo1(nums []int) []int {
    n := len(nums) + 2
    sum := n*(1+n)>>1
    for _, v := range nums {
        sum -= v
    }
    t := sum / 2
    missingSum := (t + 1) * t / 2
    for _, v := range nums {
        if v <= t {
            missingSum -= v
        }
    }
    return []int{ missingSum, sum - missingSum }
}

func main() {
    // Example 1:
    // Input: [1]
    // Output: [2,3]
    fmt.Println(missingTwo([]int{1})) // [2,3]
    // Example 2:
    // Input: [2,3]
    // Output: [1,4]
    fmt.Println(missingTwo([]int{2,3})) // [1,4]

    fmt.Println(missingTwo([]int{1,2,3,4,5,6,7,8,9})) // [11 10]
    fmt.Println(missingTwo([]int{9,8,7,6,5,4,3,2,1})) // [11 10]

    fmt.Println(missingTwo1([]int{1})) // [2,3]
    fmt.Println(missingTwo1([]int{2,3})) // [1,4]
    fmt.Println(missingTwo1([]int{1,2,3,4,5,6,7,8,9})) // [11 10]
    fmt.Println(missingTwo1([]int{9,8,7,6,5,4,3,2,1})) // [11 10]
}