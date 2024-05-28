package main

// 360. Sort Transformed Array
// Given a sorted integer array nums and three integers a, b and c, 
// apply a quadratic function of the form f(x) = ax2 + bx + c to each element nums[i] in the array, 
// and return the array in a sorted order.

// Example 1:
// Input: nums = [-4,-2,2,4], a = 1, b = 3, c = 5
// Output: [3,9,15,33]

// Example 2:
// Input: nums = [-4,-2,2,4], a = -1, b = 3, c = 5
// Output: [-23,-5,1,7]
 
// Constraints:
//     1 <= nums.length <= 200
//     -100 <= nums[i], a, b, c <= 100
//     nums is sorted in ascending order.
 
// Follow up: Could you solve it in O(n) time?

import "fmt"

func sortTransformedArray(nums []int, a int, b int, c int) []int {
    res, cur := make([]int, len(nums)), 0
    if a > 0 {
        cur = len(res) - 1
    }
    for i, j := 0, len(nums) - 1; i <= j; {
        r1 := a * nums[i] * nums[i] + b * nums[i] + c // f(x) = ax2 + bx + c 
        r2 := a * nums[j] * nums[j] + b * nums[j] + c
        if a <= 0 {
            if r1 < r2 {
                res[cur] = r1
                i++
            } else {
                res[cur] = r2
                j--
            }
            cur++
        } else {
            if r1 > r2 {
                res[cur] = r1
                i++
            } else {
                res[cur] = r2
                j--
            }
            cur--
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [-4,-2,2,4], a = 1, b = 3, c = 5
    // Output: [3,9,15,33]
    fmt.Println(sortTransformedArray([]int{-4,-2,2,4}, 1, 3, 5)) // [3,9,15,33]
    // Example 2:
    // Input: nums = [-4,-2,2,4], a = -1, b = 3, c = 5
    // Output: [-23,-5,1,7]
    fmt.Println(sortTransformedArray([]int{-4,-2,2,4}, -1, 3, 5)) // [-23,-5,1,7]
}