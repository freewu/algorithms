package main

// 1390. Four Divisors
// Given an integer array nums, return the sum of divisors of the integers in that array that have exactly four divisors. 
// If there is no such integer in the array, return 0.

// Example 1:
// Input: nums = [21,4,7]
// Output: 32
// Explanation: 
// 21 has 4 divisors: 1, 3, 7, 21
// 4 has 3 divisors: 1, 2, 4
// 7 has 2 divisors: 1, 7
// The answer is the sum of divisors of 21 only.

// Example 2:
// Input: nums = [21,21]
// Output: 64

// Example 3:
// Input: nums = [1,2,3,4,5]
// Output: 0

// Constraints:
//     1 <= nums.length <= 10^4
//     1 <= nums[i] <= 10^5

import "fmt"
import "math"

func sumFourDivisors(nums []int) int {
    sum := 0
    calc := func(num int) int {
        if num <= 5 { return 0 }
        root := int(math.Sqrt(float64(num)))
        if root * root == num { return 0 }
        divisorCount, summary := 2, num + 1
        for i := 2; i <= root; i++ {
            if num % i == 0 {
                divisorCount += 2
                summary += i
                summary += num / i
                if divisorCount > 4 {
                    return 0
                }
            }
        }
        if divisorCount == 4 { return summary }
        return 0
    }
    for _, v := range nums {
        sum += calc(v)
    }
    return sum
}

func main() {
    // Example 1:
    // Input: nums = [21,4,7]
    // Output: 32
    // Explanation: 
    // 21 has 4 divisors: 1, 3, 7, 21
    // 4 has 3 divisors: 1, 2, 4
    // 7 has 2 divisors: 1, 7
    // The answer is the sum of divisors of 21 only.
    fmt.Println(sumFourDivisors([]int{21,4,7})) // 32
    // Example 2:
    // Input: nums = [21,21]
    // Output: 64
    fmt.Println(sumFourDivisors([]int{21,21})) // 64
    // Example 3:
    // Input: nums = [1,2,3,4,5]
    // Output: 0
    fmt.Println(sumFourDivisors([]int{1,2,3,4,5})) // 0
}