package main

// 3115. Maximum Prime Difference
// You are given an integer array nums.
// Return an integer that is the maximum distance between the indices of two (not necessarily different) prime numbers in nums.

// Example 1:
// Input: nums = [4,2,9,5,3]
// Output: 3
// Explanation: nums[1], nums[3], and nums[4] are prime. So the answer is |4 - 1| = 3.

// Example 2:
// Input: nums = [4,8,2,8]
// Output: 0
// Explanation: nums[2] is prime. Because there is just one prime number, the answer is |2 - 2| = 0.

// Constraints:
//     1 <= nums.length <= 3 * 10^5
//     1 <= nums[i] <= 100
//     The input is generated such that the number of prime numbers in the nums is at least one.

import "fmt"
import "math"

func maximumPrimeDifference(nums []int) int {
    isPrime := func(n int) bool {
        if n == 1 {  return false; }
        for i := 2; i <= int(math.Sqrt(float64(n)+1)); i++ {
            if n % i == 0 { return false; }
        }
        return true
    }
    primes := []int{} // 记录所有质数出现的位置
    for i, n := range nums { // 找出所有质数
        if isPrime(n) {
            primes = append(primes, i) 
        }
    }
    return primes[len(primes)-1] - primes[0] // 最后一个减去第一个就是最大距离
}


func maximumPrimeDifference1(nums []int) int {
    isPrime := func(n int) bool {
        if n == 1 {  return false; }
        for i := 2; i <= int(math.Sqrt(float64(n)+1)); i++ {
            if n % i == 0 { return false; }
        }
        return true
    }
    start, end := 0, len(nums) - 1
    for !isPrime(nums[start]) { start++; } // 找到第一个质数
    for !isPrime(nums[end]) { end--; } // 找到最后一个质数
    return end - start
}

func main() {
    // Example 1:
    // Input: nums = [4,2,9,5,3]
    // Output: 3
    // Explanation: nums[1], nums[3], and nums[4] are prime. So the answer is |4 - 1| = 3.
    fmt.Println(maximumPrimeDifference([]int{4,2,9,5,3})) // 3
    // Example 2:
    // Input: nums = [4,8,2,8]
    // Output: 0
    // Explanation: nums[2] is prime. Because there is just one prime number, the answer is |2 - 2| = 0.
    fmt.Println(maximumPrimeDifference([]int{4,8,2,8})) // 0

    fmt.Println(maximumPrimeDifference1([]int{4,2,9,5,3})) // 3
    fmt.Println(maximumPrimeDifference1([]int{4,8,2,8})) // 0
}