package main

// 1822. Sign of the Product of an Array
// There is a function signFunc(x) that returns:
//     1 if x is positive.
//     -1 if x is negative.
//     0 if x is equal to 0.

// You are given an integer array nums. Let product be the product of all values in the array nums.
// Return signFunc(product).

// Example 1:
// Input: nums = [-1,-2,-3,-4,3,2,1]
// Output: 1
// Explanation: The product of all values in the array is 144, and signFunc(144) = 1

// Example 2:
// Input: nums = [1,5,0,2,-3]
// Output: 0
// Explanation: The product of all values in the array is 0, and signFunc(0) = 0

// Example 3:
// Input: nums = [-1,1,-1,1,-1]
// Output: -1
// Explanation: The product of all values in the array is -1, and signFunc(-1) = -1
 
// Constraints:
//     1 <= nums.length <= 1000
//     -100 <= nums[i] <= 100

import "fmt"

func arraySign(nums []int) int {
    res := 0
    for _,v := range nums {
        // 出现 0 直接返回 
        if 0 == v {
            return 0
        }
        if v < 0 {
            res++
        }
    }
    // 偶数个负数
    if res % 2 == 0 {
        return 1
    }
    return -1
}

func arraySign1(nums []int) int {
    res := 1
    for i:= 0; i < len(nums); i++{
        if nums[i] == 0{
            return 0
        }
        if nums[i] < 0 {
            res = -res
        }
    }
    return res
}

func main() {
    fmt.Println(arraySign([]int{-1,-2,-3,-4,3,2,1})) // 1
    fmt.Println(arraySign([]int{1,5,0,2,-3})) // 0
    fmt.Println(arraySign([]int{-1,1,-1,1,-1})) // -1
    fmt.Println(arraySign([]int{1,2,3,4,5,6})) // 1

    fmt.Println(arraySign1([]int{-1,-2,-3,-4,3,2,1})) // 1
    fmt.Println(arraySign1([]int{1,5,0,2,-3})) // 0
    fmt.Println(arraySign1([]int{-1,1,-1,1,-1})) // -1
    fmt.Println(arraySign1([]int{1,2,3,4,5,6})) // 1
}