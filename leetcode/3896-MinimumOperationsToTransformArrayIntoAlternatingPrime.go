package main

// 3896. Minimum Operations to Transform Array into Alternating Prime
// You are given an integer array nums.

// An array is considered alternating prime if:
//     1. Elements at even indices (0-based) are prime numbers.
//     2. Elements at odd indices are non-prime numbers.

// In one operation, you may increment any element by 1.

// Return the minimum number of operations required to transform nums into an alternating prime array.

// A prime number is a natural number greater than 1 with only two factors, 1 and itself.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: 3
// Explanation:
// The element at index 0 must be prime. Increment nums[0] = 1 to 2, using 1 operation.
// The element at index 1 must be non-prime. Increment nums[1] = 2 to 4, using 2 operations.
// The element at index 2 is already prime.
// The element at index 3 is already non-prime.
// Total operations = 1 + 2 = 3.

// Example 2:
// Input: nums = [5,6,7,8]
// Output: 0
// Explanation:
// The elements at indices 0 and 2 are already prime.
// The elements at indices 1 and 3 are already non-prime.
// No operations are needed.

// Example 3:
// Input: nums = [4,4]
// Output: 1
// Explanation:
// The element at index 0 must be prime. Increment nums[0] = 4 to 5, using 1 operation.
// The element at index 1 is already non-prime.
// Total operations = 1.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

var MX = 100050
var isPrime = make([]bool, MX)
var nextPrime = make([]int, MX)

func init() {
    for i := 0; i < MX; i++ {
        isPrime[i] = true
    }
    isPrime[0], isPrime[1] = false, false
    for i := 2; i*i < MX; i++ {
        if isPrime[i] {
            for j := i * i; j < MX; j += i {
                isPrime[j] = false
            }
        }
    }
    lastPrime := 2
    for i := MX - 1; i >= 0; i-- {
        if isPrime[i] {
            lastPrime = i
        }
        nextPrime[i] = lastPrime
    }
}

func minOperations(nums []int) int {
    res := 0
    for i, v := range nums {
        if i % 2 == 0 {
            if !isPrime[v] {
                res += nextPrime[v] - v
            }
        } else {
            if isPrime[v] {
                if v == 2 {
                    res += 2
                } else {
                    res += 1
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4]
    // Output: 3
    // Explanation:
    // The element at index 0 must be prime. Increment nums[0] = 1 to 2, using 1 operation.
    // The element at index 1 must be non-prime. Increment nums[1] = 2 to 4, using 2 operations.
    // The element at index 2 is already prime.
    // The element at index 3 is already non-prime.
    // Total operations = 1 + 2 = 3.
    fmt.Println(minOperations([]int{1,2,3,4})) // 3
    // Example 2:
    // Input: nums = [5,6,7,8]
    // Output: 0
    // Explanation:
    // The elements at indices 0 and 2 are already prime.
    // The elements at indices 1 and 3 are already non-prime.
    // No operations are needed.
    fmt.Println(minOperations([]int{5,6,7,8})) // 0
    // Example 3:
    // Input: nums = [4,4]
    // Output: 1
    // Explanation:
    // The element at index 0 must be prime. Increment nums[0] = 4 to 5, using 1 operation.
    // The element at index 1 is already non-prime.
    // Total operations = 1.
    fmt.Println(minOperations([]int{4,4})) // 1

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // 5
}