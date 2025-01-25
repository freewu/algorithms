package main

// 2778. Sum of Squares of Special Elements
// You are given a 1-indexed integer array nums of length n.

// An element nums[i] of nums is called special if i divides n, i.e. n % i == 0.

// Return the sum of the squares of all special elements of nums.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: 21
// Explanation: There are exactly 3 special elements in nums: nums[1] since 1 divides 4, nums[2] since 2 divides 4, and nums[4] since 4 divides 4. 
// Hence, the sum of the squares of all special elements of nums is nums[1] * nums[1] + nums[2] * nums[2] + nums[4] * nums[4] = 1 * 1 + 2 * 2 + 4 * 4 = 21.  

// Example 2:
// Input: nums = [2,7,1,19,18,3]
// Output: 63
// Explanation: There are exactly 4 special elements in nums: nums[1] since 1 divides 6, nums[2] since 2 divides 6, nums[3] since 3 divides 6, and nums[6] since 6 divides 6. 
// Hence, the sum of the squares of all special elements of nums is nums[1] * nums[1] + nums[2] * nums[2] + nums[3] * nums[3] + nums[6] * nums[6] = 2 * 2 + 7 * 7 + 1 * 1 + 3 * 3 = 63. 

// Constraints:
//     1 <= nums.length == n <= 50
//     1 <= nums[i] <= 50

import "fmt"

func sumOfSquares(nums []int) int {
    res, n := 0, len(nums)
    for i := 1; i < n + 1; i++ {
        if n % i == 0 {
            res += nums[i - 1] * nums[i - 1]
        }
    }
    return res
}

func sumOfSquares1(nums []int) int {
    res, n := 0, len(nums) 
    for i, v := range nums {
        if n % (i + 1) == 0 {
            res += v * v
        }
    } 
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4]
    // Output: 21
    // Explanation: There are exactly 3 special elements in nums: nums[1] since 1 divides 4, nums[2] since 2 divides 4, and nums[4] since 4 divides 4. 
    // Hence, the sum of the squares of all special elements of nums is nums[1] * nums[1] + nums[2] * nums[2] + nums[4] * nums[4] = 1 * 1 + 2 * 2 + 4 * 4 = 21.  
    fmt.Println(sumOfSquares([]int{1,2,3,4})) // 21
    // Example 2:
    // Input: nums = [2,7,1,19,18,3]
    // Output: 63
    // Explanation: There are exactly 4 special elements in nums: nums[1] since 1 divides 6, nums[2] since 2 divides 6, nums[3] since 3 divides 6, and nums[6] since 6 divides 6. 
    // Hence, the sum of the squares of all special elements of nums is nums[1] * nums[1] + nums[2] * nums[2] + nums[3] * nums[3] + nums[6] * nums[6] = 2 * 2 + 7 * 7 + 1 * 1 + 3 * 3 = 63. 
    fmt.Println(sumOfSquares([]int{2,7,1,19,18,3})) // 63

    fmt.Println(sumOfSquares([]int{1,2,3,4,5,6,7,8,9})) // 91
    fmt.Println(sumOfSquares([]int{9,8,7,6,5,4,3,2,1})) // 131

    fmt.Println(sumOfSquares1([]int{1,2,3,4})) // 21
    fmt.Println(sumOfSquares1([]int{2,7,1,19,18,3})) // 63
    fmt.Println(sumOfSquares1([]int{1,2,3,4,5,6,7,8,9})) // 91
    fmt.Println(sumOfSquares1([]int{9,8,7,6,5,4,3,2,1})) // 131
}