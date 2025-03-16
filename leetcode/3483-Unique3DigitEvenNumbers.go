package main

// 3483. Unique 3-Digit Even Numbers
// You are given an array of digits called digits. 
// Your task is to determine the number of distinct three-digit even numbers that can be formed using these digits.

// Note: Each copy of a digit can only be used once per number, and there may not be leading zeros.

// Example 1:
// Input: digits = [1,2,3,4]
// Output: 12
// Explanation: The 12 distinct 3-digit even numbers that can be formed are 124, 132, 134, 142, 214, 234, 312, 314, 324, 342, 412, and 432. Note that 222 cannot be formed because there is only 1 copy of the digit 2.

// Example 2:
// Input: digits = [0,2,2]
// Output: 2
// Explanation: The only 3-digit even numbers that can be formed are 202 and 220. Note that the digit 2 can be used twice because it appears twice in the array.

// Example 3:
// Input: digits = [6,6,6]
// Output: 1
// Explanation: Only 666 can be formed.

// Example 4:
// Input: digits = [1,3,5]
// Output: 0
// Explanation: No even 3-digit numbers can be formed.

// Constraints:
//     3 <= digits.length <= 10
//     0 <= digits[i] <= 9

import "fmt"

func totalNumbers(digits []int) int {
    set := make(map[int]bool)
    n := len(digits)
    for i := 0; i < n; i++ {
        if digits[i] == 0 { continue }
        for j := 0; j < n; j++ {
            if j == i { continue }
            for k := 0; k < n; k++ {
                if k == i || k == j { continue }
                if digits[k] % 2 == 0 { // 偶数
                    key := digits[i] * 100 + digits[j] * 10 + digits[k] // 结尾为偶数
                    set[key] = true
                }
            }
        }
    }
    return len(set)
}

func main() {
    // Example 1:
    // Input: digits = [1,2,3,4]
    // Output: 12
    // Explanation: The 12 distinct 3-digit even numbers that can be formed are 124, 132, 134, 142, 214, 234, 312, 314, 324, 342, 412, and 432. Note that 222 cannot be formed because there is only 1 copy of the digit 2.
    fmt.Println(totalNumbers([]int{1,2,3,4})) // 12
    // Example 2:
    // Input: digits = [0,2,2]
    // Output: 2
    // Explanation: The only 3-digit even numbers that can be formed are 202 and 220. Note that the digit 2 can be used twice because it appears twice in the array.
    fmt.Println(totalNumbers([]int{0,2,2})) // 2
    // Example 3:
    // Input: digits = [6,6,6]
    // Output: 1
    // Explanation: Only 666 can be formed.
    fmt.Println(totalNumbers([]int{6,6,6})) // 1
    // Example 4:
    // Input: digits = [1,3,5]
    // Output: 0
    // Explanation: No even 3-digit numbers can be formed.
    fmt.Println(totalNumbers([]int{1,3,5})) // 0

    fmt.Println(totalNumbers([]int{1,2,3,4,5,6,7,8,9})) // 224
    fmt.Println(totalNumbers([]int{9,8,7,6,5,4,3,2,1})) // 224
}