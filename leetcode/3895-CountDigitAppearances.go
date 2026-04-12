package main

// 3895. Count Digit Appearances
// You are given an integer array nums and an integer digit.

// Return the total number of times digit appears in the decimal representation of all elements in nums.

// Example 1:
// Input: nums = [12,54,32,22], digit = 2
// Output: 4
// Explanation:
// The digit 2 appears once in 12 and 32, and twice in 22. Thus, the total number of times digit 2 appears is 4.

// Example 2:
// Input: nums = [1,34,7], digit = 9
// Output: 0
// Explanation:
// The digit 9 does not appear in the decimal representation of any element in nums, so the total number of times digit 9 appears is 0.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 10^6​​​​​​​
//     0 <= digit <= 9

import "fmt"

func countDigitOccurrences(nums []int, digit int) int {
    res := 0
    for _, v := range nums {
        for v > 0 {
            if v % 10 == digit {
                res++
            }
            v /= 10
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [12,54,32,22], digit = 2
    // Output: 4
    // Explanation:
    // The digit 2 appears once in 12 and 32, and twice in 22. Thus, the total number of times digit 2 appears is 4.
    fmt.Println(countDigitOccurrences([]int{12,54,32,22}, 2)) // 4
    // Example 2:
    // Input: nums = [1,34,7], digit = 9
    // Output: 0
    // Explanation:
    // The digit 9 does not appear in the decimal representation of any element in nums, so the total number of times digit 9 appears is 0.
    fmt.Println(countDigitOccurrences([]int{1,34,7}, 9)) // 0

    fmt.Println(countDigitOccurrences([]int{1,2,3,4,5,6,7,8,9}, 9)) // 1
    fmt.Println(countDigitOccurrences([]int{9,8,7,6,5,4,3,2,1}, 9)) // 1
}