package main

// 3591. Check if Any Element Has Prime Frequency
// You are given an integer array nums.

// Return true if the frequency of any element of the array is prime, otherwise, return false.

// The frequency of an element x is the number of times it occurs in the array.

// A prime number is a natural number greater than 1 with only two factors, 1 and itself.

// Example 1:
// Input: nums = [1,2,3,4,5,4]
// Output: true
// Explanation:
// 4 has a frequency of two, which is a prime number.

// Example 2:
// Input: nums = [1,2,3,4,5]
// Output: false
// Explanation:
// All elements have a frequency of one.

// Example 3:
// Input: nums = [2,2,2,4,4]
// Output: true
// Explanation:
// Both 2 and 4 have a prime frequency.

// Constraints:
//     1 <= nums.length <= 100
//     0 <= nums[i] <= 100

import "fmt"

func checkPrimeFrequency(nums []int) bool {
    freq := make(map[int]int)
    for _, v := range nums {
        freq[v]++
    }
    checkPrime := func(n int) bool {
        if n < 2 { return false  }
        for i := 2; i*i <= n; i++ {
            if n % i == 0 {
                return false
            }
        }
        return true
    }
    for _, v := range freq {
        if checkPrime(v) {
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5,4]
    // Output: true
    // Explanation:
    // 4 has a frequency of two, which is a prime number.
    fmt.Println(checkPrimeFrequency([]int{1,2,3,4,5,4})) // true
    // Example 2:
    // Input: nums = [1,2,3,4,5]
    // Output: false
    // Explanation:
    // All elements have a frequency of one.
    fmt.Println(checkPrimeFrequency([]int{1,2,3,4,5})) // false
    // Example 3:
    // Input: nums = [2,2,2,4,4]
    // Output: true
    // Explanation:
    // Both 2 and 4 have a prime frequency.
    fmt.Println(checkPrimeFrequency([]int{2,2,2,4,4})) // true

    fmt.Println(checkPrimeFrequency([]int{1,2,3,4,5,6,7,8,9})) // false
    fmt.Println(checkPrimeFrequency([]int{9,8,7,6,5,4,3,2,1})) // false
}